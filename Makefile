.PHONY: install install-node-modules build all lint run

TAG_NAME := $(shell git rev-parse --short HEAD)

install: install-node-modules
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.27.0

install-node-modules:
	cd frontend && npm i && cd ..

generate-static:
	rm -rf static/ && cd frontend && npm run build && cp -rp build/ ../static && cd ..

build: generate-static
	go build -race

lint:
	./bin/golangci-lint run

run: build
	ENV=prod ./club

run-debug: build
	ENV=dev ./club
