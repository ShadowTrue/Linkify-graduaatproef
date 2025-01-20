# Use the Go image
FROM golang:1.23.4

# Set the working directory for the backend
WORKDIR /app/backend

# Copy go.mod and go.sum, then download dependencies
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy backend application code
COPY backend/ .

# Expose the backend API server port (default is 5000)
EXPOSE 5000

# Default command to run the Go backend
CMD ["go", "run", "main.go"]
