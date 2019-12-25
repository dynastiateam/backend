SHELL=/bin/sh
IMAGE_TAG := $(shell git rev-parse HEAD)

export GO111MODULE=on

.PHONY: ci
ci: deps

.PHONY: deps
deps:
	rm -rf vendor
	go mod download
	go mod vendor
	go mod tidy
