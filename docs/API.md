# API Reference

## Package: citytimezones

The `citytimezones` package provides a comprehensive API for looking up timezone information by city name, with support for various search methods and flexible options.

### Types

#### CityData

```go
type CityData struct {
    Lat           float64 `json:"lat"`
    Lng           float64 `json:"lng"`
    Pop           float64 `json:"pop"`
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

Represents a city with its timezone and geographical information.

#### SearchOptions

```go
type SearchOptions struct {
    CaseSensitive bool
    ExactMatch    bool
}
```

Provides configuration for search operations.

### Functions

#### LookupViaCity

```go
func LookupViaCity(cityName string) ([]CityData, error)
```

Searches for cities by exact city name match (case-insensitive).

**Parameters:**
- `cityName` (string): The city name to search for

**Returns:**
- `[]CityData`: Array of matching cities
- `error`: Error if the search fails

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

#### FindFromCityStateProvince

```go
func FindFromCityStateProvince(searchString string) ([]CityData, error)
```

Searches for cities using partial matching across city, state, province, and country fields.

**Parameters:**
- `searchString` (string): Search term that can match city, state, province, or country

**Returns:**
- `[]CityData`: Array of matching cities
- `error`: Error if the search fails

**Example:**
```go
cities, err := citytimezones.FindFromCityStateProvince("springfield mo")
if err != nil {
    log.Fatal(err)
}
```

#### FindFromIsoCode

```go
func FindFromIsoCode(isoCode string) ([]CityData, error)
```

Searches for cities by ISO2 or ISO3 country codes.

**Parameters:**
- `isoCode` (string): ISO2 or ISO3 country code

**Returns:**
- `[]CityData`: Array of matching cities
- `error`: Error if the search fails

**Example:**
```go
cities, err := citytimezones.FindFromIsoCode("DE")
if err != nil {
    log.Fatal(err)
}
```

#### SearchCities

```go
func SearchCities(query string, options SearchOptions) ([]CityData, error)
```

Advanced search with configurable options.

**Parameters:**
- `query` (string): Search query
- `options` (SearchOptions): Search configuration

**Returns:**
- `[]CityData`: Array of matching cities
- `error`: Error if the search fails

**Example:**
```go
options := citytimezones.SearchOptions{
    CaseSensitive: true,
    ExactMatch:    false,
}
cities, err := citytimezones.SearchCities("Chicago", options)
```

#### GetCityMapping

```go
func GetCityMapping() ([]CityData, error)
```

Returns all available cities in the database.

**Returns:**
- `[]CityData`: Array of all cities
- `error`: Error if loading fails

**Example:**
```go
allCities, err := citytimezones.GetCityMapping()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Total cities: %d\n", len(allCities))
```

#### DefaultSearchOptions

```go
func DefaultSearchOptions() SearchOptions
```

Returns the default search configuration.

**Returns:**
- `SearchOptions`: Default search options

**Example:**
```go
options := citytimezones.DefaultSearchOptions()
// options.CaseSensitive = false
// options.ExactMatch = false
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

Common error scenarios:
- Data file not found
- JSON parsing errors
- Invalid search parameters

## Thread Safety

The library is thread-safe and can be used concurrently from multiple goroutines. The city data is loaded once and cached for subsequent requests.

## Performance Considerations

- City data is loaded lazily on first access
- Data is cached in memory after first load
- Search operations are O(n) with early termination
- Memory usage is optimized for the dataset size
