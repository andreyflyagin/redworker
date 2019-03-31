FROM golang:alpine AS builder
ADD . /src
RUN cd /src && go build -o worker -mod vendor

FROM alpine
WORKDIR /app
COPY --from=builder /src/worker /app/
CMD ./worker
