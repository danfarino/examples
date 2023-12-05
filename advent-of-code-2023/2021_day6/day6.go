package day6

import (
	"strconv"
	"strings"
)

func process(s string, days int) int {
	buckets := make([]int, 9)

	for _, numStr := range strings.Split(strings.TrimSpace(s), ",") {
		buckets[mustParseInt(numStr)]++
	}

	for i := 0; i < days; i++ {
		zeroCount := buckets[0]
		for j := 1; j < len(buckets); j++ {
			buckets[j-1] = buckets[j]
		}
		buckets[6] += zeroCount
		buckets[8] = zeroCount
	}

	sum := 0
	for _, n := range buckets {
		sum += n
	}

	return sum
}

func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
