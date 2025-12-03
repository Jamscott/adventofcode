package day3

import (
	"aoc2025/solver"
	"aoc2025/utils"
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
		line = strings.TrimSpace(line)
		if line != "" {
			result = append(result, line)
		}
	}

	return result
}

func (s Solution) solve(input string, digitCount int) (int, error) {
	lines := parseInput(input)
	total := 0

	for _, line := range lines {
		total += buildMaxNumberFromDigits(line, digitCount)
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

func buildMaxNumberFromDigits(line string, numDigits int) int {
	result := 0
	lastIndex := -1

	for digitsRemaining := numDigits; digitsRemaining > 0; digitsRemaining-- {
		digit, index := findLargestDigitInRange(line, lastIndex+1, digitsRemaining-1)

		if digit == -1 {
			digit = 0
		}

		result = result*10 + digit
		lastIndex = index
	}

	return result
}
