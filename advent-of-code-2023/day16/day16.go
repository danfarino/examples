package day16

import "strings"

type Point struct {
	X int
	Y int
}

type PointHeading struct {
	Point
	Heading Heading
}

func (ph PointHeading) Turn(h Heading) PointHeading {
	ph.Heading = h
	return ph
}

func (ph PointHeading) Advance() PointHeading {
	switch ph.Heading {
	case Up:
		ph.Y -= 1
	case Down:
		ph.Y += 1
	case Left:
		ph.X -= 1
	case Right:
		ph.X += 1
	}

	return ph
}

type Heading int

const (
	Up Heading = iota
	Down
	Left
	Right
)

func processPart1(s string) int {
	return calcEnergy(parseInput(s), 0, 0, Right)
}

func processPart2(s string) int {
	lines := parseInput(s)
	maxEnergy := 0

	for x := 0; x < len(lines[0]); x++ {
		maxEnergy = max(maxEnergy, calcEnergy(lines, x, 0, Down))
		maxEnergy = max(maxEnergy, calcEnergy(lines, x, len(lines)-1, Up))
	}

	for y := 0; y < len(lines); y++ {
		maxEnergy = max(maxEnergy, calcEnergy(lines, 0, y, Right))
		maxEnergy = max(maxEnergy, calcEnergy(lines, len(lines[0])-1, y, Left))
	}

	return maxEnergy
}

func parseInput(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

func calcEnergy(lines []string, x, y int, heading Heading) int {
	seenFrom := map[PointHeading]struct{}{}
	seenPoints := map[Point]struct{}{}
	pending := []PointHeading{{Point{x, y}, heading}}

	for len(pending) > 0 {
		ph := pending[len(pending)-1]
		pending = pending[:len(pending)-1]

		if ph.X < 0 || ph.Y < 0 || ph.X >= len(lines[0]) || ph.Y >= len(lines) {
			continue
		}

		if _, ok := seenFrom[ph]; ok {
			continue
		}

		seenFrom[ph] = struct{}{}
		seenPoints[ph.Point] = struct{}{}

		c := lines[ph.Y][ph.X]

		switch {
		case c == '|' && (ph.Heading == Left || ph.Heading == Right):
			pending = append(pending, ph.Turn(Up).Advance())
			pending = append(pending, ph.Turn(Down).Advance())
		case c == '-' && (ph.Heading == Up || ph.Heading == Down):
			pending = append(pending, ph.Turn(Left).Advance())
			pending = append(pending, ph.Turn(Right).Advance())
		case c == '/':
			switch ph.Heading {
			case Up:
				pending = append(pending, ph.Turn(Right).Advance())
			case Down:
				pending = append(pending, ph.Turn(Left).Advance())
			case Left:
				pending = append(pending, ph.Turn(Down).Advance())
			case Right:
				pending = append(pending, ph.Turn(Up).Advance())
			}
		case c == '\\':
			switch ph.Heading {
			case Down:
				pending = append(pending, ph.Turn(Right).Advance())
			case Up:
				pending = append(pending, ph.Turn(Left).Advance())
			case Right:
				pending = append(pending, ph.Turn(Down).Advance())
			case Left:
				pending = append(pending, ph.Turn(Up).Advance())
			}
		default:
			pending = append(pending, ph.Advance())
		}
	}

	return len(seenPoints)
}
