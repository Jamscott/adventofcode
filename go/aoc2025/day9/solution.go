package day9

import (
	"aoc2025/solver"
	"aoc2025/utils"
	"slices"
	"strconv"
	"strings"
)

func init() {
	solver.Register(9, Solution{})
}

type Solution struct{}

type rect struct {
	p1   utils.Position
	p2   utils.Position
	area int
}

type span struct {
	begin int
	end   int
	x     int
}

func parseInput(input string) []utils.Position {
	var points []utils.Position

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}

		x, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			continue
		}
		y, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			continue
		}

		points = append(points, utils.Position{X: x, Y: y})
	}

	return points
}

func hittest(lines *[]span, p utils.Position) int {
	n := 0

	for _, s := range *lines {
		if s.x <= p.X {
			continue
		}

		if s.begin <= p.Y && p.Y < s.end {
			n++
		}
	}

	return n
}

func (s Solution) Part1(input string) (int, error) {
	points := parseInput(input)
	sol := 0

	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]

			w := p2.X - p1.X
			if w < 0 {
				w = -w
			}
			h := p2.Y - p1.Y
			if h < 0 {
				h = -h
			}
			w++
			h++
			area := w * h

			if area > sol {
				sol = area
			}
		}
	}

	return sol, nil
}

func (s Solution) Part2(input string) (int, error) {
	points := parseInput(input)
	if len(points) == 0 {
		return 0, nil
	}

	var horLines []span
	var vertLines []span

	last := points[len(points)-1]
	for _, p := range points {
		if last.X == p.X {
			vertLines = append(vertLines, span{
				begin: min(p.Y, last.Y),
				end:   max(p.Y, last.Y),
				x:     p.X,
			})
		} else if last.Y == p.Y {
			horLines = append(horLines, span{
				begin: min(p.X, last.X),
				end:   max(p.X, last.X),
				x:     p.Y,
			})
		}
		last = p
	}

	var rects []rect
	for i, p1 := range points {
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]

			w := p2.X - p1.X
			if w < 0 {
				w = -w
			}
			h := p2.Y - p1.Y
			if h < 0 {
				h = -h
			}
			w++
			h++
			area := w * h

			rects = append(rects, rect{p1, p2, area})
		}
	}

	slices.SortFunc(rects, func(a, b rect) int {
		return b.area - a.area
	})

	for _, rect := range rects {
		p1 := rect.p1
		p2 := rect.p2

		minx := min(p1.X, p2.X)
		maxx := max(p1.X, p2.X)
		miny := min(p1.Y, p2.Y)
		maxy := max(p1.Y, p2.Y)

		if hittest(&vertLines, utils.Position{X: minx, Y: miny})%2 != 1 {
			continue
		}

		valid := true

		for _, s := range vertLines {
			if minx < s.x && s.x < maxx {
				if max(s.begin, miny) < min(s.end, maxy) {
					valid = false
					break
				}
			}
		}

		if !valid {
			continue
		}

		for _, s := range horLines {
			if miny < s.x && s.x < maxy {
				if max(s.begin, minx) < min(s.end, maxx) {
					valid = false
					break
				}
			}
		}

		if !valid {
			continue
		}

		return rect.area, nil
	}

	return -1, nil
}
