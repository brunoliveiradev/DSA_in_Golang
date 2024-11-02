package stringsdsa

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
