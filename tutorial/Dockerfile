# Stage 1: Build
FROM golang:1.25-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o tutorial main.go

# Stage 2: Runtime
FROM alpine:latest
RUN apk add --no-cache ca-certificates curl

WORKDIR /
# ENTRYPOINT ["/bin/sh", "-c", "pwd; ls -la; ls -la ./resources; exec tail -f /dev/null"]
COPY --from=builder /app/tutorial .
COPY config ./config
RUN pwd

ENV ENV=dev
EXPOSE 8080

# HEALTHCHECK --interval=30s --timeout=5s --retries=3 \
#   CMD curl --fail http://localhost:8080/health || exit 1

CMD ["./tutorial"]






# # Stage 1: Build
# FROM golang:1.21-alpine AS builder

# WORKDIR /app

# # Install git for dependencies
# RUN apk add --no-cache git

# # Copy go.mod and go.sum first (for caching dependencies)
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy source
# COPY . .

# # Build the binary
# RUN CGO_ENABLED=0 GOOS=linux go build -o tutorial main.go

# # Stage 2: Minimal image
# FROM alpine:latest
# RUN apk add --no-cache ca-certificates

# WORKDIR /root/
# COPY --from=builder /app/tutorial .

# # Environment variable for config
# ENV ENV=dev

# EXPOSE 8080

# CMD ["./tutorial"]