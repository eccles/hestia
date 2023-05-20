#!/bin/bash
#
# Finds installed tools
#

for f in go golangci-lint protoc mockery protoc-gen-go protoc-gen-go-grpc protoc-gen-validate nvm
do
        (. ./scripts/tools/source/${f} && ${f}_which)
done
