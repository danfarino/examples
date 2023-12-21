package day17

import (
	"math"
	"slices"

	"github.com/danfarino/examples/advent-of-code-2023/priorityqueue"
)

func processPart1(s string) int {
	grid := NewGrid(s)
	return getShortestPathCost(grid, 1, 3)
}

func processPart2(s string) int {
	grid := NewGrid(s)
	return getShortestPathCost(grid, 4, 10)
}

type Direction int

const (
	Up Direction = 1 << iota
	Down
	Left
	Right
)

type Point struct {
	X int
	Y int
}

type PathInfo struct {
	Cost           int
	Path           []Point
	AvailableMoves Direction
}

type CheapestKey struct {
	Point
	XRun int
	YRun int
}

func getShortestPathCost(grid Grid, minMove, maxRun int) int {
	// adapted from: https://www.redblobgames.com/pathfinding/a-star/introduction.html

	xSize, ySize := grid.Size()
	start := Point{0, 0}
	goal := Point{xSize - 1, ySize - 1}

	cheapestGoalCost := math.MaxInt
	cheapest := map[CheapestKey]int{}
	pq := priorityqueue.New[PathInfo]()
	pq.Push(PathInfo{0, []Point{start}, Up | Down | Left | Right}, 0)

	for pq.Len() > 0 {
		pendingPoint := pq.Pop()

		cur := pendingPoint.Path[len(pendingPoint.Path)-1]

		for _, dir := range []Direction{Down, Right, Up, Left} {
			if pendingPoint.AvailableMoves&dir == 0 {
				continue
			}

			availableMoves := Up | Down | Left | Right

			neighbor := Point{cur.X, cur.Y}
			switch dir {
			case Up:
				if neighbor.Y == 0 {
					continue
				}
				neighbor.Y--
				availableMoves &^= Down
			case Down:
				if neighbor.Y == ySize-1 {
					continue
				}
				neighbor.Y++
				availableMoves &^= Up
			case Left:
				if neighbor.X == 0 {
					continue
				}
				neighbor.X--
				availableMoves &^= Right
			case Right:
				if neighbor.X == xSize-1 {
					continue
				}
				neighbor.X++
				availableMoves &^= Left
			default:
				panic("bug")
			}

			newCost := pendingPoint.Cost + grid.Get(neighbor)

			if newCost > cheapestGoalCost {
				continue
			}

			neighborPath := append(slices.Clone(pendingPoint.Path), neighbor)

			xRun, yRun := getRuns(neighborPath, maxRun)
			switch {
			case xRun == maxRun:
				availableMoves &^= Right
			case xRun == -maxRun:
				availableMoves &^= Left
			case yRun == maxRun:
				availableMoves &^= Down
			case yRun == -maxRun:
				availableMoves &^= Up
			}

			switch {
			case xRun > 0 && xRun < minMove:
				availableMoves = Right
			case xRun < 0 && -xRun < minMove:
				availableMoves = Left
			case yRun > 0 && yRun < minMove:
				availableMoves = Down
			case yRun < 0 && -yRun < minMove:
				availableMoves = Up
			}

			if neighbor == goal {
				cheapestGoalCost = newCost
				continue
			}

			if availableMoves == 0 {
				continue
			}

			key := CheapestKey{neighbor, xRun, yRun}
			if prevCost, ok := cheapest[key]; !ok || newCost < prevCost {
				cheapest[key] = newCost
				pq.Push(PathInfo{
					Cost:           newCost,
					Path:           neighborPath,
					AvailableMoves: availableMoves,
				}, goal.X-neighbor.X+goal.Y-neighbor.Y)
			}
		}
	}

	return cheapestGoalCost
}

func getRuns(path []Point, maxRun int) (int, int) {
	xRun := 0
	yRun := 0

	for i := 0; i < maxRun && i < len(path)-1; i++ {
		p1 := path[len(path)-1-i]
		p2 := path[len(path)-2-i]

		if p1.X != p2.X {
			if yRun != 0 {
				break
			}
			if p1.X > p2.X {
				xRun++
			} else {
				xRun--
			}
		} else {
			if xRun != 0 {
				break
			}
			if p1.Y > p2.Y {
				yRun++
			} else {
				yRun--
			}
		}
	}

	return xRun, yRun
}
