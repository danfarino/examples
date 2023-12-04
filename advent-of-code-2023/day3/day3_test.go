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

func TestExample(t *testing.T) {
	part1, part2 := process(example)
	assert.Equal(t, 4361, part1)
	assert.Equal(t, 467835, part2)
}

func TestInput(t *testing.T) {
	part1, part2 := process(input)
	assert.Equal(t, 532445, part1)
	assert.Equal(t, 79842967, part2)
}

func BenchmarkInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process(input)
	}
}
