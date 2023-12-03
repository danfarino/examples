package day3

import (
	"regexp"
	"strconv"
	"strings"
)

var numRegexp = regexp.MustCompile(`\d+`)

func processPart1(s string) int {
	lines := strings.Split(strings.TrimSpace(s), "\n")

	total := 0

	for i := range lines {
		numRanges := numRegexp.FindAllStringIndex(lines[i], -1)

		for _, numRange := range numRanges {
			sym, _, _ := findSymbolPos(false, lines, i, numRange[0], numRange[1])
			if sym != 0 {
				numStr := lines[i][numRange[0]:numRange[1]]
				total += mustParseInt(numStr)
			}
		}
	}

	return total
}

func processPart2(s string) int {
	lines := strings.Split(strings.TrimSpace(s), "\n")

	type Position struct {
		x int
		y int
	}
	numsByStar := map[Position][]string{}

	for i := range lines {
		numRanges := numRegexp.FindAllStringIndex(lines[i], -1)

		for _, numRange := range numRanges {
			sym, symX, symY := findSymbolPos(true, lines, i, numRange[0], numRange[1])
			if sym != 0 {
				numStr := lines[i][numRange[0]:numRange[1]]
				key := Position{symX, symY}
				numsByStar[key] = append(numsByStar[key], numStr)
			}
		}
	}

	total := 0

	for _, numStrs := range numsByStar {
		if len(numStrs) == 2 {
			total += mustParseInt(numStrs[0]) * mustParseInt(numStrs[1])
		}
	}

	return total
}

func findSymbolPos(starOnly bool, lines []string, lineIdx, numRangeStart, numRangePastEnd int) (uint8, int, int) {
	left := max(numRangeStart-1, 0)
	right := min(numRangePastEnd, len(lines[lineIdx])-1)
	top := max(lineIdx-1, 0)
	bottom := min(lineIdx+1, len(lines)-1)

	for x := left; x <= right; x++ {
		for y := top; y <= bottom; y++ {
			c := lines[y][x]
			if (starOnly && c == '*') || (!starOnly && (c != '.' && (c < '0' || c > '9'))) {
				return c, x, y
			}
		}
	}

	return 0, 0, 0
}

func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
