# Use an official Go runtime as the base image
FROM golang:1.17 AS builder

# Set the working directory
WORKDIR /app

# Copy the entire source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o attendance-app main.go

# Use a lightweight Alpine Linux image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder image
COPY --from=builder /app/attendance-app .

# Expose the port your application listens on
EXPOSE 3000

# Command to run the application
CMD ["./attendance-app"]
