# A development environment in an image to run tests on, build, etc.
FROM golang:1.14.13-stretch

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GO111MODULE on
ENV GOPATH /go
ENV GOLANGCI_LINT_VERSION 1.21.0
ENV KUBEBUILDER_VERSION 2.2.0
ENV KUBEBUILDER_DIR /usr/local/kubebuilder
ENV PATH ${KUBEBUILDER_DIR}/bin:${PATH}

WORKDIR /workspace

# Install golangci-lint
RUN curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v${GOLANGCI_LINT_VERSION}

# Install kubebuilder + dependencies
RUN echo "installing kubebuilder@$KUBEBUILDER_VERSION, kustomize@$KUSTOMIZE_VERSION" && \
    mkdir -p ${KUBEBUILDER_DIR} && \
    curl -sL https://go.kubebuilder.io/dl/${KUBEBUILDER_VERSION}/${GOOS}/${GOARCH} | tar -xz -C ${KUBEBUILDER_DIR} && \
    DL_NAME=kubebuilder_${KUBEBUILDER_VERSION}_${GOOS}_${GOARCH} && \
    mv $KUBEBUILDER_DIR/${DL_NAME}/bin $KUBEBUILDER_DIR/bin && \
    rm -rf ${KUBEBUILDER_DIR}/${DL_NAME}

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Download the necessessary controller-gen and kustomize
COPY Makefile Makefile
RUN make controller-gen kustomize

# Copy the dev / test files
COPY hack/ hack/
COPY Dockerfile Dockerfile
COPY .golangci.yml .golangci.yml
COPY PROJECT PROJECT
COPY config/ config/

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/
COPY pkg/ pkg/
COPY mocks/ mocks/
