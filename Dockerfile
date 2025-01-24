# Use the official Go image to build the application
FROM golang:1.23.2 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules and dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/server/main.go

# Use a minimal image for running the app
FROM alpine:3.18

# Set up a working directory
WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/main .

# Expose the port your app runs on
EXPOSE 4001

# Command to run the executable
CMD ["./main"]
