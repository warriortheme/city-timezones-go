package city

import (
	"errors"
	"testing"
)

func TestDataLoadError(t *testing.T) {
	t.Run("Error message", func(t *testing.T) {
		originalErr := errors.New("file not found")
		err := NewDataLoadError("load operation", originalErr)

		expectedMsg := "failed to load city data during load operation: file not found"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("Unwrap", func(t *testing.T) {
		originalErr := errors.New("file not found")
		err := NewDataLoadError("load operation", originalErr)

		unwrapped := err.Unwrap()
		if unwrapped != originalErr {
			t.Errorf("Expected unwrapped error to be original error")
		}
	})
}

func TestSearchError(t *testing.T) {
	t.Run("Error message", func(t *testing.T) {
		originalErr := errors.New("invalid query")
		err := NewSearchError("chicago", "search operation", originalErr)

		expectedMsg := "search error for query 'chicago' during search operation: invalid query"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("Unwrap", func(t *testing.T) {
		originalErr := errors.New("invalid query")
		err := NewSearchError("chicago", "search operation", originalErr)

		unwrapped := err.Unwrap()
		if unwrapped != originalErr {
			t.Errorf("Expected unwrapped error to be original error")
		}
	})
}

func TestValidationError(t *testing.T) {
	t.Run("Error message", func(t *testing.T) {
		err := NewValidationError("input", "contains invalid characters", "test@input")

		expectedMsg := "validation error for field 'input': contains invalid characters (value: test@input)"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("Error message with empty value", func(t *testing.T) {
		err := NewValidationError("field", "required field", "")
		expectedMsg := "validation error for field 'field': required field (value: )"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("Error message with nil value", func(t *testing.T) {
		err := NewValidationError("field", "invalid value", nil)
		expectedMsg := "validation error for field 'field': invalid value"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})
}

func TestCacheError(t *testing.T) {
	t.Run("Error message", func(t *testing.T) {
		originalErr := errors.New("cache miss")
		err := NewCacheError("get", "test_key", originalErr)

		expectedMsg := "cache error during get for key 'test_key': cache miss"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("Unwrap", func(t *testing.T) {
		originalErr := errors.New("cache miss")
		err := NewCacheError("get", "test_key", originalErr)

		unwrapped := err.Unwrap()
		if unwrapped != originalErr {
			t.Errorf("Expected unwrapped error to be original error")
		}
	})
}
