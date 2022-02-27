
# environment vars
PWD := $(shell pwd )
export GOROOT := ${HOME}/.local/go
export GOBIN := $(PWD)/bin

# show available options
.PHONY: help
help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: tools-install
tools-install: ## Ensure that all tools are available
	@./scripts/tools/install.sh

# actual code handling
.PHONY: generate
generate: ## generate code such as proto stuff
	@go generate ./...

.PHONY: qa
qa: ## quality check all source code
	@go mod vendor
	@gofmt -l -s -w $(shell find . -type f -name '*.go'| grep -v "/vendor/\|/.git/")
	@golangci-lint run
	@go mod tidy
	@go mod verify

.PHONY: unittests
unittests: ## unittest all source code
	@go test -v ./...

.PHONY: build
build: ## Compile all source code
	@go install -v ./...
