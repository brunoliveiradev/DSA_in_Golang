package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinStartValue_PositiveAndNegativeNumbers(t *testing.T) {
	nums := []int{-3, 2, -3, 4, 2}
	expected := 5
	result := MinStartValue(nums)
	assert.Equal(t, expected, result, "Positive and negative numbers test failed")
}

func TestMinStartValue_AllPositiveNumbers(t *testing.T) {
	nums := []int{1, 2}
	expected := 1
	result := MinStartValue(nums)
	assert.Equal(t, expected, result, "All positive numbers test failed")
}

func TestMinStartValue_MixedNumbers(t *testing.T) {
	nums := []int{1, -2, -3}
	expected := 5
	result := MinStartValue(nums)
	assert.Equal(t, expected, result, "Mixed numbers test failed")
}

func TestMinStartValue_SingleNegativeNumber(t *testing.T) {
	nums := []int{-1}
	expected := 2
	result := MinStartValue(nums)
	assert.Equal(t, expected, result, "Single negative number test failed")
}

func TestMinStartValue_SinglePositiveNumber(t *testing.T) {
	nums := []int{1}
	expected := 1
	result := MinStartValue(nums)
	assert.Equal(t, expected, result, "Single positive number test failed")
}

func TestMinStartValue_AllNegativeNumbers(t *testing.T) {
	nums := []int{-1, -2, -3, -4}
	expected := 11
	result := MinStartValue(nums)
	assert.Equal(t, expected, result, "All negative numbers test failed")
}

func TestMinStartValue_EmptyArray(t *testing.T) {
	nums := []int{}
	expected := 1
	result := MinStartValue(nums)
	assert.Equal(t, expected, result, "Empty array test failed")
}
