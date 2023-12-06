import run from "aocrunner";
import "../utils/helpers.js";

type Range = { destination: number; source: number; range: number };

const parseChunk = (maps: string[]): Range[] => {
  const result: Range[] = maps
    .map((destinationResourceLine: string) => {
      const [destination, source, range] = destinationResourceLine
        .split(" ")
        .mapToNumber();
      return { destination, source, range };
    })
    .sort((a: Range, b: Range) => a.source - b.source);

  for (let i = 0; i < result.length; i++) {
    const val = result[i];
    const nextVal = result[i + 1];

    if (nextVal) {
      if (nextVal.source <= val.source + val.range) {
        val.range -= nextVal.source - (val.source + val.range) + 1;
      }
    }
  }

  return result;
};

const mapResources = (rawInput: string) => {
  const lines = rawInput.splitNewLines();

  const seeds = lines[0]
    .split(" ")
    .map((seed) => parseInt(seed))
    .filter((seed) => !isNaN(seed));

  const mappedResourceRanges: [string, Range[]][] = rawInput
    .split("\n\n")
    .splice(1)
    .map((chunk) => chunk.splitNewLines())
    .map(([resourceName, ...resourceLines]) => {
      const name = resourceName.split(" ")[0];
      return [name, parseChunk(resourceLines)];
    });

  return {
    seeds,
    mappedResourceRanges,
  };
};

const intersectsRange = (range: Range, index: number) => {
  return range.source <= index && range.source + range.range >= index;
};

const getValue = (
  ranges: Range[],
  index: number,
  chunkName: string,
  lastValues: Record<string, Range>,
): number => {
  const prev = lastValues[chunkName];

  if (prev && intersectsRange(prev, index)) {
    return prev.destination + (index - prev.source); 
  }

  for (let i = 0; i < ranges.length; i++) {
    const range = ranges[i];

    if (range.source > index) {
      return index;
    }

    if (intersectsRange(range, index)) {
      lastValues[chunkName] = range;
      return range.destination + (index - range.source) 
    }
  }

  return index;
};

const part1 = (rawInput: string) => {
  const { seeds, mappedResourceRanges } = mapResources(rawInput);

  const lastValuesMappings: Record<string, Range> = {};

  return seeds.reduce((acc, seed) => {
    let index = mappedResourceRanges.reduce((acc, [name, ranges]) => {
      return getValue(ranges, acc, name, lastValuesMappings);
    }, seed);

    if (index < acc) {
      return index;
    }

    return acc;
  }, Infinity); 
};

const part2 = (rawInput: string) => {
  const { seeds, mappedResourceRanges } = mapResources(rawInput);

  const lastValuesMappings: Record<string, Range> = {};

  let closest: number = Infinity

  //[79-14] [55-13]
  for (let i = 0; i < seeds.length; i += 2) {
    const start = seeds[i];
    const range = seeds[i + 1];

    for (let seed = start; seed < start + range; seed++) {
      let index = seed;

      for (let x = 0; x < mappedResourceRanges.length; x++) {
        const [name, ranges] = mappedResourceRanges[x];
        index = getValue(ranges, index, name, lastValuesMappings);
      }

      if (index < closest) {
        closest = index;
      }
    }
  }

  return closest;
};

run({
  part1: {
    tests: [
      {
        input: `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`,
        expected: 35,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`,
        expected: 46,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: true,
});
