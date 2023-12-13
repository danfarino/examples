package day11

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
	result := process(example, 2)
	assert.Equal(t, 374, result)
}

func TestInputPart1(t *testing.T) {
	result := process(input, 2)
	assert.Equal(t, 9647174, result)
}

func TestExamplePart2_10(t *testing.T) {
	result := process(example, 10)
	assert.Equal(t, 1030, result)
}

func TestExamplePart2_100(t *testing.T) {
	result := process(example, 100)
	assert.Equal(t, 8410, result)
}

func TestInputPart2(t *testing.T) {
	result := process(input, 1_000_000)
	assert.Equal(t, 377318892554, result)
}

func BenchmarkInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process(input, 1_000_000)
	}
}
