package day19

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
	assert.Equal(t, 19114, result)
}

func TestInputPart1(t *testing.T) {
	result := processPart1(input)
	assert.Equal(t, 421983, result)
}

//func BenchmarkInputPart2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		processPart2(input)
//	}
//}
