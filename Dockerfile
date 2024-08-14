# Stage 1: Build the Go binary
FROM golang:1.22.6-bookworm AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the rest of the source code
COPY . ./

# Initialize the Go module
RUN go mod init ascii-art-web

# Tidy up dependencies
RUN go mod tidy

# Build the Go binary
RUN go build -o server server.go

# Stage 2: Create a lightweight image for running the binary
FROM debian:latest

# Set the working directory inside the container
WORKDIR /application

# Metadata for the final image
LABEL maintainer="Hezborn Shikuku, emmail:shikukuhezborn@gmail.com"
LABEL maintainer="Rabin Otieno, email:rabinyitzahck@gmail.com"
LABEL maintainer="Rodgers Ogwel, email:tidings@outlook.com"
LABEL description="ascii-art-web Go application"
LABEL version="1.0.0"
LABEL application="ascii-art-web"

# Copy the binary and necessary files from the builder stage
COPY --from=builder /app/. ./ 

# Clean up unnecessary files
RUN rm -f go.mod server.go

# Expose the application port
EXPOSE 8080

# Command to run the binary
CMD ["./server"]