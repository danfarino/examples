package day19

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/danfarino/examples/advent-of-code-2023/shared"
)

func processPart1(s string) int {
	workflows, parts := parseInput(s)

	total := 0

	for _, part := range parts {
		currentWorkflow := "in"
		for currentWorkflow != "A" && currentWorkflow != "R" {
			for _, rule := range workflows[currentWorkflow] {
				if rule.Matches(part) {
					currentWorkflow = rule.Target
					break
				}
			}
		}

		if currentWorkflow == "A" {
			for _, value := range part {
				total += value
			}
		}
	}

	return total
}

func processPart2(s string) int {
	workflows, _ := parseInput(s)
	partRanges := [4]Range{
		{1, 4000},
		{1, 4000},
		{1, 4000},
		{1, 4000},
	}
	return recursePart2(workflows, "in", partRanges)
}

func letterIndex(letter byte) int {
	return strings.IndexByte("xmas", letter)
}

func recursePart2(workflows map[string][]Rule, name string, partRanges [4]Range) int {
	if name == "R" {
		return 0
	}

	if name == "A" {
		total := 1
		for _, letterRange := range partRanges {
			total *= letterRange.Max - letterRange.Min + 1
		}
		return total
	}

	total := 0
	for _, rule := range workflows[name] {
		if rule.Letter == 0 {
			total += recursePart2(workflows, rule.Target, partRanges)
			continue
		}

		matchedPartRanges := partRanges
		matchedLetterRange := &matchedPartRanges[letterIndex(rule.Letter)]

		unmatchedPartRanges := partRanges
		unmatchedLetterRange := &unmatchedPartRanges[letterIndex(rule.Letter)]

		switch rule.Operator {
		case "<":
			matchedLetterRange.Max = min(matchedLetterRange.Max, rule.Value-1)
			unmatchedLetterRange.Min = max(unmatchedLetterRange.Min, rule.Value)
		case ">":
			matchedLetterRange.Min = max(matchedLetterRange.Min, rule.Value+1)
			unmatchedLetterRange.Max = min(unmatchedLetterRange.Max, rule.Value)
		}

		if matchedLetterRange.Min <= matchedLetterRange.Max {
			total += recursePart2(workflows, rule.Target, matchedPartRanges)
			partRanges = unmatchedPartRanges
		}
	}

	return total
}

type Range struct {
	Min int
	Max int
}

type Part [4]int

type Rule struct {
	Letter   byte
	Operator string
	Value    int
	Target   string
}

func (r Rule) Matches(part Part) bool {
	if r.Operator != "" {
		letterValue := part[letterIndex(r.Letter)]
		switch r.Operator {
		case "<":
			if letterValue >= r.Value {
				return false
			}
		case ">":
			if letterValue <= r.Value {
				return false
			}
		}
	}

	return true
}

var workflowRuleRegexp = regexp.MustCompile(`[{,](?:([xmas])([<>])(\d+):)?([^,}]+)`)

func parseInput(s string) (map[string][]Rule, []Part) {
	scn := bufio.NewScanner(strings.NewReader(s))
	workflows := map[string][]Rule{}
	for scn.Scan() {
		line := scn.Text()
		if line == "" {
			break
		}

		name, _, _ := strings.Cut(line, "{")
		var rules []Rule
		for _, m := range workflowRuleRegexp.FindAllStringSubmatch(line, -1) {
			var char byte
			if len(m[1]) > 0 {
				char = m[1][0]
			}

			var value int
			if len(m[3]) > 0 {
				value = shared.MustParseInt(m[3])
			}
			rules = append(rules, Rule{
				Letter:   char,
				Operator: m[2],
				Value:    value,
				Target:   m[4],
			})
		}

		workflows[name] = rules
	}
	var parts []Part
	for scn.Scan() {
		var part Part
		_, err := fmt.Sscanf(scn.Text(), "{x=%d,m=%d,a=%d,s=%d}", &part[0], &part[1], &part[2], &part[3])
		if err != nil {
			panic(err)
		}
		parts = append(parts, part)
	}

	return workflows, parts
}
