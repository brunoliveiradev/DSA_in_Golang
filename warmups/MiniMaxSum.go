package warmups

import (
	"fmt"
)

// miniMaxSum takes an array of 5 integers as input and calculates the sum of the four smallest and four largest integers in the array.
// Then print the respective minimum and maximum values as a single line of two space-separated long integers.
// Time Complexity: O(n)
// Space Complexity: O(1)
func miniMaxSum(arr []int32) {
	// Write your code here
	var totalSum int64 = 0
	var minVal int32 = arr[0]
	var maxVal int32 = arr[0]

	// Find total sum, minimum and maximum values in one pass
	for _, value := range arr {
		totalSum += int64(value)
		if value < minVal {
			minVal = value
		}
		if value > maxVal {
			maxVal = value
		}
	}

	// Minimum sum = total - maximum element
	// Maximum sum = total - minimum element
	minSum := totalSum - int64(maxVal)
	maxSum := totalSum - int64(minVal)

	fmt.Print(minSum, " ", maxSum)
}
