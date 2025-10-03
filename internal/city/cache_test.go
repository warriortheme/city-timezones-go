package city

import (
	"testing"
)

func TestSearchCache(t *testing.T) {
	t.Run("NewSearchCache", func(t *testing.T) {
		cache := NewSearchCache()
		if cache == nil {
			t.Error("cache should not be nil")
		}
		if cache.Size() != 0 {
			t.Errorf("initial cache size should be 0, got %d", cache.Size())
		}
	})

	t.Run("Set and Get", func(t *testing.T) {
		cache := NewSearchCache()
		testData := []CityData{
			{City: "Chicago", ISO2: "US", Timezone: "America/Chicago"},
		}

		cache.Set("chicago", testData)
		if cache.Size() != 1 {
			t.Errorf("cache size should be 1, got %d", cache.Size())
		}

		result, exists := cache.Get("chicago")
		if !exists {
			t.Error("key should exist")
		}
		if len(result) != 1 {
			t.Errorf("result length should be 1, got %d", len(result))
		}
		if result[0].City != "Chicago" {
			t.Errorf("expected Chicago, got %s", result[0].City)
		}
	})

	t.Run("Clear", func(t *testing.T) {
		cache := NewSearchCache()
		testData := []CityData{{City: "Chicago"}}

		cache.Set("chicago", testData)
		cache.Clear()

		if cache.Size() != 0 {
			t.Errorf("cache should be empty, got %d", cache.Size())
		}
	})
}

func TestGlobalCache(t *testing.T) {
	t.Run("Global cache operations", func(t *testing.T) {
		// Clear cache first
		ClearCache()
		if CacheSize() != 0 {
			t.Errorf("cache should be empty, got %d", CacheSize())
		}

		// Test data
		testData := []CityData{
			{City: "Chicago", ISO2: "US", Timezone: "America/Chicago"},
		}

		// Set cached result
		SetCachedResult("chicago", testData)
		if CacheSize() != 1 {
			t.Errorf("cache should have 1 item, got %d", CacheSize())
		}

		// Get cached result
		result, exists := GetCachedResult("chicago")
		if !exists {
			t.Error("key should exist")
		}
		if len(result) != 1 {
			t.Errorf("result length should be 1, got %d", len(result))
		}
		if result[0].City != "Chicago" {
			t.Errorf("expected Chicago, got %s", result[0].City)
		}
	})
}
