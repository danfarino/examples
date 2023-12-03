package day3

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
	actual := processPart1(example)
	assert.Equal(t, 4361, actual)
}

func TestInputPart1(t *testing.T) {
	actual := processPart1(input)
	assert.Equal(t, 532445, actual)
}

func TestExamplePart2(t *testing.T) {
	actual := processPart2(example)
	assert.Equal(t, 467835, actual)
}

func TestInputPart2(t *testing.T) {
	actual := processPart2(input)
	assert.Equal(t, 79842967, actual)
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
