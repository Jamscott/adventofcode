import run from "aocrunner";
import { NEWLINE } from "../utils/regex.js";
import "../utils/helpers.js";
import { log } from "console";

const parseInput = (rawInput: string) => rawInput.split(NEWLINE)

const words: Record<string, string> = {
  one: "1",
  two: "2",
  three: "3",
  four: "4",
  five: "5",
  six: "6",
  seven: "7",
  eight: "8",
  nine: "9",
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput)
    .map((line) => line.replace(/[a-z]/g, ''))
    .splitChars()
    .map((line) => line.map((num) => parseInt(num, 10)))
    .map((line) => Number(`${line[0]}${line[line.length - 1]}`))
    .sum()
  return input
};

let digit = /one|two|three|four|five|six|seven|eight|nine|\d/g
let reverseDigit = /enin|thgie|neves|xis|evif|ruof|eerht|owt|eno|\d/g

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput)
    .map((line) => {
      const firstDigit = line.match(digit)?.[0]
      const lastDigit = line.reverse().match(reverseDigit)?.[0]

      const mappedDigit = firstDigit && (words[firstDigit] || firstDigit)
      const mappedReverseDigit = lastDigit && (words[lastDigit.reverse()] || lastDigit)

      return Number(`${mappedDigit}${mappedReverseDigit}`)
    })
    .sum()
  return input
};

run({
  part1: {
    tests: [
      {
        input: `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet      
`,
        expected: 142,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`,
        expected: 281,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
