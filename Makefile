# Variables
BINARY_NAME=ulsidor
BUILD_DIR=build

# Default target
all: build

# Build the project
build:
	@echo "Building..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/ulsidor_cli

# Run the project
run: build
	@echo "Running..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	go clean
	rm -rf $(BUILD_DIR)

# Test the project
test:
	@echo "Testing..."
	go test -v ./...

# Format the code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Check for linting issues
lint:
	@echo "Linting code..."
	golint ./...

# Install necessary tools (like golint)
install-tools:
	go get -u golang.org/x/lint/golint

# Full build: format, lint, test, and then build
full-build: fmt lint test build

.PHONY: all build run clean test fmt lint install-tools full-build
