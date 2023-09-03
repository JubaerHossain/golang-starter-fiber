# Use a lightweight base image
FROM golang:1.21 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy only the Go module files and download dependencies (if go.mod or go.sum changes)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your Go application code into the container
COPY . .

# Build your Go application with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Create a minimal image
FROM alpine:latest

# Install CA certificates for SSL/TLS compatibility
RUN apk --no-cache add ca-certificates

# Set the working directory inside the final container
WORKDIR /app

# Copy the built Go binary from the previous stage
COPY --from=build /app/app .

# Expose the port your application will listen on (adjust if needed)
EXPOSE 8080

# Optionally, create a non-root user for running the application
# USER myuser

# Run your Go application (replace with your binary name if different)
CMD ["./app"]
