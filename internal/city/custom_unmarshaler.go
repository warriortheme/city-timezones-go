package city

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// CityDataRaw represents the raw JSON structure for flexible unmarshaling
type CityDataRaw struct {
	Lat           float64     `json:"lat"`
	Lng           float64     `json:"lng"`
	Pop           interface{} `json:"pop"` // Can be int or float
	City          string      `json:"city"`
	ISO2          interface{} `json:"iso2"` // Can be string or number
	ISO3          interface{} `json:"iso3"` // Can be string or number
	Country       string      `json:"country"`
	Timezone      string      `json:"timezone"`
	Province      string      `json:"province"`
	ExactCity     string      `json:"exactCity"`
	CityASCII     string      `json:"city_ascii"`
	StateANSI     string      `json:"state_ansi"`
	ExactProvince string      `json:"exactProvince"`
}

// ToCityData converts the raw structure to the final CityData structure
func (raw *CityDataRaw) ToCityData() CityData {
	return CityData{
		Lat:           raw.Lat,
		Lng:           raw.Lng,
		Pop:           convertToFloat64(raw.Pop),
		City:          raw.City,
		ISO2:          convertToString(raw.ISO2),
		ISO3:          convertToString(raw.ISO3),
		Country:       raw.Country,
		Timezone:      raw.Timezone,
		Province:      raw.Province,
		ExactCity:     raw.ExactCity,
		CityASCII:     raw.CityASCII,
		StateANSI:     raw.StateANSI,
		ExactProvince: raw.ExactProvince,
	}
}

// convertToFloat64 converts various numeric types to float64
func convertToFloat64(value interface{}) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
		return 0
	default:
		return 0
	}
}

// convertToString converts various types to string
func convertToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// UnmarshalCityData unmarshals JSON data into CityData slice with flexible type handling
func UnmarshalCityData(data []byte) ([]CityData, error) {
	var rawData []CityDataRaw
	if err := json.Unmarshal(data, &rawData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal raw city data: %w", err)
	}

	cities := make([]CityData, len(rawData))
	for i, raw := range rawData {
		cities[i] = raw.ToCityData()
	}

	return cities, nil
}
