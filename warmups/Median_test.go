package warmups

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMedian(t *testing.T) {
	assert.Equal(t, int32(13), FindMedian([]int32{12, 15, 10, 20, 25, 30, 5, 8}))
	assert.Equal(t, int32(30), FindMedian([]int32{10, 20, 30, 40, 50}))
	assert.Equal(t, int32(3), FindMedian([]int32{3, 5, 4, 1, 2}))

}

func TestFindMedianEvenLength(t *testing.T) {
	assert.Equal(t, int32(16), FindMedian([]int32{12, 15, 10, 20, 25, 30, 5, 8, 18, 22}))
	assert.Equal(t, int32(10), FindMedian([]int32{1, 7, 3, 5, 11, 9, 15, 13, 19, 17}))
}
