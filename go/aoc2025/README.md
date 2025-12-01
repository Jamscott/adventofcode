# Advent of Code 2025 - Go Runner

An interactive command-line tool for running Advent of Code solutions in Go.

## Project Structure

```
aoc2025/
├── main.go              # Interactive CLI runner
├── runner/
│   └── runner.go        # Day discovery and runner utilities
├── template/            # Template for new days
│   ├── solution.go      # Solution template with Part 1 & 2
│   ├── input.txt        # Actual puzzle input
│   ├── input1.txt       # Test input for Part 1
│   └── input2.txt       # Test input for Part 2
└── dayX/                # Individual day folders (copy from template)
    ├── solution.go
    ├── input.txt
    ├── input1.txt
    └── input2.txt
```

## Getting Started

### Running the Application

```bash
cd go/aoc2025
go run main.go
```

The application will:
1. Auto-discover all day folders in the directory
2. Show you a menu of available days
3. Let you select which day to run
4. Run both test cases and actual solutions for that day

### Creating a New Day

1. **Copy the template folder:**
   ```bash
   cp -r template day1
   ```

2. **Update the package name** in `day1/solution.go`:
   ```go
   package day1  // Change from 'package template'
   ```

3. **Update the file loading path** in the `loadInput` function:
   ```go
   data, err := os.ReadFile("day1/" + filename)  // Change from "template/"
   ```

4. **Add your test inputs:**
   - `input1.txt` - Example input for Part 1 (from the puzzle description)
   - `input2.txt` - Example input for Part 2
   - `input.txt` - Your actual puzzle input

5. **Update expected test results** in `TestPart1()` and `TestPart2()`:
   ```go
   expectedResult := 42  // Update with the expected answer from the example
   ```

6. **Register the day in `main.go`:**
   ```go
   import (
       // ... existing imports
       "aoc2025/day1"
   )

   dayRunners := map[string]runner.DayRunner{
       // ... existing entries
       "day1": {
           TestPart1: day1.TestPart1,
           RunPart1:  day1.RunPart1,
           TestPart2: day1.TestPart2,
           RunPart2:  day1.RunPart2,
       },
   }
   ```

7. **Implement your solution** in the `Part1()` and `Part2()` functions

## How It Works

### File Loading
- File loading is abstracted in the `loadInput()` function in each day's solution
- Tests automatically load from `inputX.txt` files
- Solutions load from `input.txt`

### Test-Driven Development
When you run a day, it:
1. ✅ Runs Part 1 test with `input1.txt`
2. 🚀 If test passes, runs Part 1 with actual `input.txt`
3. ✅ Runs Part 2 test with `input2.txt`
4. 🚀 If test passes, runs Part 2 with actual `input.txt`

If any test fails, it won't run the actual solution for that part.

## Example Workflow

1. Start the runner:
   ```bash
   go run main.go
   ```

2. You'll see:
   ```
   ==================================================
        Advent of Code 2025 - Day Runner
   ==================================================

   Available days:
     0 - Template
     1 - day1

   Commands:
     <number>     - Run tests and solutions for a specific day
     q or quit    - Exit the program

   Enter day number:
   ```

3. Enter a day number (e.g., `1`) and press Enter

4. The runner executes:
   ```
   --------------- Running Day 1 ---------------

   📝 Running Part 1 Test...
   ✓ Part 1 test passed!

   🚀 Running Part 1 Solution...
   Part 1 Result: 12345

   📝 Running Part 2 Test...
   ✓ Part 2 test passed!

   🚀 Running Part 2 Solution...
   Part 2 Result: 67890

   -------------- Completed Day 1 --------------
   ```

## Tips

- Use `0` to test the template implementation
- Copy the template for each new day to maintain consistency
- Update expected test results as soon as you know what they should be
- The file loading is abstracted, so you only need to implement the solution logic
- Days are auto-discovered from the filesystem - just make sure to register them in `main.go`

## Building

To build a standalone executable:

```bash
go build -o aoc2025
./aoc2025
```

