package warmups

// countingSort counts the frequency of integers in the range [0, 99].
// It returns a slice of length 100, where each index represents the count
// of that integer from the input.
//
// Time Complexity: O(n)
// Space Complexity: O(1) — fixed output size (100)
func countingSort(arr []int32) []int32 {
	// Write your code here
	freq := make([]int32, 100)

	// Count the frequency of each integer in the input array
	// The input is guaranteed to be in the range [0, 99]
	for _, val := range arr {
		freq[val]++
	}

	return freq

}
