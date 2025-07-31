# Start from the official Go image to build the binary
FROM golang:1.24-alpine AS build

WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY ../../go.mod ./
RUN go mod download

# Copy the source code
COPY ../../ ./

# Build the binary (adjust main.go if needed)
RUN go build -o /date-diff-api ./cmd/http

# Use a minimal image to run the binary
FROM alpine:latest

# Add ca-certificates for HTTPS support if needed
RUN apk --no-cache add ca-certificates

# Copy the binary from the build stage
COPY --from=build /date-diff-api /date-diff-api

# Expose the port your app listens on (default 8080)
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/date-diff-api"]
