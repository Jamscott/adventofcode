package utils

import (
	"fmt"
	"os"
	"strings"
)

func LoadInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", filename, err)
	}
	return string(data), nil
}

func NewLineSplit(input string) []string {
	return strings.Split(input, "\n")
}

func SpaceSplit(input string) []string {
	return strings.Split(input, " ")
}
