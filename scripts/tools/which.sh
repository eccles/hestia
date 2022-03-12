#!/bin/bash
#
# Finds installed tools
#
(. ./scripts/tools/source/go && go_which)
(. ./scripts/tools/source/golangci-lint && golangci-lint_which)
(. ./scripts/tools/source/protoc && protoc_which)
(. ./scripts/tools/source/mockery && mockery_which)
(. ./scripts/tools/source/protoc-gen-go && protoc-gen-go_which)
(. ./scripts/tools/source/protoc-gen-go-grpc && protoc-gen-go-grpc_which)
(. ./scripts/tools/source/protoc-gen-validate && protoc-gen-validate_which)
