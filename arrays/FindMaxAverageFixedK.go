package arrays

// FindMaxAverageFixedK finds the maximum average value of a contiguous subarray of length k.
// Technique: Sliding Window with fixed size
// Time Complexity: O(n) - We traverse the array once, adjusting the subarray sum for each new window of size k.
// Space Complexity: O(1) - Only a few variables are used, independent of the input size.
func FindMaxAverageFixedK(nums []int, k int) float64 {
	// If the array length is less than k, we cannot calculate the average, so return 0.
	if len(nums) < k {
		return 0
	}

	curr := 0
	// Build the initial sum of the first window of size k.
	for i := 0; i < k; i++ {
		curr += nums[i]
	}
	// Initialize the answer with the sum of the first window.
	ans := float64(curr)

	// Slide the window across the array, adjusting the sum by removing the leftmost element and adding the new one.
	for i := k; i < len(nums); i++ {
		curr += nums[i] - nums[i-k]
		// Update the maximum sum found.
		ans = maxFloat64(ans, float64(curr))
	}

	// Return the maximum average found.
	return ans / float64(k)
}

// maxFloat64 returns the greater value between two float64 numbers.
func maxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
