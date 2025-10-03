package city

import (
	"testing"
)

func TestUnmarshalCityData(t *testing.T) {
	t.Run("Valid JSON data", func(t *testing.T) {
		jsonData := `[
			{
				"lat": 41.82999066,
				"lng": -87.75005497,
				"pop": 5915976,
				"city": "Chicago",
				"iso2": "US",
				"iso3": "USA",
				"country": "United States of America",
				"timezone": "America/Chicago",
				"province": "Illinois",
				"exactCity": "Chicago",
				"city_ascii": "Chicago",
				"state_ansi": "IL",
				"exactProvince": "IL"
			}
		]`

		cities, err := UnmarshalCityData([]byte(jsonData))
		if err != nil {
			t.Errorf("Should not have error: %v", err)
		}
		if len(cities) != 1 {
			t.Errorf("Should have 1 city, got %d", len(cities))
		}
		if cities[0].City != "Chicago" {
			t.Errorf("Expected Chicago, got %s", cities[0].City)
		}
	})

	t.Run("Empty array", func(t *testing.T) {
		jsonData := `[]`

		cities, err := UnmarshalCityData([]byte(jsonData))
		if err != nil {
			t.Errorf("Should not have error: %v", err)
		}
		if len(cities) != 0 {
			t.Errorf("Should have 0 cities, got %d", len(cities))
		}
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		jsonData := `invalid json`

		_, err := UnmarshalCityData([]byte(jsonData))
		if err == nil {
			t.Error("Should have error for invalid JSON")
		}
	})
}

