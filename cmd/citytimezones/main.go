package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/richoandika/city-timezones-go/pkg/citytimezones"
)

func main() {
	var (
		cityName     = flag.String("city", "", "Search by city name")
		searchString = flag.String("search", "", "Search by city, state, province, or country")
		isoCode      = flag.String("iso", "", "Search by ISO2 or ISO3 country code")
		timezone     = flag.String("timezone", "", "Filter by timezone")
		country      = flag.String("country", "", "Filter by country")
		output       = flag.String("output", "table", "Output format: table, json")
		limit        = flag.Int("limit", 10, "Limit number of results")
		help         = flag.Bool("help", false, "Show help")
	)

	flag.Parse()

	if *help {
		showHelp()
		return
	}

	// Load all cities
	allCities, err := citytimezones.GetCityMapping()
	if err != nil {
		log.Fatal("Failed to load city data:", err)
	}

	var results []citytimezones.CityData

	// Perform search based on flags
	if *cityName != "" {
		results, err = citytimezones.LookupViaCity(*cityName)
	} else if *searchString != "" {
		results, err = citytimezones.FindFromCityStateProvince(*searchString)
	} else if *isoCode != "" {
		results, err = citytimezones.FindFromIsoCode(*isoCode)
	} else {
		// No search criteria provided, show all cities
		results = allCities
	}

	if err != nil {
		log.Fatal("Search failed:", err)
	}

	// Apply additional filters
	if *timezone != "" {
		results = filterByTimezone(results, *timezone)
	}
	if *country != "" {
		results = filterByCountry(results, *country)
	}

	// Limit results
	if *limit > 0 && len(results) > *limit {
		results = results[:*limit]
	}

	// Output results
	if *output == "json" {
		outputJSON(results)
	} else {
		outputTable(results)
	}
}

func showHelp() {
	fmt.Println("City Timezones Go - Command Line Tool")
	fmt.Println("====================================")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  citytimezones [options]")
	fmt.Println()
	fmt.Println("Search Options (use one):")
	fmt.Println("  -city string")
	fmt.Println("        Search by city name")
	fmt.Println("  -search string")
	fmt.Println("        Search by city, state, province, or country")
	fmt.Println("  -iso string")
	fmt.Println("        Search by ISO2 or ISO3 country code")
	fmt.Println()
	fmt.Println("Filter Options:")
	fmt.Println("  -timezone string")
	fmt.Println("        Filter by timezone")
	fmt.Println("  -country string")
	fmt.Println("        Filter by country")
	fmt.Println()
	fmt.Println("Output Options:")
	fmt.Println("  -output string")
	fmt.Println("        Output format: table, json (default: table)")
	fmt.Println("  -limit int")
	fmt.Println("        Limit number of results (default: 10)")
	fmt.Println("  -help")
	fmt.Println("        Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  citytimezones -city Chicago")
	fmt.Println("  citytimezones -search 'springfield mo'")
	fmt.Println("  citytimezones -iso DE -limit 5")
	fmt.Println("  citytimezones -timezone 'America/New_York' -output json")
}

func filterByTimezone(cities []citytimezones.CityData, timezone string) []citytimezones.CityData {
	var filtered []citytimezones.CityData
	for _, city := range cities {
		if strings.Contains(strings.ToLower(city.Timezone), strings.ToLower(timezone)) {
			filtered = append(filtered, city)
		}
	}
	return filtered
}

func filterByCountry(cities []citytimezones.CityData, country string) []citytimezones.CityData {
	var filtered []citytimezones.CityData
	for _, city := range cities {
		if strings.Contains(strings.ToLower(city.Country), strings.ToLower(country)) {
			filtered = append(filtered, city)
		}
	}
	return filtered
}

func outputJSON(cities []citytimezones.CityData) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(cities); err != nil {
		log.Fatal("Failed to encode JSON:", err)
	}
}

func outputTable(cities []citytimezones.CityData) {
	if len(cities) == 0 {
		fmt.Println("No cities found.")
		return
	}

	fmt.Printf("Found %d cities:\n\n", len(cities))
	fmt.Printf("%-20s %-15s %-20s %-15s %-20s %-10s %-10s\n",
		"City", "Province", "Country", "Timezone", "ISO2/ISO3", "Lat", "Lng")
	fmt.Println(strings.Repeat("-", 120))

	for _, city := range cities {
		fmt.Printf("%-20s %-15s %-20s %-15s %-10s %-10.4f %-10.4f\n",
			truncateString(city.City, 20),
			truncateString(city.Province, 15),
			truncateString(city.Country, 20),
			truncateString(city.Timezone, 15),
			city.ISO2+"/"+city.ISO3,
			city.Lat,
			city.Lng)
	}
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
