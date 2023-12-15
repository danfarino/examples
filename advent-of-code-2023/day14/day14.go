package day14

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"strings"
)

func processPart1(s string) int {
	grid := parseGrid(s)
	tilt(grid, 0)
	return getWeight(grid)
}

func processPart2(s string, tiltCycleCount int) int {
	grid := parseGrid(s)

	seen := map[string]int{}
	var seenWeights []int

	for loopCycleLen := 0; loopCycleLen < tiltCycleCount; loopCycleLen++ {
		doTiltCycle(grid)

		hash := hashGrid(grid)
		if loopCycleStart, ok := seen[hash]; ok {
			loopCycleLen -= loopCycleStart
			return seenWeights[loopCycleStart+(tiltCycleCount-loopCycleStart-1)%loopCycleLen]
		}

		seen[hash] = loopCycleLen
		seenWeights = append(seenWeights, getWeight(grid))
	}

	panic("BAD")
}

func parseGrid(s string) [][]byte {
	return bytes.Split([]byte(strings.TrimSpace(s)), []byte("\n"))
}

func getWeight(grid [][]byte) int {
	gridSize := len(grid[0])
	total := 0
	for y, line := range grid {
		for _, c := range line {
			if c == 'O' {
				total += gridSize - y
			}
		}
	}
	return total
}

func hashGrid(grid [][]byte) string {
	h := sha1.New()
	for _, line := range grid {
		h.Write(line)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func doTiltCycle(grid [][]byte) {
	tilt(grid, 0)
	tilt(grid, 1)
	tilt(grid, 2)
	tilt(grid, 3)
}

func tilt(grid [][]byte, dir int) {
	gridSize := len(grid[0])

	for x := 0; x < gridSize; x++ {
		rollTo := 0

		for y := 0; y < gridSize; y++ {
			x2, y2 := translatePoint(dir, gridSize, x, y)

			switch grid[y2][x2] {
			case 'O':
				if rollTo < y {
					newX, newY := translatePoint(dir, gridSize, x, rollTo)
					grid[newY][newX] = 'O'
					grid[y2][x2] = '.'
				}
				rollTo++
			case '#':
				rollTo = y + 1
			}
		}
	}
}

func translatePoint(dir, gridSize, x, y int) (int, int) {
	switch dir {
	case 0:
		return x, y
	case 1:
		return y, gridSize - x - 1
	case 2:
		return gridSize - x - 1, gridSize - y - 1
	case 3:
		return gridSize - y - 1, x
	}

	panic("BAD")
}

func printGrid(grid [][]byte) {
	gridSize := len(grid)
	var sb strings.Builder
	for y := 0; y < gridSize; y++ {
		sb.WriteString(fmt.Sprintf("% 3d  ", y+1))
		for x := 0; x < gridSize; x++ {
			c := grid[y][x]
			sb.WriteByte(c)
			sb.WriteByte(' ')
		}
		sb.WriteByte('\n')
	}
	fmt.Print(sb.String())
}
