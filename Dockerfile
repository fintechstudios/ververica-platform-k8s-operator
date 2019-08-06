ARG BUILD_IMG

# Build the manager binary from the builder docker image
FROM ${BUILD_IMG} as builder

# Build
RUN make ARGS="-a"

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:latest
WORKDIR /
COPY --from=builder /workspace/bin/manager .
ENTRYPOINT ["/manager"]
