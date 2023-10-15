# Use the official Golang image to build the Go application
FROM golang:latest AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the Go application source code
COPY cmd/marvel-webhook/ .

# Build the Go application
RUN go build -o marvel-webhook .

# Run the Go binary
CMD ["./marvel-webhook"]
