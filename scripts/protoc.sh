#!/bin/bash
#
# Uses protoc to generate pb.go files
# This is executed in a subdirectory so we must cd into the root of the
# repo
#
# #1 - must be the relative directory path - protoc is unable to 
#      figure this out from the $(pwd) command.
#
cd $( dirname $( dirname $0))
. scripts/source/environment

log_info "Generate protobuf for ${1}"
protoc -I . \
    --go_out=paths=source_relative:. \
    --go-grpc_out=paths=source_relative:. \
    ${1}/*.proto
