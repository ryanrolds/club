.PHONY: install build all

TAG_NAME := $(shell git rev-parse --short HEAD)

install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.27.0

build:
	go build -race

lint:
	./bin/golangci-lint run
