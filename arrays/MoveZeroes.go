package arrays

// MoveZeroes moves all zeroes in the input slice to the end while maintaining the relative order of the non-zero elements.
// nums: the input slice of integers.
//
// Space complexity: O(1) - in-place modification
//
// Time complexity: O(n) - where n is the len of the array nums
func MoveZeroes(nums []int) {
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
