package day3

import (
	"regexp"
	"strconv"
	"strings"
)

var numRegexp = regexp.MustCompile(`\d+`)

type Symbol struct {
	Char uint8
	X    int
	Y    int
}

func process(s string) (int, int) {
	lines := strings.Split(strings.TrimSpace(s), "\n")

	part1Total := 0
	numsByStar := map[Symbol][]string{}

	for i := range lines {
		numRanges := numRegexp.FindAllStringIndex(lines[i], -1)

		for _, numRange := range numRanges {
			numStr := lines[i][numRange[0]:numRange[1]]

			syms := findSymbols(lines, i, numRange[0], numRange[1])
			if len(syms) > 0 {
				part1Total += mustParseInt(numStr)
			}

			for _, sym := range syms {
				if sym.Char == '*' {
					numsByStar[sym] = append(numsByStar[sym], numStr)
				}
			}
		}
	}

	part2Total := 0

	for _, numStrs := range numsByStar {
		if len(numStrs) == 2 {
			part2Total += mustParseInt(numStrs[0]) * mustParseInt(numStrs[1])
		}
	}

	return part1Total, part2Total
}

func findSymbols(lines []string, lineIdx, numRangeStart, numRangePastEnd int) []Symbol {
	left := max(numRangeStart-1, 0)
	right := min(numRangePastEnd, len(lines[lineIdx])-1)
	top := max(lineIdx-1, 0)
	bottom := min(lineIdx+1, len(lines)-1)

	var results []Symbol

	for x := left; x <= right; x++ {
		for y := top; y <= bottom; y++ {
			c := lines[y][x]
			if c != '.' && (c < '0' || c > '9') {
				results = append(results, Symbol{
					Char: c,
					X:    x,
					Y:    y,
				})
			}
		}
	}

	return results
}

func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
