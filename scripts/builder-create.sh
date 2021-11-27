#!/bin/bash
#
# Makes builder image.
#
docker build \
	-t hestia:latest \
	-f Dockerfile  \
	.
