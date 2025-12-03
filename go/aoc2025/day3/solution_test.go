package day3

import (
	"aoc2025/solver"
	"testing"
)

const (
	expectedPart1 = 357
	expectedPart2 = 3121910778619
)

func TestSolution(t *testing.T) {
	solver.RunTests(t, solver.TestConfig{
		Solver:        Solution{},
		Part1Expected: expectedPart1,
		Part2Expected: expectedPart2,
	})
}

func BenchmarkSolution(b *testing.B) {
	solver.RunBenchmarks(b, solver.BenchmarkConfig{
		Solver: Solution{},
	})
}
