import run from "aocrunner";
import '../utils/helpers.js';

const parseInput = (rawInput: string) => rawInput
  .splitNewLines()

const symbolMatch = /([^\w^.])/g;
const numberMatch = /(\d{1,9})/g;

type Match =  {
  x: number
  y: number
}

type NumberMatch = Match & {
  match: number
}

type StringMatch = Match & {
  match: string
}

const matchNumber = (match: RegExp, text: string, y?: number): NumberMatch[] => {
  return Array
    .from(text.matchAll(match))
    .map(m => ({ 
      match: parseInt(m[0]),
      x: m.index || 0,
      y: y || 0 
    }))
}  

const matchString = (match: RegExp, text: string, y?: number): StringMatch[] => {
  return Array
    .from(text.matchAll(match))
    .map(m => ({ 
      match: m[0].toString(),
      x: m.index || 0, 
      y: y || 0 
    }))
}

type Direction = {
  x: number, 
  y: number
}

const directions: Direction[] =  [
  { x: 1, y: 0 },
  { x: 1, y: -1 },
  { x: 0, y: -1 },
  { x: -1, y: -1 },
  { x: -1, y: 0 },
  { x: -1, y: 1 },
  { x: 0, y: 1 },
  { x: 1, y: 1 }
]      

const filterNumbersWithSymbol = (number: NumberMatch, symbols: StringMatch[]) => {
  const length = number.match.toString().length;

  for (let i = 0; i < length; i++) {
      const found = directions
        .some((direction: Direction) => symbols
        .find(symbol => symbol.x === number.x + i + direction.x && symbol.y === number.y + direction.y))
      
      if (found) { 
        return true;
      }
  }
  return false;
}

const part1 = (rawInput: string) => {
  const input = parseInput(rawInput)

  let symbols = input
    .map((line, index) => matchString(symbolMatch, line, index))
    .flat()

  let numbers = input
    .map((line, index) => matchNumber(/(\d{1,9})/g, line, index))
    .flat()
    .filter(number => filterNumbersWithSymbol(number, symbols))
    
  return numbers
    .map((number: any) => number.match)
    .sum();
};

type SymbolMap = {
  parts: number[];
  match: string;
  x: number;
  y: number;
}[]

function filterNumbersPart2(number: NumberMatch, symbols: SymbolMap) {
  const length = number.match.toString().length;

  for (let i = 0; i < length; i++) {
      const found = directions.some((direction: Direction)=> {
          const symbol = symbols.find(symbol => symbol.x === number.x + i + direction.x && symbol.y === number.y + direction.y)
          if (symbol) {
              symbol.parts.push(number.match)
              return true;
          }
          return false;
      })
      if (found) return true;
  }
  return false;
}

const part2 = (rawInput: string) => {
  const input = parseInput(rawInput)
  
  let symbols = input
    .map((line, index) => matchString(symbolMatch, line, index))
    .flat()
    .map(symbol => { return {...symbol, parts: [] as number[]}} )

  input
    .map((line, index) => matchNumber(numberMatch, line, index))
    .flat()
    .filter(number => filterNumbersPart2(number, symbols))
  
  return symbols
    .filter(symbol => symbol.match === '*' && symbol.parts.length === 2)
    .map(symbol=> symbol.parts.power())
    .sum();
};

run({
  part1: {
    tests: [
      {
        input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
        expected: 4361,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
        expected: 467835,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
