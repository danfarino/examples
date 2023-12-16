package day16

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
	assert.Equal(t, 46, result)
}

func TestInputPart1(t *testing.T) {
	result := processPart1(input)
	assert.Equal(t, 8539, result)
}

func TestExamplePart2(t *testing.T) {
	result := processPart2(example)
	assert.Equal(t, 51, result)
}

func TestInputPart2(t *testing.T) {
	result := processPart2(input)
	assert.Equal(t, 8674, result)
}

func BenchmarkInputPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processPart1(input)
	}
}

func BenchmarkInputPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processPart2(input)
	}
}
