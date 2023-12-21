package day17

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example
var example string

//go:embed input
var input string

func TestExamplePart1(t *testing.T) {
	result := processPart1(example)
	assert.Equal(t, 102, result)
}

func TestExamplePart2(t *testing.T) {
	result := processPart2(example)
	assert.Equal(t, 94, result)
}

func TestInputPart1(t *testing.T) {
	result := processPart1(input)
	assert.Equal(t, 814, result)
}

func TestInputPart2(t *testing.T) {
	result := processPart2(input)
	assert.Equal(t, 974, result)
}

func BenchmarkInputPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processPart2(input)
	}
}

func BenchmarkInputPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processPart2(input)
	}
}

func TestRuns(t *testing.T) {
	tests := []struct {
		path []Point
		xRun int
		yRun int
	}{
		{
			path: []Point{
				{0, 0},
				{1, 0},
			},
			xRun: 1,
			yRun: 0,
		},
		{
			path: []Point{
				{0, 0},
				{1, 0},
				{2, 0},
			},
			xRun: 2,
			yRun: 0,
		},
		{
			path: []Point{
				{0, 0},
				{1, 0},
				{2, 0},
				{3, 0},
				{4, 0},
			},
			xRun: 3,
			yRun: 0,
		},
		{
			path: []Point{
				{0, 0},
				{1, 0},
				{2, 0},
				{3, 0},
				{3, 1},
			},
			xRun: 0,
			yRun: 1,
		},
		{
			path: []Point{
				{10, 10},
				{9, 10},
				{8, 10},
				{7, 10},
			},
			xRun: -3,
			yRun: 0,
		},
		{
			path: []Point{
				{10, 10},
				{10, 9},
				{10, 8},
				{10, 7},
			},
			xRun: 0,
			yRun: -3,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			xRun, yRun := getRuns(test.path, 3)
			assert.Equal(t, test.xRun, xRun)
			assert.Equal(t, test.yRun, yRun)
		})
	}
}
