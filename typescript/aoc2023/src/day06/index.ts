import run from "aocrunner";

import "../utils/helpers.js";
import { matchNumber } from "../utils/regex.js";

const parseInput = (rawInput: string) => rawInput
  .splitNewLines()

const calculatePossibleOutcomes = (time: number, record: number) => {
  let winningRaces: number[] = [];

  for (let speed = 0; speed < time; speed++) {
    const myDistance = (time - speed) * speed;
    
    if (myDistance > record) {
      winningRaces.push(myDistance);
    }
  }
  return winningRaces.length;
}

const part1 = (rawInput: string) => {
  const [times, distances] = parseInput(rawInput)
    .map((line) => matchNumber(/(\d{1,9})/g, line))
    
  const possibleOutcomesPerRace = times
    .map((t, i) => calculatePossibleOutcomes(t, distances[i]))
    .power()

  return possibleOutcomesPerRace
};

const part2 = (rawInput: string) => {
  const [time, record] = parseInput(rawInput)
    .map((line) => matchNumber(/(\d{1,9})/g, line))
    .map((line) => parseInt(line.join("")))
    .flat()

  return calculatePossibleOutcomes(time, record)
};

run({
  part1: {
    tests: [
      {
        input: `Time:  7  15   30
Distance:  9  40  200`,
        expected: 288,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `Time:      71530
Distance:  940200`,
        expected: 71503,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
