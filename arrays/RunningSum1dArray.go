package arrays

// RunningSum1dArray calculates the running sum of a 1-dimensional array.
// The running sum at each index i is the sum of all elements from index 0 to i.
//
// Time Complexity: O(n), where n is the number of elements in the input array.
// This is because the function iterates through the array once.
//
// Space Complexity: O(1), as the function modifies the input array in place and does not use any additional space
// proportional to the input size.
func RunningSum1dArray(nums []int) []int {
	var runningSum int = 0

	// Iterate through the array to calculate the running sum
	for i, num := range nums {
		runningSum += num
		nums[i] = runningSum
	}

	return nums
}
