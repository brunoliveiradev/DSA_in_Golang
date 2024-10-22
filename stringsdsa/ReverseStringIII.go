package stringsdsa

// ReverseWords reverses the characters of each word in a string s, preserving word order and spaces.
// The input string does not contain leading or trailing spaces.
// Space and Time Complexity: O(n) - where n is the len of s, and also because of an extra slice to revert
func ReverseWords(s string) string {
	if len(s) == 1 {
		return s
	}

	runes := []rune(s)
	wordStartIndex := 0

	for wordStartIndex < len(runes) {
		// Find the wordEndIndex of the current word
		wordEndIndex := findWordEnd(runes, wordStartIndex)

		// Reverse the current word
		reverseWord(runes, wordStartIndex, wordEndIndex-1)

		// Move to the next word (skip the space)
		wordStartIndex = wordEndIndex + 1
	}

	return string(runes)
}

// findWordEnd finds the index of the end of the current word, stopping at a space or the end of the slice.
func findWordEnd(runes []rune, start int) int {
	end := start
	for end < len(runes) && runes[end] != ' ' {
		end++
	}
	return end
}

// reverseWord reverses the elements in the runes slice from index start to index end.
func reverseWord(runes []rune, start, end int) {
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}
