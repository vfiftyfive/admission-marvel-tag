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
RUN CGO_ENABLED=0 go build -ldflags "-linkmode external -extldflags -static" -o marvel-webhook .

# Use Distroless as the base image for the runtime
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go binary from the build stage
COPY --from=build /app/marvel-webhook .

# Run the Go binary
ENTRYPOINT ["./marvel-webhook"]