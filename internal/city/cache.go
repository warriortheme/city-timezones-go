package city

import (
	"sync"
)

// SearchCache provides thread-safe caching for search results
type SearchCache struct {
	mu    sync.RWMutex
	cache map[string][]CityData
}

// NewSearchCache creates a new search cache
func NewSearchCache() *SearchCache {
	return &SearchCache{
		cache: make(map[string][]CityData),
	}
}

// Get retrieves a cached result
func (c *SearchCache) Get(key string) ([]CityData, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	result, exists := c.cache[key]
	return result, exists
}

// Set stores a result in the cache
func (c *SearchCache) Set(key string, result []CityData) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = result
}

// Clear clears the cache
func (c *SearchCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache = make(map[string][]CityData)
}

// Size returns the number of cached entries
func (c *SearchCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.cache)
}

// Global cache instance
var searchCache = NewSearchCache()

// GetCachedResult retrieves a cached search result
func GetCachedResult(key string) ([]CityData, bool) {
	return searchCache.Get(key)
}

// SetCachedResult stores a search result in cache
func SetCachedResult(key string, result []CityData) {
	searchCache.Set(key, result)
}

// ClearCache clears the global search cache
func ClearCache() {
	searchCache.Clear()
}

// CacheSize returns the size of the global cache
func CacheSize() int {
	return searchCache.Size()
}
