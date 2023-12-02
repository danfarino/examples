package day2

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
	actual := processPart1(example, map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	})
	assert.Equal(t, 8, actual)
}

func TestInputPart1(t *testing.T) {
	actual := processPart1(input, map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	})
	assert.Equal(t, 2156, actual)
}

func TestExamplePart2(t *testing.T) {
	actual := processPart2(example)
	assert.Equal(t, 2286, actual)
}

func TestInputPart2(t *testing.T) {
	actual := processPart2(input)
	assert.Equal(t, 66909, actual)
}

func BenchmarkInputPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processPart1(input, map[string]int{
			"red":   12,
			"green": 13,
			"blue":  14,
		})
	}
}

func BenchmarkInputPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processPart2(input)
	}
}
