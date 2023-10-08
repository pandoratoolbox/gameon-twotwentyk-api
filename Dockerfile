# FROM golang:bullseye AS build
FROM --platform=linux/arm64 golang:1.21.2-bullseye AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy the local package files to the container’s workspace.
COPY . .

RUN go mod tidy

# Build the Go app
RUN go build

# Run the binary program produced by `go install`
CMD ["gameon-twotwentyk-api"]