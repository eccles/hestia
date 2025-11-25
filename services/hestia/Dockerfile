# syntax=docker/dockerfile:1.3

FROM --platform=linux/amd64 gcr.io/distroless/static

WORKDIR /service

COPY bin/whoamiv1 .

COPY service/whoamiv1/files files

ENV PATH=$PATH:/service

ENTRYPOINT ["/service/whoamiv1"]
