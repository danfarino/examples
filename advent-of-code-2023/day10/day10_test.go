package day10

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example
var example string

//go:embed example2
var example2 string

//go:embed input
var input string

func TestExamplePart1(t *testing.T) {
	result, _ := process(example)
	assert.Equal(t, 8, result)
}

func TestInputPart1(t *testing.T) {
	result, _ := process(input)
	assert.Equal(t, 6979, result)
}

func TestExamplePart2(t *testing.T) {
	_, result := process(example2)
	assert.Equal(t, 10, result)
}

func TestInputPart2(t *testing.T) {
	_, result := process(input)
	assert.Equal(t, 443, result)
}

func BenchmarkInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process(input)
	}
}
