.PHONY: build test clean run-examples run-basic run-advanced run-cli help

# Build the CLI tool
build:
	@echo "Building citytimezones CLI..."
	@go build -o bin/citytimezones ./cmd/citytimezones

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -cover ./...

# Run comprehensive test suite
test-comprehensive:
	@echo "Running comprehensive test suite..."
	@./scripts/test_runner.sh

# Run unit tests only
test-unit:
	@echo "Running unit tests..."
	@go test -v ./internal/city

# Run integration tests only
test-integration:
	@echo "Running integration tests..."
	@go test -v ./internal/city -run TestDataLoadingIntegration
	@go test -v ./internal/city -run TestSearchIntegration
	@go test -v ./internal/city -run TestCachingIntegration
	@go test -v ./internal/city -run TestConcurrencyIntegration

# Run public API tests
test-api:
	@echo "Running public API tests..."
	@go test -v ./pkg/citytimezones

# Run security tests
test-security:
	@echo "Running security tests..."
	@go test -v ./internal/city -run TestValidation
	@go test -v ./internal/city -run TestErrorHandlingIntegration

# Run performance tests
test-performance:
	@echo "Running performance tests..."
	@go test -bench=. -benchmem ./internal/city
	@go test -bench=. -benchmem -cpu=1,2,4,8 ./internal/city

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	@go test -bench=. ./...

# Run benchmarks with memory profiling
bench-mem:
	@echo "Running benchmarks with memory profiling..."
	@go test -bench=. -benchmem ./internal/city

# Run race condition tests
test-race:
	@echo "Running race condition tests..."
	@go test -race ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/

# Run basic example
run-basic:
	@echo "Running basic example..."
	@go run ./examples/basic

# Run advanced example
run-advanced:
	@echo "Running advanced example..."
	@go run ./examples/advanced

# Run examples
run-examples: run-basic run-advanced

# Run CLI with help
run-cli: build
	@echo "Running CLI with help..."
	@./bin/citytimezones -help

# Run CLI examples
run-cli-examples: build
	@echo "Running CLI examples..."
	@echo "Searching for Chicago:"
	@./bin/citytimezones -city Chicago
	@echo "\nSearching for German cities:"
	@./bin/citytimezones -iso DE -limit 3
	@echo "\nSearching for cities in New York timezone:"
	@./bin/citytimezones -timezone "America/New_York" -limit 3

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go mod download

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	@go vet ./...

# Run all checks
check: fmt lint test

# Prepare for release
release-prep:
	@echo "Preparing for release..."
	@./scripts/release.sh

# Show help
help:
	@echo "Available targets:"
	@echo "  build          - Build the CLI tool"
	@echo "  test           - Run basic tests"
	@echo "  test-coverage  - Run tests with coverage"
	@echo "  test-comprehensive - Run comprehensive test suite"
	@echo "  test-unit      - Run unit tests only"
	@echo "  test-integration - Run integration tests only"
	@echo "  test-api       - Run public API tests"
	@echo "  test-security  - Run security tests"
	@echo "  test-performance - Run performance tests"
	@echo "  test-race      - Run race condition tests"
	@echo "  bench          - Run benchmarks"
	@echo "  bench-mem      - Run benchmarks with memory profiling"
	@echo "  clean          - Clean build artifacts"
	@echo "  run-basic      - Run basic example"
	@echo "  run-advanced   - Run advanced example"
	@echo "  run-examples   - Run all examples"
	@echo "  run-cli        - Run CLI with help"
	@echo "  run-cli-examples - Run CLI examples"
	@echo "  deps           - Install dependencies"
	@echo "  fmt            - Format code"
	@echo "  lint           - Lint code"
	@echo "  check          - Run all checks (fmt, lint, test)"
	@echo "  release-prep   - Prepare for release (run all checks)"
	@echo "  help           - Show this help"
