package utils

import "slices"

func ReverseString(input string) string {
	runes := []rune(input)
	slices.Reverse(runes)
	return string(runes)
}
