package day12

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example
var example string

//go:embed input
var input string

func TestExampleSlow(t *testing.T) {
	result := processSlow(example)
	assert.Equal(t, 21, result)
}

func TestExampleFast(t *testing.T) {
	result := processFast(example, false)
	assert.Equal(t, 21, result)
}

func TestExampleUnfolded(t *testing.T) {
	result := processFast(example, true)
	assert.Equal(t, 525152, result)
}

func TestInputSlow(t *testing.T) {
	result := processSlow(input)
	assert.Equal(t, 7402, result)
}

func TestInputFast(t *testing.T) {
	result := processFast(input, false)
	assert.Equal(t, 7402, result)
}

func TestInputFastUnfolded(t *testing.T) {
	result := processFast(input, true)
	assert.Equal(t, 3384337640277, result)
}

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processSlow(input)
	}
}

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processFast(input, false)
	}
}

func BenchmarkFastUnfolded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processFast(input, true)
	}
}
