package stringsdsa

import "fmt"

// TwoPointers checks if the given string is a palindrome using the two pointers technique.
// This technique compares characters from the beginning and end of the string, moving
// towards the center. If all characters match, the function returns true, indicating
// the string is a palindrome; otherwise, it returns false.
//
// Time complexity: O(n) - n is the length of the string.
//
// Space complexity: O(1) - only a constant amount of space is used.
func TwoPointers(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			fmt.Printf("%s is not a Palindrome\n", s)
			return false
		}
		left++
		right--
	}
	fmt.Printf("%s is a Palindrome \n", s)
	return true
}
