package day12

import (
	"aoc2025/solver"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	solver.Register(12, Solution{})
}

type Solution struct{}

type testCase struct {
	width    int
	height   int
	patterns []int
}

var pattern = regexp.MustCompile(`^(\d+)x(\d+): (\d+) (\d+) (\d+) (\d+) (\d+) (\d+)$`)

func parseInput(input string) []testCase {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Reverse the lines and take while not empty
	var testLines []string
	for i := len(lines) - 1; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			break
		}
		testLines = append(testLines, line)
	}

	var cases []testCase
	for _, line := range testLines {
		matches := pattern.FindStringSubmatch(line)
		if matches == nil {
			continue
		}

		width, _ := strconv.Atoi(matches[1])
		height, _ := strconv.Atoi(matches[2])

		patterns := make([]int, 0, 6)
		for i := 3; i <= 8; i++ {
			val, _ := strconv.Atoi(matches[i])
			patterns = append(patterns, val)
		}

		cases = append(cases, testCase{
			width:    width,
			height:   height,
			patterns: patterns,
		})
	}

	return cases
}

func (s Solution) Part1(input string) (int, error) {
	cases := parseInput(input)
	count := 0

	for _, tc := range cases {
		area := tc.width * tc.height

		required := 0
		for _, p := range tc.patterns {
			required += p
		}

		if area >= required*(3*3) {
			count++
		}
	}

	return count, nil
}

func (s Solution) Part2(input string) (int, error) {
	return -1, nil
}
