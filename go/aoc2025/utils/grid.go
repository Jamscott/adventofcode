package utils

import "strings"

type Position struct {
	X int
	Y int
}

type Grid [][]rune

var (
	DirectionsCardinal = []Position{
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
	}

	DirectionsAll = []Position{
		{X: -1, Y: -1},
		{X: 0, Y: -1},
		{X: 1, Y: -1},
		{X: 1, Y: 0},
		{X: 1, Y: 1},
		{X: 0, Y: 1},
		{X: -1, Y: 1},
		{X: -1, Y: 0},
	}
)

func ParseGrid(input string) Grid {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make(Grid, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		row := []rune(line)
		grid = append(grid, row)
	}

	return grid
}

func (g Grid) Height() int {
	return len(g)
}

func (g Grid) Width() int {
	if len(g) == 0 {
		return 0
	}
	return len(g[0])
}

func (g Grid) InBounds(x, y int) bool {
	return y >= 0 && y < len(g) && x >= 0 && x < len(g[y])
}

func (g Grid) InBoundsPos(pos Position) bool {
	return g.InBounds(pos.X, pos.Y)
}

func (g Grid) Get(x, y int, defaultVal rune) rune {
	if !g.InBounds(x, y) {
		return defaultVal
	}
	return g[y][x]
}

func (g Grid) GetPos(pos Position, defaultVal rune) rune {
	return g.Get(pos.X, pos.Y, defaultVal)
}

func (g Grid) Set(x, y int, val rune) {
	if g.InBounds(x, y) {
		g[y][x] = val
	}
}

func (g Grid) SetPos(pos Position, val rune) {
	g.Set(pos.X, pos.Y, val)
}

func (g Grid) GetNeighbors(x, y int, includeDiagonals bool) []Position {
	directions := DirectionsCardinal
	if includeDiagonals {
		directions = DirectionsAll
	}

	neighbors := make([]Position, 0, len(directions))
	for _, dir := range directions {
		nx, ny := x+dir.X, y+dir.Y
		if g.InBounds(nx, ny) {
			neighbors = append(neighbors, Position{X: nx, Y: ny})
		}
	}
	return neighbors
}

func (g Grid) GetNeighborsPos(pos Position, includeDiagonals bool) []Position {
	return g.GetNeighbors(pos.X, pos.Y, includeDiagonals)
}

func (g Grid) CountNeighbors(x, y int, includeDiagonals bool, predicate func(rune) bool) int {
	count := 0
	neighbors := g.GetNeighbors(x, y, includeDiagonals)
	for _, pos := range neighbors {
		if predicate(g[pos.Y][pos.X]) {
			count++
		}
	}
	return count
}

func (g Grid) CountNeighborsPos(pos Position, includeDiagonals bool, predicate func(rune) bool) int {
	return g.CountNeighbors(pos.X, pos.Y, includeDiagonals, predicate)
}

func (g Grid) FindAll(predicate func(rune) bool) []Position {
	positions := make([]Position, 0)
	for y, row := range g {
		for x, cell := range row {
			if predicate(cell) {
				positions = append(positions, Position{X: x, Y: y})
			}
		}
	}
	return positions
}
