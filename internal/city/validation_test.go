package city

import (
	"testing"
)

func TestValidateSearchInput(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		maxLength int
		want      string
		wantErr   bool
	}{
		{
			name:      "Valid input",
			input:     "Chicago",
			maxLength: 100,
			want:      "Chicago",
			wantErr:   false,
		},
		{
			name:      "Empty input",
			input:     "",
			maxLength: 100,
			want:      "",
			wantErr:   false,
		},
		{
			name:      "Whitespace input",
			input:     "   Chicago   ",
			maxLength: 100,
			want:      "Chicago",
			wantErr:   false,
		},
		{
			name:      "Input too long",
			input:     "This is a very long input that exceeds the maximum length allowed",
			maxLength: 10,
			want:      "",
			wantErr:   true,
		},
		{
			name:      "Path traversal attempt",
			input:     "../etc/passwd",
			maxLength: 100,
			want:      "",
			wantErr:   true,
		},
		{
			name:      "XSS attempt",
			input:     "<script>alert('xss')</script>",
			maxLength: 100,
			want:      "",
			wantErr:   true,
		},
		{
			name:      "JavaScript protocol",
			input:     "javascript:alert('xss')",
			maxLength: 100,
			want:      "",
			wantErr:   true,
		},
		{
			name:      "Data protocol",
			input:     "data:text/html,<script>alert('xss')</script>",
			maxLength: 100,
			want:      "",
			wantErr:   true,
		},
		{
			name:      "Code injection - eval",
			input:     "eval(malicious_code)",
			maxLength: 100,
			want:      "",
			wantErr:   true,
		},
		{
			name:      "Code injection - exec",
			input:     "exec('rm -rf /')",
			maxLength: 100,
			want:      "",
			wantErr:   true,
		},
		{
			name:      "Code injection - system",
			input:     "system('rm -rf /')",
			maxLength: 100,
			want:      "",
			wantErr:   true,
		},
		{
			name:      "Valid special characters",
			input:     "São Paulo",
			maxLength: 100,
			want:      "São Paulo",
			wantErr:   false,
		},
		{
			name:      "Valid numbers and symbols",
			input:     "City-123",
			maxLength: 100,
			want:      "City-123",
			wantErr:   false,
		},
		{
			name:      "Input exactly at max length",
			input:     "Chicago",
			maxLength: 7,
			want:      "Chicago",
			wantErr:   false,
		},
		{
			name:      "Input one character over max length",
			input:     "Chicago",
			maxLength: 6,
			want:      "",
			wantErr:   true,
		},
		{
			name:      "Input with only whitespace",
			input:     "   ",
			maxLength: 100,
			want:      "",
			wantErr:   false,
		},
		{
			name:      "Input with tabs and newlines",
			input:     "\t\n\r",
			maxLength: 100,
			want:      "",
			wantErr:   false,
		},
		{
			name:      "Input with mixed whitespace",
			input:     "  \t  \n  ",
			maxLength: 100,
			want:      "",
			wantErr:   false,
		},
		{
			name:      "Input with unicode characters",
			input:     "São Paulo",
			maxLength: 100,
			want:      "São Paulo",
			wantErr:   false,
		},
		{
			name:      "Input with emoji",
			input:     "City 🏙️",
			maxLength: 100,
			want:      "City 🏙️",
			wantErr:   false,
		},
		{
			name:      "Input with backslashes",
			input:     "City\\Name",
			maxLength: 100,
			want:      "City\\Name",
			wantErr:   false,
		},
		{
			name:      "Input with quotes",
			input:     "City\"Name\"",
			maxLength: 100,
			want:      "City\"Name\"",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateSearchInput(tt.input, tt.maxLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSearchInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateSearchInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsSuspiciousPatterns(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Clean input",
			input: "Chicago",
			want:  false,
		},
		{
			name:  "Path traversal",
			input: "../etc/passwd",
			want:  true,
		},
		{
			name:  "XSS script tag",
			input: "<script>alert('xss')</script>",
			want:  true,
		},
		{
			name:  "JavaScript protocol",
			input: "javascript:alert('xss')",
			want:  true,
		},
		{
			name:  "Data protocol",
			input: "data:text/html,<script>alert('xss')</script>",
			want:  true,
		},
		{
			name:  "Code injection - eval",
			input: "eval(malicious_code)",
			want:  true,
		},
		{
			name:  "Code injection - exec",
			input: "exec('rm -rf /')",
			want:  true,
		},
		{
			name:  "Code injection - system",
			input: "system('rm -rf /')",
			want:  true,
		},
		{
			name:  "Case insensitive detection",
			input: "JAVASCRIPT:alert('xss')",
			want:  true,
		},
		{
			name:  "Mixed case detection",
			input: "JaVaScRiPt:alert('xss')",
			want:  true,
		},
		{
			name:  "Valid input with similar patterns",
			input: "script.js file",
			want:  false,
		},
		{
			name:  "Valid input with data",
			input: "data analysis",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsSuspiciousPatterns(tt.input); got != tt.want {
				t.Errorf("containsSuspiciousPatterns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateISOCode(t *testing.T) {
	tests := []struct {
		name    string
		isoCode string
		want    string
		wantErr bool
	}{
		{
			name:    "Valid ISO2 code",
			isoCode: "US",
			want:    "US",
			wantErr: false,
		},
		{
			name:    "Valid ISO2 code lowercase",
			isoCode: "us",
			want:    "US",
			wantErr: false,
		},
		{
			name:    "Valid ISO3 code",
			isoCode: "USA",
			want:    "USA",
			wantErr: false,
		},
		{
			name:    "Valid ISO3 code lowercase",
			isoCode: "usa",
			want:    "USA",
			wantErr: false,
		},
		{
			name:    "Empty ISO code",
			isoCode: "",
			want:    "",
			wantErr: false,
		},
		{
			name:    "Invalid length - 1 character",
			isoCode: "U",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid length - 4 characters",
			isoCode: "USAA",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid characters - numbers",
			isoCode: "U1",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid characters - symbols",
			isoCode: "U@",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Whitespace",
			isoCode: "  US  ",
			want:    "US",
			wantErr: false,
		},
		{
			name:    "Valid ISO2 with mixed case",
			isoCode: "uS",
			want:    "US",
			wantErr: false,
		},
		{
			name:    "Valid ISO3 with mixed case",
			isoCode: "uSa",
			want:    "USA",
			wantErr: false,
		},
		{
			name:    "Invalid characters - special symbols",
			isoCode: "U@",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid characters - spaces in middle",
			isoCode: "U S",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid characters - numbers",
			isoCode: "U1",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateISOCode(tt.isoCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateISOCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateISOCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidISO2Code(t *testing.T) {
	tests := []struct {
		name string
		code string
		want bool
	}{
		{
			name: "Valid ISO2",
			code: "US",
			want: true,
		},
		{
			name: "Invalid length",
			code: "U",
			want: false,
		},
		{
			name: "Invalid length long",
			code: "USA",
			want: false,
		},
		{
			name: "Invalid characters",
			code: "U1",
			want: false,
		},
		{
			name: "Lowercase",
			code: "us",
			want: false,
		},
		{
			name: "Empty string",
			code: "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidISO2Code(tt.code); got != tt.want {
				t.Errorf("isValidISO2Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidISO3Code(t *testing.T) {
	tests := []struct {
		name string
		code string
		want bool
	}{
		{
			name: "Valid ISO3",
			code: "USA",
			want: true,
		},
		{
			name: "Invalid length short",
			code: "US",
			want: false,
		},
		{
			name: "Invalid length long",
			code: "USAA",
			want: false,
		},
		{
			name: "Invalid characters",
			code: "US1",
			want: false,
		},
		{
			name: "Lowercase",
			code: "usa",
			want: false,
		},
		{
			name: "Empty string",
			code: "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidISO3Code(tt.code); got != tt.want {
				t.Errorf("isValidISO3Code() = %v, want %v", got, tt.want)
			}
		})
	}
}
