package day3

import (
	"aoc2025/solver"
	"aoc2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	solver.Register(3, Solution{})
}

type Solution struct{}

const (
	part1DigitCount = 2
	part2DigitCount = 12
)

func (s Solution) Part1(input string) (int, error) {
	return s.solve(input, part1DigitCount)
}

func (s Solution) Part2(input string) (int, error) {
	return s.solve(input, part2DigitCount)
}

// [ "987654321111111" ]
// [ "811111111111119" ]
// [ "234234234234278" ]
// [ "818181911112111" ]
func parseInput(input string) []string {
	lines := utils.NewLineSplit(input)
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		if trimmed := strings.TrimSpace(line); trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func (s Solution) solve(input string, digitCount int) (int, error) {
	lines := parseInput(input)
	total := 0

	for _, line := range lines {
		output := buildMaxNumberFromDigits(line, digitCount)

		joltage, err := strconv.Atoi(output)
		if err != nil {
			return 0, fmt.Errorf("failed to parse joltage: %w", err)
		}

		total += joltage
	}

	return total, nil
}

func findLargestDigitInRange(line string, from, to int) (digit int, position int) {
	max := -1
	index := -1
	end := len(line) - to

	if end <= from {
		return max, index
	}

	for i := from; i < end; i++ {

		if line[i] < '0' || line[i] > '9' {
			continue
		}

		num := int(line[i] - '0')

		if num > max {
			max = num
			index = i
		}

	}
	return max, index
}

func buildMaxNumberFromDigits(line string, numDigits int) string {
	lastIndex := -1
	var joltage strings.Builder
	joltage.Grow(numDigits)

	for i := 1; i <= numDigits; i++ {
		value, index := findLargestDigitInRange(line, lastIndex+1, numDigits-i)
		joltage.WriteByte(byte(value + '0'))
		lastIndex = index
	}

	return joltage.String()
}
