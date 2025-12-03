package utils

import "strings"

func NewLineSplit(input string) []string {
	return strings.Split(input, "\n")
}

func SpaceSplit(input string) []string {
	return strings.Split(input, " ")
}

func CommaSplit(input string) []string {
	return strings.Split(input, ",")
}

func CharSplit(input string) []string {
	return strings.Split(input, "")
}

func ReverseString(input string) string {
	chars := CharSplit(input)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return strings.Join(chars, "")
}
