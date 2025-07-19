package warmups

import "fmt"

// PlusMinus calculate the ratios of its elements that are positive, negative, and zero. Print the decimal value of each fraction on a new line with  places after the decimal.
// The function takes a slice of int32 as input and prints three lines:
//   - The first line contains the ratio of positive numbers.
//   - The second line contains the ratio of negative numbers.
//   - The third line contains the ratio of zeros.
//
// Each ratio is printed with six digits after the decimal point.
//
// Time Complexity: O(n), where n is the number of elements in the input array.
// Space Complexity: O(1), as we are using a constant amount of space for counters.
func PlusMinus(arr []int32) {
	var positive, negative, zero float32

	for _, n := range arr {
		if n > 0 {
			positive++
		} else if n < 0 {
			negative++
		} else {
			zero++
		}
	}

	total := float32(len(arr))
	fmt.Printf("%.6f\n", positive/total)
	fmt.Printf("%.6f\n", negative/total)
	fmt.Printf("%.6f\n", zero/total)
}
