package city

import (
	"strings"
	"testing"
)

func TestLookupViaCity(t *testing.T) {
	t.Run("Find Chicago", func(t *testing.T) {
		cities, err := LookupViaCity("Chicago")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) == 0 {
			t.Error("Should find Chicago")
		}
		if len(cities) > 0 && cities[0].City != "Chicago" {
			t.Errorf("Should be Chicago, got %s", cities[0].City)
		}
	})

	t.Run("Find non-existent city", func(t *testing.T) {
		cities, err := LookupViaCity("NonExistentCity")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) != 0 {
			t.Errorf("Should not find non-existent city, got %d results", len(cities))
		}
	})

	t.Run("Empty city name", func(t *testing.T) {
		cities, err := LookupViaCity("")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) != 0 {
			t.Errorf("Should not find anything for empty name, got %d results", len(cities))
		}
	})

	t.Run("Whitespace only city name", func(t *testing.T) {
		cities, err := LookupViaCity("   ")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) != 0 {
			t.Errorf("Should not find anything for whitespace name, got %d results", len(cities))
		}
	})

	t.Run("Case insensitive search", func(t *testing.T) {
		cities, err := LookupViaCity("chicago")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) == 0 {
			t.Error("Should find Chicago with lowercase")
		}
	})

	t.Run("Search with special characters", func(t *testing.T) {
		cities, err := LookupViaCity("São Paulo")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		// São Paulo might not be in the data, but should not error
		_ = cities // Just check it doesn't panic
	})

	t.Run("Search with numbers", func(t *testing.T) {
		cities, err := LookupViaCity("City123")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		// City123 might not be in the data, but should not error
		_ = cities // Just check it doesn't panic
	})

	t.Run("Search with unicode characters", func(t *testing.T) {
		cities, err := LookupViaCity("São Paulo")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		// São Paulo might not be in the data, but should not error
		_ = cities // Just check it doesn't panic
	})

	t.Run("Search with emoji", func(t *testing.T) {
		cities, err := LookupViaCity("City 🏙️")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		// City with emoji might not be in the data, but should not error
		_ = cities // Just check it doesn't panic
	})

	t.Run("Search with special symbols", func(t *testing.T) {
		cities, err := LookupViaCity("City-Name")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		// City-Name might not be in the data, but should not error
		_ = cities // Just check it doesn't panic
	})

	t.Run("Search with very long input", func(t *testing.T) {
		longInput := strings.Repeat("a", 101) // Exceeds 100 char limit
		cities, err := LookupViaCity(longInput)
		// The validation should error for input too long
		if err == nil {
			t.Error("Should error for input too long")
		}
		if cities != nil {
			t.Errorf("Should not return cities for invalid input, got %d results", len(cities))
		}
	})

	t.Run("Search with malicious input", func(t *testing.T) {
		maliciousInput := "<script>alert('xss')</script>"
		cities, err := LookupViaCity(maliciousInput)
		// The validation should error for malicious input
		if err == nil {
			t.Error("Should error for malicious input")
		}
		if cities != nil {
			t.Errorf("Should not return cities for malicious input, got %d results", len(cities))
		}
	})
}

