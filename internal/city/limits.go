package city

import (
	"sync"
	"time"
)

// RateLimiter provides rate limiting functionality
type RateLimiter struct {
	mu       sync.Mutex
	requests map[string][]time.Time
	limit    int
	window   time.Duration
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

// Allow checks if a request is allowed for the given key
func (rl *RateLimiter) Allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-rl.window)

	// Clean old requests
	if requests, exists := rl.requests[key]; exists {
		var validRequests []time.Time
		for _, req := range requests {
			if req.After(cutoff) {
				validRequests = append(validRequests, req)
			}
		}
		rl.requests[key] = validRequests
	}

	// Check if we're under the limit
	if len(rl.requests[key]) < rl.limit {
		rl.requests[key] = append(rl.requests[key], now)
		return true
	}

	return false
}

// Global rate limiter (100 requests per minute per IP)
var globalRateLimiter = NewRateLimiter(100, time.Minute)

// CheckRateLimit checks if a request is allowed
func CheckRateLimit(key string) bool {
	return globalRateLimiter.Allow(key)
}

// ResourceManager manages resource usage
type ResourceManager struct {
	mu           sync.RWMutex
	maxMemoryMB  int
	currentMB    int
	activeSearch int
	maxSearches  int
}

// NewResourceManager creates a new resource manager
func NewResourceManager(maxMemoryMB, maxSearches int) *ResourceManager {
	return &ResourceManager{
		maxMemoryMB: maxMemoryMB,
		maxSearches: maxSearches,
	}
}

// CanAllocate checks if resources can be allocated
func (rm *ResourceManager) CanAllocate() bool {
	rm.mu.RLock()
	defer rm.mu.RUnlock()

	return rm.currentMB < rm.maxMemoryMB && rm.activeSearch < rm.maxSearches
}

// Allocate allocates resources
func (rm *ResourceManager) Allocate(memoryMB int) bool {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if rm.currentMB+memoryMB > rm.maxMemoryMB || rm.activeSearch >= rm.maxSearches {
		return false
	}

	rm.currentMB += memoryMB
	rm.activeSearch++
	return true
}

// Release releases resources
func (rm *ResourceManager) Release(memoryMB int) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	if rm.currentMB >= memoryMB {
		rm.currentMB -= memoryMB
	}
	if rm.activeSearch > 0 {
		rm.activeSearch--
	}
}

// Global resource manager
var globalResourceManager = NewResourceManager(500, 50) // 500MB, 50 concurrent searches

// CheckResourceAvailability checks if resources are available
func CheckResourceAvailability() bool {
	return globalResourceManager.CanAllocate()
}

// AllocateResources allocates resources
func AllocateResources(memoryMB int) bool {
	return globalResourceManager.Allocate(memoryMB)
}

// ReleaseResources releases resources
func ReleaseResources(memoryMB int) {
	globalResourceManager.Release(memoryMB)
}
