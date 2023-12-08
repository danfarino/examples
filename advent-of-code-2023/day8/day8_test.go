package day8

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example1
var example1 string

//go:embed example2
var example2 string

//go:embed example3
var example3 string

//go:embed input
var input string

func TestExample1Part1(t *testing.T) {
	result := processPart1(example1)
	assert.Equal(t, 2, result)
}

func TestExample2Part1(t *testing.T) {
	result := processPart1(example2)
	assert.Equal(t, 6, result)
}

func TestInputPart1(t *testing.T) {
	result := processPart1(input)
	assert.Equal(t, 19783, result)
}

func TestExample3Part2(t *testing.T) {
	result := processPart2(example3)
	assert.Equal(t, 6, result)
}

func TestInputPart2(t *testing.T) {
	result := processPart2(input)
	assert.Equal(t, 9177460370549, result)
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
