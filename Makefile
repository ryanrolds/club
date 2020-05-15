.PHONY: install build all lint run

TAG_NAME := $(shell git rev-parse --short HEAD)

all: build test run

install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.27.0

build:
	go build -race

fakes:
	go generate ./...

test:
	go test ./...

coverage:
	go test -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html
	open cover.html

lint:
	./bin/golangci-lint run

run: build
	ENV=prod ./club

run-debug: build
	ENV=dev ./club
