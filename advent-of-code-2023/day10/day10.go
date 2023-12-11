package day10

import (
	"strings"
)

type Point struct {
	X int
	Y int
}

type PointSet = map[Point]struct{}

func process(s string) (int, int) {
	lines := strings.Split(strings.TrimSpace(s), "\n")

	curvePoints := findCurvePoints(lines)
	insideCount := countInsidePoints(lines, curvePoints)

	return len(curvePoints) / 2, insideCount
}

func findCurvePoints(lines []string) PointSet {
	var startPos Point
	for i, line := range lines {
		sIndex := strings.IndexByte(line, 'S')
		if sIndex >= 0 {
			startPos = Point{sIndex, i}
			break
		}
	}

	curvePoints := PointSet{}
	curvePoints[startPos] = struct{}{}

	curPos := startPos
	nextPos := findFirstMove(lines, curPos)

	for {
		curvePoints[nextPos] = struct{}{}
		nextCurPos := nextPos

		switch lines[nextPos.Y][nextPos.X] {
		case '-':
			nextPos = Point{2*nextPos.X - curPos.X, curPos.Y}
		case '|':
			nextPos = Point{curPos.X, 2*nextPos.Y - curPos.Y}
		case 'F':
			if curPos.X == nextPos.X {
				nextPos = Point{curPos.X + 1, curPos.Y - 1}
			} else {
				nextPos = Point{curPos.X - 1, curPos.Y + 1}
			}
		case '7':
			if curPos.X == nextPos.X {
				nextPos = Point{curPos.X - 1, curPos.Y - 1}
			} else {
				nextPos = Point{curPos.X + 1, curPos.Y + 1}
			}
		case 'J':
			if curPos.X == nextPos.X {
				nextPos = Point{curPos.X - 1, curPos.Y + 1}
			} else {
				nextPos = Point{curPos.X + 1, curPos.Y - 1}
			}
		case 'L':
			if curPos.X == nextPos.X {
				nextPos = Point{curPos.X + 1, curPos.Y + 1}
			} else {
				nextPos = Point{curPos.X - 1, curPos.Y - 1}
			}
		default:
			panic("invalid character")
		}

		curPos = nextCurPos

		if nextPos == startPos {
			break
		}
	}

	return curvePoints
}

func countInsidePoints(lines []string, curvePoints PointSet) int {
	isOnPath := func(x, y int) bool {
		_, ok := curvePoints[Point{x, y}]
		return ok
	}

	insideCount := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			if isOnPath(x, y) {
				continue
			}

			delta := -1
			// If we see 'S' to the left of us, we'll search to the right instead (so that we don't have to understand
			// what connector 'S' is masquerading as). Also, search to the right if it's shorter.
			if x > (len(lines[y])/2) || strings.IndexByte(lines[y][0:x], 'S') >= 0 {
				delta = 1
			}

			crossCount := 0 // how many times did we cross an inside/outside boundary
			cornerY := 0    // for corners, keep track of the direction (so we can determine if the next corner is considered a cross)
			for x2 := x + delta; x2 >= 0 && x2 < len(lines[y]); x2 += delta {
				if isOnPath(x2, y) {
					switch lines[y][x2] {
					case '|':
						crossCount++
					case 'F', '7':
						if cornerY == 0 {
							cornerY = 1
						} else {
							if cornerY == -1 {
								crossCount++
							}
							cornerY = 0
						}
					case 'L', 'J':
						if cornerY == 0 {
							cornerY = -1
						} else {
							if cornerY == 1 {
								crossCount++
							}
							cornerY = 0
						}
					}
				}
			}

			if crossCount%2 == 1 {
				insideCount++
			}
		}
	}

	return insideCount
}

func findFirstMove(lines []string, curPos Point) Point {
	offsets := []int{-1, 1}

	for _, dx := range offsets {
		for dy := range offsets {
			nextPos := Point{curPos.X + dx, curPos.Y + dy}

			if nextPos.X < 0 || nextPos.Y < 0 || nextPos.X == len(lines[0]) || nextPos.Y == len(lines) {
				continue
			}

			c := lines[nextPos.Y][nextPos.X]
			switch {
			case c == '|' && dx == 0:
				return nextPos
			case c == '-' && dy == 0:
				return nextPos
			case c == 'F' && ((dx == -1 && dy == 0) || dx == 0 && dy == -1):
				return nextPos
			case c == '7' && ((dx == 1 && dy == 0) || (dx == 0 && dy == -1)):
				return nextPos
			case c == 'J' && ((dx == 0 && dy == 1) || (dx == 1 && dy == 0)):
				return nextPos
			case c == 'L' && ((dx == 0 && dy == 1) || (dx == -1 && dy == 0)):
				return nextPos
			}
		}
	}

	panic("failed to find first move")
}
