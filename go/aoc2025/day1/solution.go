package day1

import (
	"aoc2025/solver"
	"aoc2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	solver.Register(1, Solution{})
}

type Solution struct{}

func (s Solution) Part1(input string) (int, error) {
	lines := utils.NewLineSplit(input)

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

		position = utils.Mod(position+distance, 100)

		if position == 0 {
			count++
		}
	}

	return count, nil
}

func (s Solution) Part2(input string) (int, error) {
	lines := utils.NewLineSplit(input)

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
