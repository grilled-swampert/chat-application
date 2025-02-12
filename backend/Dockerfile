# Use the latest Golang Alpine base image
FROM golang:1.21-alpine

# Install required dependencies
RUN apk add --no-cache bash ca-certificates git build-base

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules manifests and download dependencies first for caching optimization
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go application
RUN go build -o chat-application

# List files (optional, for debugging)
RUN ls -l

# Command to run the application
CMD ["/app/chat-application"]
