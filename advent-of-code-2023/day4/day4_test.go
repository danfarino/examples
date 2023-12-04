package day4

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
	assert.Equal(t, 13, part1)
	assert.Equal(t, 30, part2)
}

func TestInput(t *testing.T) {
	part1, part2 := process(input)
	assert.Equal(t, 26346, part1)
	assert.Equal(t, 8467762, part2)
}

func BenchmarkInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process(input)
	}
}
