package day4

import (
	"strings"
)

func process(s string) (int, int) {
	winCounts := parseWinCounts(s)

	part1Total := 0

	for _, winCount := range winCounts {
		part1Total += 1 << winCount >> 1
	}

	part2Total := part2(winCounts, winCounts, 0)

	return part1Total, part2Total
}

func part2(winCounts, batch []int, batchOffset int) int {
	total := len(batch)

	for i, winCount := range batch {
		idx := batchOffset + i + 1
		total += part2(winCounts, winCounts[idx:idx+winCount], idx)
	}

	return total
}

func parseWinCounts(s string) []int {
	var results []int

	for _, line := range strings.Split(strings.TrimSpace(s), "\n") {
		_, withoutCardHeader, _ := strings.Cut(line, ": ")
		winningNumsStr, myNumsStr, _ := strings.Cut(withoutCardHeader, " | ")

		winningNums := map[string]struct{}{}

		for _, numStr := range strings.Fields(winningNumsStr) {
			winningNums[numStr] = struct{}{}
		}

		winCount := 0
		for _, num := range strings.Fields(myNumsStr) {
			if _, ok := winningNums[num]; ok {
				winCount++
			}
		}

		results = append(results, winCount)
	}

	return results
}
