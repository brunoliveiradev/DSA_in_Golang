package arrays

import "math"

// MinSubArrayLen finds the minimal length of a contiguous subarray of which the sum is at least `target`.
// If there is no such subarray, it returns 0.
//
// Time Complexity: O(n), where n is the length of the input array `nums`.
// Space Complexity: O(1), as the space used does not depend on the input size.
//
// Whenever you see a problem that involves finding contiguous sub-array (size or the sub-array itself) always consider using sliding window.
func MinSubArrayLen(target int, nums []int) int {
	left, right, sumOfCurrentWindow := 0, 0, 0
	answer := math.MaxInt64

	for right = 0; right < len(nums); right++ {
		sumOfCurrentWindow = sumOfCurrentWindow + nums[right]

		for sumOfCurrentWindow >= target {
			answer = minInt(answer, right-left+1)

			sumOfCurrentWindow = sumOfCurrentWindow - nums[left]
			left++
		}
	}

	if answer == math.MaxInt64 {
		return 0
	}

	return answer
}

// minInt returns the minimum value between two int numbers.
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
