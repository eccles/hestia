#!/usr/bin/env just --justfile
#
name := "hestia"

default:
	@just --list --unsorted --justfile {{justfile()}} | grep -v default

# Install grpc plugins and other go tools
tools:
	#!/usr/bin/env bash
	set -euxo pipefail
	source ./scripts/environment
	log_info "Install go tools"
	which go
	go version
	go get -tool google.golang.org/protobuf/cmd/protoc-gen-go
	go get -tool google.golang.org/grpc/cmd/protoc-gen-go-grpc	
	go install tool

# generate all code
generate:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/environment
	log_info "Generate code"
	which go
	go generate ./...

# QA all code
qa:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/environment
	log_info "Check go.mod and lint code"
	which go
	go mod tidy
	go mod verify
	log_info "Vetting"
	go vet ./...
	log_info "Formatting"
	golangci-lint fmt ./...
	log_info "Linting"
	golangci-lint run ./...
	log_info "Vulnerability checking"
	go run golang.org/x/vuln/cmd/govulncheck@latest --show verbose ./...

# check if there are ny uncommitted artifacts
check:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/environment
	log_info "Check for uncommitted artifacts"
	git diff --exit-code

# unittest all code
unittest:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/environment
	log_info "Run unittests"
	which go
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# build
build:
	#!/usr/bin/env bash
	set -euo pipefail
	source ./scripts/environment
	log_info "Build binariers"
	which go
	go build -o bin/ ./...
