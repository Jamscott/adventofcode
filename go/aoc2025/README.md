# Advent of Code 2025 - Go Solutions

This project contains solutions for [Advent of Code 2025](https://adventofcode.com/2025) written in Go.

## 🚀 Quick Start

### Run All Days
```bash
go run .
```

### Run Specific Day
```bash
go run . -day 1
go run . -day 2
```

### Run Tests
```bash
# Test all days
go test ./...

# Test specific day
go test ./day1

# Run with verbose output
go test -v ./day1
```

### Run Benchmarks
```bash
# Benchmark all days
go test -bench=. ./...

# Benchmark specific day
go test -bench=. ./day1

# With memory stats
go test -bench=. -benchmem ./day1
```

## 📁 Project Structure

```
aoc2025/
├── main.go              # CLI entry point with day selection
├── solver/              # Shared solver framework
│   ├── solver.go        # Solver interface and registry
│   └── testhelper.go    # Shared test and benchmark helpers
├── utils/               # Shared utility functions
│   └── utils.go         # Input loading, string splitting, etc.
├── template/            # Template for new days
│   └── dayN/
│       ├── solution.go
│       └── solution_test.go
├── day1/                # Day 1 solution
│   ├── solution.go
│   ├── solution_test.go
│   ├── input.txt        # Actual puzzle input
│   ├── input1.txt       # Part 1 test input
│   └── input2.txt       # Part 2 test input
├── day2/                # Day 2 solution
│   └── ...
└── newday.ps1           # Script to create new day from template
```

## 🆕 Creating a New Day

### Using the Script (Recommended)

**Windows (PowerShell):**
```powershell
.\newday.ps1 3
```

**Linux/macOS:**
```bash
chmod +x newday.sh
./newday.sh 3
```

### Manual Setup

1. Copy the template:
   ```bash
   cp -r template/dayN day3
   ```

2. Update `day3/solution.go`:
   - Replace `package dayN` with `package day3`
   - Replace `solver.Register(N, ...)` with `solver.Register(3, ...)`

3. Update `day3/solution_test.go`:
   - Replace `package dayN` with `package day3`

4. Add import to `main.go`:
   ```go
   _ "aoc2025/day3"
   ```

5. Fill in your input files and implement the solution!

## 📝 Writing Solutions

Each day implements the `solver.Solver` interface:

```go
package day3

import "aoc2025/solver"

func init() {
    solver.Register(3, Solution{})
}

type Solution struct{}

func (s Solution) Part1(input string) (int, error) {
    // Your solution here
    return 0, nil
}

func (s Solution) Part2(input string) (int, error) {
    // Your solution here
    return 0, nil
}
```

## 🧪 Writing Tests

The framework provides helpers to eliminate boilerplate:

```go
package day3

import (
    "aoc2025/solver"
    "testing"
)

const (
    expectedPart1 = 42
    expectedPart2 = 100
)

func TestSolution(t *testing.T) {
    solver.RunTests(t, solver.TestConfig{
        Solver:        Solution{},
        Part1Expected: expectedPart1,
        Part2Expected: expectedPart2,
    })
}

func BenchmarkSolution(b *testing.B) {
    solver.RunBenchmarks(b, solver.BenchmarkConfig{
        Solver: Solution{},
    })
}
```

### Custom Table-Driven Tests

For additional test cases:

```go
func TestCustomCases(t *testing.T) {
    tests := []solver.TestCase{
        {
            Name:     "Example 1",
            Input:    "1 2 3",
            Expected: 6,
            Part:     1,
        },
        {
            Name:     "Example 2",
            Input:    "4 5 6",
            Expected: 15,
            Part:     1,
        },
    }
    solver.RunTableTests(t, Solution{}, tests)
}
```

## 🛠️ Utilities

The `utils` package provides common functions:

```go
import "aoc2025/utils"

// Load input file
input, err := utils.LoadInput("day1/input.txt")

// Split by newlines
lines := utils.NewLineSplit(input)

// Split by spaces
words := utils.SpaceSplit(line)
```

## 🎯 Benefits of This Structure

1. **No Manual main.go Edits**: Days auto-register via `init()`
2. **CLI Day Selection**: Run specific days with `-day` flag
3. **Minimal Test Boilerplate**: Shared test helpers
4. **Standardized Benchmarks**: Consistent across all days
5. **Easy Day Creation**: Scripts generate boilerplate
6. **Type Safety**: Interface ensures consistent API
7. **Timing Built-in**: Automatic performance measurement

## 📊 Example Output

```
$ go run . -day 1

=== Day 1 ===
Part 1: 3 (took 45.23µs)
Part 2: 0 (took 12.34µs)
```

```
$ go run .

Running all 2 registered days...

=== Day 1 ===
Part 1: 3 (took 45.23µs)
Part 2: 0 (took 12.34µs)

=== Day 2 ===
Part 1: 1227775554 (took 123.45ms)
Part 2: 4174379265 (took 234.56ms)
```

