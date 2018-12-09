# SOCKS5 Proxy Server

[![Go Report](https://goreportcard.com/badge/github.com/tetafro/socks5)](https://goreportcard.com/report/github.com/tetafro/socks5)
[![Release](https://img.shields.io/github/tag/tetafro/socks5.svg)](https://github.com/tetafro/socks5/releases)

SOCKS5 proxy server with password authentication. This is a simple wrapper
around [go-socks5](https://github.com/armon/go-socks5) library.

## Run

Build and run (server will listen on `0.0.0.0:1080` by default)
```sh
go get -u github.com/tetafro/socks5
socks5 -user bob -password qwerty
```

You can also use `-host` and `-port` flags to bind the server to particular address.
```sh
socks5 -host 127.0.0.1 -port 8080 -user bob -password qwerty
```

## Run docker

Get docker image and run proxy server in container on port 8088
```sh
docker run -d -p 8088:1080 tetafro/socks5 \
    sh -c '/app/socks5 -username bob -password qwerty'
```
