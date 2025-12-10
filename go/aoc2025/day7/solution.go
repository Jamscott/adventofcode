package day7

import (
	"aoc2025/solver"
	"aoc2025/utils"
	"strings"
)

func init() {
	solver.Register(7, Solution{})
}

type Solution struct{}

type data struct {
	startPos utils.Position
	grid     utils.Grid
}

func parseInput(input string) data {
	grid := utils.ParseGrid(input)
	startPos := grid.FindRune('S')
	return data{startPos, grid}
}

var visited = make(map[utils.Position]bool)

func part1DFS(input data, pos utils.Position) {
	if visited[pos] {
		return
	}

	if !input.grid.InBoundsPos(pos) {
		return
	}

	c := input.grid.GetPos(pos, '#')
	if c != '^' {
		part1DFS(input, utils.Position{X: pos.X, Y: pos.Y + 1})
		return
	}

	visited[pos] = true
	part1DFS(input, utils.Position{X: pos.X - 1, Y: pos.Y})
	part1DFS(input, utils.Position{X: pos.X + 1, Y: pos.Y})
}

var cache = make(map[utils.Position]int)

func part2DFS(input data, pos utils.Position) int {
	if count, hit := cache[pos]; hit {
		return count
	}

	if !input.grid.InBoundsPos(pos) {
		return 1
	}

	c := input.grid.GetPos(pos, '#')
	if c != '^' {
		res := part2DFS(input, utils.Position{X: pos.X, Y: pos.Y + 1})
		cache[pos] = res
		return res
	}

	res := part2DFS(input, utils.Position{X: pos.X - 1, Y: pos.Y}) +
		part2DFS(input, utils.Position{X: pos.X + 1, Y: pos.Y})
	cache[pos] = res
	return res
}

func (s Solution) Part1(input string) (int, error) {
	visited = make(map[utils.Position]bool)
	data := parseInput(strings.TrimSpace(input))
	part1DFS(data, data.startPos)
	return len(visited), nil
}

func (s Solution) Part2(input string) (int, error) {
	cache = make(map[utils.Position]int)
	data := parseInput(strings.TrimSpace(input))
	return part2DFS(data, data.startPos), nil
}
