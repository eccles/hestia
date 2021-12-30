#!/bin/bash
#
# Installs necessary tools
#
. ./scripts/source/log
. ./scripts/source/environment
. ./scripts/tools/source/golang
. ./scripts/tools/source/golangci-lint
. ./scripts/tools/source/protoc
. ./scripts/tools/source/protoc-gen-go
. ./scripts/tools/source/protoc-gen-go-grpc

golang_install
golang_lint_install
protoc_install
protoc_gen_go_install
protoc_gen_go_grpc_install
