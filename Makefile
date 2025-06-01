# Go parameters
BINARY_NAME=pqc-poc
MAIN_FILE=main.go
GO=go

.PHONY: all build clean run test deps help

all: deps build

build:
	@echo "Building..."
	$(GO) build -o $(BINARY_NAME) $(MAIN_FILE)

run: build
	@echo "Running..."
	./$(BINARY_NAME)

test:
	@echo "Running tests..."
	$(GO) test -v ./...

clean:
	@echo "Cleaning..."
	$(GO) clean
	rm -f $(BINARY_NAME)

deps:
	@echo "Downloading dependencies..."
	$(GO) mod tidy

help:
	@echo "Available targets:"
	@echo "  all    - Download dependencies and build the application"
	@echo "  build  - Build the application"
	@echo "  run    - Build and run the application"
	@echo "  test   - Run tests"
	@echo "  clean  - Remove build artifacts"
	@echo "  deps   - Download dependencies"