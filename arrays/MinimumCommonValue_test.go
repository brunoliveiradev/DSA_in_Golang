package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumCommonValueFirstExample_Ok(t *testing.T) {
	n1 := []int{1, 2, 3}
	n2 := []int{2, 4}
	expected := 2
	result := GetCommon(n1, n2)
	assert.Equal(t, expected, result, "Should be: 2")
}

func TestMinimumCommonValueSecondExample_Ok(t *testing.T) {
	n1 := []int{1, 2, 3, 6}
	n2 := []int{2, 3, 4, 5}
	expected := 2
	result := GetCommon(n1, n2)
	assert.Equal(t, expected, result, "Should be: 2")
}

func TestMinimumCommonValueThridExample_Ok(t *testing.T) {
	n1 := []int{1, 2, 3, 6}
	n2 := []int{7, 8, 9, 10}
	expected := -1
	result := GetCommon(n1, n2)
	assert.Equal(t, expected, result, "Should be: 2")
}

func TestMinimumCommonValueFourthExample_Ok(t *testing.T) {
	n1 := []int{1}
	n2 := []int{2}
	expected := -1
	result := GetCommon(n1, n2)
	assert.Equal(t, expected, result, "Should be: 2")
}
