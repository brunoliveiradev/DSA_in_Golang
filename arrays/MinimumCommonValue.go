package arrays

// GetCommon is a function where given two integer arrays nums1 and nums2, sorted in non-decreasing order, return the minimum integer common to both arrays. If there is no common integer amongst nums1 and nums2, return -1.
//
// # This has solved using two pointers technique
//
// Space complexity: O(1) - since no new structure is created
// Time Complexity: O(n + m), where 'n' is the length of nums1 and 'm' is the length of nums2.
func GetCommon(nums1 []int, nums2 []int) int {
	a, b := 0, 0

	// The loop continues as long as both pointers are within the bounds of the arrays.
	for a < len(nums1) && b < len(nums2) {

		if nums1[a] == nums2[b] {
			// The first common element found will be the smallest since the arrays are sorted.
			return nums1[a]
		} else if nums1[a] < nums2[b] {
			// If the element in 'nums1[a]' is smaller than 'nums2[b]', increment pointer 'a'.
			// This means we are trying to match by increasing the value in 'nums1'.
			a++
		} else {
			// Otherwise, if 'nums2[b]' is smaller than 'nums1[a]', increment pointer 'b'.
			// Thus, we increase the value in 'nums2' to seek a possible match.
			b++
		}
	}
	// This means there are no common elements in the arrays.
	return -1
}
