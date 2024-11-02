package stringsdsa

import (
	"unicode"
)

// ReverseOnlyLetters given a string s, reverse the string according to the following rules:
//
// - all the characters that are not English letters remain in the same position
//
// - all the English letters (lowercase or uppercase) should be reversed
//
// Space and Time Complexity: O(n) where n is the len of the string s
func ReverseOnlyLetters(s string) string {
	letter := []rune(s)
	left, right := 0, len(s)-1

	for left < right {
		if !unicode.IsLetter(letter[left]) {
			// Move left pointer if not a letter
			left++
		} else if !unicode.IsLetter(letter[right]) {
			// Move right pointer if not a letter
			right--
		} else {
			// Swap letters at left and right pointers
			letter[left], letter[right] = letter[right], letter[left]
			left++
			right--
		}
	}

	return string(letter)
}
