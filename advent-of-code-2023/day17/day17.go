package day17

import (
	"math"

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
	Point          Point
	XRun           int
	YRun           int
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
	pq.Push(PathInfo{0, start, 0, 0, Up | Down | Left | Right}, 0)

	for pq.Len() > 0 {
		queueItem := pq.Pop()

		cur := queueItem.Point

		for _, dir := range []Direction{Down, Right, Up, Left} {
			if queueItem.AvailableMoves&dir == 0 {
				continue
			}

			availableMoves := Up | Down | Left | Right
			xRun := queueItem.XRun
			yRun := queueItem.YRun

			neighbor := Point{cur.X, cur.Y}

			switch dir {
			case Up:
				if neighbor.Y == 0 {
					continue
				}
				neighbor.Y--
				availableMoves &^= Down
				if yRun < 0 {
					yRun--
				} else {
					yRun = -1
				}
				xRun = 0
			case Down:
				if neighbor.Y == ySize-1 {
					continue
				}
				neighbor.Y++
				availableMoves &^= Up
				if yRun > 0 {
					yRun++
				} else {
					yRun = 1
				}
				xRun = 0
			case Left:
				if neighbor.X == 0 {
					continue
				}
				neighbor.X--
				availableMoves &^= Right
				if xRun < 0 {
					xRun--
				} else {
					xRun = -1
				}
				yRun = 0
			case Right:
				if neighbor.X == xSize-1 {
					continue
				}
				neighbor.X++
				availableMoves &^= Left
				if xRun > 0 {
					xRun++
				} else {
					xRun = 1
				}
				yRun = 0
			default:
				panic("bug")
			}

			newCost := queueItem.Cost + grid.Get(neighbor)

			if newCost > cheapestGoalCost {
				continue
			}

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
					Point:          neighbor,
					XRun:           xRun,
					YRun:           yRun,
					AvailableMoves: availableMoves,
				}, goal.X-neighbor.X+goal.Y-neighbor.Y)
			}
		}
	}

	return cheapestGoalCost
}
