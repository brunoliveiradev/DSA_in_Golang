package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsUniqueIntegerWhenAllOthersAreDuplicated(t *testing.T) {
	arr := []int32{1, 2, 3, 2, 1}
	result := LonelyInteger(arr)
	assert.Equal(t, int32(3), result)
}

func TestReturnsUniqueIntegerForSingleElementArray(t *testing.T) {
	arr := []int32{42}
	result := LonelyInteger(arr)
	assert.Equal(t, int32(42), result)
}

func TestReturnsUniqueIntegerWithNegativeNumbers(t *testing.T) {
	arr := []int32{-1, -2, -2}
	result := LonelyInteger(arr)
	assert.Equal(t, int32(-1), result)
}

func TestReturnsUniqueIntegerWithZeroIncluded(t *testing.T) {
	arr := []int32{0, 7, 7}
	result := LonelyInteger(arr)
	assert.Equal(t, int32(0), result)
}

func TestReturnsUniqueIntegerWhenUniqueIsAtEnd(t *testing.T) {
	arr := []int32{5, 5, 8}
	result := LonelyInteger(arr)
	assert.Equal(t, int32(8), result)
}
