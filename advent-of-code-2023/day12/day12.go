package day12

import (
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func processFast(s string, unfold bool) int {
	records := parseInput(s)

	if unfold {
		unfoldRecords(records)
	}

	total := 0
	for _, record := range records {
		total += getCountFast(record)
	}
	return total
}

func processSlow(s string) int {
	records := parseInput(s)

	total := 0
	for _, record := range records {
		total += getCountSlow(record)
	}
	return total
}

type CacheKey struct {
	SizesLen  int
	LayoutLen int
}

func getCountFast(record Record) int {
	return countFastRecurse([]byte(record.Layout), record.Sizes, map[CacheKey]int{})
}

func countFastRecurse(layout []byte, sizes []int, cache map[CacheKey]int) int {
	if len(sizes) == 0 {
		for _, c := range layout {
			if c == '#' {
				return 0
			}
		}

		return 1
	}

	if len(layout) == 0 {
		return 0
	}

	minNeeded := lo.Sum(sizes) + len(sizes) - 1
	if minNeeded > len(layout) {
		return 0
	}

	if layout[0] == '.' {
		cacheKey := CacheKey{
			SizesLen:  len(sizes),
			LayoutLen: len(layout),
		}

		if cached, ok := cache[cacheKey]; ok {
			return cached
		}

		n := countFastRecurse(layout[1:], sizes, cache)

		cache[cacheKey] = n
		return n
	}

	springLen := 0

CountSprings:
	for i, c := range layout {
		switch c {
		case '.':
			break CountSprings
		case '#':
			springLen++
			if springLen > sizes[0] {
				return 0
			}
		case '?':
			layout[i] = '.'
			n1 := countFastRecurse(layout, sizes, cache)
			layout[i] = '#'
			n2 := countFastRecurse(layout, sizes, cache)
			layout[i] = '?'
			return n1 + n2
		}
	}

	if springLen == sizes[0] {
		return countFastRecurse(layout[springLen:], sizes[1:], cache)
	}

	return 0
}

func getCountSlow(record Record) int {
	count := 0
	springMask := 0
	emptyMask := 0

	for i := 0; i < len(record.Layout); i++ {
		switch record.Layout[len(record.Layout)-i-1] {
		case '#':
			springMask |= 1 << i
		case '.':
			emptyMask |= 1 << i
		}
	}

	for i := 1; i < 1<<len(record.Layout); i++ {
		if i&springMask != springMask || i&emptyMask != 0 {
			continue
		}

		n := i
		for n&1 == 0 {
			n >>= 1
		}

		var sizes []int
		for n != 0 {
			for n&1 == 0 {
				n >>= 1
			}

			ct := 0
			for n&1 == 1 {
				ct++
				n >>= 1
			}
			sizes = append(sizes, ct)
		}
		slices.Reverse(sizes)

		if slices.Equal(sizes, record.Sizes) {
			count++
		}
	}

	return count
}

func unfoldRecords(records []Record) {
	for i, record := range records {
		for n := 1; n < 5; n++ {
			records[i] = Record{
				Layout: records[i].Layout + "?" + record.Layout,
				Sizes:  append(records[i].Sizes, record.Sizes...),
			}
		}
	}
}

type Record struct {
	Layout string
	Sizes  []int
}

func parseInput(s string) []Record {
	var records []Record

	for _, line := range strings.Split(strings.TrimSpace(s), "\n") {
		layout, sizesStr, _ := strings.Cut(line, " ")

		record := Record{Layout: layout}

		for _, numStr := range strings.Split(sizesStr, ",") {
			record.Sizes = append(record.Sizes, mustParseInt(numStr))
		}

		records = append(records, record)
	}

	return records
}

func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
