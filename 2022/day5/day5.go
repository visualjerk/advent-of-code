package main

import (
	"fmt"
	"strings"

	"aoc.io/utils"
)

func parseRawStacks(rawStacks string) [][]string {
	lines := strings.Split(rawStacks, "\n")

	indexLine := lines[len(lines)-1]
	indexes := strings.Split(strings.TrimSpace(indexLine), "   ")

	stacks := [][]string{}
	for i := 0; i < len(indexes); i++ {
		stacks = append(stacks, []string{})
	}

	crateLines := lines[:len(lines)-1]
	for i := 0; i < len(crateLines); i++ {
		crates := strings.Split(strings.ReplaceAll(crateLines[i], "    ", " "), " ")

		for u := 0; u < len(stacks); u++ {
			if crates[u] != "" {
				stacks[u] = append(stacks[u], crates[u])
			}
		}
	}

	return stacks
}

func main() {
	data := utils.LoadData()
	dataParts := strings.Split(data, "\n\n")
	stacks := parseRawStacks(dataParts[0])

	fmt.Println(stacks)
}
