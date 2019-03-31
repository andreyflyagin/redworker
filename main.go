package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	electionKey          = "election"
	masterNameKey        = "master"
	messageQueue         = "mchannel"
	errorQueue           = "errors"
	electionKeyExpire    = time.Millisecond * 700
	messageInterval      = time.Millisecond * 500
	slaveConsumeInterval = time.Millisecond * 400
	redisAddr            = "redis:6379"
)

func getUnique() string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 10)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

type nodeSrv struct {
	client     *redis.Client
	ctx        context.Context
	hostname   string
	masterName string
}

// return 5% error
func (p *nodeSrv) hasError() bool {
	min := 1
	max := 20
	return rand.Intn(max-min)+min == 1
}

func (p *nodeSrv) master() {
	log.Printf("[info] master: %s, %s\n", p.hostname, p.masterName)

	var (
		err error
		i   int64
	)

	_, err = p.client.Set(masterNameKey, p.masterName, 0).Result()
	if err != nil {
		panic(err)
	}

	for {
		// in case message generation + sleeping takes too long and another node became a master, check it
		if v, err := p.client.Get(masterNameKey).Result(); err != nil || v != p.masterName {
			return
		}

		// hello    10 from a57c2d6bdc9e false
		msg := fmt.Sprintf("hello %5d from %s", i, p.hostname)
		i++

		_, err = p.client.RPush(messageQueue, msg).Result()
		if err != nil {
			panic(err)
		}

		// small node tries to master forever
		_, err = p.client.Set(electionKey, 1, electionKeyExpire).Result()
		if err != nil {
			panic(err)
		}

		select {
		case <-p.ctx.Done():
			return
		case <-time.After(messageInterval):
		}
	}
}

func (p *nodeSrv) slave() {
	log.Printf("[info] slave: %s\n", p.hostname)

	for {
		_, err := p.client.Get(electionKey).Result()
		if err != nil {
			return
		}

		v, err := p.client.LPop(messageQueue).Result()
		if err == nil {
			hasError := p.hasError()

			log.Printf("[info] %s, received: %s %v\n", p.hostname, v, hasError)
			if hasError {
				p.client.RPush(errorQueue, v)
			}
		}
		select {
		case <-time.After(slaveConsumeInterval):
		case <-p.ctx.Done():
			return
		}
	}
}

func (p *nodeSrv) run() {
	for {
		select {
		case <-p.ctx.Done():
			return
		default:
		}

		electKey, err := p.client.Incr(electionKey).Result()
		if err != nil {
			panic(err)
		}

		if electKey == 1 {
			p.master()
		} else {
			p.slave()
		}
	}
}

func getRedClient(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       1,
	})

	for {
		if err := client.Ping().Err(); err != nil {
			log.Printf("[info] failed ping\n")
			select {
			case <-time.After(time.Second):
				continue
			case <-ctx.Done():
			}
		}
		break
	}
	return client
}

func processError(ctx context.Context, client *redis.Client, hostname string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		v, err := client.LPop(errorQueue).Result()
		if err != nil {
			return
		}
		log.Printf("[info] %s, process error: %s \n", hostname, v)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	rand.Seed(time.Now().UnixNano())

	log.Printf("[info] started\n")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-sigChan
		cancel()
	}()

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	log.Printf("[info] hostname: %s\n", hostname)

	client := getRedClient(ctx)
	defer func() {
		err := client.Close()
		if err != nil {
			log.Printf("[error] failed close redis\n")
		}
	}()

	if len(os.Args) > 1 && os.Args[1] == "getErrors" {
		processError(ctx, client, hostname)
	} else {
		node := nodeSrv{
			client:     client,
			ctx:        ctx,
			hostname:   hostname,
			masterName: getUnique(),
		}
		node.run()
	}

	log.Printf("[info] exited\n")
}
