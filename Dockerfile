# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Copy src folder to into go/src folder of image
ADD src /go/src 

# Set working dir in image
WORKDIR /go/src

# Build the go app
RUN go build -o sabr0 .

# Command to run the executable
CMD ["./sabr0"]