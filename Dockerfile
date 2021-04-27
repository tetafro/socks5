FROM golang:1.16-alpine3.13 AS build

WORKDIR /go/src/github.com/tetafro/socks5

COPY . .
RUN go build -o ./bin/socks5

FROM alpine:3.13

COPY --from=build /go/src/github.com/tetafro/socks5/bin/socks5 /app/

EXPOSE 1080

CMD /app/socks5
