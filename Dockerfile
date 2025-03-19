# Use Golang image
FROM golang:1.21-alpine

# Set environment
WORKDIR /app

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the app
RUN go build -o golang-log-service

# Run the application
CMD ["./golang-log-service"]
