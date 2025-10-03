package city

// CityData represents a city with its timezone and geographical information
type CityData struct {
	Lat           float64 `json:"lat"`
	Lng           float64 `json:"lng"`
	Pop           float64 `json:"pop"` // Changed to float64 to handle decimal values
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

// SearchOptions provides configuration for search operations
type SearchOptions struct {
	CaseSensitive bool
	ExactMatch    bool
}

// DefaultSearchOptions returns the default search configuration
func DefaultSearchOptions() SearchOptions {
	return SearchOptions{
		CaseSensitive: false,
		ExactMatch:    false,
	}
}
