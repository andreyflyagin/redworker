#### usage
init swarm 

    docker swarm init

build & run

    docker stack rm red
    docker-compose build && docker stack deploy -c docker-compose.yaml red && docker service logs red_worker -f

get errors

    docker run --rm --network red_default rednode ./worker getErrors


#### kill node
    red_worker.21.0srntxbisnc3@linuxkit-025000000001    | 2019/03/31 18:36:18.774903 [info] bc2dc9eab9bf, received: hello     6 from a57c2d6bdc9e false 

kill master

    docker kill a57c2d6bdc9e

kill slave

    docker kill bc2dc9eab9bf

#### simple output 
    Successfully built b7dce10227ed
    Successfully tagged rednode:latest
    redis uses an image, skipping
    Ignoring unsupported options: build
    
    Creating network red_default
    Creating service red_redis
    Creating service red_worker
    red_worker.4.k8cz7yupfbek@linuxkit-025000000001     | 2019/03/31 19:33:35.036857 [info] started
    red_worker.4.k8cz7yupfbek@linuxkit-025000000001     | 2019/03/31 19:33:35.037430 [info] hostname: 5af9532bf83f
    red_worker.4.k8cz7yupfbek@linuxkit-025000000001     | 2019/03/31 19:33:35.045016 [info] failed ping
    red_worker.23.s7d9l5288bz4@linuxkit-025000000001    | 2019/03/31 19:33:35.324358 [info] started
    red_worker.23.s7d9l5288bz4@linuxkit-025000000001    | 2019/03/31 19:33:35.324688 [info] hostname: 019dba4f997a
    red_worker.23.s7d9l5288bz4@linuxkit-025000000001    | 2019/03/31 19:33:35.330359 [info] failed ping
    red_worker.4.k8cz7yupfbek@linuxkit-025000000001     | 2019/03/31 19:33:36.053554 [info] failed ping
    red_worker.23.s7d9l5288bz4@linuxkit-025000000001    | 2019/03/31 19:33:36.339483 [info] failed ping
    red_worker.4.k8cz7yupfbek@linuxkit-025000000001     | 2019/03/31 19:33:37.062048 [info] failed ping
    red_worker.24.9k1zazs0uwr4@linuxkit-025000000001    | 2019/03/31 19:33:37.065237 [info] started
    red_worker.24.9k1zazs0uwr4@linuxkit-025000000001    | 2019/03/31 19:33:37.065610 [info] hostname: 53dad3520e31
    red_worker.24.9k1zazs0uwr4@linuxkit-025000000001    | 2019/03/31 19:33:37.081622 [info] failed ping
    red_worker.23.s7d9l5288bz4@linuxkit-025000000001    | 2019/03/31 19:33:37.346678 [info] failed ping
    red_worker.4.k8cz7yupfbek@linuxkit-025000000001     | 2019/03/31 19:33:38.113324 [info] failed ping
    red_worker.24.9k1zazs0uwr4@linuxkit-025000000001    | 2019/03/31 19:33:38.117625 [info] failed ping
    red_worker.23.s7d9l5288bz4@linuxkit-025000000001    | 2019/03/31 19:33:38.352968 [info] failed ping
    red_worker.18.b9f9yk24wsik@linuxkit-025000000001    | 2019/03/31 19:33:38.571850 [info] started
    red_worker.18.b9f9yk24wsik@linuxkit-025000000001    | 2019/03/31 19:33:38.571981 [info] hostname: 19a7a730badf
    red_worker.18.b9f9yk24wsik@linuxkit-025000000001    | 2019/03/31 19:33:38.579113 [info] failed ping
    red_worker.4.k8cz7yupfbek@linuxkit-025000000001     | 2019/03/31 19:33:39.122698 [info] failed ping
    red_worker.24.9k1zazs0uwr4@linuxkit-025000000001    | 2019/03/31 19:33:39.128766 [info] failed ping
    red_worker.17.yepse9l8u512@linuxkit-025000000001    | 2019/03/31 19:33:39.305089 [info] started
    red_worker.17.yepse9l8u512@linuxkit-025000000001    | 2019/03/31 19:33:39.305557 [info] hostname: 5208ccf9aff8
    red_worker.17.yepse9l8u512@linuxkit-025000000001    | 2019/03/31 19:33:39.312074 [info] failed ping
    red_worker.23.s7d9l5288bz4@linuxkit-025000000001    | 2019/03/31 19:33:39.361657 [info] failed ping
    red_worker.18.b9f9yk24wsik@linuxkit-025000000001    | 2019/03/31 19:33:39.587947 [info] failed ping
    red_worker.4.k8cz7yupfbek@linuxkit-025000000001     | 2019/03/31 19:33:40.130105 [info] failed ping
    red_worker.24.9k1zazs0uwr4@linuxkit-025000000001    | 2019/03/31 19:33:40.135267 [info] failed ping
    red_worker.17.yepse9l8u512@linuxkit-025000000001    | 2019/03/31 19:33:40.319034 [info] failed ping
    red_worker.23.s7d9l5288bz4@linuxkit-025000000001    | 2019/03/31 19:33:40.370531 [info] failed ping
    red_worker.18.b9f9yk24wsik@linuxkit-025000000001    | 2019/03/31 19:33:40.597927 [info] failed ping
    red_worker.5.ld7ffzkboyxo@linuxkit-025000000001     | 2019/03/31 19:33:40.791120 [info] started
    red_worker.5.ld7ffzkboyxo@linuxkit-025000000001     | 2019/03/31 19:33:40.791554 [info] hostname: 4a2577b1f544
    red_worker.5.ld7ffzkboyxo@linuxkit-025000000001     | 2019/03/31 19:33:40.797809 [info] failed ping
    red_worker.4.k8cz7yupfbek@linuxkit-025000000001     | 2019/03/31 19:33:41.143751 [info] failed ping
    red_worker.24.9k1zazs0uwr4@linuxkit-025000000001    | 2019/03/31 19:33:41.147930 [info] failed ping
    red_worker.11.43xj5asdssr9@linuxkit-025000000001    | 2019/03/31 19:33:41.127298 [info] started
    red_worker.11.43xj5asdssr9@linuxkit-025000000001    | 2019/03/31 19:33:41.127806 [info] hostname: 164906e9f6ad
    red_worker.11.43xj5asdssr9@linuxkit-025000000001    | 2019/03/31 19:33:41.143758 [info] failed ping
    red_worker.19.xnt3m09vvw4q@linuxkit-025000000001    | 2019/03/31 19:33:41.181659 [info] started
    red_worker.19.xnt3m09vvw4q@linuxkit-025000000001    | 2019/03/31 19:33:41.182609 [info] hostname: 8d4ab763b4e9
    red_worker.19.xnt3m09vvw4q@linuxkit-025000000001    | 2019/03/31 19:33:41.190371 [info] failed ping
    red_worker.17.yepse9l8u512@linuxkit-025000000001    | 2019/03/31 19:33:41.325730 [info] failed ping
    red_worker.16.0hyygjyy7s98@linuxkit-025000000001    | 2019/03/31 19:33:41.341632 [info] started
    red_worker.16.0hyygjyy7s98@linuxkit-025000000001    | 2019/03/31 19:33:41.342166 [info] hostname: e069a7574696
    red_worker.16.0hyygjyy7s98@linuxkit-025000000001    | 2019/03/31 19:33:41.350327 [info] failed ping
    red_worker.23.s7d9l5288bz4@linuxkit-025000000001    | 2019/03/31 19:33:41.376526 [info] failed ping
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:33:41.515081 [info] started
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:33:41.515713 [info] hostname: c07b6a9f1da7
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:33:41.521975 [info] failed ping
    red_worker.18.b9f9yk24wsik@linuxkit-025000000001    | 2019/03/31 19:33:41.605713 [info] failed ping
    red_worker.5.ld7ffzkboyxo@linuxkit-025000000001     | 2019/03/31 19:33:41.805633 [info] failed ping
    red_worker.6.usmgqho90f6s@linuxkit-025000000001     | 2019/03/31 19:33:41.840526 [info] started
    red_worker.6.usmgqho90f6s@linuxkit-025000000001     | 2019/03/31 19:33:41.841049 [info] hostname: 8b7bd72f076c
    red_worker.6.usmgqho90f6s@linuxkit-025000000001     | 2019/03/31 19:33:41.847174 [info] failed ping
    red_worker.21.dxnsszbki7fv@linuxkit-025000000001    | 2019/03/31 19:33:42.031452 [info] started
    red_worker.21.dxnsszbki7fv@linuxkit-025000000001    | 2019/03/31 19:33:42.031813 [info] hostname: d65535628ef2
    red_worker.21.dxnsszbki7fv@linuxkit-025000000001    | 2019/03/31 19:33:42.038177 [info] failed ping
    red_worker.11.43xj5asdssr9@linuxkit-025000000001    | 2019/03/31 19:33:42.159834 [info] failed ping
    red_worker.24.9k1zazs0uwr4@linuxkit-025000000001    | 2019/03/31 19:33:42.160081 [info] failed ping
    red_worker.4.k8cz7yupfbek@linuxkit-025000000001     | 2019/03/31 19:33:42.159577 [info] failed ping
    red_worker.19.xnt3m09vvw4q@linuxkit-025000000001    | 2019/03/31 19:33:42.200510 [info] failed ping
    red_worker.8.zgtqii4fju70@linuxkit-025000000001     | 2019/03/31 19:33:42.204902 [info] started
    red_worker.8.zgtqii4fju70@linuxkit-025000000001     | 2019/03/31 19:33:42.208919 [info] hostname: 9ad92264b3ec
    red_worker.8.zgtqii4fju70@linuxkit-025000000001     | 2019/03/31 19:33:42.217327 [info] failed ping
    red_worker.13.shkfpecqh0ru@linuxkit-025000000001    | 2019/03/31 19:33:42.237032 [info] started
    red_worker.13.shkfpecqh0ru@linuxkit-025000000001    | 2019/03/31 19:33:42.237478 [info] hostname: b5a582372da6
    red_worker.13.shkfpecqh0ru@linuxkit-025000000001    | 2019/03/31 19:33:42.245542 [info] failed ping
    red_worker.17.yepse9l8u512@linuxkit-025000000001    | 2019/03/31 19:33:42.336498 [info] failed ping
    red_worker.16.0hyygjyy7s98@linuxkit-025000000001    | 2019/03/31 19:33:42.356578 [info] failed ping
    red_worker.23.s7d9l5288bz4@linuxkit-025000000001    | 2019/03/31 19:33:42.383409 [info] failed ping
    red_worker.15.psxg7ip9wmoc@linuxkit-025000000001    | 2019/03/31 19:33:42.372892 [info] started
    red_worker.15.psxg7ip9wmoc@linuxkit-025000000001    | 2019/03/31 19:33:42.373441 [info] hostname: cb64fdd780b2
    red_worker.15.psxg7ip9wmoc@linuxkit-025000000001    | 2019/03/31 19:33:42.382803 [info] failed ping
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:33:42.524203 [info] slave: c07b6a9f1da7
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:33:42.524695 [info] c07b6a9f1da7, received: hello     0 from 86b24accd0bb false
    red_worker.12.lend53sl6l18@linuxkit-025000000001    | 2019/03/31 19:33:42.510965 [info] started
    red_worker.12.lend53sl6l18@linuxkit-025000000001    | 2019/03/31 19:33:42.511303 [info] hostname: 86b24accd0bb
    red_worker.12.lend53sl6l18@linuxkit-025000000001    | 2019/03/31 19:33:42.514488 [info] master: 86b24accd0bb, 4OHOTICW70
    red_worker.18.b9f9yk24wsik@linuxkit-025000000001    | 2019/03/31 19:33:42.607495 [info] slave: 19a7a730badf
    red_worker.7.t72g0v0oe0mx@linuxkit-025000000001     | 2019/03/31 19:33:42.608721 [info] started
    red_worker.7.t72g0v0oe0mx@linuxkit-025000000001     | 2019/03/31 19:33:42.611151 [info] hostname: 7fb28ac71cf7
    red_worker.7.t72g0v0oe0mx@linuxkit-025000000001     | 2019/03/31 19:33:42.613634 [info] slave: 7fb28ac71cf7
    red_worker.10.vvmeaed6tfna@linuxkit-025000000001    | 2019/03/31 19:33:42.737718 [info] started
    red_worker.10.vvmeaed6tfna@linuxkit-025000000001    | 2019/03/31 19:33:42.738097 [info] hostname: 79d90819a453
    red_worker.10.vvmeaed6tfna@linuxkit-025000000001    | 2019/03/31 19:33:42.740393 [info] slave: 79d90819a453
    red_worker.5.ld7ffzkboyxo@linuxkit-025000000001     | 2019/03/31 19:33:42.809729 [info] slave: 4a2577b1f544
    red_worker.6.usmgqho90f6s@linuxkit-025000000001     | 2019/03/31 19:33:42.850452 [info] slave: 8b7bd72f076c
    red_worker.21.dxnsszbki7fv@linuxkit-025000000001    | 2019/03/31 19:33:43.041259 [info] slave: d65535628ef2
    red_worker.21.dxnsszbki7fv@linuxkit-025000000001    | 2019/03/31 19:33:43.044017 [info] d65535628ef2, received: hello     1 from 86b24accd0bb false
    red_worker.11.43xj5asdssr9@linuxkit-025000000001    | 2019/03/31 19:33:43.162449 [info] slave: 164906e9f6ad
    red_worker.24.9k1zazs0uwr4@linuxkit-025000000001    | 2019/03/31 19:33:43.177696 [info] slave: 53dad3520e31
    red_worker.4.k8cz7yupfbek@linuxkit-025000000001     | 2019/03/31 19:33:43.177124 [info] slave: 5af9532bf83f
    red_worker.19.xnt3m09vvw4q@linuxkit-025000000001    | 2019/03/31 19:33:43.203180 [info] slave: 8d4ab763b4e9
    red_worker.8.zgtqii4fju70@linuxkit-025000000001     | 2019/03/31 19:33:43.221801 [info] slave: 9ad92264b3ec
    red_worker.25.vgxbmsj22kv7@linuxkit-025000000001    | 2019/03/31 19:33:43.216919 [info] started
    red_worker.25.vgxbmsj22kv7@linuxkit-025000000001    | 2019/03/31 19:33:43.218750 [info] hostname: a49cb5b64664
    red_worker.25.vgxbmsj22kv7@linuxkit-025000000001    | 2019/03/31 19:33:43.231270 [info] slave: a49cb5b64664
    red_worker.13.shkfpecqh0ru@linuxkit-025000000001    | 2019/03/31 19:33:43.248476 [info] slave: b5a582372da6
    red_worker.1.n01eh5hr1jdu@linuxkit-025000000001     | 2019/03/31 19:33:43.246900 [info] started
    red_worker.1.n01eh5hr1jdu@linuxkit-025000000001     | 2019/03/31 19:33:43.247622 [info] hostname: 06706143d0e8
    red_worker.1.n01eh5hr1jdu@linuxkit-025000000001     | 2019/03/31 19:33:43.254737 [info] slave: 06706143d0e8
    red_worker.17.yepse9l8u512@linuxkit-025000000001    | 2019/03/31 19:33:43.339940 [info] slave: 5208ccf9aff8
    red_worker.16.0hyygjyy7s98@linuxkit-025000000001    | 2019/03/31 19:33:43.359481 [info] slave: e069a7574696
    red_worker.23.s7d9l5288bz4@linuxkit-025000000001    | 2019/03/31 19:33:43.387024 [info] slave: 019dba4f997a
    red_worker.15.psxg7ip9wmoc@linuxkit-025000000001    | 2019/03/31 19:33:43.387881 [info] slave: cb64fdd780b2
    red_worker.20.zufzrx7kms8v@linuxkit-025000000001    | 2019/03/31 19:33:43.467658 [info] started
    red_worker.20.zufzrx7kms8v@linuxkit-025000000001    | 2019/03/31 19:33:43.468231 [info] hostname: 71715458d4dc
    red_worker.20.zufzrx7kms8v@linuxkit-025000000001    | 2019/03/31 19:33:43.471369 [info] slave: 71715458d4dc
    red_worker.2.xdh888x8qofr@linuxkit-025000000001     | 2019/03/31 19:33:43.526340 [info] started
    red_worker.2.xdh888x8qofr@linuxkit-025000000001     | 2019/03/31 19:33:43.527267 [info] hostname: 70bee845fd66
    red_worker.2.xdh888x8qofr@linuxkit-025000000001     | 2019/03/31 19:33:43.531174 [info] slave: 70bee845fd66
    red_worker.2.xdh888x8qofr@linuxkit-025000000001     | 2019/03/31 19:33:43.531803 [info] 70bee845fd66, received: hello     2 from 86b24accd0bb false
    red_worker.14.nakl25epbxa3@linuxkit-025000000001    | 2019/03/31 19:33:43.647620 [info] started
    red_worker.14.nakl25epbxa3@linuxkit-025000000001    | 2019/03/31 19:33:43.648635 [info] hostname: 37b09afcac45
    red_worker.14.nakl25epbxa3@linuxkit-025000000001    | 2019/03/31 19:33:43.653736 [info] slave: 37b09afcac45
    red_worker.3.lwrtg8inxyht@linuxkit-025000000001     | 2019/03/31 19:33:43.646392 [info] started
    red_worker.3.lwrtg8inxyht@linuxkit-025000000001     | 2019/03/31 19:33:43.647060 [info] hostname: 2a9ff237a04c
    red_worker.3.lwrtg8inxyht@linuxkit-025000000001     | 2019/03/31 19:33:43.652932 [info] slave: 2a9ff237a04c
    red_worker.22.k1qvjr3zfoid@linuxkit-025000000001    | 2019/03/31 19:33:43.678042 [info] started
    red_worker.22.k1qvjr3zfoid@linuxkit-025000000001    | 2019/03/31 19:33:43.679230 [info] hostname: 3c93e155d546
    red_worker.22.k1qvjr3zfoid@linuxkit-025000000001    | 2019/03/31 19:33:43.686356 [info] slave: 3c93e155d546
    red_worker.8.zgtqii4fju70@linuxkit-025000000001     | 2019/03/31 19:33:44.025979 [info] 9ad92264b3ec, received: hello     3 from 86b24accd0bb false
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:33:44.533440 [info] c07b6a9f1da7, received: hello     4 from 86b24accd0bb false
    red_worker.21.dxnsszbki7fv@linuxkit-025000000001    | 2019/03/31 19:33:45.055083 [info] d65535628ef2, received: hello     5 from 86b24accd0bb false
    red_worker.2.xdh888x8qofr@linuxkit-025000000001     | 2019/03/31 19:33:45.541949 [info] 70bee845fd66, received: hello     6 from 86b24accd0bb false
    red_worker.8.zgtqii4fju70@linuxkit-025000000001     | 2019/03/31 19:33:46.039226 [info] 9ad92264b3ec, received: hello     7 from 86b24accd0bb false
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:33:46.548185 [info] c07b6a9f1da7, received: hello     8 from 86b24accd0bb false
    red_worker.21.dxnsszbki7fv@linuxkit-025000000001    | 2019/03/31 19:33:47.067328 [info] d65535628ef2, received: hello     9 from 86b24accd0bb false
    red_worker.2.xdh888x8qofr@linuxkit-025000000001     | 2019/03/31 19:33:47.551862 [info] 70bee845fd66, received: hello    10 from 86b24accd0bb false
    red_worker.8.zgtqii4fju70@linuxkit-025000000001     | 2019/03/31 19:33:48.050602 [info] 9ad92264b3ec, received: hello    11 from 86b24accd0bb false
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:33:48.559054 [info] c07b6a9f1da7, received: hello    12 from 86b24accd0bb false
    red_worker.21.dxnsszbki7fv@linuxkit-025000000001    | 2019/03/31 19:33:49.077441 [info] d65535628ef2, received: hello    13 from 86b24accd0bb false
    red_worker.2.xdh888x8qofr@linuxkit-025000000001     | 2019/03/31 19:33:49.564791 [info] 70bee845fd66, received: hello    14 from 86b24accd0bb false
    red_worker.8.zgtqii4fju70@linuxkit-025000000001     | 2019/03/31 19:33:50.060250 [info] 9ad92264b3ec, received: hello    15 from 86b24accd0bb false
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:33:50.569069 [info] c07b6a9f1da7, received: hello    16 from 86b24accd0bb false
    red_worker.21.dxnsszbki7fv@linuxkit-025000000001    | 2019/03/31 19:33:51.087607 [info] d65535628ef2, received: hello    17 from 86b24accd0bb false
    red_worker.2.xdh888x8qofr@linuxkit-025000000001     | 2019/03/31 19:33:51.576785 [info] 70bee845fd66, received: hello    18 from 86b24accd0bb true
    red_worker.25.vgxbmsj22kv7@linuxkit-025000000001    | 2019/03/31 19:33:52.070750 [info] a49cb5b64664, received: hello    19 from 86b24accd0bb false
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:33:52.577611 [info] c07b6a9f1da7, received: hello    20 from 86b24accd0bb false
    red_worker.21.dxnsszbki7fv@linuxkit-025000000001    | 2019/03/31 19:33:53.097328 [info] d65535628ef2, received: hello    21 from 86b24accd0bb false
    red_worker.2.xdh888x8qofr@linuxkit-025000000001     | 2019/03/31 19:33:53.589857 [info] 70bee845fd66, received: hello    22 from 86b24accd0bb false
    red_worker.8.zgtqii4fju70@linuxkit-025000000001     | 2019/03/31 19:33:54.079850 [info] 9ad92264b3ec, received: hello    23 from 86b24accd0bb false
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:33:54.587705 [info] c07b6a9f1da7, received: hello    24 from 86b24accd0bb false
    red_worker.21.dxnsszbki7fv@linuxkit-025000000001    | 2019/03/31 19:33:55.104752 [info] d65535628ef2, received: hello    25 from 86b24accd0bb false
    red_worker.2.xdh888x8qofr@linuxkit-025000000001     | 2019/03/31 19:33:55.600632 [info] 70bee845fd66, received: hello    26 from 86b24accd0bb false
    red_worker.25.vgxbmsj22kv7@linuxkit-025000000001    | 2019/03/31 19:33:56.090143 [info] a49cb5b64664, received: hello    27 from 86b24accd0bb false
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:33:56.600669 [info] c07b6a9f1da7, received: hello    28 from 86b24accd0bb false
    red_worker.18.b9f9yk24wsik@linuxkit-025000000001    | 2019/03/31 19:33:57.085643 [info] 19a7a730badf, received: hello    29 from 86b24accd0bb false
    red_worker.2.xdh888x8qofr@linuxkit-025000000001     | 2019/03/31 19:33:57.618407 [info] 70bee845fd66, received: hello    30 from 86b24accd0bb false
    red_worker.5.ld7ffzkboyxo@linuxkit-025000000001     | 2019/03/31 19:33:58.088959 [info] 4a2577b1f544, received: hello    31 from 86b24accd0bb false
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:33:58.613452 [info] c07b6a9f1da7, received: hello    32 from 86b24accd0bb false
    red_worker.7.t72g0v0oe0mx@linuxkit-025000000001     | 2019/03/31 19:33:59.096801 [info] 7fb28ac71cf7, received: hello    33 from 86b24accd0bb false
    red_worker.10.vvmeaed6tfna@linuxkit-025000000001    | 2019/03/31 19:33:59.629598 [info] 79d90819a453, received: hello    34 from 86b24accd0bb false
    red_worker.25.vgxbmsj22kv7@linuxkit-025000000001    | 2019/03/31 19:34:00.108772 [info] a49cb5b64664, received: hello    35 from 86b24accd0bb false
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:34:00.588848 [info] c07b6a9f1da7, received: hello    36 from 86b24accd0bb false
    red_worker.18.b9f9yk24wsik@linuxkit-025000000001    | 2019/03/31 19:34:01.073313 [info] 19a7a730badf, received: hello    37 from 86b24accd0bb false
    red_worker.10.vvmeaed6tfna@linuxkit-025000000001    | 2019/03/31 19:34:01.609235 [info] 79d90819a453, received: hello    38 from 86b24accd0bb false
    red_worker.25.vgxbmsj22kv7@linuxkit-025000000001    | 2019/03/31 19:34:02.085261 [info] a49cb5b64664, received: hello    39 from 86b24accd0bb false
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:34:02.600954 [info] c07b6a9f1da7, received: hello    40 from 86b24accd0bb false
    red_worker.18.b9f9yk24wsik@linuxkit-025000000001    | 2019/03/31 19:34:03.085530 [info] 19a7a730badf, received: hello    41 from 86b24accd0bb false
    red_worker.10.vvmeaed6tfna@linuxkit-025000000001    | 2019/03/31 19:34:03.620516 [info] 79d90819a453, received: hello    42 from 86b24accd0bb false
    red_worker.8.zgtqii4fju70@linuxkit-025000000001     | 2019/03/31 19:34:04.095845 [info] 9ad92264b3ec, received: hello    43 from 86b24accd0bb false
    red_worker.9.xg0wckv7gplc@linuxkit-025000000001     | 2019/03/31 19:34:04.609665 [info] c07b6a9f1da7, received: hello    44 from 86b24accd0bb false
    red_worker.18.b9f9yk24wsik@linuxkit-025000000001    | 2019/03/31 19:34:05.097597 [info] 19a7a730badf, received: hello    45 from 86b24accd0bb false
    red_worker.2.xdh888x8qofr@linuxkit-025000000001     | 2019/03/31 19:34:05.632649 [info] 70bee845fd66, received: hello    46 from 86b24accd0bb false


