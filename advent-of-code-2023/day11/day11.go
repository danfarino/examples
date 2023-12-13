package day11

import (
	"strings"
)

type XY struct {
	X int
	Y int
}

func process(s string, expansion int) int {
	seenX := map[int]struct{}{}
	seenY := map[int]struct{}{}
	var galaxies []XY

	lines := strings.Split(strings.TrimSpace(s), "\n")

	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				galaxies = append(galaxies, XY{x, y})
				seenX[x] = struct{}{}
				seenY[y] = struct{}{}
			}
		}
	}

	var expandX []int
	var expandY []int

	for x := 0; x < len(lines[0]); x++ {
		if _, ok := seenX[x]; !ok {
			expandX = append(expandX, x)
		}
	}

	for y := 0; y < len(lines); y++ {
		if _, ok := seenY[y]; !ok {
			expandY = append(expandY, y)
		}
	}

	for i, galaxy := range galaxies {
		for _, x := range expandX {
			if x < galaxy.X {
				galaxies[i].X += expansion - 1
			} else {
				break
			}
		}

		for _, y := range expandY {
			if y < galaxy.Y {
				galaxies[i].Y += expansion - 1
			} else {
				break
			}
		}
	}

	distance := 0
	for i, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies[i+1:] {
			dx := galaxy1.X - galaxy2.X
			dy := galaxy1.Y - galaxy2.Y
			if dx < 0 {
				dx = -dx
			}
			if dy < 0 {
				dy = -dy
			}
			distance += dx + dy
		}
	}

	return distance
}
