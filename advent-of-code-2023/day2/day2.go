package day2

import (
	"strconv"
	"strings"
)

type Game struct {
	Num     int
	Subsets []map[string]int
}

func parseGames(s string) []Game {
	var games []Game

	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}

		before, after, found := strings.Cut(line, ": ")
		if !found {
			panic(line)
		}

		gameNum, err := strconv.Atoi(strings.TrimPrefix(before, "Game "))
		if err != nil {
			panic(line)
		}

		game := Game{Num: gameNum}

		for _, subset := range strings.Split(after, "; ") {
			subsetMap := map[string]int{}

			for _, forColor := range strings.Split(subset, ", ") {
				colorCountStr, color, found := strings.Cut(forColor, " ")
				if !found {
					panic(forColor)
				}

				colorCount, err := strconv.Atoi(colorCountStr)
				if err != nil {
					panic(colorCountStr)
				}

				subsetMap[color] = colorCount
			}

			game.Subsets = append(game.Subsets, subsetMap)
		}

		games = append(games, game)
	}

	return games
}

func processPart1(s string, bag map[string]int) int {
	result := 0

GAME:
	for _, game := range parseGames(s) {
		for _, subset := range game.Subsets {
			for color, colorCount := range subset {
				if bag[color] < colorCount {
					continue GAME
				}
			}
		}

		result += game.Num
	}

	return result
}

func processPart2(s string) int {
	totalPower := 0

	for _, game := range parseGames(s) {
		maxByColor := map[string]int{}

		for _, subset := range game.Subsets {
			for color, colorCount := range subset {
				if curMax := maxByColor[color]; curMax < colorCount {
					maxByColor[color] = colorCount
				}
			}
		}

		gamePower := 1
		for _, curMax := range maxByColor {
			gamePower *= curMax
		}

		totalPower += gamePower
	}

	return totalPower
}
