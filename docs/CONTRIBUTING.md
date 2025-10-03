# Contributing to City Timezones Go

Thank you for your interest in contributing to City Timezones Go! This document provides guidelines for contributing to the project.

## Development Setup

### Prerequisites

- Go 1.21 or later
- Git
- Make (optional, for using Makefile commands)

### Getting Started

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/your-username/city-timezones-go.git
   cd city-timezones-go
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run tests to ensure everything works:
   ```bash
   make test
   ```

## Project Structure

The project follows the standard Go project layout:

```
city-timezones-go/
├── cmd/                    # Command-line applications
│   └── citytimezones/     # CLI tool
├── examples/              # Example applications
│   ├── basic/             # Basic usage examples
│   └── advanced/          # Advanced usage examples
├── internal/              # Private application code
│   └── city/              # Internal city package
├── pkg/                   # Public library code
│   └── citytimezones/     # Public API package
├── data/                  # Data files
│   └── cityMap.json       # City data
├── docs/                  # Documentation
├── Makefile              # Build automation
└── README.md             # Project documentation
```

## Development Guidelines

### Code Style

- Follow Go standard formatting (`go fmt`)
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions focused and small
- Use proper error handling

### Testing

- Write tests for all new functionality
- Maintain or improve test coverage
- Use table-driven tests where appropriate
- Test both success and error cases

### Documentation

- Update README.md for user-facing changes
- Update API.md for API changes
- Add inline comments for complex logic
- Update examples if needed

## Making Changes

### 1. Create a Feature Branch

```bash
git checkout -b feature/your-feature-name
```

### 2. Make Your Changes

- Write your code following the guidelines above
- Add tests for new functionality
- Update documentation as needed

### 3. Run Quality Checks

```bash
# Format code
make fmt

# Lint code
make lint

# Run tests
make test

# Run all checks
make check
```

### 4. Commit Your Changes

```bash
git add .
git commit -m "Add your descriptive commit message"
```

Use conventional commit messages:
- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation changes
- `test:` for test additions/changes
- `refactor:` for code refactoring

### 5. Push and Create Pull Request

```bash
git push origin feature/your-feature-name
```

Then create a pull request on GitHub.

## Testing

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run benchmarks
make bench
```

### Test Structure

- Unit tests go in `*_test.go` files
- Use table-driven tests for multiple test cases
- Test both success and error scenarios
- Mock external dependencies when needed

### Example Test

```go
func TestLookupViaCity(t *testing.T) {
    tests := []struct {
        name     string
        cityName string
        expected int
        hasError bool
    }{
        {
            name:     "Find Chicago",
            cityName: "Chicago",
            expected: 1,
            hasError: false,
        },
        // ... more test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            results, err := LookupViaCity(tt.cityName)
            // ... assertions
        })
    }
}
```

## Building and Running

### Build Commands

```bash
# Build CLI tool
make build

# Run examples
make run-examples

# Run CLI examples
make run-cli-examples
```

### Manual Build

```bash
# Build CLI
go build -o bin/citytimezones ./cmd/citytimezones

# Run examples
go run ./examples/basic
go run ./examples/advanced
```

## Data Updates

If you need to update the city data:

1. Ensure the JSON structure remains compatible
2. Update tests if the data changes affect expected results
3. Document any breaking changes
4. Consider versioning for major data updates

## Performance Considerations

- Profile code for performance bottlenecks
- Use benchmarks to measure improvements
- Consider memory usage for large datasets
- Optimize search algorithms if needed

## Release Process

1. Update version numbers in relevant files
2. Update CHANGELOG.md
3. Create a release tag
4. Update documentation

## Code Review Process

- All changes require code review
- Reviewers should check:
  - Code quality and style
  - Test coverage
  - Documentation updates
  - Performance implications
  - Breaking changes

## Questions and Support

- Open an issue for questions or bug reports
- Use discussions for general questions
- Check existing issues before creating new ones

## License

By contributing, you agree that your contributions will be licensed under the same license as the project (MIT License).