func TestFindFromCityStateProvince(t *testing.T) {
	t.Run("Find Springfield MO", func(t *testing.T) {
		cities, err := FindFromCityStateProvince("springfield mo")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) == 0 {
			t.Error("Should find Springfield MO")
		}
		if len(cities) > 0 && cities[0].City != "Springfield" {
			t.Errorf("Should be Springfield, got %s", cities[0].City)
		}
	})

	t.Run("Find London", func(t *testing.T) {
		cities, err := FindFromCityStateProvince("london")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) == 0 {
			t.Error("Should find London")
		}
		if len(cities) > 0 && cities[0].City != "London" {
			t.Errorf("Should be London, got %s", cities[0].City)
		}
	})

	t.Run("Empty search string", func(t *testing.T) {
		cities, err := FindFromCityStateProvince("")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) != 0 {
			t.Errorf("Should not find anything for empty search, got %d results", len(cities))
		}
	})

	t.Run("Non-existent search", func(t *testing.T) {
		cities, err := FindFromCityStateProvince("NonExistentCity")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) != 0 {
			t.Errorf("Should not find non-existent city, got %d results", len(cities))
		}
	})

	t.Run("Whitespace only search", func(t *testing.T) {
		cities, err := FindFromCityStateProvince("   ")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) != 0 {
			t.Errorf("Should not find anything for whitespace search, got %d results", len(cities))
		}
	})

	t.Run("Case insensitive search", func(t *testing.T) {
		cities, err := FindFromCityStateProvince("LONDON")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) == 0 {
			t.Error("Should find London with uppercase")
		}
	})

	t.Run("Search with partial match", func(t *testing.T) {
		cities, err := FindFromCityStateProvince("spring")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		// Should find Springfield
		if len(cities) == 0 {
			t.Error("Should find Springfield with partial match")
		}
	})

	t.Run("Search with state only", func(t *testing.T) {
		cities, err := FindFromCityStateProvince("illinois")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		// Should find cities in Illinois
		_ = cities // Just check it doesn't panic
	})

	t.Run("Search with unicode characters", func(t *testing.T) {
		cities, err := FindFromCityStateProvince("São Paulo")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		// São Paulo might not be in the data, but should not error
		_ = cities // Just check it doesn't panic
	})

	t.Run("Search with emoji", func(t *testing.T) {
		cities, err := FindFromCityStateProvince("City 🏙️")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		// City with emoji might not be in the data, but should not error
		_ = cities // Just check it doesn't panic
	})

	t.Run("Search with special symbols", func(t *testing.T) {
		cities, err := FindFromCityStateProvince("City-Name")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		// City-Name might not be in the data, but should not error
		_ = cities // Just check it doesn't panic
	})

	t.Run("Search with very long input", func(t *testing.T) {
		longInput := strings.Repeat("a", 101) // Exceeds 100 char limit
		cities, err := FindFromCityStateProvince(longInput)
		// The validation should pass but the search should work
		if err != nil {
			t.Errorf("Should not error for long input: %v", err)
		}
		// Should not find anything for very long input
		_ = cities // Just check it doesn't panic
	})

	t.Run("Search with malicious input", func(t *testing.T) {
		maliciousInput := "<script>alert('xss')</script>"
		cities, err := FindFromCityStateProvince(maliciousInput)
		// The validation should error for malicious input
		if err == nil {
			t.Error("Should error for malicious input")
		}
		if cities != nil {
			t.Errorf("Should not return cities for malicious input, got %d results", len(cities))
		}
	})
}

func TestFindFromIsoCode(t *testing.T) {
	t.Run("Find by ISO2 DE", func(t *testing.T) {
		cities, err := FindFromIsoCode("DE")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) == 0 {
			t.Error("Should find German cities")
		}
		if len(cities) > 0 && cities[0].ISO2 != "DE" {
			t.Errorf("Should be German cities, got ISO2 %s", cities[0].ISO2)
		}
	})

	t.Run("Find by ISO3 DEU", func(t *testing.T) {
		cities, err := FindFromIsoCode("DEU")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) == 0 {
			t.Error("Should find German cities")
		}
		if len(cities) > 0 && cities[0].ISO3 != "DEU" {
			t.Errorf("Should be German cities, got ISO3 %s", cities[0].ISO3)
		}
	})

	t.Run("Find by ISO2 DE uppercase", func(t *testing.T) {
		cities, err := FindFromIsoCode("DE")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) == 0 {
			t.Error("Should find German cities")
		}
	})

	t.Run("Find non-existent ISO code", func(t *testing.T) {
		cities, err := FindFromIsoCode("XX")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) != 0 {
			t.Errorf("Should not find non-existent ISO code, got %d results", len(cities))
		}
	})

	t.Run("Empty ISO code", func(t *testing.T) {
		cities, err := FindFromIsoCode("")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) != 0 {
			t.Errorf("Should not find anything for empty ISO code, got %d results", len(cities))
		}
	})

	t.Run("Whitespace ISO code", func(t *testing.T) {
		cities, err := FindFromIsoCode("   ")
		// This should error due to validation
		if err == nil {
			t.Error("Should error for whitespace ISO code")
		}
		if len(cities) != 0 {
			t.Errorf("Should not find anything for whitespace ISO code, got %d results", len(cities))
		}
	})

	t.Run("Invalid ISO code format", func(t *testing.T) {
		cities, err := FindFromIsoCode("INVALID")
		// This should error due to validation
		if err == nil {
			t.Error("Should error for invalid ISO code format")
		}
		if len(cities) != 0 {
			t.Errorf("Should not find anything for invalid ISO code, got %d results", len(cities))
		}
	})

	t.Run("Find by ISO2 GB", func(t *testing.T) {
		cities, err := FindFromIsoCode("GB")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) == 0 {
			t.Error("Should find British cities")
		}
		if len(cities) > 0 && cities[0].ISO2 != "GB" {
			t.Errorf("Should be British cities, got ISO2 %s", cities[0].ISO2)
		}
	})

	t.Run("Find by ISO3 GBR", func(t *testing.T) {
		cities, err := FindFromIsoCode("GBR")
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) == 0 {
			t.Error("Should find British cities")
		}
		if len(cities) > 0 && cities[0].ISO3 != "GBR" {
			t.Errorf("Should be British cities, got ISO3 %s", cities[0].ISO3)
		}
	})
}

