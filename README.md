# SOCKS5 Proxy Server

[![Release](https://img.shields.io/github/tag/tetafro/socks5.svg)](https://github.com/tetafro/socks5/releases)
[![CI](https://github.com/tetafro/socks5/actions/workflows/tag.yml/badge.svg)](https://github.com/tetafro/socks5/actions)

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

```sh
docker run --detach \
    --publish 1080:1080 \
    --name socks5 \
    --env USERNAME=bob \
    --env PASSWORD=qwerty \
    ghcr.io/tetafro/socks5
```
