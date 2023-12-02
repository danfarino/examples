package day1

import (
	"strings"
)

func process(input string, useWords bool) int {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		total += processLine(line, useWords)
	}
	return total
}

var numWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func processLine(line string, useWords bool) int {
	num := 0

	for i := 0; i < len(line); i++ {
		if n, ok := getNum(line[i:], useWords); ok {
			num = n * 10
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if n, ok := getNum(line[i:], useWords); ok {
			num += n
			break
		}
	}

	return num
}

func getNum(s string, useWords bool) (int, bool) {
	c := s[0]
	if c >= '0' && c <= '9' {
		return int(c - '0'), true
	}

	if useWords {
		for i, word := range numWords {
			if strings.HasPrefix(s, word) {
				return i + 1, true
			}
		}
	}

	return 0, false
}
