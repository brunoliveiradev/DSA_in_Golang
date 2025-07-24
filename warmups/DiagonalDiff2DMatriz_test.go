package warmups

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiagonalDifferenceReturnsCorrectValueForMatrixWithEqualDiagonals(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 1},
	}
	result := diagonalDifference(matrix)
	assert.Equal(t, 8, result)
}

func TestDiagonalDifferenceCalculatesCorrectlyForMatrixWithDifferentDiagonals(t *testing.T) {
	matrix := [][]int{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	result := diagonalDifference(matrix)
	assert.Equal(t, 15, result)
}

func TestDiagonalDifferenceHandlesSingleElementMatrix(t *testing.T) {
	matrix := [][]int{
		{5},
	}
	result := diagonalDifference(matrix)
	assert.Equal(t, 0, result)
}

func TestDiagonalDifferenceHandlesMatrixWithNegativeNumbers(t *testing.T) {
	matrix := [][]int{
		{-1, -2, -3},
		{-4, -5, -6},
		{-7, -8, -9},
	}
	result := diagonalDifference(matrix)
	assert.Equal(t, 0, result)
}

func TestDiagonalDifferenceHandlesMatrixWithMixedPositiveAndNegativeNumbers(t *testing.T) {
	matrix := [][]int{
		{1, -2, 3},
		{-4, 5, -6},
		{7, -8, 9},
	}
	result := diagonalDifference(matrix)
	assert.Equal(t, 0, result)
}
