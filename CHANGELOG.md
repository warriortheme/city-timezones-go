# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- GitHub Actions CI/CD pipeline
- Security scanning with Gosec and Trivy
- Code quality analysis with CodeQL
- Automated dependency updates with Dependabot
- Comprehensive issue and PR templates
- Contributing guidelines

### Changed
- Improved project documentation
- Enhanced error handling
- Better test coverage

## [1.0.0] - 2024-01-01

### Added
- Initial release of city-timezones-go
- Core lookup functions:
  - `LookupViaCity()` - Search by exact city name
  - `FindFromCityStateProvince()` - Partial matching across fields
  - `FindFromIsoCode()` - Search by ISO country codes
  - `SearchCities()` - Advanced search with options
  - `GetCityMapping()` - Get all available cities
- CLI tool with comprehensive options
- Thread-safe operations with sync.Once
- Comprehensive test coverage
- Performance optimizations
- Support for 7,000+ cities with timezone data
- Multiple output formats (table, JSON)
- Filtering by timezone and country
- Result limiting
- Memory-efficient lazy loading
- Go standard project layout

### Features
- **Fast City Lookup**: Find cities by exact name match
- **Partial Matching**: Search across city, state, province, and country fields
- **ISO Code Support**: Look up cities by ISO2 or ISO3 country codes
- **Flexible Search**: Case-sensitive and exact match options
- **Comprehensive Data**: Over 7,000 cities with timezone information
- **Thread-Safe**: Safe for concurrent use
- **CLI Tool**: Command-line interface for quick lookups
- **Well-Tested**: Comprehensive test coverage

### Technical Details
- Go 1.21+ support
- MIT License
- Zero external dependencies
- Cross-platform support (Linux, macOS, Windows)
- ARM64 and AMD64 architectures
- Memory efficient
- Fast lookups with O(n) search complexity

### CLI Usage
```bash
# Search by city name
./citytimezones -city Chicago

# Search with partial matching
./citytimezones -search "springfield mo"

# Search by ISO code
./citytimezones -iso DE -limit 5

# Filter by timezone
./citytimezones -timezone "America/New_York" -limit 3

# Output as JSON
./citytimezones -city Chicago -output json
```

### API Usage
```go
package main

import (
    "fmt"
    "github.com/richoandika/city-timezones-go/pkg/citytimezones"
)

func main() {
    // Look up a city by name
    cities, err := citytimezones.LookupViaCity("Chicago")
    if err != nil {
        log.Fatal(err)
    }
    
    for _, city := range cities {
        fmt.Printf("%s, %s - %s\n", city.City, city.Province, city.Timezone)
    }
}
```

## [0.1.0] - 2024-01-01

### Added
- Initial development version
- Basic city lookup functionality
- Core data structures
- Initial test suite
- Basic CLI implementation

---

## Legend

- **Added** for new features
- **Changed** for changes in existing functionality
- **Deprecated** for soon-to-be removed features
- **Removed** for now removed features
- **Fixed** for any bug fixes
- **Security** for vulnerability fixes