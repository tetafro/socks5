FROM golang:1.10-alpine AS build

WORKDIR /go/src/github.com/tetafro/socks5

COPY . .
RUN go build -o ./bin/socks5

FROM alpine:3.7

COPY --from=build /go/src/github.com/tetafro/socks5/bin/socks5 /app/

EXPOSE 1080

CMD /app/socks5
