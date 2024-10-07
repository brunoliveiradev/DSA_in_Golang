package arrays

// MinStartValue calculates the minimum positive start value needed to ensure that the
// cumulative sum of the start value and the elements in the array `nums` is always at least 1.
//
// Given an array of integers `nums`, the goal is to find a start value such that when we
// iterate over `nums`, adding each element to the cumulative sum, the sum is never less than 1.
//
// Space Complexity: O(1), as the function only uses a constant amount of extra space.
// Time Complexity: O(n), where n is the number of elements in the input array.
func MinStartValue(nums []int) int {
	// `sum` is used to store the cumulative sum as we iterate through `nums`.
	// `minValue` is used to track the minimum cumulative sum encountered during the iteration.
	sum := 0
	minValue := 0

	// Iterate over each element `num` in `nums`. We are using `range` to get the value directly.
	for _, num := range nums {
		// Add the current element `num` to `sum`.
		sum += num

		// Update `minValue` if the current cumulative sum is less than `minValue`.
		// This helps us keep track of the lowest point the cumulative sum reaches.
		if sum < minValue {
			minValue = sum
		}
	}

	// Calculate and return the minimum positive start value needed.
	// Since we need the cumulative sum to never be less than 1, we return `1 - minValue`.
	// If `minValue` is negative, this ensures `startValue` is large enough to offset any dips below 1.
	return 1 - minValue
}
