#!/bin/bash
#
# Installs necessary tools
#
. ./scripts/source/log
. ./scripts/source/environment
. ./scripts/tools/source/golang
. ./scripts/tools/source/golangci-lint

golang_install
golang_lint_install
