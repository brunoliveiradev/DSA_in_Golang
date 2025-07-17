package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountPairsUnsortedArray_MultiplePairs(t *testing.T) {
	arr := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	n := int32(len(arr))
	result := CountPairsUnsortedArray(arr, n)
	assert.Equal(t, int32(3), result)
}

func TestCountPairsUnsortedArray_NoPairs(t *testing.T) {
	arr := []int32{1, 2, 3, 4, 5}
	n := int32(len(arr))
	result := CountPairsUnsortedArray(arr, n)
	assert.Equal(t, int32(0), result)
}

func TestCountPairsUnsortedArray_AllPairs(t *testing.T) {
	arr := []int32{1, 1, 2, 2, 3, 3}
	n := int32(len(arr))
	result := CountPairsUnsortedArray(arr, n)
	assert.Equal(t, int32(3), result)
}

func TestCountPairsUnsortedArray_EmptyArray(t *testing.T) {
	arr := []int32{}
	n := int32(len(arr))
	result := CountPairsUnsortedArray(arr, n)
	assert.Equal(t, int32(0), result)
}

func TestCountPairsUnsortedArray_SingleSock(t *testing.T) {
	arr := []int32{42}
	n := int32(len(arr))
	result := CountPairsUnsortedArray(arr, n)
	assert.Equal(t, int32(0), result)
}

func TestCountPairsUnsortedArray_MismatchedLengthPanics(t *testing.T) {
	arr := []int32{1, 2, 3}
	n := int32(2)
	assert.Panics(t, func() {
		CountPairsUnsortedArray(arr, n)
	})
}
