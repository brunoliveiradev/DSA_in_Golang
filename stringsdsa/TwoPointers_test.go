package stringsdsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoPointers(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		// Basic Palindrome Cases
		{"racecar", true},
		{"madam", true},
		{"level", true},

		// Non-Palindrome Cases
		{"hello", false},
		{"world", false},
		{"golang", false},

		// Edge Cases
		{"", true},    // An empty string is a palindrome
		{"a", true},   // A single character string is a palindrome
		{"ab", false}, // Two different characters are not a palindrome
		{"aa", true},  // Two identical characters are a palindrome

		// Special Cases
		{"A man a plan a canal Panama", false}, // Case-sensitive and with spaces
		{"12321", true},                        // Numeric palindrome
		{"12345", false},                       // Non-palindromic numeric string
		{"!@##@!", true},                       // Palindrome with special characters
		{"!@#a#@!", true},                      // Palindrome with mixed special characters and letters
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := TwoPointers(tt.input)
			assert.Equal(t, tt.expected, result, "TwoPointers(%s) = %v; want %v", tt.input, result, tt.expected)
		})
	}
}
