#!/bin/bash

# Set variables
IMAGE_NAME="ascii-art-web"
IMAGE_VERSION="1.0.0"
CONTAINER_NAME="ascii-art-web-container"
PORT=8080

# Step 1: Build the Docker image
echo "Building the Docker image..."
docker build -t ${IMAGE_NAME}:${IMAGE_VERSION} .

# Check if the image was built successfully
if [ $? -ne 0 ]; then
    echo "Failed to build the Docker image."
    exit 1
fi

# Step 2: Remove any existing container with the same name
echo "Removing any existing container with the name ${CONTAINER_NAME}..."
docker rm -f ${CONTAINER_NAME} 2>/dev/null

# Step 3: Run the Docker container
echo "Running the Docker container..."
docker run -d --name ${CONTAINER_NAME} -p ${PORT}:8080 ${IMAGE_NAME}:${IMAGE_VERSION}

# Check if the container is running
if [ $? -ne 0 ]; then
    echo "Failed to start the Docker container."
    exit 1
fi

echo "Docker container started successfully."
echo "Application is running at http://localhost:${PORT}"
