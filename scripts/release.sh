#!/bin/bash

# Release preparation script for city-timezones-go
# This script helps prepare the project for publication on pkg.go.dev

set -e

echo "🚀 Preparing city-timezones-go for release..."

# Check if we're in a git repository
if [ ! -d ".git" ]; then
    echo "❌ Not in a git repository. Please initialize git first."
    exit 1
fi

# Check if go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or later."
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | cut -d' ' -f3 | sed 's/go//')
REQUIRED_VERSION="1.21"
if [ "$(printf '%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V | head -n1)" != "$REQUIRED_VERSION" ]; then
    echo "❌ Go version $GO_VERSION is not supported. Please install Go $REQUIRED_VERSION or later."
    exit 1
fi

echo "✅ Go version $GO_VERSION is supported"

# Clean and tidy dependencies
echo "🧹 Cleaning and tidying dependencies..."
go mod tidy
go mod download

# Run tests
echo "🧪 Running tests..."
go test -v ./...

# Run benchmarks
echo "🏃 Running benchmarks..."
go test -bench=. ./...

# Check code formatting
echo "🎨 Checking code formatting..."
if ! go fmt ./...; then
    echo "❌ Code formatting issues found. Please run 'go fmt ./...' and fix issues."
    exit 1
fi

# Check for linting issues
echo "🔍 Running linter..."
if ! go vet ./...; then
    echo "❌ Linting issues found. Please fix them before release."
    exit 1
fi

# Build examples
echo "🔨 Building examples..."
go build -o /tmp/citytimezones-basic ./examples/basic
go build -o /tmp/citytimezones-advanced ./examples/advanced
go build -o /tmp/citytimezones-cli ./cmd/citytimezones

# Clean up build artifacts
rm -f /tmp/citytimezones-*

echo "✅ All checks passed!"

# Check if we're ready for pkg.go.dev
echo ""
echo "📦 Ready for pkg.go.dev publication!"
echo ""
echo "Next steps:"
echo "1. Create a GitHub repository: https://github.com/richoandika/city-timezones-go"
echo "2. Push your code to GitHub"
echo "3. Create a git tag: git tag v1.0.0"
echo "4. Push the tag: git push origin v1.0.0"
echo "5. The package will automatically appear on pkg.go.dev within a few minutes"
echo ""
echo "Package URL: https://pkg.go.dev/github.com/richoandika/city-timezones-go"
echo ""
echo "🎉 Happy coding!"
