declare global {
  interface Array<T> {
    /** Sums all numbers in the array */
    sum(): number;
    /** Splits each string in the array into an array of characters */
    splitChars(): string[][];
    /** Splits the array into an array of arrays, each representing a line */
    splitNewLines(): string[][];
    /** times each number in the array by each number in the array */
    power(): number;

    mapToNumber(): number[];
  }
  interface String {
    /** Splits a string into an array of characters */
    splitChars(): string[];
    /** Splits a string into an array of arrays, each representing a line */
    splitNewLines(): string[];
    /** Reverses a string */
    reverse(): string;
  }
  interface Map<K, V> {
    /** Merges the highest value of each key into the map
     * This replaces the value of the current map with the value of the new map if the new value is higher
     */
    mergeHighestValue(map: Map<K, V>): Map<K, V>;
    /** Sums all values in the map */
    sumValues(): number;
    /** times the value of each key by each value */
    powerValues(): number;
  }
}

Map.prototype.mergeHighestValue = function (map: Map<any, any>) {
  for (const [key, value] of map.entries()) {
    if (this.has(key)) {
      const current = this.get(key);
      if (current < value) {
        this.set(key, value);
      }
    } else {
      this.set(key, value);
    }
  }
  return this;
};

Map.prototype.sumValues = function () {
  return Array.from(this.values()).sum();
};

Map.prototype.powerValues = function () {
  return Array.from(this.values()).power();
};

String.prototype.splitChars = function () {
  return this.split("");
};

String.prototype.splitNewLines = function () {
  return this.split("\n");
};

String.prototype.reverse = function () {
  return this.splitChars().reverse().join("");
};

Array.prototype.power = function () {
  return this.reduce((acc, number) => {
    return acc * number;
  }, 1);
};

Array.prototype.sum = function () {
  return this.reduce((a, b) => a + b, 0);
};

Array.prototype.splitChars = function () {
  return this.map((line) => line.split(""));
};

Array.prototype.splitNewLines = function () {
  return this.map((line) => line.split("\n"));
};

Array.prototype.mapToNumber = function () {
  return this.map((v) => Number(v));
};

export {};
