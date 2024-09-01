package arrays

import (
	"fmt"
)

// SortedSquares takes a sorted integer array `nums` (which can contain negative numbers)
// and returns a new array containing the squares of each number, sorted in ascending order.
// The function efficiently computes the result by using the "Two Pointers" technique.
// It starts with two pointers: `left` at the beginning and `right` at the end of the array.
// Since the input array is sorted, the largest squared value will either be at `left` or `right`.
// The function compares the squares of the numbers at these pointers, placing the larger
// one at the end of the result array `ans`, and then moves the corresponding pointer inward.
// This process is repeated until all elements are processed.
//
// Time Complexity: O(n), where n is the number of elements in the array.
// Space Complexity: O(n), for the output array `ans`.
func SortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	ans := make([]int, len(nums))
	currPosition := len(nums) - 1 // tracks the position to insert the next largest square

	for left <= right {
		leftSquare, rightSquare := nums[left]*nums[left], nums[right]*nums[right]
		// Compare the squares of the values at left and right pointers and place the larger square at currPosition in ans
		if leftSquare > rightSquare {
			ans[currPosition] = leftSquare
			left++
		} else {
			ans[currPosition] = rightSquare
			right--
		}
		// Move the corresponding pointer (left or right) inward and decrement currPosition
		currPosition--
	}

	fmt.Printf("ans: %v\n", ans)
	return ans
}
