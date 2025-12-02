# Script to create a new day's solution from the template
# Usage: .\newday.ps1 <day_number>
# Example: .\newday.ps1 3

param(
    [Parameter(Mandatory=$true)]
    [int]$Day
)

$DayDir = "day$Day"

if (Test-Path $DayDir) {
    Write-Error "Error: $DayDir already exists!"
    exit 1
}

Write-Host "Creating $DayDir..." -ForegroundColor Green
New-Item -ItemType Directory -Path $DayDir | Out-Null

# Copy template files
Copy-Item "template/dayN/solution.go" "$DayDir/solution.go"
Copy-Item "template/dayN/solution_test.go" "$DayDir/solution_test.go"

# Replace N with the actual day number
(Get-Content "$DayDir/solution.go") -replace 'dayN', "day$Day" `
    -replace 'Register\(N,', "Register($Day," `
    -replace 'day N', "day $Day" | Set-Content "$DayDir/solution.go"

(Get-Content "$DayDir/solution_test.go") -replace 'dayN', "day$Day" | Set-Content "$DayDir/solution_test.go"

# Create empty input files
New-Item -ItemType File -Path "$DayDir/input.txt" | Out-Null
New-Item -ItemType File -Path "$DayDir/input1.txt" | Out-Null
New-Item -ItemType File -Path "$DayDir/input2.txt" | Out-Null

Write-Host "✓ Created $DayDir with:" -ForegroundColor Green
Write-Host "  - solution.go"
Write-Host "  - solution_test.go"
Write-Host "  - input.txt (empty)"
Write-Host "  - input1.txt (empty)"
Write-Host "  - input2.txt (empty)"
Write-Host ""
Write-Host "Next steps:" -ForegroundColor Yellow
Write-Host "1. Add the import to main.go: _ `"aoc2025/day$Day`""
Write-Host "2. Fill in the input files"
Write-Host "3. Implement Part1 and Part2 in solution.go"
Write-Host "4. Update expected answers in solution_test.go"
Write-Host "5. Run with: go run . -day $Day"

