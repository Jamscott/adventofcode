package day6

import (
	"aoc2025/solver"
	"strconv"
	"strings"
)

func init() {
	solver.Register(6, Solution{})
}

type Solution struct{}

type data struct {
	nums [][]int
	ops  []string
}

func sumFold(nums []int) int {
	n := 0
	for _, num := range nums {
		n += num
	}
	return n
}

func multFold(nums []int) int {
	n := 1
	for _, num := range nums {
		n *= num
	}
	return n
}

func process(nums []int, op string) int {
	switch op {
	case "+":
		return sumFold(nums)
	case "*":
		return multFold(nums)
	default:
		return 0
	}
}

func parseInput(input string) data {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var nums [][]int
	for _, s := range lines[:len(lines)-1] {
		var line []int
		for _, field := range strings.Fields(s) {
			num, err := strconv.Atoi(field)
			if err != nil {
				continue
			}
			line = append(line, num)
		}
		nums = append(nums, line)
	}
	ops := strings.Fields(lines[len(lines)-1])
	return data{nums, ops}
}

func stringToCharMatrix(input string) [][]rune {
	lines := strings.Split(input, "\n")
	var chars [][]rune
	for i, line := range lines {
		for j, c := range line {
			for len(chars) <= j {
				chars = append(chars, []rune{})
			}
			for len(chars[j]) <= i {
				chars[j] = append(chars[j], '_')
			}
			chars[j][i] = c
		}
	}
	return chars
}

func (s Solution) Part1(input string) (int, error) {
	inp := parseInput(input)
	c := 0
	for n := 0; n < len(inp.nums[0]); n++ {
		op := inp.ops[n]
		var nums []int
		for _, row := range inp.nums {
			nums = append(nums, row[n])
		}
		c += process(nums, op)
	}
	return c, nil
}

func (s Solution) Part2(input string) (int, error) {
	trimmed := strings.TrimSpace(input)
	chars := stringToCharMatrix(trimmed)

	var transposedLines []string
	for _, l := range chars {
		transposedLines = append(transposedLines, strings.TrimSpace(string(l)))
	}

	transposedText := strings.Join(transposedLines, "\n")

	sol := 0
	groups := strings.Split(transposedText, "\n\n")
	for _, group := range groups {
		groupLines := strings.Split(group, "\n")
		if len(groupLines) == 0 {
			continue
		}
		fstLine := groupLines[0]
		if len(fstLine) == 0 {
			continue
		}
		op := fstLine[len(fstLine)-1]

		var nums []int
		for _, line := range groupLines {
			numString := strings.TrimSpace(strings.TrimRight(line, "+*"))
			if len(numString) == 0 {
				continue
			}
			n, err := strconv.Atoi(numString)
			if err != nil {
				continue
			}
			nums = append(nums, n)
		}

		sol += process(nums, string(op))
	}
	return sol, nil
}
