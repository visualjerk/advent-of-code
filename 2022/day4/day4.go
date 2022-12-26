package main

import (
	"aoc.io/utils"
	"fmt"
	sf "github.com/sa-/slicefunk"
	"strconv"
	"strings"
)

func safeStringToInt(input string) int {
	marks, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}
	return marks
}

func parseRawSection(rawSection string) [2]int {
	section := sf.Map(strings.Split(rawSection, "-"), safeStringToInt)
	return [2]int{section[0], section[1]}
}

func parseRawPairSection(rawPairSection string) [2][2]int {
	rawSections := strings.Split(rawPairSection, ",")
	sections := sf.Map(rawSections[:2], parseRawSection)
	return [2][2]int{sections[0], sections[1]}
}

func getPairSections(input string) [][2][2]int {
	rawPairSections := strings.Split(input, "\n")
	return sf.Map(rawPairSections, parseRawPairSection)
}

func sectionsAreOverlapping(section1 [2]int, section2 [2]int) bool {
	return !(section1[1] < section2[0] || section2[1] < section1[0])
}

func hasOverlappingPair(pairSection [2][2]int) bool {
	section1 := pairSection[0]
	section2 := pairSection[1]

	return sectionsAreOverlapping(section1, section2)
}

func getOverlappingPairs(pairSections [][2][2]int) [][2][2]int {
	return sf.Filter(pairSections, hasOverlappingPair)
}

func main() {
	data := utils.LoadData()
	pairSections := getPairSections(data)
	fullyOverlappingPairs := getOverlappingPairs(pairSections)

	fmt.Println(len(fullyOverlappingPairs))
}
