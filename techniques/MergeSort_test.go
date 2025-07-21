package techniques

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateRandomArrayReturnsCorrectLength(t *testing.T) {
	arr := GenerateRandomArray(10)
	assert.Equal(t, 10, len(arr))
}

func TestGenerateRandomArrayReturnsEmptySliceForZeroLength(t *testing.T) {
	arr := GenerateRandomArray(0)
	assert.Equal(t, 0, len(arr))
}

func TestGenerateRandomArrayReturnsEmptySliceForNegativeLength(t *testing.T) {
	arr := GenerateRandomArray(-5)
	assert.Equal(t, 0, len(arr))
}

func TestMergeSortSortsAlreadySortedArray(t *testing.T) {
	arr := []int32{1, 2, 3, 4, 5}
	expected := []int32{1, 2, 3, 4, 5}
	MergeSort(arr)
	assert.Equal(t, expected, arr)
}

func TestMergeSortSortsReverseArray(t *testing.T) {
	arr := []int32{5, 4, 3, 2, 1}
	expected := []int32{1, 2, 3, 4, 5}
	MergeSort(arr)
	assert.Equal(t, expected, arr)
}

func TestMergeSortSortsArrayWithDuplicates(t *testing.T) {
	arr := []int32{3, 1, 2, 3, 1}
	expected := []int32{1, 1, 2, 3, 3}
	MergeSort(arr)
	assert.Equal(t, expected, arr)
}

func TestMergeSortHandlesSingleElementArray(t *testing.T) {
	arr := []int32{42}
	expected := []int32{42}
	MergeSort(arr)
	assert.Equal(t, expected, arr)
}

func TestMergeSortHandlesEmptyArray(t *testing.T) {
	arr := []int32{}
	expected := []int32{}
	MergeSort(arr)
	assert.Equal(t, expected, arr)
}
