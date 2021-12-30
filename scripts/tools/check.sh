#!/bin/bash
#
# Checks that all tools are available
. scripts/source/log
. scripts/source/environment
. scripts/tools/source/golang
. scripts/tools/source/golangci-lint
. scripts/tools/source/protoc
. scripts/tools/source/protoc-gen-go
. scripts/tools/source/protoc-gen-go-grpc

golang_check
golang_lint_check
protoc_check
protoc_gen_go_check
protoc_gen_go_grpc_check
