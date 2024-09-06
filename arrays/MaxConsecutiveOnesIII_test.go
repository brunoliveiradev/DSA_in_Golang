package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 1, 0, 0}, 0, 3},
		{[]int{1, 1, 1, 1, 1}, 0, 5},
		{[]int{0, 0, 0, 0, 0}, 5, 5},
		{[]int{0, 0, 0, 1, 1, 1, 0, 0, 0}, 3, 6},
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
		{[]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, 5, 10},
		{[]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0}, 0, 1},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 0, 10},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 10, 10},
	}

	for _, test := range tests {
		result := LongestOnes(test.nums, test.k)
		assert.Equal(t, test.expected, result, "For nums %v and k %d", test.nums, test.k)
	}
}
