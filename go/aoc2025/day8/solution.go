package day8

import (
	"aoc2025/solver"
	"aoc2025/utils"
	"slices"
	"strconv"
	"strings"
)

func init() {
	solver.Register(8, Solution{})
}

type Solution struct{}

type position struct {
	index int
	vec   utils.Vec3
}

type distance struct {
	pos1 position
	pos2 position
	dist float64
}

type DisjointSet struct {
	parent   map[int]int
	elements map[int]position
}

func (ds *DisjointSet) makeSet(pos position) {
	if _, exists := ds.parent[pos.index]; !exists {
		ds.parent[pos.index] = pos.index
		ds.elements[pos.index] = pos
	}
}

func (ds *DisjointSet) find(index int) int {
	if ds.parent[index] == index {
		return index
	}
	ds.parent[index] = ds.find(ds.parent[index])
	return ds.parent[index]
}

func (ds *DisjointSet) union(x int, y int) {
	rootX := ds.find(x)
	rootY := ds.find(y)

	if rootX != rootY {
		ds.parent[rootX] = rootY
	}
}

func parseInput(input string) []position {
	var positions []position

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for index, line := range lines {
		line = strings.TrimSpace(line) // Remove \r and any whitespace
		coords := strings.Split(line, ",")
		if len(coords) != 3 {
			continue
		}

		x, err := strconv.Atoi(strings.TrimSpace(coords[0]))
		if err != nil {
			continue
		}

		y, err := strconv.Atoi(strings.TrimSpace(coords[1]))
		if err != nil {
			continue
		}

		z, err := strconv.Atoi(strings.TrimSpace(coords[2]))
		if err != nil {
			continue
		}

		pos := position{
			index: index,
			vec:   utils.Vec3{X: x, Y: y, Z: z},
		}
		positions = append(positions, pos)
	}

	return positions
}

func calculateDistances(positions []position) []distance {
	distancesLength := len(positions) * (len(positions) - 1) / 2
	distances := make([]distance, 0, distancesLength)

	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			pos1 := positions[i]
			pos2 := positions[j]
			d := pos1.vec.Distance(pos2.vec)

			dist := distance{pos1: pos1, pos2: pos2, dist: d}
			distances = append(distances, dist)
		}
	}

	slices.SortFunc(distances, func(a, b distance) int {
		if a.dist < b.dist {
			return -1
		}
		if a.dist > b.dist {
			return 1
		}
		return 0
	})

	return distances
}

func (s Solution) Part1(input string) (int, error) {
	positions := parseInput(input)
	distances := calculateDistances(positions)

	ds := &DisjointSet{
		parent:   make(map[int]int),
		elements: make(map[int]position),
	}

	for i := 0; i < len(positions); i++ {
		ds.makeSet(positions[i])
	}

	connectionsToMake := min(1000, len(distances))
	for i := 0; i < connectionsToMake; i++ {
		pos1 := distances[i].pos1
		pos2 := distances[i].pos2
		ds.union(pos1.index, pos2.index)
	}

	circuitSizes := make(map[int]int)
	for i := 0; i < len(positions); i++ {
		representative := ds.find(i)
		circuitSizes[representative]++
	}

	var sizes []int
	for _, size := range circuitSizes {
		sizes = append(sizes, size)
	}

	slices.SortFunc(sizes, func(a, b int) int { return b - a })

	product := 1
	for i := 0; i < min(3, len(sizes)); i++ {
		product *= sizes[i]
	}

	return product, nil
}

func (s Solution) Part2(input string) (int, error) {
	positions := parseInput(input)
	if len(positions) == 0 {
		return 0, nil
	}

	distances := calculateDistances(positions)

	ds := &DisjointSet{
		parent:   make(map[int]int),
		elements: make(map[int]position),
	}

	for i := 0; i < len(positions); i++ {
		ds.makeSet(positions[i])
	}

	for i := 0; i < len(distances); i++ {
		pos1 := distances[i].pos1
		pos2 := distances[i].pos2

		ds.union(pos1.index, pos2.index)

		representative := ds.find(0)
		allSame := true

		for j := 1; j < len(positions); j++ {
			if representative != ds.find(j) {
				allSame = false
				break
			}
		}

		if allSame {
			return pos1.vec.X * pos2.vec.X, nil
		}
	}

	return 0, nil
}
