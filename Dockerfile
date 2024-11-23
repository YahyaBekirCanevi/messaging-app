# Start with the official Go image for development
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o messaging-app

# Use a minimal base image for production
FROM debian:bullseye-slim

# Set the working directory for the final image
WORKDIR /app

# Copy the built application binary from the builder stage
COPY --from=builder /app/messaging-app .

# Expose the application's port
EXPOSE 8080

# Run the application
CMD ["./messaging-app"]
