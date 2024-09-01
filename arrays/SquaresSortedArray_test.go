package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			input:    []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
		{
			input:    []int{-5, -3, -2, -1},
			expected: []int{1, 4, 9, 25},
		},
		{
			input:    []int{0, 1, 2, 3, 4},
			expected: []int{0, 1, 4, 9, 16},
		},
		{
			input:    []int{-2, -1},
			expected: []int{1, 4},
		},
	}

	for _, test := range tests {
		output := SortedSquares(test.input)
		assert.Equal(t, test.expected, output, "SortedSquares(%v) = %v; want %v", test.input, output, test.expected)
	}
}
