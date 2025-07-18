package arrays

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func captureOutput(f func()) string {
	reader, writer, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = writer
	f()
	writer.Close()
	os.Stdout = old
	var buf bytes.Buffer
	buf.ReadFrom(reader)
	return buf.String()
}

func TestPlusMinus_AllPositive(t *testing.T) {
	arr := []int32{1, 2, 3, 4, 5}
	output := captureOutput(func() { PlusMinus(arr) })
	assert.Equal(t, "1.000000\n0.000000\n0.000000\n", output)
}

func TestPlusMinus_AllNegative(t *testing.T) {
	arr := []int32{-1, -2, -3, -4, -5}
	output := captureOutput(func() { PlusMinus(arr) })
	assert.Equal(t, "0.000000\n1.000000\n0.000000\n", output)
}

func TestPlusMinus_AllZero(t *testing.T) {
	arr := []int32{0, 0, 0, 0, 0}
	output := captureOutput(func() { PlusMinus(arr) })
	assert.Equal(t, "0.000000\n0.000000\n1.000000\n", output)
}

func TestPlusMinus_MixedValues(t *testing.T) {
	arr := []int32{1, 1, 0, -1, -1}
	output := captureOutput(func() { PlusMinus(arr) })
	assert.Equal(t, "0.400000\n0.400000\n0.200000\n", output)
}

func TestPlusMinus_EmptyArray(t *testing.T) {
	arr := []int32{}
	output := captureOutput(func() { PlusMinus(arr) })
	assert.Contains(t, output, "NaN")
}
