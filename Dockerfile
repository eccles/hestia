
FROM golang:1.17.3-buster

ENV DEBIAN_FRONTEND noninteractive

RUN env | sort && echo && go env && echo

WORKDIR ${GOPATH}

RUN apt-get update \
  && apt-get upgrade -y --no-install-recommends \
  && apt-get install -y \
        autoconf \
        automake \
        pkg-config \
        build-essential \
        libffi-dev \
        libtool \
        gettext-base \
        jq \
        libssl-dev \
        python3 \
        python3-dev \
        python3-pip \
        python3-sha3 \
        unzip \
        vim-nox \
        yamllint \
  && apt-get autoremove \
  && apt-get autoclean \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

ENV PROTOC_VERSION=3.13.0
RUN curl -fsSOL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
  && unzip -q -o protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local bin/protoc \
  && chmod +x /usr/local/bin/protoc \
  && unzip -q -o protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local include/* \
  && chmod -R +rx /usr/local/include \
  && rm -rf protoc-${PROTOC_VERSION}-linux-x86_64.zip

# common google api proto files.
RUN curl -fsSOL https://github.com/googleapis/api-common-protos/archive/master.zip \
  && unzip -q -o master.zip -d /usr/local api-common-protos-master/google/* \
  && curl -fsSOL https://raw.githubusercontent.com/grpc/grpc-proto/master/grpc/health/v1/health.proto \
  && mkdir -p /usr/local/api-common-protos-master/grpc/health/v1 \
  && mv health.proto /usr/local/api-common-protos-master/grpc/health/v1/ \
  && ls -l /usr/local \
  && rm -rf master.zip

# Additional proto utilities
ENV PROTO_GEN_DOC_VERSION=v1.4.1
ENV PROTO_GEN_VALIDATE_VERSION=v0.6.1
RUN go install github.com/envoyproxy/protoc-gen-validate@${PROTO_GEN_VALIDATE_VERSION} \
  && go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@${PROTO_GEN_DOC_VERSION}

# Standard tools
ENV GRPC_HEALTH_PROBE_VERSION=v0.3.5
RUN curl -fSsOL https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 \
  && mv grpc_health_probe-linux-amd64 /usr/local/bin/grpc_health_probe \
  && chmod a+rx /usr/local/bin/grpc_health_probe

# go linter
RUN curl -fsSOL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh \
  && sh ./golangci-lint.sh v1.36.0 \
  && rm -rf golangci-lint.sh

ENV GOLICENSE_VERSION=0.2.0
RUN curl -fSsOL https://github.com/mitchellh/golicense/releases/download/v${GOLICENSE_VERSION}/golicense_${GOLICENSE_VERSION}_linux_x86_64.tar.gz \
  && tar xvzf golicense_${GOLICENSE_VERSION}_linux_x86_64.tar.gz -C /usr/local/bin \
  && rm golicense_${GOLICENSE_VERSION}_linux_x86_64.tar.gz

RUN go install github.com/jstemmer/go-junit-report@v0.9.1 \
  && go install golang.org/x/tools/cmd/goimports@v0.1.7 \
  && go install github.com/go-delve/delve/cmd/dlv@v1.7.2 \
  && go install github.com/psampaz/go-mod-outdated@v0.7 \
  && go install rsc.io/goversion@v1.2

ENV GO_TEST_SUM_VERSION=0.6.0
ENV GO_TEST_TOOLS_VERSION=3.0.3
RUN go install github.com/axw/gocov/gocov@v1.0.0 \
  && go install github.com/AlekSi/gocov-xml@latest \
  && go install github.com/matm/gocov-html@latest \
  && go install gotest.tools/v3/assert/cmd/gty-migrate-from-testify@v${GO_TEST_TOOLS_VERSION} \
  && curl -fSsOL https://github.com/gotestyourself/gotestsum/releases/download/v${GO_TEST_SUM_VERSION}/gotestsum_${GO_TEST_SUM_VERSION}_linux_amd64.tar.gz \
  && tar xvzf gotestsum_${GO_TEST_SUM_VERSION}_linux_amd64.tar.gz -C /usr/local/bin \
  && rm gotestsum_${GO_TEST_SUM_VERSION}_linux_amd64.tar.gz
