# Makefile for flate

BINARY_NAME=flate
GO_VERSION=1.21
LDFLAGS=-ldflags="-s -w"

# Default target
.PHONY: all
all: build

# Build the binary
.PHONY: build
build:
	go build $(LDFLAGS) -o $(BINARY_NAME) .

# Run tests
.PHONY: test
test:
	go test -v ./...

# Clean build artifacts
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)_test
	rm -f flate  # Remove old binary name
	rm -f flate_test  # Remove old test binary name
	rm -rf dist/

# Build for all platforms
.PHONY: build-all
build-all: clean
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-arm64 .
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe .

# Install the binary to GOPATH/bin
.PHONY: install
install:
	go install $(LDFLAGS) .

# Run linter
.PHONY: lint
lint:
	go vet ./...
	go fmt ./...

# Run the binary with test data
.PHONY: demo
demo: build
	@echo "Testing compression (default raw deflate):"
	@echo "hello world!" | ./$(BINARY_NAME) | ./$(BINARY_NAME) -d
	@echo ""
	@echo "Testing zlib compression:"
	@echo "hello world!" | ./$(BINARY_NAME) -z | ./$(BINARY_NAME) -d -z

# Show help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build      - Build the binary"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean build artifacts"
	@echo "  build-all  - Build for all platforms"
	@echo "  install    - Install to GOPATH/bin"
	@echo "  lint       - Run linter and formatter"
	@echo "  demo       - Run demo with test data"
	@echo "  help       - Show this help message"
