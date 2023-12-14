package day14

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
	result := processPart1(example)
	assert.Equal(t, 136, result)
}

func TestInputPart1(t *testing.T) {
	result := processPart1(input)
	assert.Equal(t, 109939, result)
}

func TestExamplePart2(t *testing.T) {
	result := processPart2(example, 1_000_000_000)
	assert.Equal(t, 64, result)
}

func TestInputPart2(t *testing.T) {
	result := processPart2(input, 1_000_000_000)
	assert.Equal(t, 101010, result)
}

func BenchmarkInputPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processPart2(input, 1_000_000_000)
	}
}
