#!/bin/bash
#
# Builds all golang binaries
. scripts/source/log
. scripts/source/environment

go install -v ./...
