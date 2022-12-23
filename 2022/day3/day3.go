package main

import (
	"aoc.io/utils"
	"fmt"
	"strings"
)

func getRucksacks(data string) [][2][]string {
	lines := strings.Split(data, "\n")
	rucksacks := [][2][]string{}

	for i := 0; i < len(lines); i++ {
		items := strings.Split(lines[i], "")
		half := len(items) / 2
		compartments := [2][]string{items[:half], items[half:]}
		rucksacks = append(rucksacks, compartments)
	}

	return rucksacks
}

func getWrongItem(rucksack [2][]string) string {
	firstCompartment := rucksack[0]
	secondCompartment := strings.Join(rucksack[1], "")

	var wrongItem string

	for i := 0; i < len(firstCompartment); i++ {
		item := firstCompartment[i]
		if strings.Contains(secondCompartment, item) {
			wrongItem = item
		}
	}

	return wrongItem
}

func getWrongItems(rucksacks [][2][]string) []string {
	wrongItems := []string{}

	for i := 0; i < len(rucksacks); i++ {
		wrongItem := getWrongItem(rucksacks[i])
		wrongItems = append(wrongItems, wrongItem)
	}

	return wrongItems
}

func getItemPriority(item string) int {
	itemCharPos := int([]rune(item)[0])

	// Convert lowercase chars to 1 - 26
	if itemCharPos > 96 {
		return itemCharPos - 96
	}
	// Convert uppercase chars to 27 - 52
	return itemCharPos - 38
}

func getTotalPriority(items []string) int {
	totalPriority := 0

	for i := 0; i < len(items); i++ {
		totalPriority += getItemPriority(items[i])
	}

	return totalPriority
}

func main() {
	data := utils.LoadData()
	rucksacks := getRucksacks(data)
	wrongItems := getWrongItems(rucksacks)
	totalPriority := getTotalPriority(wrongItems)

	fmt.Println(totalPriority)
}
