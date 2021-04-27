.PHONY: dep
dep:
	go mod tidy && go mod verify

.PHONY: build
build:
	go build -o ./bin/socks5

.PHONY: lint
lint:
	@ golangci-lint run

.PHONY: docker
docker:
	docker build -t tetafro/socks5 .
