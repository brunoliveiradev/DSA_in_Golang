package arrays

// MoveZeroes moves all zeroes in the input slice to the end while maintaining the relative order of the non-zero elements.
// nums: the input slice of integers.
//
// Uses two pointers technique.
// Left pointer: left moves forward whenever it is on a non-zero number, keeping it fixed only in positions where there is a zero that needs to be moved.
// Right pointer: right move through the entire array, always looking for the next non-zero number to swap with left, ensuring that all zeros are pushed to the end.
//
// Space complexity: O(1) - in-place modification
//
// Time complexity: O(n) - where n is the len of the array nums
func MoveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}

	left := 0

	for right := 1; right < len(nums); right++ {
		if nums[left] == 0 && nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		} else if nums[left] != 0 {
			left++
		}
	}
}
