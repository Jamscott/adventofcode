package day5

import (
	"aoc2025/solver"
	"aoc2025/utils"
	"strconv"
	"strings"
)

func init() {
	solver.Register(5, Solution{})
}

type Solution struct{}

func parseInput(input string) (utils.RangeSet[int64], []int64) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	parsingRanges := true
	var ranges utils.RangeSet[int64]
	var items []int64

	for _, line := range lines {
		if line = strings.TrimSpace(line); line == "" {
			parsingRanges = false
			continue
		}

		if parsingRanges {
			if r, err := utils.ParseInt64Range(line); err == nil {
				ranges = append(ranges, r)
			}
		} else {
			if n, err := strconv.ParseInt(line, 10, 64); err == nil {
				items = append(items, n)
			}
		}
	}

	ranges.Sort()
	return ranges, items
}

func (s Solution) Part1(input string) (int, error) {
	ranges, items := parseInput(input)
	count := 0
	for _, item := range items {
		if ranges.Contains(item) {
			count++
		}
	}
	return count, nil
}

func (s Solution) Part2(input string) (int, error) {
	ranges, _ := parseInput(input)
	return int(mergeTotalLength(ranges)), nil
}

func mergeTotalLength(ranges utils.RangeSet[int64]) int64 {
	if len(ranges) == 0 {
		return 0
	}

	var total, prevEnd int64
	for i, r := range ranges {
		if i == 0 || prevEnd < r.Start {
			total += r.End - r.Start + 1
			prevEnd = r.End
		} else if prevEnd < r.End {
			total += r.End - prevEnd
			prevEnd = r.End
		}
	}
	return total
}
