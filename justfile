#!/usr/bin/env just --justfile
#
name := "hestie"
all := "go golangci-lint protoc mockery protoc-gen-go protoc-gen-go-grpc protoc-gen-validate nvm"
#
# show current available options
default:
	@just --list

# Install or which current tools e.g. just tools install|which
tools SUB:
	#!/usr/bin/env bash
	for f in $(ls scripts/tools/source | sort | cut -c3-)
	do
		(. ./scripts/tools/source/*${f} && ${f}_{{SUB}})
	done

# Install one or more tools
install +TOOLS:
	#!/usr/bin/env bash
	for f in {{TOOLS}}
	do
		(. ./scripts/tools/source/*${f} && ${f}_install)
	done

# Find one or more tools
which +TOOLS:
	#!/usr/bin/env bash
	for f in {{TOOLS}}
	do
		(. ./scripts/tools/source/*${f} && ${f}_which)
	done

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
	golangci-lint run

# unittest all code
unittests:
	go test -v ./...

# build
build:
	go install -v ./...
