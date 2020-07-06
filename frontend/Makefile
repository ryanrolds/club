.PHONY: install build lint lint-fix format test ci run run-debug all

install:
	yarn install
	yarn global add react-scripts

build:
	npm run build

lint:
	npm run lint

lint-fix:
	npm run lint-fix

format:
	npm run format

test:
	npm run test

ci: install lint build

run:
	npm run start

run-debug:
	npm run dev
