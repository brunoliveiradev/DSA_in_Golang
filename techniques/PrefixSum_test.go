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
