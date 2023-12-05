package day6

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

var example = "3,4,3,1,2"

//go:embed input
var input string

func TestExampleSmall(t *testing.T) {
	result := process(example, 80)
	assert.Equal(t, 5934, result)
}

func TestExampleLarge(t *testing.T) {
	result := process(example, 256)
	assert.Equal(t, 26984457539, result)
}

func TestInputSmall(t *testing.T) {
	result := process(input, 80)
	assert.Equal(t, 355386, result)
}

func TestInputLarge(t *testing.T) {
	result := process(input, 256)
	assert.Equal(t, 1613415325809, result)
}

func BenchmarkInput256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process(input, 256)
	}
}
