package day7

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func process(s string, jokers bool) int {
	hands := parseHands(s, jokers)

	sort.Slice(hands, func(i, j int) bool {
		h1 := hands[i]
		h2 := hands[j]

		if h1.Type < h2.Type {
			return true
		}

		if h1.Type == h2.Type {
			for k := 0; k < len(h1.Cards); k++ {
				if h1.Cards[k] < h2.Cards[k] {
					return true
				} else if h1.Cards[k] > h2.Cards[k] {
					return false
				}
			}
		}

		return false
	})

	winnings := 0
	for i := 0; i < len(hands); i++ {
		winnings += hands[i].Bid * (i + 1)
	}

	return winnings
}

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var part1CardOrder = "23456789TJQKA"
var part2CardOrder = "J23456789TQKA"

type Hand struct {
	Type  int
	Cards []int
	Bid   int
}

func parseHands(s string, jokers bool) []Hand {
	var results []Hand

	cardOrder := part1CardOrder
	if jokers {
		cardOrder = part2CardOrder
	}

	for _, line := range strings.Split(strings.TrimSpace(s), "\n") {
		var hand Hand

		var cardsStr string
		_, err := fmt.Sscanln(line, &cardsStr, &hand.Bid)
		if err != nil {
			panic(err)
		}

		cardCounts := map[int]int{}
		jokerCount := 0

		for _, r := range cardsStr {
			card := strings.IndexRune(cardOrder, r)
			hand.Cards = append(hand.Cards, card)
			if jokers && r == 'J' {
				jokerCount++
			} else {
				cardCounts[card]++
			}
		}

		var allCounts []int
		for _, count := range cardCounts {
			allCounts = append(allCounts, count)
		}

		if len(allCounts) > 0 {
			slices.Sort(allCounts)
			allCounts[len(allCounts)-1] += jokerCount
		}

		switch {
		case len(allCounts) == 0 || slices.Contains(allCounts, 5):
			hand.Type = FiveOfAKind
		case slices.Contains(allCounts, 4):
			hand.Type = FourOfAKind
		case slices.Contains(allCounts, 3):
			if slices.Contains(allCounts, 2) {
				hand.Type = FullHouse
			} else {
				hand.Type = ThreeOfAKind
			}
		default:
			pairCount := 0
			for _, n := range allCounts {
				if n == 2 {
					pairCount++
				}
			}
			if pairCount == 2 {
				hand.Type = TwoPair
			} else if pairCount == 1 {
				hand.Type = OnePair
			}
		}

		results = append(results, hand)
	}

	return results
}