#### simple output error colleting
    2019/03/31 19:35:43.468893 [info] started
    2019/03/31 19:35:43.469653 [info] hostname: 3b2caa61e102
    2019/03/31 19:35:43.472511 [info] 3b2caa61e102, process error: hello    18 from 86b24accd0bb
    2019/03/31 19:35:43.472654 [info] 3b2caa61e102, process error: hello    76 from 86b24accd0bb
    2019/03/31 19:35:43.472842 [info] 3b2caa61e102, process error: hello   129 from 86b24accd0bb
    2019/03/31 19:35:43.473133 [info] 3b2caa61e102, process error: hello   148 from 86b24accd0bb
    2019/03/31 19:35:43.473345 [info] 3b2caa61e102, process error: hello   167 from 86b24accd0bb
    2019/03/31 19:35:43.473508 [info] 3b2caa61e102, process error: hello   192 from 86b24accd0bb
    2019/03/31 19:35:43.473712 [info] 3b2caa61e102, process error: hello   193 from 86b24accd0bb
    2019/03/31 19:35:43.473983 [info] 3b2caa61e102, process error: hello   227 from 86b24accd0bb
    2019/03/31 19:35:43.474204 [info] 3b2caa61e102, process error: hello   231 from 86b24accd0bb
    2019/03/31 19:35:43.474367 [info] exited