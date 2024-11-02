package stringsdsa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqualSubstring(t *testing.T) {
	assert.Equal(t, 3, EqualSubstring("abcd", "bcdf", 3))
	assert.Equal(t, 1, EqualSubstring("abcd", "cdef", 3))
	assert.Equal(t, 1, EqualSubstring("abcd", "acde", 0))
}
