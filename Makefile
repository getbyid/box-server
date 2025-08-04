BIN_NAME=box-server
SRC_DIR=.
BUILD_DIR=./bin

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	@echo "Building the application"
	@mkdir -p "$(BUILD_DIR)"
	@go build -ldflags "-s -w" -o "$(BUILD_DIR)/$(BIN_NAME)" "$(SRC_DIR)"
	@env GOOS=android GOARCH=arm64 go build -ldflags "-s -w" -o "$(BUILD_DIR)/$(BIN_NAME)-android" "$(SRC_DIR)"

# Install into local user bin directory
.PHONY: install
install:
	@echo "Copying to .local/bin"
	@cp "$(BUILD_DIR)/$(BIN_NAME)" ~/.local/bin

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning up"
	@rm -rf "$(BUILD_DIR)"

# Help command
.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  make build   - Build the application"
	@echo "  make install - Install into local user bin directory"
	@echo "  make clean   - Clean build artifacts"
	@echo "  make help    - Show this help message"
