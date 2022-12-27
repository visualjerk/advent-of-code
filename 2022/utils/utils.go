package utils

import (
	"os"
	"strconv"
)

func LoadData() string {
	file := "input"

	if os.Getenv("TEST") != "" {
		file = "test_input"
	}

	data, error := os.ReadFile(file)

	if error != nil {
		return ""
	}

	return string(data)
}

func SafeStringToInt(input string) int {
	marks, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}
	return marks
}
