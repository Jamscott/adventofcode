package day4

import (
	"aoc2025/solver"
	"aoc2025/utils"
)

func init() {
	solver.Register(4, Solution{})
}

type Solution struct{}

const (
	roleSymbol  = '@'
	emptySymbol = '.'
)

func (s Solution) Part1(input string) (int, error) {
	grid := utils.ParseGrid(input)
	roles := findAccessibleRoles(grid)
	return len(roles), nil
}

func (s Solution) Part2(input string) (int, error) {
	grid := utils.ParseGrid(input)
	totalRoles := 0

	for {
		roles := findAccessibleRoles(grid)
		if len(roles) == 0 {
			break
		}

		totalRoles += len(roles)

		for _, pos := range roles {
			grid.SetPos(pos, emptySymbol)
		}
	}

	return totalRoles, nil
}

func findAccessibleRoles(grid utils.Grid) []utils.Position {
	var accessible []utils.Position

	for y, row := range grid {
		for x, cell := range row {
			if cell == roleSymbol {
				neighborCount := grid.CountNeighbors(x, y, true, func(r rune) bool {
					return r == roleSymbol
				})

				if neighborCount < 4 {
					accessible = append(accessible, utils.Position{X: x, Y: y})
				}
			}
		}
	}

	return accessible
}
