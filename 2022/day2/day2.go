package main

import (
	"fmt"
	"os"
)

func loadData() string {
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

func main() {
	data := loadData()

	fmt.Println(data)
}
