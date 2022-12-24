package main

import (
	"aoc.io/utils"
	"fmt"
	sf "github.com/sa-/slicefunk"
	"strings"
)

func getRucksacks(data string) [][]string {
	lines := strings.Split(data, "\n")
	rucksacks := [][]string{}

	for i := 0; i < len(lines); i++ {
		items := strings.Split(lines[i], "")
		rucksacks = append(rucksacks, items)
	}

	return rucksacks
}

func getSharedItem(itemGroups [][]string) string {
	firstItemGroup := itemGroups[0]
	remainingItemGroups := sf.Map(itemGroups[1:], func(group []string) string { return strings.Join(group, "") })

	sharedItems := sf.Filter(firstItemGroup, func(item string) bool {
		foundInGroups := sf.Filter(remainingItemGroups, func(group string) bool {
			return strings.Contains(group, item)
		})
		return len(foundInGroups) == len(remainingItemGroups)
	})

	return sharedItems[0]
}

func getSharedItems(rucksacks [][][]string) []string {
	return sf.Map(rucksacks, getSharedItem)
}

func getRucksackGroups(rucksacks [][]string, groupCount int) [][][]string {
	groups := [][][]string{}
	groupIndex := 0

	for i := 0; i < len(rucksacks); i++ {
		if len(groups) < groupIndex+1 {
			groups = append(groups, [][]string{})
		}
		groups[groupIndex] = append(groups[groupIndex], rucksacks[i])

		if (i+1)%groupCount == 0 {
			groupIndex++
		}
	}

	return groups
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
	rucksackGroups := getRucksackGroups(rucksacks, 3)
	sharedItems := getSharedItems(rucksackGroups)
	totalPriority := getTotalPriority(sharedItems)

	fmt.Println(totalPriority)
}