func TestSearchCities(t *testing.T) {
	t.Run("Case insensitive search", func(t *testing.T) {
		options := SearchOptions{
			CaseSensitive: false,
			ExactMatch:    false,
		}

		cities, err := SearchCities("chicago", options)
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) == 0 {
			t.Error("Should find Chicago")
		}
	})

	t.Run("Case sensitive search", func(t *testing.T) {
		options := SearchOptions{
			CaseSensitive: true,
			ExactMatch:    false,
		}

		cities, err := SearchCities("Chicago", options)
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) == 0 {
			t.Error("Should find Chicago")
		}
	})

	t.Run("Exact match search", func(t *testing.T) {
		options := SearchOptions{
			CaseSensitive: false,
			ExactMatch:    true,
		}

		cities, err := SearchCities("Chicago", options)
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) == 0 {
			t.Error("Should find Chicago")
		}
	})

	t.Run("Empty search string", func(t *testing.T) {
		options := SearchOptions{
			CaseSensitive: false,
			ExactMatch:    false,
		}

		cities, err := SearchCities("", options)
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) != 0 {
			t.Errorf("Should not find anything for empty search, got %d results", len(cities))
		}
	})

	t.Run("Whitespace search string", func(t *testing.T) {
		options := SearchOptions{
			CaseSensitive: false,
			ExactMatch:    false,
		}

		cities, err := SearchCities("   ", options)
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) != 0 {
			t.Errorf("Should not find anything for whitespace search, got %d results", len(cities))
		}
	})

	t.Run("Non-existent search", func(t *testing.T) {
		options := SearchOptions{
			CaseSensitive: false,
			ExactMatch:    false,
		}

		cities, err := SearchCities("NonExistentCity", options)
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		if len(cities) != 0 {
			t.Errorf("Should not find non-existent city, got %d results", len(cities))
		}
	})

	t.Run("Search with special characters", func(t *testing.T) {
		options := SearchOptions{
			CaseSensitive: false,
			ExactMatch:    false,
		}

		cities, err := SearchCities("São Paulo", options)
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		// São Paulo might not be in the data, but should not error
		_ = cities // Just check it doesn't panic
	})

	t.Run("Search with numbers and symbols", func(t *testing.T) {
		options := SearchOptions{
			CaseSensitive: false,
			ExactMatch:    false,
		}

		cities, err := SearchCities("City-123", options)
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		// City-123 might not be in the data, but should not error
		_ = cities // Just check it doesn't panic
	})

	t.Run("Search with very long string", func(t *testing.T) {
		options := SearchOptions{
			CaseSensitive: false,
			ExactMatch:    false,
		}

		longString := "This is a very long search string that might exceed normal limits"
		cities, err := SearchCities(longString, options)
		if err != nil {
			t.Errorf("Should not error: %v", err)
		}
		// Should not find anything for very long string
		_ = cities // Just check it doesn't panic
	})

	t.Run("Search with very long input that exceeds limit", func(t *testing.T) {
		options := SearchOptions{
			CaseSensitive: false,
			ExactMatch:    false,
		}

		longInput := strings.Repeat("a", 101) // Exceeds 100 char limit
		cities, err := SearchCities(longInput, options)
		// SearchCities doesn't validate input, so it should work
		if err != nil {
			t.Errorf("Should not error for long input: %v", err)
		}
		// Should not find anything for very long input
		_ = cities // Just check it doesn't panic
	})

	t.Run("Search with malicious input", func(t *testing.T) {
		options := SearchOptions{
			CaseSensitive: false,
			ExactMatch:    false,
		}

		maliciousInput := "<script>alert('xss')</script>"
		cities, err := SearchCities(maliciousInput, options)
		// SearchCities doesn't validate input, so it should work
		if err != nil {
			t.Errorf("Should not error for malicious input: %v", err)
		}
		// Should not find anything for malicious input
		_ = cities // Just check it doesn't panic
	})
}
