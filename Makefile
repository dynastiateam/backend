SHELL=/bin/bash
IMAGE_TAG := $(shell git rev-parse HEAD)

export GO111MODULE=on

.PHONY: ci
ci: deps lint test dockerise

.PHONY: deps
deps:
	go mod download
	go mod vendor

.PHONY: dockerise
dockerise: deps
	docker build -t "dynastiateam/backend:${IMAGE_TAG}" .

.PHONY: lint
lint:
	GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.17.1
	golangci-lint run

.PHONY: test
test:
	go test -v -cover ./... -count=1