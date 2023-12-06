const DOUBLENEWLINE = /\r?\n\r?\n/;
const NEWLINE = /\r?\n/;
const WHITESPACE = /\s/g;
const DIGIT = /\d/g;
const LETTER = /[a-z]/g;
const ALPHANUMERIC = /[a-z0-9]/g;

export const matchNumber = (match: RegExp, text: string): number[] => {
    return Array
      .from(text.matchAll(match))
      .map(m =>  parseInt(m[0]))
  }  

export {
    DOUBLENEWLINE,
    NEWLINE,
    WHITESPACE,
    DIGIT,
    LETTER,
    ALPHANUMERIC,
}