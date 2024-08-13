# Start from the official Golang image for building the binary
FROM golang:1.22.1 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the rest of the source code
COPY . .

# Build the Go application binary
RUN go build -o /app/main server.go

# Create a smaller image for the final executable
FROM debian:bullseye-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .
COPY /app/handlers .
COPY /app/ascii-art .
COPY /app/static .
COPY /app/templates .

# Expose the application port
EXPOSE 8080

# Command to run the binary
CMD ["./main"]
