# SOCKS5 Proxy Server

[![Go Report](https://goreportcard.com/badge/github.com/tetafro/socks5)](https://goreportcard.com/report/github.com/tetafro/socks5)
[![Release](https://img.shields.io/github/tag/tetafro/socks5.svg)](https://github.com/tetafro/socks5/releases)

SOCKS5 proxy server with password authentication. This is a simple wrapper
around [go-socks5](https://github.com/armon/go-socks5) library.

## Run

Build and run (server will listen on `0.0.0.0:1080` by default)
```sh
go get -u github.com/tetafro/socks5
USERNAME=bob PASSWORD=qwerty socks5
```

Use `HOST` and `PORT` to bind the server to a particular address.
```sh
HOST=127.0.0.1 PORT=8080 \
USERNAME=bob PASSWORD=qwerty \
socks5
```

## Run docker

Get docker [image](https://hub.docker.com/r/tetafro/socks5/) and run proxy
server in container on port 8088
```sh
docker run --detach \
    --publish 8088:1080 \
    --name socks5 \
    --env USERNAME=bob \
    --env PASSWORD=qwerty \
    ghcr.io/tetafro/socks5
```
