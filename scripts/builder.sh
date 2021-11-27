#!/bin/sh
#
# Executes a command inside the builder container
#
docker run \
    --rm -it \
    -v $(pwd):/home/hestia \
    -w /home/hestia \
    -u $(id -u):$(id -g) \
    -e HOME=/home/hestia \
    -e USER=hestia \
    -e GOPATH=/home/hestia \
    hestia:latest\
    "$@"
