# Use the latest Golang image as the base image
FROM golang:1.18-bullseye

# Set the working directory to /app
WORKDIR /app

# Copy the entire project to the container
COPY . .

# Build the Go project
RUN go build .

## Expose port 8080 for the app to listen on
#EXPOSE 8080

# Run the compiled binary when the container starts
CMD ["./main"]
