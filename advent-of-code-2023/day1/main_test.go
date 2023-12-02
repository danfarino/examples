package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed test1.txt
var test1Input string

func Test1(t *testing.T) {
	result := process(test1Input)
	assert.Equal(t, 281, result)
}

func TestPuzzle(t *testing.T) {
	result := process(puzzle)
	assert.Equal(t, 55429, result)
}

func BenchmarkProcessLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process(puzzle)
	}
}
