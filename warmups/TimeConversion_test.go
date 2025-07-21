package warmups

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertsPMTimeCorrectly(t *testing.T) {
	assert.Equal(t, "19:05:45", TimeConversion("07:05:45PM"))
}

func TestConvertsAMTimeCorrectly(t *testing.T) {
	assert.Equal(t, "07:05:45", TimeConversion("07:05:45AM"))
}

func TestConvertsMidnightTo0000(t *testing.T) {
	assert.Equal(t, "00:00:00", TimeConversion("12:00:00AM"))
}

func TestConvertsNoonTo1200(t *testing.T) {
	assert.Equal(t, "12:00:00", TimeConversion("12:00:00PM"))
}

func TestConvertsSingleDigitHourAM(t *testing.T) {
	assert.Equal(t, "01:00:00", TimeConversion("01:00:00AM"))
}

func TestConvertsSingleDigitHourPM(t *testing.T) {
	assert.Equal(t, "13:00:00", TimeConversion("01:00:00PM"))
}
