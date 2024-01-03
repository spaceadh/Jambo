# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod .
COPY go.sum .

# Download dependencies and update go.mod
RUN go mod tidy
RUN go mod download

# Copy the local source files to the container's working directory
COPY . .

RUN go mod tidy
# Build the Go application
RUN go build -o jambo main.go

# Expose the port on which the application will run
EXPOSE 8080

# Set the entry point for the container
ENTRYPOINT ["./jambo"]

# You can also specify default arguments to your application
# CMD ["arg1", "arg2"]