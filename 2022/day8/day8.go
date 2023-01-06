package main

import (
	"aoc.io/utils"
	"strings"
)

type Tree struct {
	height int
	hidden bool
}

type TreeGrid [][]Tree

func (grid *TreeGrid) addTree(rowIndex, height int) {
	if len(*grid) < rowIndex+1 {
		*grid = append(*grid, []Tree{})
	}
	(*grid)[rowIndex] = append((*grid)[rowIndex], Tree{
		height: height,
	})
}

func parseData(data string) TreeGrid {
	grid := TreeGrid{}
	rows := strings.Split(data, "\n")

	for rowIndex, row := range rows {
		trees := strings.Split(row, "")
		for _, rawHeight := range trees {
			height := utils.SafeStringToInt(rawHeight)
			grid.addTree(rowIndex, height)
		}
	}

	return grid
}

func main() {
	data := utils.LoadData()
	treeGrid := parseData(data)

	println(treeGrid)
}
