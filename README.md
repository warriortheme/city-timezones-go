# City Timezones Go

[![Go Reference](https://pkg.go.dev/badge/github.com/richoandika/city-timezones-go.svg)](https://pkg.go.dev/github.com/richoandika/city-timezones-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/richoandika/city-timezones-go)](https://goreportcard.com/report/github.com/richoandika/city-timezones-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A fast and efficient Go library for looking up timezone information by city name, with support for partial matching, ISO code lookups, and flexible search options.

> **Note**: This is a Go port of the original [city-timezones](https://github.com/kevinroberts/city-timezones) JavaScript library by [Kevin Roberts](https://github.com/kevinroberts). This Go implementation maintains API compatibility while providing significant performance improvements and type safety.

## Features

- **Fast City Lookup**: Find cities by exact name match
- **Partial Matching**: Search across city, state, province, and country fields
- **ISO Code Support**: Look up cities by ISO2 or ISO3 country codes
- **Flexible Search**: Case-sensitive and exact match options
- **Comprehensive Data**: Over 7,000 cities with timezone information
- **Thread-Safe**: Safe for concurrent use
- **CLI Tool**: Command-line interface for quick lookups
- **Well-Tested**: Comprehensive test coverage

## Installation

```bash
go get github.com/richoandika/city-timezones-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/richoandika/city-timezones-go/pkg/citytimezones"
)

func main() {
    // Look up a city by name
    cities, err := citytimezones.LookupViaCity("Chicago")
    if err != nil {
        log.Fatal(err)
    }
    
    if len(cities) > 0 {
        city := cities[0]
        fmt.Printf("Found: %s, %s\n", city.City, city.Province)
        fmt.Printf("Timezone: %s\n", city.Timezone)
        fmt.Printf("Coordinates: %.4f, %.4f\n", city.Lat, city.Lng)
    }
}
```

## API Reference

### Core Functions

#### `LookupViaCity(cityName string) ([]CityData, error)`
Searches for cities by exact city name match (case-insensitive).

```go
cities, err := citytimezones.LookupViaCity("Chicago")
```

#### `FindFromCityStateProvince(searchString string) ([]CityData, error)`
Searches for cities using partial matching across city, state, province, and country fields.

```go
cities, err := citytimezones.FindFromCityStateProvince("springfield mo")
```

#### `FindFromIsoCode(isoCode string) ([]CityData, error)`
Searches for cities by ISO2 or ISO3 country codes.

```go
cities, err := citytimezones.FindFromIsoCode("DE")
```

#### `SearchCities(query string, options SearchOptions) ([]CityData, error)`
Advanced search with configurable options.

```go
options := citytimezones.SearchOptions{
    CaseSensitive: true,
    ExactMatch:    false,
}
cities, err := citytimezones.SearchCities("Chicago", options)
```

#### `GetCityMapping() ([]CityData, error)`
Returns all available cities in the database.

```go
allCities, err := citytimezones.GetCityMapping()
```

### Data Structure

```go
type CityData struct {
    Lat           float64 `json:"lat"`
    Lng           float64 `json:"lng"`
    Pop           int     `json:"pop"`
    City          string  `json:"city"`
    ISO2          string  `json:"iso2"`
    ISO3          string  `json:"iso3"`
    Country       string  `json:"country"`
    Timezone      string  `json:"timezone"`
    Province      string  `json:"province"`
    ExactCity     string  `json:"exactCity"`
    CityASCII     string  `json:"city_ascii"`
    StateANSI     string  `json:"state_ansi"`
    ExactProvince string  `json:"exactProvince"`
}
```

## Examples

### Basic Usage

```go
// Find Chicago
cities, err := citytimezones.LookupViaCity("Chicago")
if err != nil {
    log.Fatal(err)
}

for _, city := range cities {
    fmt.Printf("%s, %s - %s\n", city.City, city.Province, city.Timezone)
}
```

### Partial Matching

```go
// Find Springfield, Missouri
cities, err := citytimezones.FindFromCityStateProvince("springfield mo")
if err != nil {
    log.Fatal(err)
}

for _, city := range cities {
    fmt.Printf("%s, %s - %s\n", city.City, city.Province, city.Timezone)
}
```

### ISO Code Lookup

```go
// Find all German cities
cities, err := citytimezones.FindFromIsoCode("DE")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Found %d German cities\n", len(cities))
```

### Advanced Search

```go
// Case-sensitive search
options := citytimezones.SearchOptions{
    CaseSensitive: true,
    ExactMatch:    false,
}
cities, err := citytimezones.SearchCities("Chicago", options)
```

## Command Line Interface

The package includes a CLI tool for quick lookups:

```bash
# Build the CLI
make build

# Search for a city
./bin/citytimezones -city Chicago

# Search with partial matching
./bin/citytimezones -search "springfield mo"

# Search by ISO code
./bin/citytimezones -iso DE -limit 5

# Filter by timezone
./bin/citytimezones -timezone "America/New_York" -limit 3

# Output as JSON
./bin/citytimezones -city Chicago -output json
```

## Development

### Project Structure

```
city-timezones-go/
├── cmd/
│   └── citytimezones/          # CLI application
├── examples/
│   ├── basic/                  # Basic usage examples
│   └── advanced/               # Advanced usage examples
├── internal/
│   └── city/                   # Internal city package
├── pkg/
│   └── citytimezones/          # Public API package
├── data/
│   └── cityMap.json           # City data
├── Makefile                   # Build automation
└── README.md
```

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run benchmarks
make bench
```

### Building

```bash
# Build CLI tool
make build

# Run examples
make run-examples

# Run CLI examples
make run-cli-examples
```

## Performance

The library is optimized for performance:

- **Lazy Loading**: City data is loaded only when needed
- **Thread-Safe**: Safe for concurrent use with sync.Once
- **Memory Efficient**: Minimal memory footprint
- **Fast Lookups**: O(n) search with early termination

## Error Handling

All functions return errors that should be checked:

```go
cities, err := citytimezones.LookupViaCity("Chicago")
if err != nil {
    log.Printf("Lookup failed: %v", err)
    return
}
```

## Thread Safety

The library is thread-safe and can be used concurrently from multiple goroutines.

## Attribution

This Go library is a port of the original [city-timezones](https://github.com/kevinroberts/city-timezones) JavaScript library created by [Kevin Roberts](https://github.com/kevinroberts). The original library has been widely adopted with over 150 stars and 3,300+ users on GitHub.

### Original Project
- **Repository**: [kevinroberts/city-timezones](https://github.com/kevinroberts/city-timezones)
- **NPM Package**: [city-timezones](https://www.npmjs.com/package/city-timezones)
- **Author**: [Kevin Roberts](https://github.com/kevinroberts)

### Go Port Improvements
- **Type Safety**: Compile-time error checking
- **Performance**: Faster execution and lower memory usage
- **Concurrency**: Thread-safe operations
- **Error Handling**: Explicit error handling
- **Maintainability**: Clear separation of concerns
- **Testing**: Comprehensive test coverage

## License

MIT License - see LICENSE file for details.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Run `make check`
6. Submit a pull request

## Changelog

### v1.0.0
- Initial release
- Core lookup functions
- CLI tool
- Comprehensive test coverage
- Go standard project layout