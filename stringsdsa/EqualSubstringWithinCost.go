package stringsdsa

import "math"

// EqualSubstring Return the maximum length of a substring of s that can be changed to be the same as the corresponding substring of t with a cost less than or equal to maxCost. If there is no substring from s that can be changed to its corresponding substring from t, return 0.
//
// Time complexity: O(n) - linear, where n is the length of string inputs
// Space complexity: O(n) - linear, since we needed to use auxiliar variables to convert to rune
//
// Solved using sliding window technique
func EqualSubstring(s string, t string, maxCost int) int {
	// convert string to slice
	// conversion to rune is necessary to correctly handle characters in Go, especially in mathematical operations involving different ASCII values.
	sRunes, tRunes := []rune(s), []rune(t)
	maxChanges, currentCostToChange := 0, 0
	left := 0

	for right := 0; right < len(sRunes); right++ {
		// diff between sRunes[right] e tRunes[right]
		diff := int(math.Abs(float64(sRunes[right] - tRunes[right])))
		currentCostToChange += diff

		// If the window is broken (cost > maxCost), remove from left
		if currentCostToChange > maxCost {
			leftDiff := int(math.Abs(float64(sRunes[left] - tRunes[left])))
			currentCostToChange -= leftDiff

			left++
		}

		maxChanges = maxInt(maxChanges, right-left+1)
	}
	return maxChanges
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
