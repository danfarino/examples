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
		var trace []string
		currentRule := "in"
		for currentRule != "A" && currentRule != "R" {
			trace = append(trace, currentRule)
			for _, rule := range workflows[currentRule] {
				if rule.Operator != "" {
					lhs := part.Get(rule.Lhs)
					switch rule.Operator {
					case "<":
						if lhs >= rule.Rhs {
							continue
						}
					case ">":
						if lhs <= rule.Rhs {
							continue
						}
					}
				}

				currentRule = rule.Target
				break
			}
		}

		trace = append(trace, currentRule)
		//fmt.Println(trace)

		if currentRule == "A" {
			total += part.X + part.M + part.A + part.S
		}
	}

	return total
}

type Part struct {
	X int
	M int
	A int
	S int
}

func (p Part) Get(letter string) int {
	switch letter {
	case "x":
		return p.X
	case "m":
		return p.M
	case "a":
		return p.A
	case "s":
		return p.S
	default:
		panic("bug")
	}
}

type Rule struct {
	Lhs      string
	Operator string
	Rhs      int
	Target   string
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
			var rhs int
			if len(m[3]) > 0 {
				rhs = shared.MustParseInt(m[3])
			}
			rules = append(rules, Rule{
				Lhs:      m[1],
				Operator: m[2],
				Rhs:      rhs,
				Target:   m[4],
			})
		}

		workflows[name] = rules
	}
	var parts []Part
	for scn.Scan() {
		var part Part
		_, err := fmt.Sscanf(scn.Text(), "{x=%d,m=%d,a=%d,s=%d}", &part.X, &part.M, &part.A, &part.S)
		if err != nil {
			panic(err)
		}
		parts = append(parts, part)
	}

	return workflows, parts
}
