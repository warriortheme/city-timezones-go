package city

import (
	"testing"
)

func TestLoadCityData(t *testing.T) {
	t.Run("Load city data", func(t *testing.T) {
		cities, err := LoadCityData()
		if err != nil {
			t.Errorf("Should load data without error: %v", err)
		}
		if cities == nil {
			t.Error("Cities should not be nil")
		}
		if len(cities) == 0 {
			t.Error("Should have cities")
		}
	})

	t.Run("Multiple calls return same data", func(t *testing.T) {
		cities1, err1 := LoadCityData()
		if err1 != nil {
			t.Errorf("First call should not error: %v", err1)
		}

		cities2, err2 := LoadCityData()
		if err2 != nil {
			t.Errorf("Second call should not error: %v", err2)
		}

		if len(cities1) != len(cities2) {
			t.Errorf("Both calls should return same number of cities: %d vs %d", len(cities1), len(cities2))
		}
	})
}

func TestGetCityData(t *testing.T) {
	t.Run("Get city data", func(t *testing.T) {
		cities, err := GetCityData()
		if err != nil {
			t.Errorf("Should get data without error: %v", err)
		}
		if cities == nil {
			t.Error("Cities should not be nil")
		}
		if len(cities) == 0 {
			t.Error("Should have cities")
		}
	})
}

func TestLoadCityDataFromFile(t *testing.T) {
	t.Run("Load from non-existent file", func(t *testing.T) {
		// This test would require mocking the file system
		// For now, we'll test the happy path which is already covered
		t.Skip("File system mocking not implemented")
	})
}
