package day2

import (
	"aoc2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	fmt.Println("Day 2")
	input, err := utils.LoadInput("day2/input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to load input: %w", err))
	}

	result, err := Part1(input)
	if err != nil {
		panic(fmt.Errorf("failed to run Part1: %w", err))
	}
	fmt.Printf("Part 1 Result: %d\n", result)

	result, err = Part2(input)
	if err != nil {
		panic(fmt.Errorf("failed to run Part2: %w", err))
	}
	fmt.Printf("Part 2 Result: %d\n", result)
}

type numRange struct {
	min int
	max int
}

func parseInput(input string) []numRange {
	if input == "" {
		return nil
	}

	parts := strings.Split(input, ",")
	ranges := make([]numRange, 0, len(parts))

	for _, p := range parts {

		p = strings.TrimSpace(p)

		if p == "" {
			continue
		}

		bounds := strings.Split(p, "-")

		if len(bounds) != 2 {
			continue
		}

		min, err1 := strconv.Atoi(strings.TrimSpace(bounds[0]))
		max, err2 := strconv.Atoi(strings.TrimSpace(bounds[1]))
		if err1 != nil || err2 != nil {
			fmt.Printf("Failed to parse range: %s\n", p)
			continue
		}

		if min > max {
			fmt.Printf("Invalid range: %s\n", p)
			continue
		}

		ranges = append(ranges, numRange{min: min, max: max})
	}

	return ranges
}

func hasRepeatingHalves(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)

	if length%2 != 0 {
		return false
	}

	mid := length / 2
	return str[:mid] == str[mid:]
}

func hasRepeatingPattern(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)

	for patternLen := 1; patternLen <= length/2; patternLen++ {
		if length%patternLen != 0 {
			continue
		}

		pattern := str[:patternLen]
		isRepeating := true

		for j := patternLen; j < length; j += patternLen {
			if str[j:j+patternLen] != pattern {
				isRepeating = false
				break
			}
		}

		if isRepeating {
			return true
		}
	}

	return false
}

func sumMatchingNumbers(ranges []numRange, predicate func(int) bool) int {
	sum := 0
	for _, r := range ranges {
		for i := r.min; i <= r.max; i++ {
			if predicate(i) {
				sum += i
			}
		}
	}
	return sum
}

func Part1(input string) (int, error) {
	ranges := parseInput(input)
	return sumMatchingNumbers(ranges, hasRepeatingHalves), nil
}

func Part2(input string) (int, error) {
	ranges := parseInput(input)
	return sumMatchingNumbers(ranges, hasRepeatingPattern), nil
}
