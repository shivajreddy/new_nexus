# Use official Golang image
FROM golang:1.22

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the code
COPY . .

# Build the Go app
RUN go build -o server

# Command to run the server
CMD ["./server"]
