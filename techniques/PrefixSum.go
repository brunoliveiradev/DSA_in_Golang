package techniques

// AnswerQueries processes a list of queries on a given array of integers and determines
// if the sum of elements within specified subarrays is less than a given limit.
//
// Returns:
//   - A slice of boolean values, where each boolean corresponds to a query and indicates
//     whether the sum of the specified subarray is less than the given limit.
//
// The Prefix Sum technique allows the function to efficiently calculate the sum of any
// subarray in constant time after an initial linear time preprocessing step.
// This makes the AnswerQueries function much more efficient than recalculating the sum
// for each query from scratch.
func AnswerQueries(nums []int, queries [][]int, limit int) []bool {
	// Create a prefix sum array to store the cumulative sum of elements
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]

	// Fill the prefix sum array
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	// Initialize the result slice to store the results of each query
	result := make([]bool, len(queries))

	// Process each query
	for i, query := range queries {
		x, y := query[0], query[1]
		var curr int
		if x == 0 {
			// If the subarray starts from the beginning, use the prefix sum directly
			curr = prefix[y]
		} else {
			// Otherwise, calculate the sum of the subarray using the prefix sum array
			curr = prefix[y] - prefix[x-1]
		}
		result[i] = curr < limit
	}

	return result
}
