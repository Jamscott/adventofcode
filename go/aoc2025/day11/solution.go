package day11

import (
	"aoc2025/solver"
	"strings"
)

func init() {
	solver.Register(11, Solution{})
}

type Solution struct{}

type cacheKey struct {
	start string
	nodes string // comma separated sorted node names
}

type graph map[string][]string

func parseInput(input string) graph {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	g := make(graph)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		node := strings.TrimSpace(parts[0])
		connections := strings.Fields(strings.TrimSpace(parts[1]))
		g[node] = connections
	}

	return g
}

func setToString(nodes map[string]bool) string {
	if len(nodes) == 0 {
		return ""
	}

	var nodeList []string
	for node := range nodes {
		nodeList = append(nodeList, node)
	}

	for i := 0; i < len(nodeList); i++ {
		for j := i + 1; j < len(nodeList); j++ {
			if nodeList[i] > nodeList[j] {
				nodeList[i], nodeList[j] = nodeList[j], nodeList[i]
			}
		}
	}

	return strings.Join(nodeList, ",")
}

func removeFromSet(nodes map[string]bool, node string) map[string]bool {
	newSet := make(map[string]bool)
	for k, v := range nodes {
		if k != node {
			newSet[k] = v
		}
	}
	return newSet
}

func numPathsWithNodes(g graph, cache map[cacheKey]int64, start string, nodes map[string]bool) int64 {
	key := cacheKey{start: start, nodes: setToString(nodes)}

	if cached, ok := cache[key]; ok {
		return cached
	}

	if start == "out" && len(nodes) == 0 {
		return 1
	}

	connections, exists := g[start]
	if !exists {
		cache[key] = 0
		return 0
	}

	var total int64 = 0
	for _, next := range connections {
		newNodes := removeFromSet(nodes, start)
		total += numPathsWithNodes(g, cache, next, newNodes)
	}

	cache[key] = total
	return total
}

func (s Solution) Part1(input string) (int, error) {
	g := parseInput(input)
	cache := make(map[cacheKey]int64)
	nodes := make(map[string]bool)
	result := numPathsWithNodes(g, cache, "you", nodes)
	return int(result), nil
}

func (s Solution) Part2(input string) (int, error) {
	g := parseInput(input)
	cache := make(map[cacheKey]int64)
	nodes := map[string]bool{
		"dac": true,
		"fft": true,
	}
	result := numPathsWithNodes(g, cache, "svr", nodes)
	return int(result), nil
}
