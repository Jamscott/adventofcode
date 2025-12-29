package day10

import (
	"aoc2025/solver"
	"aoc2025/utils"
	"strconv"
	"strings"
)

func init() {
	solver.Register(10, Solution{})
}

type Solution struct{}

type Machine struct {
	indicators    int
	buttons       []int
	buttonIndices [][]int
	joltages      []int
}

func parseInput(input string) []Machine {
	var machines []Machine

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		machine := parseLine(line)
		machines = append(machines, machine)
	}

	return machines
}

// Indicators:
// Index:     0  1  2  3  4  left  right
// Character: .  #  #  .  #
// Bitmask:
// Indicator string: [.##.#]
// Index left right: 0  1  2  3  4
// Character:          .  #  #  .  #
// Binary value:       0  1  1  0  1
// Bit position:  4  3  2  1  0  right  left
// Binary:        0  1  1  0  1
// bit_position = (length - 1 - index)

func parseLine(line string) Machine {
	parts := strings.Fields(line)

	indicatorString := strings.Trim(parts[0], "[]")
	indicatorString = strings.ReplaceAll(indicatorString, ".", "0")
	indicatorString = strings.ReplaceAll(indicatorString, "#", "1")
	indicatorBitmask, _ := strconv.ParseInt(indicatorString, 2, 64)

	var buttonBitmasks []int
	var buttonIndicesList [][]int
	var joltages []int

	for partIndex := 1; partIndex < len(parts); partIndex++ {
		part := parts[partIndex]

		if strings.HasPrefix(part, "{") {
			joltageString := strings.Trim(part, "{}")
			joltageStrings := strings.Split(joltageString, ",")
			for _, joltStr := range joltageStrings {
				jolt, _ := strconv.Atoi(strings.TrimSpace(joltStr))
				joltages = append(joltages, jolt)
			}
			break
		}

		buttonString := strings.Trim(part, "()")
		buttonIndicesStrings := strings.Split(buttonString, ",")

		buttonBitmask := 0
		indicatorLength := len(indicatorString)
		var indices []int

		for _, indexString := range buttonIndicesStrings {
			indicatorIndex, err := strconv.Atoi(strings.TrimSpace(indexString))
			if err != nil {
				continue
			}

			bitPosition := indicatorLength - 1 - indicatorIndex
			buttonBitmask |= (1 << bitPosition)
			indices = append(indices, indicatorIndex)
		}

		buttonBitmasks = append(buttonBitmasks, buttonBitmask)
		buttonIndicesList = append(buttonIndicesList, indices)
	}

	return Machine{
		indicators:    int(indicatorBitmask),
		buttons:       buttonBitmasks,
		buttonIndices: buttonIndicesList,
		joltages:      joltages,
	}
}

func (s Solution) Part1(input string) (int, error) {
	machines := parseInput(input)
	totalButtonPresses := 0

	for _, machine := range machines {
		buttonSubsets := utils.Subsets(machine.buttons)

		for _, buttonSubset := range buttonSubsets {
			combinedBitmask := 0
			for _, buttonBitmask := range buttonSubset {
				combinedBitmask ^= buttonBitmask
			}

			if combinedBitmask == machine.indicators {
				totalButtonPresses += len(buttonSubset)
				// winner winner chicken dinner
				break
			}
		}
	}

	return totalButtonPresses, nil
}

func (s Solution) Part2(input string) (int, error) {
	return 0, nil
}
