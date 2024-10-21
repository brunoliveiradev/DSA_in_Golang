package stringsdsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWordsMultipleWords(t *testing.T) {
	input := "Let's take LeetCode contest"
	expected := "s'teL ekat edoCteeL tsetnoc"
	result := ReverseWords(input)
	assert.Equal(t, expected, result, "Deveria inverter cada palavra corretamente")
}

func TestReverseWordsSingleWord(t *testing.T) {
	input := "Hello"
	expected := "olleH"
	result := ReverseWords(input)
	assert.Equal(t, expected, result, "Deveria inverter a palavra corretamente")
}

func TestReverseWordsMultipleWordsWithSpace(t *testing.T) {
	input := "Mr Ding"
	expected := "rM gniD"
	result := ReverseWords(input)
	assert.Equal(t, expected, result, "Deveria inverter as palavras corretamente")
}

func TestReverseWordsSingleLetter(t *testing.T) {
	input := "A"
	expected := "A"
	result := ReverseWords(input)
	assert.Equal(t, expected, result, "Deveria retornar a mesma letra, pois não há o que inverter")
}

func TestReverseWordsNumbersAndSpecialCharacters(t *testing.T) {
	input := "123 456 789"
	expected := "321 654 987"
	result := ReverseWords(input)
	assert.Equal(t, expected, result, "Deveria inverter corretamente os números")

	input2 := "Go! is fun!"
	expected2 := "!oG si !nuf"
	result2 := ReverseWords(input2)
	assert.Equal(t, expected2, result2, "Deveria inverter corretamente palavras com pontuação")
}

func TestReverseWordsOnlySpaces(t *testing.T) {
	input := "    "
	expected := "    "
	result := ReverseWords(input)
	assert.Equal(t, expected, result, "Deveria retornar a mesma string com espaços, já que não há palavras para inverter")
}
