import run from "aocrunner";
import '../utils/helpers.js';

const parseInput = (rawInput: string) => rawInput
  .splitNewLines()

type BlockCollection = Map<string, number>[]

const limits: Record<string, number> = {
  red: 12,
  green: 13,
  blue: 14,
}

class Game {
  id: number;
  blocks: BlockCollection;
  input: string;

  constructor(input: string) {
    this.input = input;
    const [id, rounds] = input.split(':')
    this.id = Number(id.replace(/\D/g, '')); 
    this.blocks = this.parseRounds(rounds);
  }

  private parseRounds(blocks: string): BlockCollection {
    const rolls = blocks
      .split(';')
      .map((round) => round
        .split(',')
        .map((roll) => roll.trim())
        .map((roll) => roll.split(' ')))
       
    const blocksCollection = rolls
      .map((round) => round
      .reduce((acc, [number, colour]) => {
        acc.set(colour, Number(number));
        return acc;
      }, new Map<string, number>()));

    return blocksCollection
  }

  public isPossible() {
    return this.blocks
      .every((round) => Array.from(round.entries())
      .every(([colour, number]) => number <= limits[colour]));
  }

  public power() {
    const highestColours = this.blocks
      .reduce((acc, round) => {
        return acc.mergeHighestValue(round);
      }, new Map<string, number>())

    return highestColours.powerValues();
  }
}

const part1 = (rawInput: string) => {
  return parseInput(rawInput)  
    .map((line) => new Game(line))
    .filter((game) => game.isPossible())
    .map((game) => game.id)
    .sum();
};

const part2 = (rawInput: string) => {
  return parseInput(rawInput)
    .map((line) => new Game(line))
    .map((game) => game.power())
    .sum()
};

run({
  part1: {
    tests: [
      {
        input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
        expected: 8,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
        expected: 2286,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
