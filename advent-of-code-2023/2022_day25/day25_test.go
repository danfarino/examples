package day25

import (
	_ "embed"
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example
var example string

//go:embed input
var input string

//go:embed encode_test
var encodeTest string

func TestExample(t *testing.T) {
	result, encoded := process(example)
	assert.Equal(t, 4890, result)
	assert.Equal(t, "2=-1=0", encoded)
}

func TestInput(t *testing.T) {
	result, encoded := process(input)
	assert.Equal(t, 34818266939311, result)
	assert.Equal(t, "2-1=10=1=1==2-1=-221", encoded)
}

func TestEncode(t *testing.T) {
	for _, line := range strings.Split(encodeTest, "\n") {
		var input int
		var output string
		_, err := fmt.Sscanln(line, &output, &input)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}

		assert.Equal(t, output, encodeSnafu(input))
	}
}

func BenchmarkInput256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		process(input)
	}
}
