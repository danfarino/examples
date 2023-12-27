package day18

import (
	_ "embed"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example
var example string

//go:embed input
var input string

func TestExamplePart1(t *testing.T) {
	result := processPart1(example)
	assert.Equal(t, 62, result)
}

func TestExamplePart2(t *testing.T) {
	result := processPart2(example)
	assert.Equal(t, 952408144115, result)
}

func TestInputPart1(t *testing.T) {
	result := processPart1(input)
	assert.Equal(t, 40761, result)
}

func TestInputPart2(t *testing.T) {
	result := processPart2(input)
	assert.Equal(t, 106920098354636, result)
}

func TestSpans_Subtract(t *testing.T) {
	tests := []struct {
		input    Spans
		hole     Span
		expected Spans
	}{
		0: {
			Spans{
				{5, 10},
				{15, 20},
			},
			Span{11, 14},
			Spans{
				{5, 10},
				{15, 20},
			},
		},
		1: {
			Spans{
				{5, 10},
				{15, 20},
			},
			Span{8, 18},
			Spans{
				{5, 7},
				{19, 20},
			},
		},
		2: {
			Spans{
				{5, 10},
				{15, 20},
			},
			Span{11, 14},
			Spans{
				{5, 10},
				{15, 20},
			},
		},
		3: {
			Spans{
				{5, 10},
				{15, 20},
			},
			Span{10, 15},
			Spans{
				{5, 9},
				{16, 20},
			},
		},
		4: {
			Spans{
				{5, 10},
				{15, 20},
			},
			Span{0, 20},
			Spans{},
		},
		5: {
			Spans{
				{5, 10},
				{15, 20},
			},
			Span{6, 8},
			Spans{
				{5, 5},
				{9, 10},
				{15, 20},
			},
		},
		6: {
			Spans{
				{5, 10},
				{15, 20},
			},
			Span{10, 10},
			Spans{
				{5, 9},
				{15, 20},
			},
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := test.input.Subtract(test.hole)
			assert.Equal(t, test.expected, actual, "hole: %v\ninput: %v", test.hole, test.input)
		})
	}
}

func TestSpans_Union(t *testing.T) {
	tests := []struct {
		input    Spans
		add      Span
		expected Spans
	}{
		0: {
			Spans{
				{5, 10},
				{15, 20},
			},
			Span{5, 10},
			Spans{
				{5, 10},
				{15, 20},
			},
		},
		1: {
			Spans{
				{5, 10},
				{15, 20},
			},
			Span{11, 14},
			Spans{
				{5, 20},
			},
		},
		2: {
			Spans{
				{5, 10},
				{15, 20},
			},
			Span{0, 25},
			Spans{
				{0, 25},
			},
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := test.input.Union(test.add)
			assert.Equal(t, test.expected, actual, "add: %v\ninput: %v", test.add, test.input)
		})
	}
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
