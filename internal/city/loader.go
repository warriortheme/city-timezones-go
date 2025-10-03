package city

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	cityData  []CityData
	loadOnce  sync.Once
	loadError error
)

// LoadCityData loads the city data from the JSON file
func LoadCityData() ([]CityData, error) {
	loadOnce.Do(func() {
		cityData, loadError = loadCityDataFromFile()
	})
	return cityData, loadError
}

// loadCityDataFromFile loads city data from the data/cityMap.json file
func loadCityDataFromFile() ([]CityData, error) {
	// Get the path to the data file relative to this source file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("failed to get current file path")
	}

	// Navigate to the project root and find the data file
	projectRoot := filepath.Join(filepath.Dir(filename), "..", "..")
	dataPath := filepath.Join(projectRoot, "data", "cityMap.json")

	// Verify the file exists and is readable
	if _, err := os.Stat(dataPath); err != nil {
		return nil, fmt.Errorf("city data file not found at %s: %w", dataPath, err)
	}

	data, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read city data file: %w", err)
	}

	cities, err := UnmarshalCityData(data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal city data: %w", err)
	}

	return cities, nil
}

// GetCityData returns the loaded city data
func GetCityData() ([]CityData, error) {
	return LoadCityData()
}
