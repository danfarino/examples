package day13

import (
	"math/bits"
	"strings"
)

func process(s string, smudge bool) int {
	total := 0

	for _, pattern := range parsePatterns(s) {
		mirror := findMirror(pattern.Rows, smudge)
		if mirror >= 0 {
			total += 100 * mirror
			continue
		}

		mirror = findMirror(pattern.Cols, smudge)
		if mirror >= 0 {
			total += mirror
			continue
		}

		panic("no mirror found")
	}

	return total
}

type Pattern struct {
	Rows []int
	Cols []int
}

func parsePatterns(s string) []Pattern {
	var patterns []Pattern

	for _, patternStr := range strings.Split(strings.TrimSpace(s), "\n\n") {
		lines := strings.Split(patternStr, "\n")

		patterns = append(patterns, Pattern{
			Rows: getRowVals(lines),
			Cols: getColVals(lines),
		})
	}

	return patterns
}

func getRowVals(lines []string) []int {
	var rowVals []int

	for _, line := range lines {
		rowVal := 0

		for x, c := range line {
			if c == '#' {
				rowVal |= 1 << x
			}
		}

		rowVals = append(rowVals, rowVal)
	}

	return rowVals
}

func getColVals(lines []string) []int {
	var colVals []int

	for x := 0; x < len(lines[0]); x++ {
		colVal := 0

		for y := 0; y < len(lines); y++ {
			if lines[y][x] == '#' {
				colVal |= 1 << y
			}
		}

		colVals = append(colVals, colVal)
	}

	return colVals
}

func findMirror(ns []int, smudge bool) int {
	for i := 1; i < len(ns); i++ {
		i2 := i
		smudged := false

		for j := i - 1; j >= 0 && i2 < len(ns); j-- {
			n1 := ns[j]
			n2 := ns[i2]

			if n1 != n2 {
				if (smudge && !smudged) && bits.OnesCount(uint(n1)^uint(n2)) == 1 {
					smudged = true
				} else {
					break
				}
			}

			if (j == 0 || i2 == len(ns)-1) && (!smudge || smudged) {
				return i
			}

			i2++
		}
	}

	return -1
}
