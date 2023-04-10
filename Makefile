.PHONY: run

build:
	go build -o ./.bin/sbro cmd/sbro/main.go

run: build
	./.bin/sbro

.DEFAULT_GOAL := run