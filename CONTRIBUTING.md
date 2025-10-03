# Contributing to City Timezones Go

Thank you for your interest in contributing to City Timezones Go! This document provides guidelines and information for contributors.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Contributing Process](#contributing-process)
- [Code Style Guidelines](#code-style-guidelines)
- [Testing Guidelines](#testing-guidelines)
- [Documentation Guidelines](#documentation-guidelines)
- [Release Process](#release-process)

## Code of Conduct

This project follows the [Contributor Covenant Code of Conduct](https://www.contributor-covenant.org/version/2/1/code_of_conduct/). By participating, you are expected to uphold this code.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/your-username/city-timezones-go.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Test your changes
6. Commit your changes: `git commit -m "Add your feature"`
7. Push to your fork: `git push origin feature/your-feature-name`
8. Create a Pull Request

## Development Setup

### Prerequisites

- Go 1.21 or later
- Git
- Make (optional, for using Makefile commands)

### Setup Steps

1. **Clone the repository**
   ```bash
   git clone https://github.com/richoandika/city-timezones-go.git
   cd city-timezones-go
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run tests**
   ```bash
   make test
   ```

4. **Build the project**
   ```bash
   make build
   ```

## Contributing Process

### Types of Contributions

We welcome several types of contributions:

- **Bug fixes**: Fix issues and improve stability
- **New features**: Add new functionality
- **Documentation**: Improve documentation and examples
- **Performance improvements**: Optimize code and algorithms
- **Tests**: Add or improve test coverage
- **Examples**: Add usage examples

### Pull Request Process

1. **Create an issue** (for significant changes) to discuss the proposed change
2. **Fork the repository** and create a feature branch
3. **Make your changes** following our coding standards
4. **Add tests** for new functionality
5. **Update documentation** if necessary
6. **Run the full test suite** to ensure nothing is broken
7. **Submit a pull request** with a clear description

### Pull Request Guidelines

- Use a clear, descriptive title
- Provide a detailed description of changes
- Reference any related issues
- Ensure all tests pass
- Update documentation if needed
- Keep commits focused and atomic

## Code Style Guidelines

### Go Code Style

- Follow standard Go formatting: `gofmt` and `goimports`
- Use meaningful variable and function names
- Write clear, self-documenting code
- Add comments for exported functions and types
- Keep functions focused and small
- Use interfaces where appropriate

### Code Organization

- Keep related functionality together
- Separate concerns into different packages
- Use the standard Go project layout
- Place public APIs in the `pkg/` directory
- Keep internal implementation in the `internal/` directory

### Example Code Style

```go
// Package citytimezones provides timezone lookup functionality for cities.
package citytimezones

import (
    "fmt"
    "errors"
)

// CityData represents a city with its timezone and geographical information.
type CityData struct {
    City     string  `json:"city"`
    Country  string  `json:"country"`
    Timezone string  `json:"timezone"`
    Lat      float64 `json:"lat"`
    Lng      float64 `json:"lng"`
}

// LookupViaCity searches for cities by exact city name match.
// It returns a slice of CityData and an error if the lookup fails.
func LookupViaCity(cityName string) ([]CityData, error) {
    if cityName == "" {
        return nil, errors.New("city name cannot be empty")
    }
    
    // Implementation here...
    return cities, nil
}
```

## Testing Guidelines

### Test Requirements

- All new code must include tests
- Aim for high test coverage (80%+)
- Write both unit tests and integration tests
- Include edge cases and error conditions
- Use table-driven tests where appropriate

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run specific test packages
go test ./pkg/citytimezones
go test ./internal/city

# Run benchmarks
make bench
```

### Test Structure

```go
func TestLookupViaCity(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected []CityData
        wantErr  bool
    }{
        {
            name:     "valid city name",
            input:    "Chicago",
            expected: []CityData{{City: "Chicago", Country: "United States"}},
            wantErr:  false,
        },
        {
            name:     "empty city name",
            input:    "",
            expected: nil,
            wantErr:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := LookupViaCity(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("LookupViaCity() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            // Add more assertions...
        })
    }
}
```

## Documentation Guidelines

### Code Documentation

- Document all exported functions, types, and variables
- Use clear, concise descriptions
- Include usage examples for complex functions
- Follow Go documentation conventions

### README Updates

- Keep the README up to date with new features
- Include usage examples
- Update installation instructions if needed
- Add new CLI options or API changes

### API Documentation

- Document all public APIs
- Include parameter descriptions
- Provide return value information
- Add usage examples

## Release Process

### Versioning

We follow [Semantic Versioning](https://semver.org/):

- **MAJOR** version for incompatible API changes
- **MINOR** version for new functionality in a backwards compatible manner
- **PATCH** version for backwards compatible bug fixes

### Release Checklist

- [ ] All tests pass
- [ ] Documentation is updated
- [ ] CHANGELOG.md is updated
- [ ] Version is bumped in go.mod
- [ ] Release notes are prepared
- [ ] Tag is created with proper version

## Getting Help

If you need help or have questions:

1. Check existing issues and discussions
2. Create a new issue with the "question" label
3. Join our community discussions
4. Contact the maintainers

## Recognition

Contributors will be recognized in:

- CONTRIBUTORS.md file
- Release notes
- Project documentation

Thank you for contributing to City Timezones Go! 🚀