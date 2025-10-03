package city

import (
	"testing"
)

func TestDefaultSearchOptions(t *testing.T) {
	t.Run("Default search options", func(t *testing.T) {
		options := DefaultSearchOptions()

		if options.CaseSensitive != false {
			t.Errorf("CaseSensitive should be false, got %v", options.CaseSensitive)
		}
		if options.ExactMatch != false {
			t.Errorf("ExactMatch should be false, got %v", options.ExactMatch)
		}
	})
}

func TestCityData(t *testing.T) {
	t.Run("CityData struct", func(t *testing.T) {
		city := CityData{
			Lat:           41.82999066,
			Lng:           -87.75005497,
			Pop:           5915976,
			City:          "Chicago",
			ISO2:          "US",
			ISO3:          "USA",
			Country:       "United States of America",
			Timezone:      "America/Chicago",
			Province:      "Illinois",
			ExactCity:     "Chicago",
			CityASCII:     "Chicago",
			StateANSI:     "IL",
			ExactProvince: "IL",
		}

		if city.City != "Chicago" {
			t.Errorf("Expected Chicago, got %s", city.City)
		}
		if city.ISO2 != "US" {
			t.Errorf("Expected US, got %s", city.ISO2)
		}
		if city.Timezone != "America/Chicago" {
			t.Errorf("Expected America/Chicago, got %s", city.Timezone)
		}
	})
}
