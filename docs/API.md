# API Documentation

This document provides comprehensive API documentation for the City Timezones Go library.

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [API Reference](#api-reference)
- [Data Structures](#data-structures)
- [Error Handling](#error-handling)
- [Performance](#performance)
- [Examples](#examples)

## Overview

The City Timezones Go library provides fast and efficient timezone lookup functionality for cities worldwide. It supports various search methods including exact name matching, partial matching, and ISO code lookups.

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

**Parameters:**
- `cityName` (string): The name of the city to search for

**Returns:**
- `[]CityData`: Slice of matching cities
- `error`: Error if the lookup fails

**Example:**
```go
cities, err := citytimezones.LookupViaCity("Chicago")
if err != nil {
    log.Fatal(err)
}

for _, city := range cities {
    fmt.Printf("%s, %s - %s\n", city.City, city.Province, city.Timezone)
}
```

#### `FindFromCityStateProvince(searchString string) ([]CityData, error)`

Searches for cities using partial matching across city, state, province, and country fields.

**Parameters:**
- `searchString` (string): Search string to match against city, state, province, or country

**Returns:**
- `[]CityData`: Slice of matching cities
- `error`: Error if the lookup fails

**Example:**
```go
cities, err := citytimezones.FindFromCityStateProvince("springfield mo")
if err != nil {
    log.Fatal(err)
}

for _, city := range cities {
    fmt.Printf("%s, %s - %s\n", city.City, city.Province, city.Timezone)
}
```

#### `FindFromIsoCode(isoCode string) ([]CityData, error)`

Searches for cities by ISO2 or ISO3 country codes.

**Parameters:**
- `isoCode` (string): ISO2 or ISO3 country code

**Returns:**
- `[]CityData`: Slice of cities in the specified country
- `error`: Error if the lookup fails

**Example:**
```go
cities, err := citytimezones.FindFromIsoCode("DE")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Found %d German cities\n", len(cities))
```

#### `SearchCities(query string, options SearchOptions) ([]CityData, error)`

Advanced search with configurable options.

**Parameters:**
- `query` (string): Search query
- `options` (SearchOptions): Search configuration options

**Returns:**
- `[]CityData`: Slice of matching cities
- `error`: Error if the lookup fails

**Example:**
```go
options := citytimezones.SearchOptions{
    CaseSensitive: true,
    ExactMatch:    false,
}
cities, err := citytimezones.SearchCities("Chicago", options)
```

#### `GetCityMapping() ([]CityData, error)`

Returns all available cities in the database.

**Returns:**
- `[]CityData`: All available cities
- `error`: Error if the lookup fails

**Example:**
```go
allCities, err := citytimezones.GetCityMapping()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Total cities: %d\n", len(allCities))
```

#### `DefaultSearchOptions() SearchOptions`

Returns the default search configuration.

**Returns:**
- `SearchOptions`: Default search options

**Example:**
```go
options := citytimezones.DefaultSearchOptions()
// Use options for search
```

## Data Structures

### CityData

Represents a city with its timezone and geographical information.

```go
type CityData struct {
    Lat           float64 `json:"lat"`            // Latitude
    Lng           float64 `json:"lng"`            // Longitude
    Pop           int     `json:"pop"`            // Population
    City          string  `json:"city"`           // City name
    ISO2          string  `json:"iso2"`           // ISO2 country code
    ISO3          string  `json:"iso3"`           // ISO3 country code
    Country       string  `json:"country"`        // Country name
    Timezone      string  `json:"timezone"`       // Timezone identifier
    Province      string  `json:"province"`       // Province/state name
    ExactCity     string  `json:"exactCity"`     // Exact city name
    CityASCII     string  `json:"city_ascii"`    // ASCII city name
    StateANSI     string  `json:"state_ansi"`    // ANSI state code
    ExactProvince string  `json:"exactProvince"` // Exact province name
}
```

### SearchOptions

Configuration options for search operations.

```go
type SearchOptions struct {
    CaseSensitive bool `json:"case_sensitive"` // Whether search is case-sensitive
    ExactMatch    bool `json:"exact_match"`    // Whether to use exact matching
}
```

## Error Handling

All functions return errors that should be checked:

```go
cities, err := citytimezones.LookupViaCity("Chicago")
if err != nil {
    log.Printf("Lookup failed: %v", err)
    return
}
```

### Common Error Types

- **Empty input**: When city name or search string is empty
- **Invalid ISO code**: When ISO code format is invalid
- **Data loading errors**: When city data cannot be loaded
- **Search errors**: When search operation fails

## Performance

The library is optimized for performance:

- **Lazy Loading**: City data is loaded only when needed
- **Thread-Safe**: Safe for concurrent use with sync.Once
- **Memory Efficient**: Minimal memory footprint
- **Fast Lookups**: O(n) search with early termination

### Performance Characteristics

- **Memory Usage**: ~2MB for full dataset
- **Lookup Speed**: < 1ms for typical searches
- **Concurrency**: Thread-safe for concurrent access
- **Caching**: Automatic caching of loaded data

## Examples

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/richoandika/city-timezones-go/pkg/citytimezones"
)

func main() {
    // Find Chicago
    cities, err := citytimezones.LookupViaCity("Chicago")
    if err != nil {
        log.Fatal(err)
    }

    for _, city := range cities {
        fmt.Printf("%s, %s - %s\n", city.City, city.Province, city.Timezone)
    }
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

### Error Handling

```go
cities, err := citytimezones.LookupViaCity("")
if err != nil {
    log.Printf("Error: %v", err)
    return
}
```

## Thread Safety

The library is thread-safe and can be used concurrently from multiple goroutines:

```go
var wg sync.WaitGroup

for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        cities, err := citytimezones.LookupViaCity("Chicago")
        if err != nil {
            log.Printf("Error: %v", err)
            return
        }
        // Process cities...
    }()
}

wg.Wait()
```

## Best Practices

1. **Always check errors**: All functions return errors that should be handled
2. **Use appropriate search methods**: Choose the right function for your use case
3. **Handle empty results**: Check if the returned slice is empty
4. **Consider performance**: Use specific search methods when possible
5. **Thread safety**: The library is safe for concurrent use

## Migration from JavaScript Version

If you're migrating from the JavaScript version:

| JavaScript | Go |
|------------|----|
| `cityTimezones.lookupViaCity()` | `citytimezones.LookupViaCity()` |
| `cityTimezones.findFromCityStateProvince()` | `citytimezones.FindFromCityStateProvince()` |
| `cityTimezones.findFromIsoCode()` | `citytimezones.FindFromIsoCode()` |
| `cityTimezones.searchCities()` | `citytimezones.SearchCities()` |
| `cityTimezones.getCityMapping()` | `citytimezones.GetCityMapping()` |

The Go version maintains API compatibility while providing better performance and type safety.