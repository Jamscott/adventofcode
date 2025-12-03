package main

import (
	"aoc2025/solver"
	"aoc2025/utils"
	"flag"
	"fmt"
	"os"

	_ "aoc2025/day1"
	_ "aoc2025/day2"
	_ "aoc2025/day3"
)

func main() {
	dayFlag := flag.Int("day", 0, "Run specific day (0 = all days)")
	flag.Parse()

	if *dayFlag == 0 {
		runAllDays()
	} else {
		runDay(*dayFlag)
	}
}

func runAllDays() {
	days := solver.All()
	if len(days) == 0 {
		fmt.Println("No days registered yet!")
		return
	}

	fmt.Printf("Running all %d registered days...\n", len(days))

	for _, day := range days {
		runDayWithSolver(day)
	}
}

func runDay(dayNumber int) {
	day, exists := solver.Get(dayNumber)
	if !exists {
		fmt.Printf("Day %d not found. Available days: ", dayNumber)
		for _, d := range solver.All() {
			fmt.Printf("%d ", d.Number)
		}
		fmt.Println()
		os.Exit(1)
	}

	runDayWithSolver(day)
}

func runDayWithSolver(day *solver.Day) {
	inputFile := fmt.Sprintf("day%d/input.txt", day.Number)
	input, err := utils.LoadInput(inputFile)

	result := day.Run(input)
	if err != nil {
		result.InputLoadErr = err
	}

	result.Print(day.Number)
}
