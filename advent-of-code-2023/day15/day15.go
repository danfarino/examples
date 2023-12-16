package day15

import (
	"regexp"
	"strconv"
	"strings"
)

func processPart1(s string) int {
	total := 0

	for _, part := range parseInput(s) {
		total += hashString(part)
	}

	return total
}

var partRegexp = regexp.MustCompile(`^([a-z]+)(=(\d+)|-)$`)

type Lens struct {
	Label       string
	FocalLength int
}

func processPart2(s string) int {
	parts := parseInput(s)

	boxes := map[int][]Lens{}

	for _, part := range parts {
		m := partRegexp.FindStringSubmatch(part)
		if m == nil {
			panic("BAD")
		}

		lensLabel := m[1]
		afterLabel := m[2]
		labelHash := hashString(lensLabel)
		focalLength := 0

		if afterLabel != "-" {
			focalLength = mustParseInt(m[3])
		}

		found := false
		for i, lens := range boxes[labelHash] {
			if lens.Label == lensLabel {
				if afterLabel == "-" {
					boxes[labelHash][i] = Lens{}
				} else {
					boxes[labelHash][i] = Lens{
						Label:       lensLabel,
						FocalLength: focalLength,
					}
					found = true
				}
				break
			}
		}

		if afterLabel != "-" && !found {
			boxes[labelHash] = append(boxes[labelHash], Lens{
				Label:       lensLabel,
				FocalLength: focalLength,
			})
		}
	}

	total := 0
	for boxNum, lenses := range boxes {
		n := 0
		for _, lens := range lenses {
			if len(lens.Label) > 0 {
				n++
				total += (boxNum + 1) * n * lens.FocalLength
			}
		}
	}

	return total
}

func parseInput(s string) []string {
	return strings.Split(strings.TrimSpace(s), ",")
}

func hashString(s string) int {
	val := 0
	for _, c := range s {
		val += int(c)
		val *= 17
		val %= 256
	}
	return val
}

func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
