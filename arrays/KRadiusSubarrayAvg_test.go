package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAveragesForKRadiusSubarray(t *testing.T) {
	// Example 1: nums = [7,4,3,9,1,8,5,2,6], k = 3
	t.Run("TestExample1", func(t *testing.T) {
		nums := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
		k := 3
		expected := []int{-1, -1, -1, 5, 4, 4, -1, -1, -1}
		result := GetAveragesForKRadiusSubarray(nums, k)
		assert.Equal(t, expected, result)
	})

	// Example 2: nums = [100000], k = 0
	t.Run("TestExample2", func(t *testing.T) {
		nums := []int{100000}
		k := 0
		expected := []int{100000}
		result := GetAveragesForKRadiusSubarray(nums, k)
		assert.Equal(t, expected, result)
	})

	// Example 3: nums = [8], k = 100000
	t.Run("TestExample3", func(t *testing.T) {
		nums := []int{8}
		k := 100000
		expected := []int{-1}
		result := GetAveragesForKRadiusSubarray(nums, k)
		assert.Equal(t, expected, result)
	})

	// Additional test case 1: Large k where no averages can be calculated
	t.Run("TestLargeKNoAverages", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		k := 10
		expected := []int{-1, -1, -1, -1, -1}
		result := GetAveragesForKRadiusSubarray(nums, k)
		assert.Equal(t, expected, result)
	})

	// Additional test case 2: nums = [1, 2, 3, 4, 5], k = 1
	t.Run("TestSmallRadius", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		k := 1
		expected := []int{-1, 2, 3, 4, -1}
		result := GetAveragesForKRadiusSubarray(nums, k)
		assert.Equal(t, expected, result)
	})

	// Additional test case 3: Empty input
	t.Run("TestEmptyInput", func(t *testing.T) {
		nums := []int{}
		k := 1
		expected := []int{}
		result := GetAveragesForKRadiusSubarray(nums, k)
		assert.Equal(t, expected, result)
	})
}
