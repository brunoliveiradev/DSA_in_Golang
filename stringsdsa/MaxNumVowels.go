package stringsdsa

// MaxVowels returns the maximum number of vowel letters in any substring of s with length k.
// Vowel letters in English are 'a', 'e', 'i', 'o', and 'u'.
//
// Used Sliding window technique to solve this challenge.
// Time Complexity: O(n), where n is the length of the string s.
// Space Complexity: O(1), as the space used by the vowels map is constant.
func MaxVowels(s string, k int) int {
	vowelCount, maxVowels := 0, 0
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	for i, char := range s {
		if vowels[char] {
			vowelCount++
		}
		if i >= k && vowels[rune(s[i-k])] {
			vowelCount--
		}
		if vowelCount > maxVowels {
			maxVowels = vowelCount
		}
	}
	return maxVowels
}
