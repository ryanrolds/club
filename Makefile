.PHONY: all build ci lint run

all:
	make -C golang all

install:
	make -C golang install

build:
	make -C golang build

ci:
	make -C golang ci

lint:
	make -C golang lint

test:
	make -C golang test

run:
	docker-compose up
