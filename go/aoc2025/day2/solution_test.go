package day2

import (
	"aoc2025/runner"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	Expected_Answer_Part1 = 1227775554
	Expected_Answer_Part2 = 4174379265
)

func TestPart1(t *testing.T) {
	input, err := runner.LoadInput("input1.txt")
	if err != nil {
		t.Fatalf("Failed to load input: %v", err)
	}

	result, err := Part1(input)
	if err != nil {
		t.Fatalf("Failed to run Part1: %v", err)
	}

	assert.Equal(t, Expected_Answer_Part1, result)
}

func TestPart2(t *testing.T) {
	input, err := runner.LoadInput("input2.txt")
	if err != nil {
		t.Fatalf("Failed to load input: %v", err)
	}

	result, err := Part2(input)
	if err != nil {
		t.Fatalf("Failed to run Part2: %v", err)
	}

	assert.Equal(t, Expected_Answer_Part2, result)
}
