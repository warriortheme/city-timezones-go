package city

import (
	"fmt"
)

// Error types for better error handling and debugging

// DataLoadError represents an error loading city data
type DataLoadError struct {
	Operation string
	Err       error
}

func (e DataLoadError) Error() string {
	return fmt.Sprintf("failed to load city data during %s: %v", e.Operation, e.Err)
}

func (e DataLoadError) Unwrap() error {
	return e.Err
}

// SearchError represents an error during search operations
type SearchError struct {
	Query     string
	Operation string
	Err       error
}

func (e SearchError) Error() string {
	return fmt.Sprintf("search error for query '%s' during %s: %v", e.Query, e.Operation, e.Err)
}

func (e SearchError) Unwrap() error {
	return e.Err
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
	Value   interface{}
}

func (e ValidationError) Error() string {
	if e.Value != nil {
		return fmt.Sprintf("validation error for field '%s': %s (value: %v)", e.Field, e.Message, e.Value)
	}
	return fmt.Sprintf("validation error for field '%s': %s", e.Field, e.Message)
}

// CacheError represents a cache operation error
type CacheError struct {
	Operation string
	Key       string
	Err       error
}

func (e CacheError) Error() string {
	return fmt.Sprintf("cache error during %s for key '%s': %v", e.Operation, e.Key, e.Err)
}

func (e CacheError) Unwrap() error {
	return e.Err
}

// Helper functions for creating specific error types

// NewDataLoadError creates a new DataLoadError
func NewDataLoadError(operation string, err error) DataLoadError {
	return DataLoadError{
		Operation: operation,
		Err:       err,
	}
}

// NewSearchError creates a new SearchError
func NewSearchError(query, operation string, err error) SearchError {
	return SearchError{
		Query:     query,
		Operation: operation,
		Err:       err,
	}
}

// NewValidationError creates a new ValidationError
func NewValidationError(field, message string, value interface{}) ValidationError {
	return ValidationError{
		Field:   field,
		Message: message,
		Value:   value,
	}
}

// NewCacheError creates a new CacheError
func NewCacheError(operation, key string, err error) CacheError {
	return CacheError{
		Operation: operation,
		Key:       key,
		Err:       err,
	}
}
