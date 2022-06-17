FROM golang:1.18-alpine3.16 AS build

WORKDIR /build

COPY . .
RUN go build -o ./bin/socks5

FROM alpine:3.16

WORKDIR /app

COPY --from=build /build/bin/socks5 /app/

RUN apk add --no-cache ca-certificates && \
    addgroup -S -g 5000 socks5 && \
    adduser -S -u 5000 -G socks5 socks5 && \
    chown -R socks5:socks5 .

USER socks5

EXPOSE 1080

CMD /app/socks5
