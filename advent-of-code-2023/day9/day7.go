package day7

import (
	"strconv"
	"strings"
)

func process(s string, negative bool) int {
	lines := parse(s)

	sum := 0
	for _, line := range lines {
		sum += predict(line, negative)
	}
	return sum
}

func predict(line []int, negative bool) int {
	diffs := make([]int, 0, len(line)-1)

	diffsAllZero := true
	for i := 1; i < len(line); i++ {
		diff := line[i] - line[i-1]
		if diff != 0 {
			diffsAllZero = false
		}
		diffs = append(diffs, diff)
	}

	if negative {
		firstNum := line[0]
		if diffsAllZero {
			return firstNum
		}
		return firstNum - predict(diffs, negative)
	}

	lastNum := line[len(line)-1]
	if diffsAllZero {
		return lastNum
	}
	return lastNum + predict(diffs, negative)
}

func parse(s string) [][]int {
	var result [][]int
	for _, line := range strings.Split(strings.TrimSpace(s), "\n") {
		numStrs := strings.Fields(line)
		nums := make([]int, 0, len(numStrs))
		for _, numStr := range numStrs {
			nums = append(nums, mustParseInt(numStr))
		}
		result = append(result, nums)
	}

	return result
}

func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
