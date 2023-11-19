# Use the official Golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Golang application
RUN go build -o main .

# Run the application when the container starts
CMD ["./main"]
