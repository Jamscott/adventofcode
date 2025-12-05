package utils

import "slices"

// ReverseString reverses a string, handling Unicode correctly
func ReverseString(input string) string {
	runes := []rune(input)
	slices.Reverse(runes)
	return string(runes)
}
