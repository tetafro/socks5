.PHONY: dep
dep:
	go mod tidy && go mod verify

.PHONY: build
build:
	go build -o ./bin/socks5

.PHONY: lint
lint: go-lint yamllint ansible-lint

.PHONY: go-lint
go-lint:
	@ echo '----------------'
	@ echo 'Running golangci-lint'
	@ echo '----------------'
	@ golangci-lint run --fix && echo OK

.PHONY: yamllint
yamllint:
	@ echo '----------------'
	@ echo 'Running yamllint'
	@ echo '----------------'
	@ yamllint --format colored --strict ./playbook.yml && echo OK

.PHONY: ansible-lint
ansible-lint:
	@ echo '--------------------'
	@ echo 'Running ansible-lint'
	@ echo '--------------------'
	@ ansible-lint ./playbook.yml && echo OK

.PHONY: docker
docker:
	docker build -t ghcr.io/tetafro/socks5 .
