package techniques

// FindMaxLength finds the length of the longest subarray whose sum is less than or equal to k.
//
// Time Complexity: O(n) - Each element is added and removed from the window at most once.
// Space Complexity: O(1) - Only a few integer variables are used.
func FindMaxLength(nums []int, k int) int {
	left, curr, ans := 0, 0, 0
	
	for right := 0; right < len(nums); right++ {
		curr += nums[right] // Add the current element to the window sum.
		// If the current sum exceeds k, shrink the window from the left.
		for curr > k {
			curr -= nums[left] // Remove the element at the left pointer from the window sum.
			left++
		}
		// Update the answer with the maximum window size found.
		ans = max(ans, right-left+1)
	}
	return ans
}

// max is a helper function that returns the larger of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// FindLength finds the length of the longest substring containing only '1', with at most one '0' that can be flipped to '1'.
//
// Time Complexity: O(n) - Each character is added and removed from the window at most once.
// Space Complexity: O(1) - Only a few integer variables are used.
func FindLength(s string) int {
	left, curr, ans := 0, 0, 0
	for right := 0; right < len(s); right++ {
		if s[right] == '0' {
			curr++ // Count the number of '0's in the current window.
		}
		// If there are more than one '0' in the window, shrink the window from the left.
		for curr > 1 {
			if s[left] == '0' {
				curr--
			}
			left++
		}
		// Update the answer with the maximum window size found.
		ans = max(ans, right-left+1)
	}
	return ans
}

// NumSubarrayProductLessThanK returns the number of subarrays where the product of all the elements is strictly less than k.
//
// Time Complexity: O(n) - Each element is multiplied into and divided out of the window product at most once.
// Space Complexity: O(1) - Only a few integer variables are used.
func NumSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	ans, left, curr := 0, 0, 1
	for right := 0; right < len(nums); right++ {
		curr *= nums[right] // Multiply the current element into the window product.
		// If the current product exceeds or equals k, shrink the window from the left.
		for curr >= k {
			curr /= nums[left] // Divide the element at the left pointer out of the window product.
			left++
		}
		// Add the number of valid subarrays ending at the current right index.
		ans += right - left + 1
	}
	return ans
}

// FindBestSubarray finds the sum of the subarray with the largest sum and of a fixed length k.
//
// Time Complexity: O(n) - Each element is added to and subtracted from the running sum once.
// Space Complexity: O(1) - Only a few integer variables are used.
func FindBestSubarray(nums []int, k int) int {
	curr := 0
	// Build the sum of the first window of size k.
	for i := 0; i < k; i++ {
		curr += nums[i]
	}
	ans := curr
	// Slide the window across the array and update the sum.
	for i := k; i < len(nums); i++ {
		curr += nums[i] - nums[i-k] // Add the new element and remove the old one from the sum.
		// Update the answer with the maximum sum found.
		ans = max(ans, curr)
	}
	return ans
}