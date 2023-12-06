package day6

import (
	"regexp"
	"strconv"
	"strings"
)

func processPart1(s string) int {
	result := 1

	for _, race := range parseInputPart1(s) {
		winCount := winCountForRace(race)
		result *= winCount
	}

	return result
}

func processPart2(s string) int {
	return winCountForRace(parseInputPart2(s))
}

func winCountForRace(race Race) int {
	winCount := 0
	for i := 1; i < race.Time; i++ {
		myDistance := i * (race.Time - i)
		if myDistance > race.Distance {
			winCount++
		}
	}
	return winCount
}

type Race struct {
	Time     int
	Distance int
}

func parseInputPart1(s string) []Race {
	lines := strings.Split(s, "\n")
	timeStrs := strings.Fields(strings.TrimPrefix(lines[0], "Time:"))
	distanceStrs := strings.Fields(strings.TrimPrefix(lines[1], "Distance:"))
	var races []Race
	for i := range timeStrs {
		races = append(races, Race{
			Time:     mustParseInt(timeStrs[i]),
			Distance: mustParseInt(distanceStrs[i]),
		})
	}
	return races
}

var nonDigitRegexp = regexp.MustCompile(`\D+`)

func parseInputPart2(s string) Race {
	lines := strings.Split(s, "\n")
	timeStr := nonDigitRegexp.ReplaceAllString(lines[0], "")
	distanceStr := nonDigitRegexp.ReplaceAllString(lines[1], "")
	return Race{
		Time:     mustParseInt(timeStr),
		Distance: mustParseInt(distanceStr),
	}
}

func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
