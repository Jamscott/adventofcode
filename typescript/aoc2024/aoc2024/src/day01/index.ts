import run from "aocrunner";
import  "../utils/helpers.js";

const parseInput = (rawInput: string) => rawInput
  .splitNewLines()
  .map((line) => line.split("   ").map((n) => parseInt(n)));

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput)
    .reduce((acc, line) => {
      const [a , b] = line;
      acc.listOne.push(a);
      acc.listTwo.push(b);
      return acc;
    }, { listOne: [], listTwo: [] } as { listOne: number[], listTwo: number[] });

  const listOne = input.listOne.sort((a, b) => a - b);
  const listTwo = input.listTwo.sort((a, b) => a - b);

  let total = 0;

  for (let i = 0; i < listOne.length; i++) {
    const diff = Math.abs(listOne[i] - listTwo[i]);
    total += diff
  }

  // console.log(total);
  

  return total
};

const calculateHowManyTimesNumberIsInArray = (array: number[], number: number) => {
  return array.filter((n) => n === number).length;
}

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput)
    .reduce((acc, line) => {
      const [a , b] = line;
      acc.listOne.push(a);
      acc.listTwo.push(b);
      return acc;
    }, { listOne: [], listTwo: [] } as { listOne: number[], listTwo: number[] });

  const listOne = input.listOne.sort((a, b) => a - b);
  const listTwo = input.listTwo.sort((a, b) => a - b);

  let similarity: number[] = [];
  
  for (let i = 0; i < listOne.length; i++) {
    const similarityNum = calculateHowManyTimesNumberIsInArray(listTwo, listOne[i]);
    console.log(similarityNum);
    similarity.push(similarityNum * listOne[i]);
  }

  return similarity.sum();
};

run({
  part1: {
    tests: [
      {
        input: `3   4
4   3
2   5
1   3
3   9
3   3`,
        expected: 11,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `3   4
4   3
2   5
1   3
3   9
3   3`,
        expected: 31,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
