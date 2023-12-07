package day5

import (
	"bufio"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func process(s string) (int, int) {
	almanac := parseInput(s)

	part1Lowest := math.MaxInt

	for _, num := range almanac.Seeds {
		loc, _ := locForNum(almanac, num, 1)
		part1Lowest = min(part1Lowest, loc)
	}

	part2Lowest := math.MaxInt

	for i := 0; i < len(almanac.Seeds); i += 2 {
		rangeStart := almanac.Seeds[i]
		rangePastEnd := rangeStart + almanac.Seeds[i+1]

		for num := rangeStart; num < rangePastEnd; {
			loc, consumedCount := locForNum(almanac, num, rangePastEnd-num)
			num += consumedCount
			part2Lowest = min(part2Lowest, loc)
		}
	}

	return part1Lowest, part2Lowest
}

func locForNum(almanac Almanac, num, remainingCount int) (int, int) {
	for name := "seed"; name != "location"; {
		mapping := almanac.MappingFrom[name]
		name = mapping.To

		for _, numRange := range mapping.NumRanges {
			offset := num - numRange.SourceStart
			if offset >= 0 && offset < numRange.Count {
				remainingCount = min(remainingCount, numRange.Count-offset)
				num = numRange.DestStart + offset
				break
			}
		}
	}

	return num, remainingCount
}

type Almanac struct {
	Seeds       []int
	MappingFrom map[string]Mapping
}

type Mapping struct {
	To        string
	NumRanges []NumRange
}

type NumRange struct {
	SourceStart int
	DestStart   int
	Count       int
}

var mapHeaderRegexp = regexp.MustCompile(`^(.+)-to-(.+) map:$`)

func parseInput(s string) Almanac {
	almanac := Almanac{MappingFrom: map[string]Mapping{}}

	scn := bufio.NewScanner(strings.NewReader(s))
	for scn.Scan() {
		if numsStr, found := strings.CutPrefix(scn.Text(), "seeds: "); found {
			for _, numStr := range strings.Fields(numsStr) {
				almanac.Seeds = append(almanac.Seeds, mustParseInt(numStr))
			}
			continue
		}

		if m := mapHeaderRegexp.FindStringSubmatch(scn.Text()); m != nil {
			fromName := m[1]

			mapping := Mapping{To: m[2]}

			for scn.Scan() {
				if scn.Text() == "" {
					break
				}

				numStrs := strings.Fields(scn.Text())
				mapping.NumRanges = append(mapping.NumRanges, NumRange{
					SourceStart: mustParseInt(numStrs[1]),
					DestStart:   mustParseInt(numStrs[0]),
					Count:       mustParseInt(numStrs[2]),
				})
			}

			sort.Slice(mapping.NumRanges, func(i, j int) bool {
				return mapping.NumRanges[i].SourceStart < mapping.NumRanges[j].SourceStart
			})

			almanac.MappingFrom[fromName] = mapping
			continue
		}
	}

	return almanac
}

func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
