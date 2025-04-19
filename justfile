# !/usr/bin/env just --justfile
#
name := "hestie"

@default:
	@just --list --unsorted --justfile {{justfile()}} | grep -v default

# Install grpc plugins and other go tools
tools:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/environment
	log_info "Install go tools"
	go get -tool google.golang.org/protobuf/cmd/protoc-gen-go
	go get -tool google.golang.org/grpc/cmd/protoc-gen-go-grpc	
	go install tool

# generate all code
generate:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/environment
	log_info "Generate code"
	go generate ./...

# QA all code
qa:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/environment
	log_info "Check go.mod and lint code"
	go mod tidy
	go mod verify
	log_info "Format code"
	gofmt -l -s -w $(find . -type f -name '*.go'| grep -v "/vendor/\|/.git/")
	log_info "Linting"
	golangci-lint run -v
	log_info "Vulnerability checking"
	go run golang.org/x/vuln/cmd/govulncheck@latest --show verbose ./...

# unittest all code
unittests:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/environment
	log_info "Run unittests"
	go test -v ./...

# build
build:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/source/environment
	log_info "Build binariers"
	go build -o bin/ ./...
