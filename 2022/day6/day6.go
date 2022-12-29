package main

import (
	"aoc.io/utils"
	"fmt"
	"strings"
)

func findIndexOfNthDistinctChars(data string, count int) int {
	characters := strings.Split(data, "")
	seenChars := map[string]int{}

	for i := 0; i < len(characters); {
		char := characters[i]
		charSeenIndex, charSeen := seenChars[char]
		if charSeen {
			seenChars = map[string]int{}
			i = charSeenIndex
		} else {
			seenChars[char] = i
		}
		if len(seenChars) == count {
			return i + 1
		}
		i++
	}
	return -1
}

func findSignalIndex(data string) int {
	return findIndexOfNthDistinctChars(data, 4)
}

func findMessageIndex(data string) int {
	return findIndexOfNthDistinctChars(data, 14)
}

func main() {
	data := utils.LoadData()
	messageIndex := findMessageIndex(data)

	fmt.Println(messageIndex)
}
