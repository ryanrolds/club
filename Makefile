.PHONY: all build ci lint run

all:
	make -C golang all
	make -C frontend all

install:
	make -C golang install
	make -C frontend install

build:
	make -C golang build
	make -C frontend build

ci:
	make -C golang ci
	make -C frontend ci

lint:
	make -C golang lint
	make -C frontend lint

test:
	make -C golang test
	make -C frontend test

run:
	docker-compose up
