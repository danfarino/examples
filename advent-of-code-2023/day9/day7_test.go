package day7

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example
var example string

//go:embed input
var input string

func TestExamplePart1(t *testing.T) {
	result := process(example, false)
	assert.Equal(t, 114, result)
}

func TestInputPart1(t *testing.T) {
	result := process(input, false)
	assert.Equal(t, 1980437560, result)
}

func TestExamplePart2(t *testing.T) {
	result := process(example, true)
	assert.Equal(t, 2, result)
}

func TestInputPart2(t *testing.T) {
	result := process(input, true)
	assert.Equal(t, 977, result)
}

func BenchmarkInputPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process(input, false)
	}
}

func BenchmarkInputPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process(input, true)
	}
}
