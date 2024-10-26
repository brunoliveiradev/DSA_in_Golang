package stringsdsa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReversePrefixWithCharacterFound(t *testing.T) {
	word := "abcdefd"
	ch := byte('d')
	expected := "dcbaefd"
	result := ReversePrefix(word, ch)
	assert.Equal(t, expected, result)
}

func TestReversePrefixWithCharacterFoundAtEnd(t *testing.T) {
	word := "xyxzxe"
	ch := byte('z')
	expected := "zxyxxe"
	result := ReversePrefix(word, ch)
	assert.Equal(t, expected, result)
}

func TestReversePrefixWithCharacterNotFound(t *testing.T) {
	word := "abcd"
	ch := byte('z')
	expected := "abcd"
	result := ReversePrefix(word, ch)
	assert.Equal(t, expected, result)
}

func TestReversePrefixWithCharacterAtStart(t *testing.T) {
	word := "rzwuktxcjfpamlonbgyieqdvhs"
	ch := byte('s')
	expected := "shvdqeiygbnolmapfjcxtkuwzr"
	result := ReversePrefix(word, ch)
	assert.Equal(t, expected, result)
}

func TestReversePrefixWithSingleCharacterWord(t *testing.T) {
	word := "a"
	ch := byte('a')
	expected := "a"
	result := ReversePrefix(word, ch)
	assert.Equal(t, expected, result)
}

func TestReversePrefixWithEmptyString(t *testing.T) {
	word := ""
	ch := byte('a')
	expected := ""
	result := ReversePrefix(word, ch)
	assert.Equal(t, expected, result)
}
