package stringsdsa

// ReverseWords reverses the characters of each word in a string s, preserving word order and spaces.
// The input string does not contain leading or trailing spaces.
func ReverseWords(s string) string {
	runes := []rune(s)
	start := 0

	for start < len(runes) {
		// Find the end of the current word
		end := findWordEnd(runes, start)

		// Reverse the current word
		reverse(runes, start, end-1)

		// Move to the next word (skip the space)
		start = end + 1
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

// reverse reverses the elements in the runes slice from index start to index end.
func reverse(runes []rune, start, end int) {
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}
