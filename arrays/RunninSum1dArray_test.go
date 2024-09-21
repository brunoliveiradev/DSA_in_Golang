package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunningSum1dArray(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4},
			expected: []int{1, 3, 6, 10},
		},
		{
			nums:     []int{1, 1, 1, 1, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			nums:     []int{3, 1, 2, 10, 1},
			expected: []int{3, 4, 6, 16, 17},
		},
		{
			nums:     []int{-1, -2, -3, -4},
			expected: []int{-1, -3, -6, -10},
		},
		{
			nums:     []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		result := RunningSum1dArray(tt.nums)
		assert.Equal(t, tt.expected, result, "RunningSum1dArray(%v)", tt.nums)
	}
}
