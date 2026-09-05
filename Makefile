.PHONY: dep
dep:
	go mod tidy && go mod verify

.PHONY: lint
lint:
	@ golangci-lint run --fix

.PHONY: build
build:
	go build -o ./bin/socks5

.PHONY: run
run:
	@ ./bin/socks5

.PHONY: docker
docker:
	docker build -t ghcr.io/tetafro/socks5 .
