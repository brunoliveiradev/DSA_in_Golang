package stringsdsa

// ReverseString reverses the input array of characters in-place with O(1) extra memory.
func ReverseString(s []byte) {
	left, right := 0, len(s)-1

	for left < right {
		// Swap the characters at positions left and right
		s[left], s[right] = s[right], s[left]

		// Move the pointers towards the center
		left++
		right--
	}
}
