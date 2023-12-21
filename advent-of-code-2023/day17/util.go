package day17

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/TwiN/go-color"
)

func renderPath(grid Grid, path []Point) {
	pathPoints := map[Point]byte{}
	for i, cur := range path {
		ch := grid.Get(cur) + '0'
		if i > 0 {
			prev := path[i-1]
			switch {
			case cur.X > prev.X:
				ch = '>'
			case cur.X < prev.X:
				ch = '<'
			case cur.Y > prev.Y:
				ch = 'V'
			case cur.Y < prev.Y:
				ch = '^'
			default:
				panic("bug")
			}
		}
		pathPoints[cur] = byte(ch)
	}

	xSize, ySize := grid.Size()
	var sb strings.Builder
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			if ch, ok := pathPoints[Point{x, y}]; ok {
				sb.WriteString(color.CyanBackground)
				sb.WriteByte(ch)
			} else {
				sb.WriteString(strconv.Itoa(grid.Get(Point{x, y})))
			}
			sb.WriteString(color.Reset)
		}
		sb.WriteByte('\n')
	}

	fmt.Print(sb.String())
}
