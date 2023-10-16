# Earthfile
VERSION 0.7

#BASE
FROM golang:latest 
# Set the working directory inside the container.
WORKDIR /app
# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download

build:
    # Build the Go app
    COPY cmd/marvel-webhook .
    RUN go build -o marvel-webhook .
    SAVE ARTIFACT marvel-webhook

test:
    # Test the Go app
    COPY cmd/marvel-webhook .
    ARG MARVEL_PRIVATE_KEY
    RUN go test -v .
    
    
docker:
    FROM golang:latest
    # Copy the compiled Go binary into the final stage container.
    COPY +build/marvel-webhook .
    # Set the environment variable for the Marvel private key.
    ARG MARVEL_PRIVATE_KEY
    # Run the Go binary.
    CMD ["./marvel-webhook"]
    SAVE IMAGE --push vfiftyfive/marvel-webhook:latest

all:
    BUILD +build
    BUILD +test
    BUILD +multi

multi: 
    BUILD --platform=linux/amd64 --platform=linux/arm64 +docker