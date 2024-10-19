package arrays

// GetAveragesForKRadiusSubarray calculates the k-radius averages for a given array.
// A k-radius average is defined as the average of the elements from index i-k to i+k inclusive,
// where i is the center index. If there are fewer than k elements before or after the index i,
// the result for that index is -1.
//
// Time complexity: O(n), where n is the length of the array. We compute the prefix sum in O(n)
// and use it to calculate each k-radius average in constant time.
//
// Space complexity: O(n), since we store the prefix sum array and the result array.
func GetAveragesForKRadiusSubarray(nums []int, k int) []int {
	// return early, because the average of only one number is the number itself
	if k == 0 {
		return nums
	}

	// the windowSize will be 2 * k + 1 because it is a radius inclusive the index
	windowSize := 2*k + 1
	n := len(nums)
	avg := make([]int, n)

	// initialized the avg array with -1
	for i := range avg {
		avg[i] = -1
	}

	// return early, because if the windowSize is greater than k, we will not have a 'k' radius
	if windowSize > n {
		return avg
	}

	// prefix will be the sum of all elements of 'nums' from index '0' to 'i'
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(nums[i])
	}

	for i := k; i < (n - k); i++ {
		leftBound := i - k
		rightBound := i + k

		subArray := prefix[rightBound+1] - prefix[leftBound]
		average := int(subArray / int64(windowSize))

		avg[i] = average
	}

	return avg
}
