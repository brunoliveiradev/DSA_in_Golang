package arrays

import "fmt"

// CombineTwoSortedArrays merges two sorted arrays with unique elements into a single 
// sorted array. The function assumes that both input arrays are already sorted in 
// ascending order. It uses two pointers, one for each array, to compare and merge 
// elements into the resulting array in a sorted manner.
//
// Time Complexity: O(n + m), where n and m are the lengths of the two arrays.
// Space Complexity: O(n + m), for storing the combined result.
func CombineTwoSortedArrays(n []int, m []int) []int {
	var ans []int
	i, j := 0, 0

	for (i < len(n)) && (j < len(m)) {
		num1 := n[i]
		num2 := m[j]

		if num1 < num2 {
			ans = append(ans, num1)
			i++
		} else {
			ans = append(ans, num2)
			j++
		}
	}

	// add leftover elements if exists
	for i < len(n) {
		ans = append(ans, n[i])
		i++
	}
	
	// add leftover elements if exists
	for j < len(m) {
		ans = append(ans, m[j])
		j++
	}

	fmt.Printf("ans: %v\n", ans)

	return ans
}
