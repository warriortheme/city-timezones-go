package city

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// ValidationError is defined in errors.go

// ValidateSearchInput validates and sanitizes search input
func ValidateSearchInput(input string, maxLength int) (string, error) {
	if input == "" {
		return "", nil
	}

	// Check for reasonable length
	if len(input) > maxLength {
		return "", ValidationError{
			Field:   "input",
			Message: fmt.Sprintf("input too long: %d characters (max: %d)", len(input), maxLength),
		}
	}

	// Check for valid UTF-8
	if !utf8.ValidString(input) {
		return "", ValidationError{
			Field:   "input",
			Message: "input contains invalid UTF-8 characters",
		}
	}

	// Trim and normalize whitespace
	normalized := strings.TrimSpace(input)

	// Check for suspicious patterns (basic security check)
	if containsSuspiciousPatterns(normalized) {
		return "", ValidationError{
			Field:   "input",
			Message: "input contains potentially suspicious patterns",
		}
	}

	return normalized, nil
}

// containsSuspiciousPatterns checks for potentially malicious input
func containsSuspiciousPatterns(input string) bool {
	// Check for common injection patterns
	suspicious := []string{
		"../", "..\\", "..%2f", "..%5c", // Path traversal
		"<script", "</script>", // XSS
		"javascript:", "data:", // Protocol injection
		"eval(", "exec(", "system(", // Code injection
	}

	lowerInput := strings.ToLower(input)
	for _, pattern := range suspicious {
		if strings.Contains(lowerInput, pattern) {
			return true
		}
	}

	return false
}

// ValidateISOCode validates ISO country codes
func ValidateISOCode(isoCode string) (string, error) {
	if isoCode == "" {
		return "", nil
	}

	normalized := strings.TrimSpace(strings.ToUpper(isoCode))

	// ISO2 codes are exactly 2 characters
	if len(normalized) == 2 {
		if !isValidISO2Code(normalized) {
			return "", ValidationError{
				Field:   "isoCode",
				Message: "invalid ISO2 country code format",
			}
		}
		return normalized, nil
	}

	// ISO3 codes are exactly 3 characters
	if len(normalized) == 3 {
		if !isValidISO3Code(normalized) {
			return "", ValidationError{
				Field:   "isoCode",
				Message: "invalid ISO3 country code format",
			}
		}
		return normalized, nil
	}

	return "", ValidationError{
		Field:   "isoCode",
		Message: "ISO code must be 2 or 3 characters",
	}
}

// isValidISO2Code checks if the string is a valid ISO2 country code format
func isValidISO2Code(code string) bool {
	if len(code) != 2 {
		return false
	}

	// Check if all characters are letters
	for _, r := range code {
		if r < 'A' || r > 'Z' {
			return false
		}
	}

	return true
}

// isValidISO3Code checks if the string is a valid ISO3 country code format
func isValidISO3Code(code string) bool {
	if len(code) != 3 {
		return false
	}

	// Check if all characters are letters
	for _, r := range code {
		if r < 'A' || r > 'Z' {
			return false
		}
	}

	return true
}
