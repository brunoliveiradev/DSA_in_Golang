package techniques

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswerQueries(t *testing.T) {
	tests := []struct {
		nums    []int
		queries [][]int
		limit   int
		want    []bool
	}{
		{
			nums:    []int{1, 2, 3, 4, 5},
			queries: [][]int{{0, 2}, {1, 3}, {0, 4}},
			limit:   6,
			want:    []bool{false, false, false},
		},
		{
			nums:    []int{10, 20, 30, 40, 50},
			queries: [][]int{{0, 1}, {1, 2}, {2, 4}},
			limit:   50,
			want:    []bool{true, false, false},
		},
		{
			nums:    []int{5, 5, 5, 5, 5},
			queries: [][]int{{0, 4}, {1, 3}, {2, 2}},
			limit:   15,
			want:    []bool{false, false, true},
		},
		{
			nums:    []int{-1, -2, -3, -4, -5},
			queries: [][]int{{0, 2}, {1, 3}, {0, 4}},
			limit:   -6,
			want:    []bool{false, true, true},
		},
		{
			nums:    []int{1, 2, 3, 4, 5},
			queries: [][]int{{0, 0}, {1, 1}, {2, 2}},
			limit:   3,
			want:    []bool{true, true, false},
		},
		{
			nums:    []int{1, 6, 3, 2, 7, 2},
			queries: [][]int{{0, 3}, {2, 5}, {2, 4}},
			limit:   13,
			want:    []bool{true, false, true},
		},
	}

	for _, tt := range tests {
		got := AnswerQueries(tt.nums, tt.queries, tt.limit)
		assert.Equal(t, tt.want, got, "AnswerQueries(%v, %v, %d)", tt.nums, tt.queries, tt.limit)
	}
}


func TestWaysToSplitArray(t *testing.T) {
	nums1 := []int{10, 4, -8, 7}
	expected1 := 2
	result1 := WaysToSplitArray(nums1)
	assert.Equal(t, expected1, result1, "Test case 1 failed")

	nums2 := []int{2, 3, 1, 0}
	expected2 := 2
	result2 := WaysToSplitArray(nums2)
	assert.Equal(t, expected2, result2, "Test case 2 failed")

	nums3 := []int{-5, -3, -1, -2}
	expected3 := 1
	result3 := WaysToSplitArray(nums3)
	assert.Equal(t, expected3, result3, "Test case 3 failed")

	nums4 := []int{-10, -5, -2, 1}
	expected4 := 0
	result4 := WaysToSplitArray(nums4)
	assert.Equal(t, expected4, result4, "Test case 4 failed")

	nums5 := []int{5, 5, 5, 5}
	expected5 := 2
	result5 := WaysToSplitArray(nums5)
	assert.Equal(t, expected5, result5, "Test case 5 failed")

	nums6 := []int{3, 1}
	expected6 := 1
	result6 := WaysToSplitArray(nums6)
	assert.Equal(t, expected6, result6, "Test case 6 failed")

	nums7 := []int{1000000, 2000000, -1000000, -500000}
	expected7 := 3
	result7 := WaysToSplitArray(nums7)
	assert.Equal(t, expected7, result7, "Test case 7 failed")	
}