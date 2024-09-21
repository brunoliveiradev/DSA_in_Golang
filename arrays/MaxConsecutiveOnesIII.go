package arrays

// LongestOnes finds the longest subarray with at most k 0s in the given binary array nums.
// It uses a sliding window approach to maintain the longest subarray with the allowed number of flips (0s to 1s).
//
// The function iterates through the array with a right pointer, and for each 0 encountered, it decreases k.
// If k becomes negative, it means the window contains more than the allowed number of 0s, so the left pointer is incremented
// to move the window forward until k is non-negative again. The length of the longest valid window is returned.
//
// Technique used: Sliding Window
// Time Complexity: O(n), where n is the length of the input array nums
// Space Complexity: O(1), as we are using a constant amount of extra space
func LongestOnes(nums []int, k int) int {
    var left, right int

    for right = range nums {
        if nums[right] == 0 {
            k--
        }

        // If k becomes negative, it means we have more than the allowed number of 0s in the window.
        // Increment the left pointer to move the window forward until k is non-negative again.
		// (If we use a for k < 0 we will shrink the window instead of moving forward)
        if k < 0 {
            if nums[left] == 0 {
                k++
            }
            left++
        }
    }

    return right - left + 1
}