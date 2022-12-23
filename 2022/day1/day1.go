package main

import (
	"aoc.io/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getBags(data string) []string {
	return strings.Split(data, "\n\n")
}

func sumCaloriesInBag(bag string) int {
	snacks := strings.Split(bag, "\n")
	calories := 0

	for i := 0; i < len(snacks); i++ {
		snackCalories, err := strconv.Atoi(snacks[i])
		if err != nil {
			return 0
		}
		calories += snackCalories
	}

	return calories
}

func getBagCalories(bags []string) []int {
	bagCalories := []int{}

	for i := 0; i < len(bags); i++ {
		calories := sumCaloriesInBag(bags[i])
		bagCalories = append(bagCalories, calories)
	}

	return bagCalories
}

func getTopCalories(calories []int, count int) []int {
	sort.Ints(calories)
	topCalories := []int{}

	for i := 0; i < count; i++ {
		topCalories = append(topCalories, calories[len(calories)-(i+1)])
	}

	return topCalories
}

func sumCalories(calories []int) int {
	sum := 0

	for i := 0; i < len(calories); i++ {
		sum += calories[i]
	}

	return sum
}

func main() {
	data := utils.LoadData()
	bags := getBags(data)
	bagCalories := getBagCalories(bags)
	topCalories := getTopCalories(bagCalories, 3)
	totalCalories := sumCalories(topCalories)

	fmt.Println(totalCalories)
}
