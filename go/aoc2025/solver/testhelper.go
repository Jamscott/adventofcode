package solver

import (
	"aoc2025/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCase struct {
	Name     string
	Input    string
	Expected int
	Part     int
}

type TestConfig struct {
	Solver        Solver
	Part1Expected int
	Part2Expected int
	Part1Input    string
	Part2Input    string
}

func RunTests(t *testing.T, config TestConfig) {
	t.Helper()

	if config.Part1Input == "" {
		config.Part1Input = "input1.txt"
	}
	if config.Part2Input == "" {
		config.Part2Input = "input2.txt"
	}

	t.Run("Part1", func(t *testing.T) {
		input, err := utils.LoadInput(config.Part1Input)
		require.NoError(t, err, "Failed to load input")

		result, err := config.Solver.Part1(input)
		require.NoError(t, err, "Part1 returned an error")

		assert.Equal(t, config.Part1Expected, result)
	})

	t.Run("Part2", func(t *testing.T) {
		input, err := utils.LoadInput(config.Part2Input)
		require.NoError(t, err, "Failed to load input")

		result, err := config.Solver.Part2(input)
		require.NoError(t, err, "Part2 returned an error")

		assert.Equal(t, config.Part2Expected, result)
	})
}

func RunTableTests(t *testing.T, solver Solver, tests []TestCase) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var result int
			var err error

			switch tt.Part {
			case 1:
				result, err = solver.Part1(tt.Input)
			case 2:
				result, err = solver.Part2(tt.Input)
			default:
				t.Fatalf("Invalid part number: %d", tt.Part)
			}

			require.NoError(t, err)
			assert.Equal(t, tt.Expected, result)
		})
	}
}

type BenchmarkConfig struct {
	Solver     Solver
	Part1Input string
	Part2Input string
}

func RunBenchmarks(b *testing.B, config BenchmarkConfig) {
	b.Helper()

	if config.Part1Input == "" {
		config.Part1Input = "input1.txt"
	}
	if config.Part2Input == "" {
		config.Part2Input = "input2.txt"
	}

	b.Run("Part1", func(b *testing.B) {
		input, err := utils.LoadInput(config.Part1Input)
		if err != nil {
			b.Fatalf("Failed to load input: %v", err)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = config.Solver.Part1(input)
		}
	})

	b.Run("Part2", func(b *testing.B) {
		input, err := utils.LoadInput(config.Part2Input)
		if err != nil {
			b.Fatalf("Failed to load input: %v", err)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = config.Solver.Part2(input)
		}
	})
}

func SkipIfNoInput(t *testing.T, filename string) {
	t.Helper()
	_, err := utils.LoadInput(filename)
	if err != nil {
		t.Skipf("Skipping: %v", err)
	}
}

func MustLoadInput(t testing.TB, filename string) string {
	t.Helper()
	input, err := utils.LoadInput(filename)
	require.NoError(t, err, fmt.Sprintf("Failed to load %s", filename))
	return input
}
