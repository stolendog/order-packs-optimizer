.PHONY: help build run test docker-build docker-run docker-stop

# Default target
help:
	@echo "Available commands:"
	@echo "  run          - Run the application locally"
	@echo "  test         - Run all tests"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run application in Docker (attaches to container)"
	@echo "  deps         - Download dependencies"

# Run the application locally
run:
	@echo "Starting application on port 9999"
	go run ./cmd/api/main.go

# Run all tests
test:
	@echo "Running tests..."
	go test -v ./...

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t order-pack-optimizer .

# Run application in Docker
docker-run:
	@echo "Starting application in Docker on port 9999"
	docker run -it --rm -p 9999:9999 --name order-pack-optimizer-container order-pack-optimizer

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod tidy