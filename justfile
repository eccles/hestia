#!/usr/bin/env just --justfile
#
name := "hestie"
#
# show current available options
default:
	@just --list

# generate all code
generate:
	go generate ./...

# QA all code
qa:
	#!/usr/bin/env bash
	go mod tidy
	go mod verify
	go mod vendor
	gofmt -l -s -w $(find . -type f -name '*.go'| grep -v "/vendor/\|/.git/")
	golangci-lint run --no-config

# unittest all code
unittests:
	go test -v ./...

# build
build:
	go build -v ./...
