#!/bin/bash
#
# Installs necessary tools
#
(. ./scripts/tools/source/go && go_install)
(. ./scripts/tools/source/golangci-lint && golangci-lint_install)
(. ./scripts/tools/source/protoc && protoc_install)
(. ./scripts/tools/source/mockery && mockery_install)
(. ./scripts/tools/source/protoc-gen-go && protoc-gen-go_install)
(. ./scripts/tools/source/protoc-gen-go-grpc && protoc-gen-go-grpc_install)
(. ./scripts/tools/source/protoc-gen-validate && protoc-gen-validate_install)
