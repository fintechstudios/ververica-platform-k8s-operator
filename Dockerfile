ARG BUILD_IMG
# Build the manager binary from the builder docker image
FROM ${BUILD_IMG} as builder
ARG GIT_COMMIT=""
ARG VERSION="latest"

# Build and force refetching, override
RUN make ARGS="-a" VERSION=${VERSION} GIT_COMMIT=${GIT_COMMIT}

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:latest
WORKDIR /
COPY --from=builder /workspace/bin/manager .
ENTRYPOINT ["/manager"]
