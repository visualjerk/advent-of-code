package main

import (
	"aoc.io/utils"
	"fmt"
	"math"
	"strings"
)

type Stack []string

type Rearrangement struct {
	crateAmount int
	fromStack   int
	toStack     int
}

func parseRawCrate(crate string) string {
	return strings.Trim(strings.Trim(crate, "["), "]")
}

func parseRawStacks(rawStacks string) []Stack {
	lines := strings.Split(rawStacks, "\n")

	indexLine := lines[len(lines)-1]
	indexes := strings.Split(strings.TrimSpace(indexLine), "   ")

	stacks := []Stack{}
	for i := 0; i < len(indexes); i++ {
		stacks = append(stacks, Stack{})
	}

	crateLines := lines[:len(lines)-1]
	for i := len(crateLines) - 1; i >= 0; i-- {
		crates := strings.Split(strings.ReplaceAll(crateLines[i], "    ", " "), " ")

		for u := 0; u < len(stacks); u++ {
			if crates[u] != "" {
				stacks[u] = append(stacks[u], parseRawCrate(crates[u]))
			}
		}
	}

	return stacks
}

func parseRawRearrangements(rawRearrangements string) []Rearrangement {
	rearrangements := []Rearrangement{}

	lines := strings.Split(rawRearrangements, "\n")
	for i := 0; i < len(lines); i++ {
		lineParts := strings.Split(lines[i], " ")
		rearrangement := Rearrangement{
			crateAmount: utils.SafeStringToInt(lineParts[1]),
			fromStack:   utils.SafeStringToInt(lineParts[3]) - 1,
			toStack:     utils.SafeStringToInt(lineParts[5]) - 1,
		}
		rearrangements = append(rearrangements, rearrangement)
	}

	return rearrangements
}

func applyRearrangement(stacks []Stack, r Rearrangement) []Stack {
	index := int(math.Max(float64(len(stacks[r.fromStack])-r.crateAmount), 0))
	stacks[r.toStack] = append(stacks[r.toStack], stacks[r.fromStack][index:]...)
	stacks[r.fromStack] = stacks[r.fromStack][:index]
	return stacks
}

func applyRearrangements(stacks []Stack, rearrangements []Rearrangement) []Stack {
	resultStacks := stacks
	for i := 0; i < len(rearrangements); i++ {
		rearrangement := rearrangements[i]
		resultStacks = applyRearrangement(resultStacks, rearrangement)
	}
	return resultStacks
}

func getTopCrateTypes(stacks []Stack) []string {
	topCrates := []string{}
	for i := 0; i < len(stacks); i++ {
		topCrates = append(topCrates, stacks[i][len(stacks[i])-1])
	}
	return topCrates
}

func main() {
	data := utils.LoadData()
	dataParts := strings.Split(data, "\n\n")

	stacks := parseRawStacks(dataParts[0])
	rearrangements := parseRawRearrangements(dataParts[1])

	finalStacks := applyRearrangements(stacks, rearrangements)
	topCrateTypes := getTopCrateTypes(finalStacks)

	fmt.Println(strings.Join(topCrateTypes, ""))
}
