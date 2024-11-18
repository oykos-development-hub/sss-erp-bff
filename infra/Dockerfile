# Stage 1: Build the Go binary
FROM golang:1.22 as builder

# Set environment variables
ENV CGO_ENABLED=0

# Set the working directory inside the builder container
WORKDIR /app

# Copy all files from the current directory to the working directory inside the container
COPY . .

# Build the Go application and output the binary to /app/cmd/bff
RUN go build -o cmd/bff .

# Stage 2: Create a minimal image with the built binary
FROM scratch

# Copy the Go binary from the builder stage
COPY --from=builder /app/cmd/bff /bff

# Copy SSL certificates to enable HTTPS support
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Set the working directory
WORKDIR /

# Command to run the Go binary
CMD ["/bff"]
