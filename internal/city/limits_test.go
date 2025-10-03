package city

import (
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	t.Run("NewRateLimiter", func(t *testing.T) {
		limiter := NewRateLimiter(10, time.Minute)
		if limiter == nil {
			t.Error("Rate limiter should not be nil")
		}
	})

	t.Run("Allow within limit", func(t *testing.T) {
		limiter := NewRateLimiter(5, time.Minute)

		// Should allow first 5 requests
		for i := 0; i < 5; i++ {
			if !limiter.Allow("test") {
				t.Errorf("Request %d should be allowed", i+1)
			}
		}
	})

	t.Run("Deny when limit exceeded", func(t *testing.T) {
		limiter := NewRateLimiter(2, time.Minute)

		// Allow first 2 requests
		if !limiter.Allow("test") {
			t.Error("First request should be allowed")
		}
		if !limiter.Allow("test") {
			t.Error("Second request should be allowed")
		}

		// Third request should be denied
		if limiter.Allow("test") {
			t.Error("Third request should be denied")
		}
	})

	t.Run("Different keys have separate limits", func(t *testing.T) {
		limiter := NewRateLimiter(1, time.Minute)

		// First key should be allowed
		if !limiter.Allow("key1") {
			t.Error("First key should be allowed")
		}

		// Second key should also be allowed (separate limit)
		if !limiter.Allow("key2") {
			t.Error("Second key should be allowed")
		}
	})
}

func TestResourceManager(t *testing.T) {
	t.Run("NewResourceManager", func(t *testing.T) {
		rm := NewResourceManager(1000, 100)
		if rm == nil {
			t.Error("Resource manager should not be nil")
		}
	})

	t.Run("CanAllocate within limits", func(t *testing.T) {
		rm := NewResourceManager(1000, 100)

		if !rm.CanAllocate() {
			t.Error("Should be able to allocate resources")
		}
	})

	t.Run("Allocate and release resources", func(t *testing.T) {
		rm := NewResourceManager(1000, 100)

		// Allocate resources
		if !rm.Allocate(500) {
			t.Error("Should be able to allocate 500MB")
		}

		// Release resources
		rm.Release(500)

		// Should be able to allocate again
		if !rm.Allocate(500) {
			t.Error("Should be able to allocate 500MB again after release")
		}
	})

	t.Run("Allocate when at limit", func(t *testing.T) {
		rm := NewResourceManager(1000, 1) // Only 1 search allowed

		// First allocation should succeed
		if !rm.Allocate(500) {
			t.Error("Should be able to allocate first time")
		}

		// Second allocation should fail (search limit reached)
		if rm.Allocate(200) {
			t.Error("Should not be able to allocate when search limit reached")
		}
	})

	t.Run("Allocate when memory limit reached", func(t *testing.T) {
		rm := NewResourceManager(500, 100) // Only 500MB memory

		// Allocate up to limit
		if !rm.Allocate(500) {
			t.Error("Should be able to allocate up to limit")
		}

		// Try to allocate more - should fail
		if rm.Allocate(100) {
			t.Error("Should not be able to allocate when memory limit reached")
		}
	})
}

func TestGlobalLimits(t *testing.T) {
	t.Run("CheckRateLimit", func(t *testing.T) {
		// This should not panic
		allowed := CheckRateLimit("test")
		_ = allowed // We can't predict the result, but it should be a boolean
	})

	t.Run("CheckResourceAvailability", func(t *testing.T) {
		// This should not panic
		available := CheckResourceAvailability()
		_ = available // We can't predict the result, but it should be a boolean
	})

	t.Run("AllocateResources", func(t *testing.T) {
		// This should not panic
		allocated := AllocateResources(100)
		_ = allocated // We can't predict the result, but it should be a boolean
	})

	t.Run("ReleaseResources", func(t *testing.T) {
		// This should not panic
		ReleaseResources(100)
	})
}
