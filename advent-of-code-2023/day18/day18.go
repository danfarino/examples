package day18

import (
	"bytes"
	"cmp"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Span struct {
	X1 int
	X2 int
}

func (s Span) Len() int {
	return s.X2 - s.X1 + 1
}

type Spans []Span

func (s Spans) Union(span Span) Spans {
	if len(s) == 0 {
		return append(s, span)
	}

	s = append(s, span)
	slices.SortFunc(s, func(a, b Span) int {
		return cmp.Compare(a.X1, b.X1)
	})

	for i := 0; i+1 < len(s); {
		if s[i].X2+1 >= s[i+1].X1 {
			s[i].X2 = max(s[i+1].X2, span.X2)
			s = slices.Delete(s, i+1, i+2)
			continue
		}
		i++
	}

	return s
}

func (s Spans) Subtract(hole Span) Spans {
	if len(s) == 0 {
		return s
	}

	for i := 0; i < len(s); {
		if s[i].X2 < hole.X1 || s[i].X1 > hole.X2 {
			i++
			continue
		}

		if s[i].X1 >= hole.X1 {
			if s[i].X2 <= hole.X2 {
				s = slices.Delete(s, i, i+1)
				continue
			}

			s[i].X1 = hole.X2 + 1
		} else if s[i].X2 >= hole.X1 {
			if s[i].X2 > hole.X2 {
				s = slices.Insert(s, i+1, Span{hole.X2 + 1, s[i].X2})
			}

			s[i].X2 = hole.X1 - 1
		}

		i++
	}

	return s
}

func processPart1(s string) int {
	steps := parseInputPart1(s)
	return process(steps)
}

func processPart2(s string) int {
	steps := parseInputPart2(s)
	return process(steps)
}

func process(steps []DigStep) int {
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	points := []Point{{0, 0}}

	x := 0
	y := 0

	byY := map[int][]Span{}

	for _, step := range steps {
		switch step.Direction {
		case 'R':
			x += step.Distance
		case 'L':
			x -= step.Distance
		case 'D':
			y += step.Distance
		case 'U':
			y -= step.Distance
		}

		minX = min(minX, x)
		maxX = max(maxX, x)
		minY = min(minY, y)
		maxY = max(maxY, y)

		points = append(points, Point{x, y})
	}

	//g := NewGrid(minX, minY, maxX, maxY)
	for i := 0; i < len(points)-1; i++ {
		p1 := points[i]
		p2 := points[i+1]

		x1 := min(p1.X, p2.X)
		x2 := max(p1.X, p2.X)
		y1 := min(p1.Y, p2.Y)
		y2 := max(p1.Y, p2.Y)

		if y1 == y2 {
			byY[y1] = append(byY[y1], Span{x1, x2})
		}

		//for x := x1; x <= x2; x++ {
		//	for y := y1; y <= y2; y++ {
		//		g.Plow('#', x, y)
		//	}
		//}
	}

	ys := lo.Keys(byY)
	slices.Sort(ys)

	total := 0
	var opens Spans
	for i, y := range ys {
		if i > 0 {
			if len(opens) == 0 {
				panic("bad")
			}
			lineTotal := 0
			for _, open := range opens {
				lineTotal += open.Len()
			}

			prevY := ys[i-1]
			total += lineTotal * (y - prevY - 1)

			//for gapY := prevY + 1; gapY < y; gapY++ {
			//	for _, open := range opens {
			//		g.PlowSpan('$', gapY, Span{open.X1 + 1, open.X2 - 1})
			//	}
			//}
		}

		spans := byY[y]

		for _, span := range spans {
			total += span.Len()

			prevLen := len(opens)
			opens = slices.DeleteFunc(opens, func(open Span) bool {
				return open == span
			})
			if len(opens) != prevLen {
				continue
			}

			found := false
			for i, open := range opens {
				switch {
				case span.X1 == open.X1:
					opens[i].X1 = span.X2
					found = true
				case span.X2 == open.X1:
					opens[i].X1 = span.X1
					found = true
				case span.X2 == open.X2:
					opens[i].X2 = span.X1
					found = true
				case span.X1 == open.X2:
					opens[i].X2 = span.X2
					found = true
				}
			}
			if !found {
				for i, open := range opens {
					if span.X1 > open.X1 && span.X2 < open.X2 {
						opens[i].X2 = span.X1
						opens = append(opens, Span{span.X2, open.X2})
						found = true
						break
					}
				}
			}
			if !found {
				opens = append(opens, span)
			}
		}

		slices.SortFunc(opens, func(a, b Span) int {
			return cmp.Compare(a.X1, b.X1)
		})

		for i := 0; i+1 < len(opens); {
			s1 := opens[i]
			s2 := opens[i+1]

			if s1.X2 >= s2.X1 {
				opens[i].X2 = s2.X2
				opens = slices.Delete(opens, i+1, i+2)
				continue
			}

			i++
		}

		tmpOpens := slices.Clone(opens)
		for _, span := range spans {
			tmpOpens = tmpOpens.Subtract(span)
		}

		for _, open := range tmpOpens {
			//g.PlowSpan('*', y, open)
			total += open.Len()
		}
	}

	//g.Dump()

	return total
}

type Grid struct {
	cells  []byte
	minX   int
	minY   int
	stride int
}

func NewGrid(minX, minY, maxX, maxY int) *Grid {
	return &Grid{
		cells:  make([]byte, (maxX-minX+1)*(maxY-minY+1)),
		minX:   minX,
		minY:   minY,
		stride: maxX - minX + 1,
	}
}

func (g *Grid) PlowedCount() int {
	count := 0
	for _, c := range g.cells {
		if c == '#' {
			count++
		}
	}
	return count
}

func (g *Grid) Plow(c byte, x, y int) {
	g.cells[x-g.minX+(g.stride*(y-g.minY))] = c
}

func (g *Grid) PlowSpan(c byte, y int, span Span) {
	for x := span.X1; x <= span.X2; x++ {
		g.Plow(c, x, y)
	}
}

func (g *Grid) Get(x, y int) byte {
	c := g.cells[x-g.minX+(g.stride*(y-g.minY))]
	if c == 0 {
		return '.'
	}
	return c
}

func (g *Grid) Dump() {
	var sb bytes.Buffer
	y := g.minY
	for i := 0; i < len(g.cells); i += g.stride {
		sb.WriteString(fmt.Sprintf("% 4d", y))
		y++
		sb.WriteString("  ")
		for _, c := range g.cells[i : i+g.stride] {
			if c != 0 {
				sb.WriteByte(c)
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('\n')
	}
	_, _ = sb.WriteTo(os.Stdout)
}

type Point struct {
	X int
	Y int
}

type DigStep struct {
	Direction byte
	Distance  int
}

func parseInputPart1(s string) []DigStep {
	var steps []DigStep
	for _, line := range strings.Split(strings.TrimSpace(s), "\n") {
		var step DigStep
		_, err := fmt.Sscanf(line, "%c %d", &step.Direction, &step.Distance)
		if err != nil {
			panic(err)
		}
		steps = append(steps, step)
	}
	return steps
}

var part2Regexp = regexp.MustCompile(`#([a-f0-9]{5})([a-f0-9])`)

func parseInputPart2(s string) []DigStep {
	var steps []DigStep
	for _, line := range strings.Split(strings.TrimSpace(s), "\n") {
		m := part2Regexp.FindStringSubmatch(line)
		n, err := strconv.ParseInt(m[1], 16, strconv.IntSize)
		if err != nil {
			panic(err)
		}
		step := DigStep{
			Distance: int(n),
		}

		switch m[2] {
		case "0":
			step.Direction = 'R'
		case "1":
			step.Direction = 'D'
		case "2":
			step.Direction = 'L'
		case "3":
			step.Direction = 'U'
		default:
			panic("bad")
		}

		steps = append(steps, step)
	}
	return steps
}
