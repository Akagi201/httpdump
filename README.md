# httpdump

Simple Golang HTTP Server which dumps the incoming client headers and body

## Build
* docker: `docker build -t httpdump .`
* `go build main.go -o httpdump`

## Run
* `--host`: default host is `0.0.0.0`.
* `--port`: default port is `2222`.
