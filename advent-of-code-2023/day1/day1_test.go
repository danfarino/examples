package day1

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed calibration1
var calibration1 string

//go:embed calibration2
var calibration2 string

//go:embed input1
var input1 string

//go:embed input2
var input2 string

func TestCalibration1(t *testing.T) {
	result := process(calibration1, false)
	assert.Equal(t, 142, result)
}

func Test1(t *testing.T) {
	result := process(input1, false)
	assert.Equal(t, 56397, result)
}

func TestCalibration2(t *testing.T) {
	result := process(calibration2, true)
	assert.Equal(t, 281, result)
}

func Test2(t *testing.T) {
	result := process(input2, true)
	assert.Equal(t, 55701, result)
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process(input1, false)
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process(input2, true)
	}
}
