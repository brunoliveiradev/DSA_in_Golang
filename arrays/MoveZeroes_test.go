package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	// Test Case 1: Basic example with interleaved zeros and non-zeros
	n1 := []int{0, 1, 0, 3, 12}
	expected1 := []int{1, 3, 12, 0, 0}
	MoveZeroes(n1)
	assert.Equal(t, expected1, n1)

	// Test Case 2: Single zero in the array
	n2 := []int{0}
	expected2 := []int{0}
	MoveZeroes(n2)
	assert.Equal(t, expected2, n2)

	// Test Case 3: No zeros in the array
	n3 := []int{1, 2, 3, 4, 5}
	expected3 := []int{1, 2, 3, 4, 5}
	MoveZeroes(n3)
	assert.Equal(t, expected3, n3)

	// Test Case 4: Zeros at the beginning and end
	n4 := []int{0, 0, 1, 2, 3, 0, 0}
	expected4 := []int{1, 2, 3, 0, 0, 0, 0}
	MoveZeroes(n4)
	assert.Equal(t, expected4, n4)

	// Test Case 5: All elements are zero
	n5 := []int{0, 0, 0, 0}
	expected5 := []int{0, 0, 0, 0}
	MoveZeroes(n5)
	assert.Equal(t, expected5, n5)

	// Test Case 6: Only one non-zero element
	n6 := []int{0, 0, 0, 5}
	expected6 := []int{5, 0, 0, 0}
	MoveZeroes(n6)
	assert.Equal(t, expected6, n6)

	// Test Case 7: Zeros and non-zeros already separated
	n7 := []int{1, 2, 3, 0, 0, 0}
	expected7 := []int{1, 2, 3, 0, 0, 0}
	MoveZeroes(n7)
	assert.Equal(t, expected7, n7)
}
