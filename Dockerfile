#Build Stage
FROM golang:1.23-alpine3.20 AS builder

# Set working directory inside the container
WORKDIR /build

# Copy go.mod and go.sum to leverage Docker cache for dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go application
RUN go build -o /app ./cmd/main.go

# Final Stage
FROM alpine:3.20

# Copy the built binary from the builder stage
COPY --from=builder /app /app

# Create .env file placeholder
RUN touch .env

# Command to run the Go application
CMD ["/app"]