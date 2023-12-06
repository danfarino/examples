package day25

import (
	"slices"
	"strings"
)

func process(s string) (int, string) {
	sum := 0
	for _, line := range strings.Split(s, "\n") {
		sum += parseSnafu(line)
	}
	return sum, encodeSnafu(sum)
}

var encodingChars = []rune{'=', '-', '0', '1', '2'}

func parseSnafu(s string) int {
	result := 0
	for _, c := range s {
		result = (result * 5) + slices.Index(encodingChars, c) - 2
	}
	return result
}

func encodeSnafu(num int) string {
	var result []rune

	carry := 0

	for num > 0 {
		n := (num % 5) + carry
		if n > 2 {
			n = n - 5
			carry = 1
		} else {
			carry = 0
		}
		result = append(result, encodingChars[n+2])
		num /= 5
	}

	if carry != 0 {
		result = append(result, encodingChars[carry+2])
	}

	slices.Reverse(result)

	return string(result)
}
