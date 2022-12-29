package main

import (
	"aoc.io/utils"
	"fmt"
	"strings"
)

func findSignalIndex(data string) int {
	characters := strings.Split(data, "")
	seenChars := map[string]bool{}

	for i := 0; i < len(characters); i++ {
		char := characters[i]
		if seenChars[char] {
			seenChars = map[string]bool{}
		}
		seenChars[char] = true
		if len(seenChars) == 4 {
			return i + 1
		}
	}
	return -1
}

func main() {
	data := utils.LoadData()
	signalIndex := findSignalIndex(data)

	fmt.Println(signalIndex)
}
