package techniques

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxLength(t *testing.T) {
	// Test case 1
	nums := []int{3, 1, 2, 7, 4, 2, 1, 1, 5}
	k := 8
	expected := 4
	result := FindMaxLength(nums, k)
	assert.Equal(t, expected, result, "Test case 1 failed")

	// Test case 2: Longest subarray with sum <= 7 is [2, 1, 5] with length 3
	nums = []int{2, 1, 5, 1, 3, 2}
	k = 7
	expected = 3
	result = FindMaxLength(nums, k)
	assert.Equal(t, expected, result, "Test case 2 failed")

	// Test case 3: Longest subarray with sum <= 15 is [2, 1, 5, 1, 3, 2] with length 6
	nums = []int{2, 1, 5, 1, 3, 2}
	k = 15
	expected = 6
	result = FindMaxLength(nums, k)
	assert.Equal(t, expected, result, "Test case 3 failed")
}

func TestFindLength(t *testing.T) {
	// Test case 1
	s := "1101100111"
	expected := 5
	result := FindLength(s)
	assert.Equal(t, expected, result, "Test case 1 failed")

	// Test case 2
	s = "11111"
	expected = 5
	result = FindLength(s)
	assert.Equal(t, expected, result, "Test case 2 failed")

	// Test case 3
	s = "00000"
	expected = 1
	result = FindLength(s)
	assert.Equal(t, expected, result, "Test case 3 failed")
}

func TestNumSubarrayProductLessThanK(t *testing.T) {
	// Test case 1: The valid subarrays are [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
	nums := []int{10, 5, 2, 6}
	k := 100
	expected := 8
	result := NumSubarrayProductLessThanK(nums, k)
	assert.Equal(t, expected, result, "Test case 1 failed")

	// Test case 2: No valid subarray since k <= 1
	nums = []int{1, 2, 3}
	k = 0
	expected = 0
	result = NumSubarrayProductLessThanK(nums, k)
	assert.Equal(t, expected, result, "Test case 2 failed")

	// Test case 3: Valid subarrays are [1], [2], [3], [1, 2]
	nums = []int{1, 2, 3}
	k = 4
	expected = 4
	result = NumSubarrayProductLessThanK(nums, k)
	assert.Equal(t, expected, result, "Test case 3 failed")

	// Test case 4: Valid subarrays are [1], [2], [3], [1, 2], [2, 3], [1, 2, 3]
	nums = []int{1, 2, 3}
	k = 10
	expected = 6
	result = NumSubarrayProductLessThanK(nums, k)
	assert.Equal(t, expected, result, "Test case 4 failed")
}

func TestFindBestSubarray(t *testing.T) {
	// Test case 1
	nums := []int{2, 1, 5, 1, 3, 2}
	k := 3
	expected := 9
	result := FindBestSubarray(nums, k)
	assert.Equal(t, expected, result, "Test case 1 failed")

	// Test case 2
	nums = []int{2, 3, 4, 1, 5}
	k = 2
	expected = 7
	result = FindBestSubarray(nums, k)
	assert.Equal(t, expected, result, "Test case 2 failed")

	// Test case 3
	nums = []int{1, 2, 3, 4, 5}
	k = 5
	expected = 15
	result = FindBestSubarray(nums, k)
	assert.Equal(t, expected, result, "Test case 3 failed")
}