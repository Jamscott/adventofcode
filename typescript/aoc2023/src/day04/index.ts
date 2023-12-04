import run from "aocrunner";
import "../utils/helpers.js";

const parseInput = (rawInput: string) => rawInput
  .splitNewLines()

const matchNumber = (match: RegExp, text: string): number[] => {
  return Array
    .from(text.matchAll(match))
    .map(m => m[0].split(" "))
    .flat()
    .filter((n) => n)
    .map((n) => parseInt(n.trim()))
}  

const winningNumberRegex = /(?<=: ).*(?= \|)/g;
const elfsNumbersRegex = /(?<=\| ).*/g; 

const processCard = (card: string) => {
  const winningNumbers = matchNumber(winningNumberRegex, card)
  const elfsNumbers = matchNumber(elfsNumbersRegex, card)
  return [winningNumbers, elfsNumbers]
}

const part1 = (rawInput: string) => {
  return parseInput(rawInput)
    .map((line) => {
      const [winningNumbers, elfsNumbers] = processCard(line)
      return elfsNumbers.reduce((acc, num) => acc + (winningNumbers.includes(num) ? 1 : 0), 0)
    })
    .reduce((acc, num) => num === 0 ? acc : acc + Math.pow(2, num - 1), 0)
};

const part2 = (rawInput: string) => {
  const lines = parseInput(rawInput);
  const scatchCardPile: number[] = Array(lines.length).fill(1);
  for (let x = 0; x < lines.length; x++) {
    const line = lines[x];
    const  [winningNumbers, elfsNumbers] = processCard(line)
    const matchingNumbers = elfsNumbers.filter((x) => winningNumbers.includes(x));
    const scatchCard = scatchCardPile[x];
    for (let i: number = 0; i < matchingNumbers.length; i++) {
      scatchCardPile[x + i + 1] = scatchCardPile[x + i + 1] + scatchCard;
    }
  }
  return scatchCardPile.sum()
};

run({
  part1: {
    tests: [
      {
        input: `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`,
        expected: 13,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 1`,
        expected: 30,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
