package template

import (
	"fmt"
	"os"
	"strings"
)

// loadInput loads the input file from the current day's directory
func loadInput(filename string) (string, error) {
	data, err := os.ReadFile("template/" + filename)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", filename, err)
	}
	return string(data), nil
}

// Part1 solves part 1 of the puzzle
func Part1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// TODO: Implement your solution here
	fmt.Printf("Processing %d lines for part 1\n", len(lines))

	return 0, nil
}

// TestPart1 runs the test case for part 1
func TestPart1() error {
	input, err := loadInput("input1.txt")
	if err != nil {
		return err
	}

	result, err := Part1(input)
	if err != nil {
		return fmt.Errorf("part1 failed: %w", err)
	}

	expectedResult := 0 // TODO: Update with expected test result
	if result != expectedResult {
		return fmt.Errorf("test failed: expected %d, got %d", expectedResult, result)
	}

	fmt.Println("✓ Part 1 test passed!")
	return nil
}

// RunPart1 runs part 1 with the actual input
func RunPart1() error {
	input, err := loadInput("input.txt")
	if err != nil {
		return err
	}

	result, err := Part1(input)
	if err != nil {
		return fmt.Errorf("part1 failed: %w", err)
	}

	fmt.Printf("Part 1 Result: %d\n", result)
	return nil
}

// Part2 solves part 2 of the puzzle
func Part2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// TODO: Implement your solution here
	fmt.Printf("Processing %d lines for part 2\n", len(lines))

	return 0, nil
}

// TestPart2 runs the test case for part 2
func TestPart2() error {
	input, err := loadInput("input2.txt")
	if err != nil {
		return err
	}

	result, err := Part2(input)
	if err != nil {
		return fmt.Errorf("part2 failed: %w", err)
	}

	expectedResult := 0 // TODO: Update with expected test result
	if result != expectedResult {
		return fmt.Errorf("test failed: expected %d, got %d", expectedResult, result)
	}

	fmt.Println("✓ Part 2 test passed!")
	return nil
}

// RunPart2 runs part 2 with the actual input
func RunPart2() error {
	input, err := loadInput("input.txt")
	if err != nil {
		return err
	}

	result, err := Part2(input)
	if err != nil {
		return fmt.Errorf("part2 failed: %w", err)
	}

	fmt.Printf("Part 2 Result: %d\n", result)
	return nil
}
