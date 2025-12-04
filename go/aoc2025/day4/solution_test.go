package day4

import (
	"aoc2025/solver"
	"testing"
)

const (
	expectedPart1 = 13
	expectedPart2 = 13
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
