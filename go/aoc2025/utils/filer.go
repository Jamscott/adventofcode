package utils

import (
	"fmt"
	"os"
)

func LoadInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", filename, err)
	}
	return string(data), nil
}
