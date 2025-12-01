import run from "aocrunner";
import '../utils/helpers.js';

const parseInput = (rawInput: string) => rawInput
  .splitNewLines();

const cards = ["2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q","K","A"];
const cards2 = ["J", "2", "3", "4", "5", "6", "7", "8", "9", "T" ,"Q","K","A"];

const countCards = (cards: string[], replace: boolean) => {
  const count = cards.reduce((acc, card) => {
      if (acc[card]) {
        acc[card] += 1;
      } else {
        acc[card] = 1;
      }
      return acc;
    }, {} as Record<string, number>)


  if (replace) {
    if (count["J"] === 5) {
      return [count["J"]];
    }

    if (count["J"]) {

      const highestCard = Object.keys(count).sort((a, b) => cards2.indexOf(b) - cards2.indexOf(a))[0];

      count[highestCard] += count["J"];
      if (highestCard !== "J") {
        delete count["J"];
      }
    }
  }
  
    
  return Object.values(count).sort((a, b) => b - a);
}

const cardHandTypeValue = (cardCounts: number[]) => {

  if (cardCounts.length === 1) {
    return 6;
  }

  if (cardCounts.length === 2) {
    if (cardCounts.includes(4)) {
      return 5;
    } else {
      return 4;
    }
  }

  if (cardCounts.length === 3) {
    if (cardCounts.includes(3)) {
      return 3;
    } else {
      return 2;
    }
  }

  if (cardCounts.length === 4) {
    return 1;
  }

  if (cardCounts.length === 5) {
    return 0;
  }

  return -1;
}

type Hand = {
  cards: string[];
  bid: number;
  value: number;
}

const getStrengthOfCard = (card: string, deck: string[]) => {
  return deck.indexOf(card) + 1
}

const parseHand = (line: string, replace: boolean): Hand => {
  const [hand, bid] = line.split(" ");

  const value = cardHandTypeValue(
    countCards(
      hand.split(""),
      replace
    )
  )

  return {
    cards: hand.split(""),
    bid: parseInt(bid),
    value
  };
}

const compareHands = (hand1: Hand, hand2: Hand, ranking: boolean) => {
  if (hand1.value > hand2.value) {
    return 1;
  }

  if (hand1.value < hand2.value) {
    return -1;
  }

  if (hand1.value === hand2.value) {
    for (let i = 0; i < hand1.cards.length; i++) {
      const card1Strength = getStrengthOfCard(hand1.cards[i], cards);
      const card2Strength = getStrengthOfCard(hand2.cards[i], cards);
      if (card1Strength > card2Strength) {
        return 1;
      }
      if (card1Strength < card2Strength) {
        return -1;
      }
    }
  }
  return 0;
}

const part1 = (rawInput: string) => {
  return parseInput(rawInput)
    .map((line) => parseHand(line, false))
    .sort((hand1, hand2) => compareHands(hand1, hand2, false))
    .map((hand, index) => hand.bid * (index + 1))
    .sum();
};

const part2 = (rawInput: string) => {
  return parseInput(rawInput)
    .map((line) => parseHand(line, true))
    .sort((hand1, hand2) => compareHands(hand1, hand2, true))
    .map((hand, index) => hand.bid * (index + 1))
    .sum();
};

run({
  part1: {
    tests: [
      {
        input: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`,
        expected: 6440,
      },
    ],
    solution: part1,
  },
  part2: {
    tests: [
      {
        input: `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`,
        expected: 5905,
      },
    ],
    solution: part2,
  },
  trimTestInputs: true,
  onlyTests: false,
});
