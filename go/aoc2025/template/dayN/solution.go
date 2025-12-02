package dayN

import (
	"aoc2025/solver"
	"aoc2025/utils"
)

func init() {
	solver.Register(N, Solution{})
}

type Solution struct{}

func (s Solution) Part1(input string) (int, error) {
	lines := utils.NewLineSplit(input)
	_ = lines
	return 0, nil
}

func (s Solution) Part2(input string) (int, error) {
	lines := utils.NewLineSplit(input)
	_ = lines
	return 0, nil
}
