package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxAverageFixedK(t *testing.T) {
	// Test case 1: Maximum average subarray of length 4 is [12, -5, -6, 50] with average 12.75
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	expected := 12.75000
	result := FindMaxAverageFixedK(nums, k)
	assert.InDelta(t, expected, result, 1e-5, "Test case 1 failed")

	// Test case 2: Only one element, so the average is the element itself
	nums = []int{5}
	k = 1
	expected = 5.00000
	result = FindMaxAverageFixedK(nums, k)
	assert.InDelta(t, expected, result, 1e-5, "Test case 2 failed")

	// Test case 3: Maximum average subarray of length 2 is [50, 3] with average 26.5
	nums = []int{1, 12, -5, -6, 50, 3}
	k = 2
	expected = 26.50000
	result = FindMaxAverageFixedK(nums, k)
	assert.InDelta(t, expected, result, 1e-5, "Test case 3 failed")

	// Test case 4: All elements are negative, check for correct handling of negative numbers
	nums = []int{-1, -12, -5, -6, -50, -3}
	k = 3
	expected = -6.00000
	result = FindMaxAverageFixedK(nums, k)
	assert.InDelta(t, expected, result, 1e-5, "Test case 4 failed")

	// Test case 5: Array contains both positive and negative numbers, with the maximum average being positive
	nums = []int{0, 4, -1, -2, -3, 5, 1, 2, -1, 4}
	k = 3
	expected = 2.66667
	result = FindMaxAverageFixedK(nums, k)
	assert.InDelta(t, expected, result, 1e-5, "Test case 5 failed")

	// Test case 6: Array with large negative and positive values to test precision
	nums = []int{-10000, 10000, 5000, -5000, 10000, -10000}
	k = 2
	expected = 7500.00000
	result = FindMaxAverageFixedK(nums, k)
	assert.InDelta(t, expected, result, 1e-5, "Test case 6 failed")
}
