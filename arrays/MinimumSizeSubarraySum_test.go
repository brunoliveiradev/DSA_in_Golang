package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimalLengthSubarraySum(t *testing.T) {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	expected := 2
	result := MinSubArrayLen(target, nums)
	assert.Equal(t, expected, result)
}

func TestMinimalLengthSubarraySumSingleElement(t *testing.T) {
	target := 4
	nums := []int{1, 4, 4}
	expected := 1
	result := MinSubArrayLen(target, nums)
	assert.Equal(t, expected, result)
}

func TestNoSubarraySum(t *testing.T) {
	target := 11
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1}
	expected := 0
	result := MinSubArrayLen(target, nums)
	assert.Equal(t, expected, result)
}

func TestMinimalLengthSubarraySumExactMatch(t *testing.T) {
	target := 15
	nums := []int{1, 2, 3, 4, 5}
	expected := 5
	result := MinSubArrayLen(target, nums)
	assert.Equal(t, expected, result)
}

func TestMinimalLengthSubarraySumEmptyArray(t *testing.T) {
	target := 5
	nums := []int{}
	expected := 0
	result := MinSubArrayLen(target, nums)
	assert.Equal(t, expected, result)
}
