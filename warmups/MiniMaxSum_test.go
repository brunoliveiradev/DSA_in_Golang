package warmups

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiniMaxSum_BasicExample(t *testing.T) {
	arr := []int32{1, 2, 3, 4, 5}
	output := captureOutput(func() { miniMaxSum(arr) })
	assert.Equal(t, "10 14", output)
}

func TestMiniMaxSum_LargerNumbers(t *testing.T) {
	arr := []int32{7, 69, 2, 221, 8974}
	output := captureOutput(func() { miniMaxSum(arr) })
	assert.Equal(t, "299 9271", output)
}

func TestMiniMaxSum_AllSameNumbers(t *testing.T) {
	arr := []int32{5, 5, 5, 5, 5}
	output := captureOutput(func() { miniMaxSum(arr) })
	assert.Equal(t, "20 20", output)
}

func TestMiniMaxSum_MinimumValues(t *testing.T) {
	arr := []int32{1, 1, 1, 1, 1}
	output := captureOutput(func() { miniMaxSum(arr) })
	assert.Equal(t, "4 4", output)
}

func TestMiniMaxSum_MaximumValues(t *testing.T) {
	arr := []int32{1000000000, 1000000000, 1000000000, 1000000000, 1000000000}
	output := captureOutput(func() { miniMaxSum(arr) })
	assert.Equal(t, "4000000000 4000000000", output)
}

func TestMiniMaxSum_MixedSmallAndLarge(t *testing.T) {
	arr := []int32{1, 1000000000, 1, 1, 1000000000}
	output := captureOutput(func() { miniMaxSum(arr) })
	assert.Equal(t, "1000000003 2000000002", output)
}

func TestMiniMaxSum_DescendingOrder(t *testing.T) {
	arr := []int32{10, 8, 6, 4, 2}
	output := captureOutput(func() { miniMaxSum(arr) })
	assert.Equal(t, "20 28", output)
}

func TestMiniMaxSum_AscendingOrder(t *testing.T) {
	arr := []int32{2, 4, 6, 8, 10}
	output := captureOutput(func() { miniMaxSum(arr) })
	assert.Equal(t, "20 28", output)
}

func TestMiniMaxSum_RandomOrder(t *testing.T) {
	arr := []int32{3, 1, 4, 1, 5}
	output := captureOutput(func() { miniMaxSum(arr) })
	assert.Equal(t, "9 13", output)
}

func TestMiniMaxSum_EdgeCaseWithZero(t *testing.T) {
	arr := []int32{0, 1, 2, 3, 4}
	output := captureOutput(func() { miniMaxSum(arr) })
	assert.Equal(t, "6 10", output)
}
