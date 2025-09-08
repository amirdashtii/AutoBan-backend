## Backend Dockerfile (Go)
# Multi-stage build: build static binary, then run on lightweight image

# 1) Builder stage
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build static binary from cmd/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o autoban ./cmd/app


# 2) Runtime stage
FROM alpine:3.20
WORKDIR /srv/app

# Install CA certificates for outbound HTTPS calls
RUN apk add --no-cache ca-certificates tzdata && \
    update-ca-certificates

# Copy binary and (optionally) config directory
COPY --from=builder /app/autoban ./autoban
COPY --from=builder /app/config ./config

# Environment
ENV GIN_MODE=release

# Expose API port
EXPOSE 8080

# Run
CMD ["./autoban"]


