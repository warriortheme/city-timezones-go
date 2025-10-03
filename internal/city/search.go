package city

import (
	"fmt"
	"strings"
)

// LookupViaCity searches for cities by exact city name match
func LookupViaCity(cityName string) ([]CityData, error) {
	// Validate and sanitize input
	validatedInput, err := ValidateSearchInput(cityName, 100) // Max 100 chars for city name
	if err != nil {
		return nil, fmt.Errorf("invalid input: %w", err)
	}

	if validatedInput == "" {
		return []CityData{}, nil
	}

	// Check cache first
	cacheKey := "city:" + strings.ToLower(validatedInput)
	if cached, exists := GetCachedResult(cacheKey); exists {
		return cached, nil
	}

	cities, err := LoadCityData()
	if err != nil {
		return nil, err
	}

	var results []CityData
	searchTerm := strings.ToLower(validatedInput)

	for _, city := range cities {
		if strings.ToLower(city.City) == searchTerm {
			results = append(results, city)
		}
	}

	// Cache the result
	SetCachedResult(cacheKey, results)

	return results, nil
}

// FindFromCityStateProvince searches for cities using partial matching
// across city, state, province, and country fields
func FindFromCityStateProvince(searchString string) ([]CityData, error) {
	// Validate and sanitize input
	validatedInput, err := ValidateSearchInput(searchString, 200) // Max 200 chars for search string
	if err != nil {
		return nil, fmt.Errorf("invalid input: %w", err)
	}

	if validatedInput == "" {
		return []CityData{}, nil
	}

	cities, err := LoadCityData()
	if err != nil {
		return nil, err
	}

	var results []CityData
	searchTerms := strings.Fields(strings.ToLower(validatedInput))

	for _, city := range cities {
		if findPartialMatch(city, searchTerms) {
			results = append(results, city)
		}
	}

	return results, nil
}

// FindFromIsoCode searches for cities by ISO2 or ISO3 country codes
func FindFromIsoCode(isoCode string) ([]CityData, error) {
	// Validate ISO code
	validatedCode, err := ValidateISOCode(isoCode)
	if err != nil {
		return nil, fmt.Errorf("invalid ISO code: %w", err)
	}

	if validatedCode == "" {
		return []CityData{}, nil
	}

	cities, err := LoadCityData()
	if err != nil {
		return nil, err
	}

	var results []CityData
	searchCode := strings.ToLower(validatedCode)

	for _, city := range cities {
		if strings.ToLower(city.ISO2) == searchCode || strings.ToLower(city.ISO3) == searchCode {
			results = append(results, city)
		}
	}

	return results, nil
}

// findPartialMatch checks if all search terms are found in the city's searchable fields
func findPartialMatch(city CityData, searchTerms []string) bool {
	// Create a combined searchable text from all relevant fields
	searchableFields := []string{
		city.City,
		city.StateANSI,
		city.Province,
		city.Country,
	}

	combinedText := strings.ToLower(strings.Join(searchableFields, " "))

	// Check if all search terms are found in the combined text
	for _, term := range searchTerms {
		if !strings.Contains(combinedText, term) {
			return false
		}
	}

	return true
}

// SearchCities provides a flexible search function with options
func SearchCities(query string, options SearchOptions) ([]CityData, error) {
	if query == "" {
		return []CityData{}, nil
	}

	cities, err := LoadCityData()
	if err != nil {
		return nil, err
	}

	var results []CityData
	searchQuery := query
	if !options.CaseSensitive {
		searchQuery = strings.ToLower(searchQuery)
	}

	for _, city := range cities {
		if matchesCity(city, searchQuery, options) {
			results = append(results, city)
		}
	}

	return results, nil
}

// matchesCity checks if a city matches the search criteria
func matchesCity(city CityData, query string, options SearchOptions) bool {
	searchableFields := []string{
		city.City,
		city.CityASCII,
		city.StateANSI,
		city.Province,
		city.Country,
		city.ISO2,
		city.ISO3,
	}

	for _, field := range searchableFields {
		fieldValue := field
		if !options.CaseSensitive {
			fieldValue = strings.ToLower(fieldValue)
		}

		if options.ExactMatch {
			if fieldValue == query {
				return true
			}
		} else {
			if strings.Contains(fieldValue, query) {
				return true
			}
		}
	}

	return false
}
