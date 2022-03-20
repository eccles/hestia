#!/bin/bash
#
# Generates a version file for embedding in application
#
# #1 - must be the relative directory path
#
cd $( dirname $( dirname $0))
. scripts/source/log

version=$(git describe --tags --abbrev=0)
log_info "Generate ${version} in version.txt for ${1}"
echo ${version} > ${1}/version.txt
