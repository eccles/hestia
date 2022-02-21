#!/bin/bash
#
# Uses protoc to generate pb.go files
# This is executed in a subdirectory so we must cd into the root of the
# repo
cd $( dirname $( dirname $0))
. scripts/source/log
. scripts/source/environment

log_info "Generate $1"
protoc -I . -I ${PROTOC_INCLUDE} \
    --go_out=paths=source_relative:. \
    --go-grpc_out=paths=source_relative:. \
    --validate_out=lang=go,paths=source_relative:. \
    $1
