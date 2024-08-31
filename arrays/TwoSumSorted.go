package arrays

import "fmt"

// TwoSumSortedArray checks if there are two unique integers in a sorted array 
// that add up to a specified target value. This function uses the "Two Pointers" 
// technique, which is only effective because the input array is sorted in ascending order. 
// The function starts with one pointer at the beginning (left) and one at the end (right) 
// of the array, and it moves the pointers towards each other based on the sum of the 
// elements at these pointers. 
// 
// Time Complexity: O(n), where n is the number of elements in the array. 
// Space Complexity: O(1), since the function uses a constant amount of extra space.
func TwoSumSortedArray(uniqueIntegers []int, target int) bool {
	left, right := 0, len(uniqueIntegers)-1

	for left < right {
		curr := uniqueIntegers[left] + uniqueIntegers[right]
		if curr == target {
			fmt.Printf("Found two numbers that result in the target: %v and %v.\n", uniqueIntegers[left], uniqueIntegers[right])
			return true
		}
		if curr > target {
			right--
		} else {
			left++
		}
	}

	fmt.Println("No sum found for the target.")
	return false
}
