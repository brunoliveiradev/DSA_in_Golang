package stringsdsa

// ReversePrefix reverses the prefix of the given word up to and including the first occurrence of the specified character `ch`.
// If the character `ch` is not found in the word, the original word is returned.
//
// Time Complexity: O(n), where n is the length of the word. This is because the function may need to scan the entire word to find the character `ch`.
// Space Complexity: O(n), where n is the length of the word. This is due to the storage required for the byte slice.
//
// Solved using two pointers technique
func ReversePrefix(word string, ch byte) string {
	if len(word) == 1 {
		return word
	}

	char := []byte(word)
	foundCh := -1 //Initialize to -1 to indicate "not found"
	for i := 0; i < len(char); i++ {
		if char[i] == ch {
			foundCh = i
			break
		}
	}

	if foundCh == -1 {
		return word
	}

	// Two pointers to reverse the section from 0 to foundCh
	left, right := 0, foundCh
	for left < right {
		char[left], char[right] = char[right], char[left]
		left++
		right--
	}

	return string(char)
}
