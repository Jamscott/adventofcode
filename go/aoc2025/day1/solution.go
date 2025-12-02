package day1

import (
	"aoc2025/runner"
	"fmt"
	"strconv"
	"strings"
	"time"

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

func formatDuration(d time.Duration) string {
	if d < time.Microsecond {
		return fmt.Sprintf("%dns", d.Nanoseconds())
	} else if d < time.Millisecond {
		return fmt.Sprintf("%.2fµs", float64(d.Nanoseconds())/1000.0)
	} else if d < time.Second {
		return fmt.Sprintf("%.2fms", float64(d.Microseconds())/1000.0)
	}
	return d.String()
}

func Run() {
	input, err := runner.LoadInput("day1/input.txt")
	if err != nil {
		panic(fmt.Errorf("failed to load input: %w", err))
	}

	// Part 1
	start := time.Now()
	result, err := Part1(input)
	if err != nil {
		panic(fmt.Errorf("failed to run Part1: %w", err))
	}
	duration := time.Since(start)
	fmt.Printf("Part 1 Result: %d (took %s)\n", result, formatDuration(duration))

	// Part 2
	start = time.Now()
	result, err = Part2(input)
	if err != nil {
		panic(fmt.Errorf("failed to run Part2: %w", err))
	}
	duration = time.Since(start)
	fmt.Printf("Part 2 Result: %d (took %s)\n", result, formatDuration(duration))
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
