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
			pending = append(pending, advance(turn(ph, Up)))
			pending = append(pending, advance(turn(ph, Down)))
		case c == '-' && (ph.Heading == Up || ph.Heading == Down):
			pending = append(pending, advance(turn(ph, Left)))
			pending = append(pending, advance(turn(ph, Right)))
		case c == '/':
			switch ph.Heading {
			case Up:
				pending = append(pending, advance(turn(ph, Right)))
			case Down:
				pending = append(pending, advance(turn(ph, Left)))
			case Left:
				pending = append(pending, advance(turn(ph, Down)))
			case Right:
				pending = append(pending, advance(turn(ph, Up)))
			}
		case c == '\\':
			switch ph.Heading {
			case Down:
				pending = append(pending, advance(turn(ph, Right)))
			case Up:
				pending = append(pending, advance(turn(ph, Left)))
			case Right:
				pending = append(pending, advance(turn(ph, Down)))
			case Left:
				pending = append(pending, advance(turn(ph, Up)))
			}
		default:
			pending = append(pending, advance(ph))
		}
	}

	return len(seenPoints)
}

func turn(ph PointHeading, heading Heading) PointHeading {
	ph.Heading = heading
	return ph
}

func advance(ph PointHeading) PointHeading {
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
