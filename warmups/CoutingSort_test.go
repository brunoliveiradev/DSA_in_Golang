package warmups

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountingSortSampleInput(t *testing.T) {
	input := []int32{
		63, 25, 73, 1, 98, 73, 56, 84, 86, 57, 16, 83, 8, 25, 81, 56, 9, 53, 98, 67,
		99, 12, 83, 89, 80, 91, 39, 86, 76, 85, 74, 39, 25, 90, 59, 10, 94, 32, 44, 3,
		89, 30, 27, 79, 46, 96, 27, 32, 18, 21, 92, 69, 81, 40, 40, 34, 68, 78, 24, 87,
		42, 69, 23, 41, 78, 22, 6, 90, 99, 89, 50, 30, 20, 1, 43, 3, 70, 95, 33, 46, 44,
		9, 69, 48, 33, 60, 65, 16, 82, 67, 61, 32, 21, 79, 75, 75, 13, 87, 70, 33,
	}

	expected := []int32{
		0, 2, 0, 2, 0, 0, 1, 0, 1, 2, 1, 0, 1, 1, 0, 0, 2, 0, 1, 0, 1, 2, 1, 1, 1, 3,
		0, 2, 0, 0, 2, 0, 3, 3, 1, 0, 0, 0, 0, 2, 2, 1, 1, 1, 2, 0, 2, 0, 1, 0, 1, 0,
		0, 1, 0, 0, 2, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 2, 1, 3, 2, 0, 0, 2, 1, 2, 1, 0,
		2, 2, 1, 2, 1, 2, 1, 1, 2, 2, 0, 3, 2, 1, 1, 0, 1, 1, 1, 0, 2, 2,
	}

	result := countingSort(input)

	assert.Equal(t, expected, result, "The counting sort result should match the expected output")
}

func TestCountingSortEmpty(t *testing.T) {
	input := []int32{}
	expected := make([]int32, 100)

	result := countingSort(input)
	assert.Equal(t, expected, result, "The counting sort result should match the expected output")
}

func TestCountingSortSingleElement(t *testing.T) {
	input := []int32{99}
	expected := make([]int32, 100)
	expected[99] = 1

	result := countingSort(input)
	assert.Equal(t, expected, result, "The counting sort result should match the expected output")
}
