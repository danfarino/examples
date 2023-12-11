package day8

import (
	"regexp"
	"strings"
)

type Node struct {
	Left  string
	Right string
}

func processPart1(s string) int {
	seq, nodes := parse(s)
	steps := 0
	current := "AAA"
	for current != "ZZZ" {
		node := nodes[current]
		if seq[steps%len(seq)] == 'L' {
			current = node.Left
		} else {
			current = node.Right
		}
		steps++
	}

	return steps
}

func processPart2(s string) int {
	seq, nodes := parse(s)

	var zNodeSteps []int

	for nodeID := range nodes {
		if !strings.HasSuffix(nodeID, "A") {
			continue
		}

		for steps := 0; ; steps++ {
			if seq[steps%len(seq)] == 'L' {
				nodeID = nodes[nodeID].Left
			} else {
				nodeID = nodes[nodeID].Right
			}

			if strings.HasSuffix(nodeID, "Z") {
				zNodeSteps = append(zNodeSteps, steps+1)
				break
			}
		}
	}

	return LCM(zNodeSteps)
}

var nodeRegexp = regexp.MustCompile(`^(...) = \((...), (...)\)`)

func parse(s string) (string, map[string]Node) {
	lines := strings.Split(strings.TrimSpace(s), "\n")

	result := map[string]Node{}
	for i := 2; i < len(lines); i++ {
		m := nodeRegexp.FindStringSubmatch(lines[i])
		result[m[1]] = Node{
			Left:  m[2],
			Right: m[3],
		}
	}

	return lines[0], result
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(nums []int) int {
	a := nums[0]
	b := nums[1]
	result := a * b / GCD(a, b)

	nums = nums[2:]
	for i := 0; i < len(nums); i++ {
		result = LCM(append([]int{result}, nums...))
	}

	return result
}
