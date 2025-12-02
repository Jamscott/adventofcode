package solver

import (
	"fmt"
	"time"
)

type Solver interface {
	Part1(input string) (int, error)
	Part2(input string) (int, error)
}

type Day struct {
	Number int
	Solver Solver
}

var registry = make(map[int]*Day)

func Register(dayNumber int, solver Solver) {
	if _, exists := registry[dayNumber]; exists {
		panic(fmt.Sprintf("day %d already registered", dayNumber))
	}
	registry[dayNumber] = &Day{
		Number: dayNumber,
		Solver: solver,
	}
}

func Get(dayNumber int) (*Day, bool) {
	day, exists := registry[dayNumber]
	return day, exists
}

func All() []*Day {
	days := make([]*Day, 0, len(registry))
	for _, day := range registry {
		days = append(days, day)
	}
	for i := 0; i < len(days)-1; i++ {
		for j := i + 1; j < len(days); j++ {
			if days[i].Number > days[j].Number {
				days[i], days[j] = days[j], days[i]
			}
		}
	}
	return days
}

type RunResult struct {
	Part1Result  int
	Part1Error   error
	Part1Time    time.Duration
	Part2Result  int
	Part2Error   error
	Part2Time    time.Duration
	InputLoadErr error
}

func formatDuration(d time.Duration) string {
	if d < time.Microsecond {
		return fmt.Sprintf("%dns", d.Nanoseconds())
	} else if d < time.Millisecond {
		return fmt.Sprintf("%.2fµs", float64(d.Nanoseconds())/1000.0)
	} else if d < time.Second {
		return fmt.Sprintf("%.2fms", float64(d.Microseconds())/1000.0)
	}
	return d.String()
}

func (d *Day) Run(input string) RunResult {
	result := RunResult{}

	start := time.Now()
	result.Part1Result, result.Part1Error = d.Solver.Part1(input)
	result.Part1Time = time.Since(start)

	start = time.Now()
	result.Part2Result, result.Part2Error = d.Solver.Part2(input)
	result.Part2Time = time.Since(start)

	return result
}

func (r *RunResult) Print(dayNumber int) {
	fmt.Printf("\n=== Day %d ===\n", dayNumber)

	if r.InputLoadErr != nil {
		fmt.Printf("❌ Failed to load input: %v\n", r.InputLoadErr)
		return
	}

	if r.Part1Error != nil {
		fmt.Printf("Part 1: Error: %v\n", r.Part1Error)
	} else {
		fmt.Printf("Part 1: %d (took %s)\n", r.Part1Result, formatDuration(r.Part1Time))
	}

	if r.Part2Error != nil {
		fmt.Printf("Part 2: Error: %v\n", r.Part2Error)
	} else {
		fmt.Printf("Part 2: %d (took %s)\n", r.Part2Result, formatDuration(r.Part2Time))
	}
}