func TestToCityData(t *testing.T) {
	t.Run("Convert raw data to CityData", func(t *testing.T) {
		raw := CityDataRaw{
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

		city := raw.ToCityData()
		if city.City != "Chicago" {
			t.Errorf("Expected Chicago, got %s", city.City)
		}
		if city.ISO2 != "US" {
			t.Errorf("Expected US, got %s", city.ISO2)
		}
	})
}

func TestConvertToFloat64(t *testing.T) {
	t.Run("Convert int to float64", func(t *testing.T) {
		result := convertToFloat64(123)
		if result != 123.0 {
			t.Errorf("Expected 123.0, got %f", result)
		}
	})

	t.Run("Convert float64 to float64", func(t *testing.T) {
		result := convertToFloat64(123.45)
		if result != 123.45 {
			t.Errorf("Expected 123.45, got %f", result)
		}
	})

	t.Run("Convert string to float64", func(t *testing.T) {
		result := convertToFloat64("123.45")
		if result != 123.45 {
			t.Errorf("Expected 123.45, got %f", result)
		}
	})

	t.Run("Convert invalid string to float64", func(t *testing.T) {
		result := convertToFloat64("invalid")
		if result != 0.0 {
			t.Errorf("Expected 0.0, got %f", result)
		}
	})

	t.Run("Convert nil to float64", func(t *testing.T) {
		result := convertToFloat64(nil)
		if result != 0.0 {
			t.Errorf("Expected 0.0, got %f", result)
		}
	})

	t.Run("Convert float32 to float64", func(t *testing.T) {
		result := convertToFloat64(float32(123.45))
		// Allow for floating point precision differences
		if result < 123.44 || result > 123.46 {
			t.Errorf("Expected ~123.45, got %f", result)
		}
	})

	t.Run("Convert int64 to float64", func(t *testing.T) {
		result := convertToFloat64(int64(123))
		if result != 123.0 {
			t.Errorf("Expected 123.0, got %f", result)
		}
	})

	t.Run("Convert bool to float64", func(t *testing.T) {
		result := convertToFloat64(true)
		if result != 0.0 {
			t.Errorf("Expected 0.0, got %f", result)
		}
	})

	t.Run("Convert false bool to float64", func(t *testing.T) {
		result := convertToFloat64(false)
		if result != 0.0 {
			t.Errorf("Expected 0.0, got %f", result)
		}
	})

	t.Run("Convert uint to float64", func(t *testing.T) {
		result := convertToFloat64(uint(123))
		// uint is not handled in the switch statement, so it returns 0.0
		if result != 0.0 {
			t.Errorf("Expected 0.0, got %f", result)
		}
	})

	t.Run("Convert int32 to float64", func(t *testing.T) {
		result := convertToFloat64(int32(123))
		if result != 123.0 {
			t.Errorf("Expected 123.0, got %f", result)
		}
	})

	t.Run("Convert string with decimal to float64", func(t *testing.T) {
		result := convertToFloat64("123.456")
		if result != 123.456 {
			t.Errorf("Expected 123.456, got %f", result)
		}
	})

	t.Run("Convert string with scientific notation to float64", func(t *testing.T) {
		result := convertToFloat64("1.23e2")
		if result != 123.0 {
			t.Errorf("Expected 123.0, got %f", result)
		}
	})

	t.Run("Convert string with negative number to float64", func(t *testing.T) {
		result := convertToFloat64("-123.45")
		if result != -123.45 {
			t.Errorf("Expected -123.45, got %f", result)
		}
	})

	t.Run("Convert string with zero to float64", func(t *testing.T) {
		result := convertToFloat64("0")
		if result != 0.0 {
			t.Errorf("Expected 0.0, got %f", result)
		}
	})

	t.Run("Convert string with empty string to float64", func(t *testing.T) {
		result := convertToFloat64("")
		if result != 0.0 {
			t.Errorf("Expected 0.0, got %f", result)
		}
	})

	t.Run("Convert string with non-numeric to float64", func(t *testing.T) {
		result := convertToFloat64("abc")
		if result != 0.0 {
			t.Errorf("Expected 0.0, got %f", result)
		}
	})
}

func TestConvertToString(t *testing.T) {
	t.Run("Convert string to string", func(t *testing.T) {
		result := convertToString("test")
		if result != "test" {
			t.Errorf("Expected 'test', got '%s'", result)
		}
	})

	t.Run("Convert int to string", func(t *testing.T) {
		result := convertToString(123)
		if result != "123" {
			t.Errorf("Expected '123', got '%s'", result)
		}
	})

	t.Run("Convert float64 to string", func(t *testing.T) {
		result := convertToString(123.45)
		if result != "123.45" {
			t.Errorf("Expected '123.45', got '%s'", result)
		}
	})

	t.Run("Convert nil to string", func(t *testing.T) {
		result := convertToString(nil)
		if result != "<nil>" {
			t.Errorf("Expected '<nil>', got '%s'", result)
		}
	})

	t.Run("Convert bool to string", func(t *testing.T) {
		result := convertToString(true)
		if result != "true" {
			t.Errorf("Expected 'true', got '%s'", result)
		}
	})

	t.Run("Convert int64 to string", func(t *testing.T) {
		result := convertToString(int64(123))
		if result != "123" {
			t.Errorf("Expected '123', got '%s'", result)
		}
	})

	t.Run("Convert float32 to string", func(t *testing.T) {
		result := convertToString(float32(123.45))
		if result != "123.45" {
			t.Errorf("Expected '123.45', got '%s'", result)
		}
	})

	t.Run("Convert false bool to string", func(t *testing.T) {
		result := convertToString(false)
		if result != "false" {
			t.Errorf("Expected 'false', got '%s'", result)
		}
	})

	t.Run("Convert uint to string", func(t *testing.T) {
		result := convertToString(uint(123))
		// uint uses the default case which formats it as "123"
		if result != "123" {
			t.Errorf("Expected '123', got '%s'", result)
		}
	})

	t.Run("Convert int32 to string", func(t *testing.T) {
		result := convertToString(int32(123))
		if result != "123" {
			t.Errorf("Expected '123', got '%s'", result)
		}
	})

	t.Run("Convert int64 to string", func(t *testing.T) {
		result := convertToString(int64(123))
		if result != "123" {
			t.Errorf("Expected '123', got '%s'", result)
		}
	})

	t.Run("Convert float64 with decimal to string", func(t *testing.T) {
		result := convertToString(123.456)
		if result != "123.456" {
			t.Errorf("Expected '123.456', got '%s'", result)
		}
	})

	t.Run("Convert float64 with scientific notation to string", func(t *testing.T) {
		result := convertToString(1.23e2)
		if result != "123" {
			t.Errorf("Expected '123', got '%s'", result)
		}
	})

	t.Run("Convert float64 with zero to string", func(t *testing.T) {
		result := convertToString(0.0)
		if result != "0" {
			t.Errorf("Expected '0', got '%s'", result)
		}
	})

	t.Run("Convert float64 with negative to string", func(t *testing.T) {
		result := convertToString(-123.45)
		if result != "-123.45" {
			t.Errorf("Expected '-123.45', got '%s'", result)
		}
	})

	t.Run("Convert complex number to string", func(t *testing.T) {
		result := convertToString(complex(1, 2))
		// Complex numbers use the default case
		if result != "(1+2i)" {
			t.Errorf("Expected '(1+2i)', got '%s'", result)
		}
	})
}
