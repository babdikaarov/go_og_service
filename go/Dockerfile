# Stage 1: Build the Go application
FROM golang AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

EXPOSE 3030
# set hostname to localhost
ENV HOSTNAME="0.0.0.0"
# Build the Go application
RUN go build -o main .

# Run the Go application
CMD ["./main"]