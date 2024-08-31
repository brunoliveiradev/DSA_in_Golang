package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombineTwoSortedArrays(t *testing.T) {
	// Teste para caso de sucesso
	result := CombineTwoSortedArrays([]int{1, 4, 7, 20}, []int{3, 5, 6})
	expected := []int{1, 3, 4, 5, 6, 7, 20}
	assert.Equal(t, expected, result, "Expected the combined array to be [1, 3, 4, 5, 6, 7, 20]")

	// Teste onde um dos arrays é vazio
	result = CombineTwoSortedArrays([]int{}, []int{1, 2, 3})
	expected = []int{1, 2, 3}
	assert.Equal(t, expected, result, "Expected the combined array to be [1, 2, 3] when one array is empty")

	result = CombineTwoSortedArrays([]int{1, 2, 3}, []int{})
	expected = []int{1, 2, 3}
	assert.Equal(t, expected, result, "Expected the combined array to be [1, 2, 3] when one array is empty")

	// Teste com arrays contendo um único elemento
	result = CombineTwoSortedArrays([]int{1}, []int{-1})
	expected = []int{-1, 1}
	assert.Equal(t, expected, result, "Expected the combined array to be [-1, 1] with single-element arrays")

	// Teste com arrays contendo números negativos
	result = CombineTwoSortedArrays([]int{-3, -1, 2}, []int{-2, 0, 3})
	expected = []int{-3, -2, -1, 0, 2, 3}
	assert.Equal(t, expected, result, "Expected the combined array to be [-3, -2, -1, 0, 2, 3] with negative numbers")

	// Teste com arrays já combinados
	result = CombineTwoSortedArrays([]int{1, 2, 3}, []int{4, 5, 6})
	expected = []int{1, 2, 3, 4, 5, 6}
	assert.Equal(t, expected, result, "Expected the combined array to be [1, 2, 3, 4, 5, 6] with sorted arrays")
}
