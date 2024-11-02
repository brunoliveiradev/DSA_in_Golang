package stringsdsa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxVowels(t *testing.T) {
	assert.Equal(t, 3, MaxVowels("abciiidef", 3))
	assert.Equal(t, 2, MaxVowels("aeiou", 2))
	assert.Equal(t, 2, MaxVowels("leetcode", 3))
	assert.Equal(t, 0, MaxVowels("bcdfg", 2))
	assert.Equal(t, 1, MaxVowels("a", 1))
	assert.Equal(t, 0, MaxVowels("b", 1))
	assert.Equal(t, 5, MaxVowels("aeiouaeiou", 5))
	assert.Equal(t, 0, MaxVowels("", 0))
	assert.Equal(t, 0, MaxVowels("xyz", 1))
	assert.Equal(t, 1, MaxVowels("a", 1))
}
