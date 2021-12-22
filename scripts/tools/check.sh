#!/bin/bash
#
# Checks that all tools are available
. scripts/source/log
. scripts/source/environment
. scripts/tools/source/golang
. scripts/tools/source/golangci-lint

golang_check
golang_lint_check
