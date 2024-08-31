package stringsdsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    []byte
		expected []byte
	}{
		{input: []byte{'h', 'e', 'l', 'l', 'o'}, expected: []byte{'o', 'l', 'l', 'e', 'h'}},          // Example 1
		{input: []byte{'H', 'a', 'n', 'n', 'a', 'h'}, expected: []byte{'h', 'a', 'n', 'n', 'a', 'H'}}, // Example 2
		{input: []byte{'a', 'b', 'c', 'd'}, expected: []byte{'d', 'c', 'b', 'a'}},                   // Even number of characters
		{input: []byte{'A'}, expected: []byte{'A'}},                                                // Single character
		{input: []byte{'a', 'b', 'c', 'b', 'a'}, expected: []byte{'a', 'b', 'c', 'b', 'a'}},        // Palindrome
	}

	for _, test := range tests {
		// Making a copy of the input to avoid modifying the original slice in case of test failure
		inputCopy := make([]byte, len(test.input))
		copy(inputCopy, test.input)

		ReverseString(inputCopy)

		// Using assert.Equal to compare the reversed string with the expected output
		assert.Equal(t, test.expected, inputCopy, "The reversed string does not match the expected output for input %v", test.input)
	}
}