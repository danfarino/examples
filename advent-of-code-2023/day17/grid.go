package day17

import "strings"

type Grid []string

func NewGrid(s string) Grid {
	return strings.Split(strings.TrimSpace(s), "\n")
}

func (g Grid) Get(p Point) int {
	return int(g[p.Y][p.X] - '0')
}

func (g Grid) Size() (int, int) {
	return len(g[0]), len(g)
}
