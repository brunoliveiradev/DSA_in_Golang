package stringsdsa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseOnlyLettersFirstExample_Ok(t *testing.T) {
	input := "ab-cd"
	expected := "dc-ba"
	result := ReverseOnlyLetters(input)
	assert.Equal(t, expected, result, "Should be: dc-ba")
}

func TestReverseOnlyLetters_Ok(t *testing.T) {
	input := "a-bC-dEf-ghIj"
	expected := "j-Ih-gfE-dCba"
	result := ReverseOnlyLetters(input)
	assert.Equal(t, expected, result, "Should be: j-Ih-gfE-dCba")
}

func TestReverseOnlyLetters_OtherExample_Ok(t *testing.T) {
	input := "Test1ng-Leet=code-Q!"
	expected := "Qedo1ct-eeLg=ntse-T!"
	result := ReverseOnlyLetters(input)
	assert.Equal(t, expected, result, "Should be: Qedo1ct-eeLg=ntse-T!")
}

func TestReverseOnlyLetters_OnlyNumbers_ok(t *testing.T) {
	input := "123-456!"
	expected := "123-456!"
	result := ReverseOnlyLetters(input)
	assert.Equal(t, expected, result, "Should be: 123-456!")
}

func TestReverseOnlyLetters_OnlyOneLetter_Ok(t *testing.T) {
	input := "a"
	expected := "a"
	result := ReverseOnlyLetters(input)
	assert.Equal(t, expected, result, "Should be: a")
}

func TestReverseOnlyLetters_OnlyOneNonLetter_Ok(t *testing.T) {
	input := "!"
	expected := "!"
	result := ReverseOnlyLetters(input)
	assert.Equal(t, expected, result, "Should be: !")
}
