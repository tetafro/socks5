.PHONY: dep
dep:
	dep ensure -v

.PHONY: build
build:
	go build -o ./bin/socks5

.PHONY: docker
docker:
	docker build -t tetafro/socks5 .
