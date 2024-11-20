package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHighestAltitudeWithPositiveGains(t *testing.T) {
	assert.Equal(t, 1, HighestAltitude([]int{-5, 1, 5, 0, -7}))
}

func TestHighestAltitudeWithNegativeGains(t *testing.T) {
	assert.Equal(t, 0, HighestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}))
}

func TestHighestAltitudeWithAllNegativeGains(t *testing.T) {
	assert.Equal(t, 0, HighestAltitude([]int{-1, -2, -3, -4}))
}

func TestHighestAltitudeWithAllPositiveGains(t *testing.T) {
	assert.Equal(t, 10, HighestAltitude([]int{1, 2, 3, 4}))
}

func TestHighestAltitudeWithZeroGains(t *testing.T) {
	assert.Equal(t, 0, HighestAltitude([]int{0, 0, 0, 0}))
}

func TestHighestAltitudeWithEmptyGainArray(t *testing.T) {
	assert.Equal(t, 0, HighestAltitude([]int{}))
}
