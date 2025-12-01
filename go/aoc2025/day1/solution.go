package day1

import (
	"aoc2025/runner"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func Mod[T constraints.Integer](n, m T) T {
	if m == 0 {
		panic("modulo by zero")
	}

	r := n % m

	if r < 0 {
		r += m
	}

	return r
}

func Run() {
	input, err := runner.LoadInput("day1/input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to load input: %w", err))
	}

	result, err := Part1(input)
	if err != nil {
		panic(fmt.Errorf("failed to run Part1: %w", err))
	}

	fmt.Println("Part 1 Result:", result)

	result, err = Part2(input)
	if err != nil {
		panic(fmt.Errorf("failed to run Part2: %w", err))
	}

	fmt.Println("Part 2 Result:", result)
}

func Part1(input string) (int, error) {
	lines := runner.NewLineSplit(input)

	position := 50
	count := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, fmt.Errorf("failed to convert distance to int: %w", err)
		}

		if line[0] == 'L' {
			distance *= -1
		}

		position = Mod(position+distance, 100)

		if position == 0 {
			count++
		}
	}

	return count, nil
}

func Part2(input string) (int, error) {
	lines := runner.NewLineSplit(input)

	position := 50
	count := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, fmt.Errorf("failed to convert distance to int: %w", err)
		}

		for range distance {
			if line[0] == 'L' {
				if position == 0 {
					position = 99
				} else {
					position -= 1
				}
			} else {
				if position == 99 {
					position = 0
				} else {
					position += 1
				}
			}

			if position == 0 {
				count++
			}
		}
	}

	return count, nil
}
