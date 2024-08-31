package stringsdsa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSubsequenceSortedString(t *testing.T) {
	// Teste para caso de sucesso
	result := IsSubsequenceSortedString("abc", "ahbgdc")
	assert.True(t, result, "Expected 'abc' to be a subsequence of 'ahbgdc'")

	// Teste onde subString é uma subsequência exata de mainString
	result = IsSubsequenceSortedString("axc", "ahbgdc")
	assert.False(t, result, "Expected 'axc' not to be a subsequence of 'ahbgdc'")

	// Teste onde mainString é uma string vazia
	result = IsSubsequenceSortedString("", "ahbgdc")
	assert.True(t, result, "Expected an empty string to be a subsequence of any string")

	// Teste onde subString é uma string vazia
	result = IsSubsequenceSortedString("abc", "")
	assert.False(t, result, "Expected 'abc' not to be a subsequence of an empty string")

	// Teste com ambas as strings vazias
	result = IsSubsequenceSortedString("", "")
	assert.True(t, result, "Expected an empty string to be a subsequence of another empty string")

	// Teste onde mainString é maior que subString
	result = IsSubsequenceSortedString("abc", "ab")
	assert.False(t, result, "Expected 'abc' not to be a subsequence of 'ab'")
}
