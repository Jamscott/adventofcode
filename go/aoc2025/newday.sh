#!/bin/bash

# Script to create a new day's solution from the template
# Usage: ./newday.sh <day_number>

if [ -z "$1" ]; then
    echo "Usage: ./newday.sh <day_number>"
    echo "Example: ./newday.sh 3"
    exit 1
fi

DAY=$1
DAY_DIR="day${DAY}"

if [ -d "$DAY_DIR" ]; then
    echo "Error: $DAY_DIR already exists!"
    exit 1
fi

echo "Creating $DAY_DIR..."
mkdir -p "$DAY_DIR"

# Copy template files
cp template/dayN/solution.go "$DAY_DIR/solution.go"
cp template/dayN/solution_test.go "$DAY_DIR/solution_test.go"

# Replace N with the actual day number
if [[ "$OSTYPE" == "darwin"* ]]; then
    # macOS
    sed -i '' "s/dayN/day${DAY}/g" "$DAY_DIR/solution.go"
    sed -i '' "s/dayN/day${DAY}/g" "$DAY_DIR/solution_test.go"
    sed -i '' "s/Register(N,/Register(${DAY},/g" "$DAY_DIR/solution.go"
    sed -i '' "s/day N/day ${DAY}/g" "$DAY_DIR/solution.go"
else
    # Linux/Windows Git Bash
    sed -i "s/dayN/day${DAY}/g" "$DAY_DIR/solution.go"
    sed -i "s/dayN/day${DAY}/g" "$DAY_DIR/solution_test.go"
    sed -i "s/Register(N,/Register(${DAY},/g" "$DAY_DIR/solution.go"
    sed -i "s/day N/day ${DAY}/g" "$DAY_DIR/solution.go"
fi

# Create empty input files
touch "$DAY_DIR/input.txt"
touch "$DAY_DIR/input1.txt"
touch "$DAY_DIR/input2.txt"

echo "✓ Created $DAY_DIR with:"
echo "  - solution.go"
echo "  - solution_test.go"
echo "  - input.txt (empty)"
echo "  - input1.txt (empty)"
echo "  - input2.txt (empty)"
echo ""
echo "Next steps:"
echo "1. Add the import to main.go: _ \"aoc2025/day${DAY}\""
echo "2. Fill in the input files"
echo "3. Implement Part1 and Part2 in solution.go"
echo "4. Update expected answers in solution_test.go"
echo "5. Run with: go run . -day ${DAY}"

