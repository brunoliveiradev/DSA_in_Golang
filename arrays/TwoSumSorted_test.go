package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSumSortedArray(t *testing.T) {
	// Teste para caso de sucesso
	result := TwoSumSortedArray([]int{1, 2, 3, 4, 5}, 9)
	assert.True(t, result, "Expected to find a pair that sums to 9")

	// Teste para caso de falha
	result = TwoSumSortedArray([]int{1, 2, 3, 4, 5}, 10)
	assert.False(t, result, "Expected not to find a pair that sums to 10")

	// Teste com array contendo números negativos
	result = TwoSumSortedArray([]int{-10, -3, 2, 4, 8}, -1)
	assert.True(t, result, "Expected to find a pair that sums to -1")

	// Teste com array contendo um único elemento
	result = TwoSumSortedArray([]int{5}, 5)
	assert.False(t, result, "Expected not to find a pair that sums to 5 with a single element array")

	// Teste com array vazio
	result = TwoSumSortedArray([]int{}, 5)
	assert.False(t, result, "Expected not to find a pair in an empty array")
}
