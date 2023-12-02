package main

import (
	_ "embed"
	"strings"
)

//go:embed puzzle.txt
var puzzle string

func main() {
	process(puzzle)
}

func process(input string) int {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		total += processLine(line)
	}
	return total
}

var numWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func processLine(line string) int {
	num := 0

	for i := 0; i < len(line); i++ {
		if n, ok := getNum(line[i:]); ok {
			num = n * 10
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if n, ok := getNum(line[i:]); ok {
			num += n
			break
		}
	}

	return num
}

func getNum(s string) (int, bool) {
	c := s[0]
	if c >= '0' && c <= '9' {
		return int(c - '0'), true
	}

	for i, word := range numWords {
		if strings.HasPrefix(s, word) {
			return i + 1, true
		}
	}

	return 0, false
}
