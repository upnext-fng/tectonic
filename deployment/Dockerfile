# Production-Ready Go 1.24 Dockerfile
# Multi-stage build for optimal image size and security

# Build stage
FROM golang:1.24-alpine AS builder

# Install essential build dependencies
RUN apk add --no-cache \
    git \
    ca-certificates \
    tzdata \
    upx && \
    update-ca-certificates

# Create appuser for security (will be used in final stage)
RUN adduser -D -g '' appuser

# Set working directory
WORKDIR /build

# Copy go mod files first for better caching
COPY go.mod go.sum ./

# Download dependencies (cached if go.mod/go.sum unchanged)
RUN go mod download && \
    go mod verify

# Copy source code
COPY . .

# Build the application with optimized flags
# CGO_ENABLED=0 for static binary
# GOOS=linux for Linux target
# -a flag forces rebuilding of packages
# -installsuffix cgo for static linking
# -ldflags for smaller binary size and version info
ARG VERSION=dev
ARG BUILD_TIME
ARG GIT_COMMIT
RUN CGO_ENABLED=0 GOOS=linux go build \
    -a -installsuffix cgo \
    -ldflags="-extldflags '-static' -s -w -X main.version=${VERSION} -X main.buildTime=${BUILD_TIME} -X main.gitCommit=${GIT_COMMIT}" \
    -o app ./main.go

# Compress binary with UPX for even smaller size
RUN upx --best --lzma app

# Final stage - distroless image for better security
FROM gcr.io/distroless/static:nonroot

# Copy timezone data from builder
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy CA certificates from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary from builder stage
COPY --from=builder /build/app /app

# Set environment variables
ENV TZ=UTC

# Use non-root user (distroless provides nonroot user)
USER nonroot:nonroot

# Expose port (adjust based on your server config)
EXPOSE 3000

# Use ENTRYPOINT and CMD properly
ENTRYPOINT ["/app"]
CMD ["serve"]