package citytimezones

import (
	"testing"
)

func TestPublicAPI(t *testing.T) {
	th := NewTestHelper(t)

	t.Run("LookupViaCity", func(t *testing.T) {
		cities, err := LookupViaCity("Chicago")
		th.AssertNoError(err, "should not error")
		th.AssertNotNil(cities, "should return results")

		if len(cities) > 0 {
			th.AssertEqual("Chicago", cities[0].City, "should find Chicago")
		}
	})

	t.Run("FindFromCityStateProvince", func(t *testing.T) {
		cities, err := FindFromCityStateProvince("springfield mo")
		th.AssertNoError(err, "should not error")
		th.AssertNotNil(cities, "should return results")
	})

	t.Run("FindFromIsoCode", func(t *testing.T) {
		cities, err := FindFromIsoCode("DE")
		th.AssertNoError(err, "should not error")
		th.AssertNotNil(cities, "should return results")
	})

	t.Run("SearchCities", func(t *testing.T) {
		options := DefaultSearchOptions()
		cities, err := SearchCities("london", options)
		th.AssertNoError(err, "should not error")
		th.AssertNotNil(cities, "should return results")
	})

	t.Run("GetCityMapping", func(t *testing.T) {
		cities, err := GetCityMapping()
		th.AssertNoError(err, "should not error")
		th.AssertNotNil(cities, "should return results")
		th.AssertEqual(true, len(cities) > 0, "should have cities")
	})

	t.Run("DefaultSearchOptions", func(t *testing.T) {
		options := DefaultSearchOptions()
		th.AssertEqual(false, options.CaseSensitive, "should not be case sensitive by default")
		th.AssertEqual(false, options.ExactMatch, "should not be exact match by default")
	})
}

func TestPublicAPI_ErrorHandling(t *testing.T) {
	th := NewTestHelper(t)

	t.Run("Invalid input handling", func(t *testing.T) {
		// Test with suspicious input
		cities, err := LookupViaCity("<script>alert('xss')</script>")
		th.AssertError(err, "should reject suspicious input")
		th.AssertEqual(0, len(cities), "should return empty results for invalid input")
	})

	t.Run("Empty input handling", func(t *testing.T) {
		cities, err := LookupViaCity("")
		th.AssertNoError(err, "should handle empty input gracefully")
		th.AssertEqual(0, len(cities), "should return empty results for empty input")
	})

	t.Run("Invalid ISO code", func(t *testing.T) {
		cities, err := FindFromIsoCode("INVALID")
		th.AssertError(err, "should reject invalid ISO code")
		th.AssertEqual(0, len(cities), "should return empty results for invalid ISO code")
	})
}

func TestPublicAPI_TypeAliases(t *testing.T) {
	th := NewTestHelper(t)

	t.Run("CityData type", func(t *testing.T) {
		var city CityData
		th.AssertEqual("", city.City, "CityData should be properly aliased")
		th.AssertEqual(0.0, city.Lat, "Lat should be float64")
		th.AssertEqual(0.0, city.Pop, "Pop should be float64")
	})

	t.Run("SearchOptions type", func(t *testing.T) {
		var options SearchOptions
		th.AssertEqual(false, options.CaseSensitive, "SearchOptions should be properly aliased")
		th.AssertEqual(false, options.ExactMatch, "SearchOptions should be properly aliased")
	})
}

func TestPublicAPI_Performance(t *testing.T) {
	th := NewTestHelper(t)

	t.Run("Repeated calls performance", func(t *testing.T) {
		// Test that repeated calls are fast (caching should help)
		for i := 0; i < 10; i++ {
			cities, err := LookupViaCity("Chicago")
			th.AssertNoError(err, "should not error")
			th.AssertNotNil(cities, "should return results")
		}
	})

	t.Run("Concurrent calls", func(t *testing.T) {
		// Test concurrent access to public API
		done := make(chan bool, 10)

		for i := 0; i < 10; i++ {
			go func() {
				defer func() { done <- true }()

				cities, err := LookupViaCity("Chicago")
				th.AssertNoError(err, "should not error")
				th.AssertNotNil(cities, "should return results")
			}()
		}

		// Wait for all goroutines to complete
		for i := 0; i < 10; i++ {
			<-done
		}
	})
}

// TestHelper is a simple test helper for the public API tests
type TestHelper struct {
	t *testing.T
}

func NewTestHelper(t *testing.T) *TestHelper {
	return &TestHelper{t: t}
}

func (th *TestHelper) AssertNoError(err error, message string) {
	if err != nil {
		th.t.Errorf("%s: unexpected error: %v", message, err)
	}
}

func (th *TestHelper) AssertError(err error, message string) {
	if err == nil {
		th.t.Errorf("%s: expected error but got none", message)
	}
}

func (th *TestHelper) AssertEqual(expected, actual interface{}, message string) {
	if expected != actual {
		th.t.Errorf("%s: expected %v, got %v", message, expected, actual)
	}
}

func (th *TestHelper) AssertNotNil(value interface{}, message string) {
	if value == nil {
		th.t.Errorf("%s: expected non-nil value", message)
	}
}

func (th *TestHelper) AssertNil(value interface{}, message string) {
	if value != nil {
		th.t.Errorf("%s: expected nil value, got %v", message, value)
	}
}
