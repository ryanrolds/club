.PHONY: install build all lint run

TAG_NAME := $(shell git rev-parse --short HEAD)

install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.27.0

build:
	go build -race

lint:
	./bin/golangci-lint run

run: build
	ENV=prod ./club

run-debug: build
	ENV=dev ./club
