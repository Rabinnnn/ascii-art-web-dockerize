#!/bin/bash

# Define variables
IMAGE_NAME="my-go-app"
TAG="latest"
CONTAINER_NAME="my-go-app-container"
PORT=8080

# Step 1: Build the Docker image
echo "Building the Docker image..."
docker build -t ${IMAGE_NAME}:${TAG} .

# Check if the build was successful
if [ $? -ne 0 ]; then
    echo "Failed to build the Docker image."
    exit 1
fi

echo "Docker image built successfully."

# Step 2: Run the Docker container
echo "Running the Docker container..."

# Check if a container with the same name already exists
if [ "$(docker ps -aq -f name=${CONTAINER_NAME})" ]; then
    # Remove the existing container
    docker rm -f ${CONTAINER_NAME}
fi

docker run -d -p ${PORT}:${PORT} --name ${CONTAINER_NAME} ${IMAGE_NAME}:${TAG}

# Check if the container started successfully
if [ $? -ne 0 ]; then
    echo "Failed to start the Docker container."
    exit 1
fi

echo "Docker container started successfully."
echo "Application is running at http://localhost:${PORT}"
