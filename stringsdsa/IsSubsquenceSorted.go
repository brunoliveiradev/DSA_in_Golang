package stringsdsa

// IsSubsequenceSortedString checks if the characters of `mainString` appear
// in the same order within `subString`, allowing gaps between characters.
// The function uses two pointers to traverse both strings. It compares the
// current character of `mainString` with the current character of `subString`.
// If they match, it moves the pointer in `mainString` forward. The pointer
// in `subString` always moves forward. The function returns true if all characters
// of `mainString` are found in the correct order within `subString`.
//
// Time Complexity: O(n + m), where n is the length of `mainString` and m is the length of `subString`.
// Space Complexity: O(1), since the function uses a constant amount of extra space.
func IsSubsequenceSortedString(mainString string, subString string) bool {
	i, j := 0, 0

	for (i < len(mainString)) && (j < len(subString)) {
		if mainString[i] == subString[j] {
			i++
		}
		j++
	}

	return i == len(mainString)
}
