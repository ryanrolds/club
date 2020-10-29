.PHONY: install build all lint run

all: build coverage

install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.27.0
	cd static/ && yarn install

build:
	go build ./cmd/backend

fakes:
	go generate ./...

test:
	go test ./...

coverage:
	go test -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html
	open cover.html

coverage-ci:
	go test -v -coverprofile cover.out ./...
	go tool cover -func=cover.out

lint:
	./bin/golangci-lint run
	cd static/ && npm run lint

ci: install lint build test coverage-ci

run: build
	ENV=prod ./backend

run-debug: build
	ENV=local ./backend

run-pretty: build
	# open http://localhost:3001
	ENV=local ./backend 2>&1 | jq --raw-output '[("# " + .level + (" " * 10))[0:10], (.func + (" " * 20))[0:20], .msg] | @tsv'


