all: test

deps:
	go mod download
.PHONY: deps

test: unit-test

unit-test: deps
	go test -v ./...
