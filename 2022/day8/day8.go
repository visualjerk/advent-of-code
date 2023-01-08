package main

import (
	"aoc.io/utils"
	"strings"
)

type Tree struct {
	row    int
	col    int
	height int
}

type TreeGroup []Tree

func (treeGroup *TreeGroup) addTree(tree Tree) {
	*treeGroup = append(*treeGroup, tree)
}

func (treeGroup *TreeGroup) getCol(col int) TreeGroup {
	result := TreeGroup{}
	for _, tree := range *treeGroup {
		if tree.col == col {
			result.addTree(tree)
		}
	}
	return result
}

func (treeGroup *TreeGroup) getRow(row int) TreeGroup {
	result := TreeGroup{}
	for _, tree := range *treeGroup {
		if tree.row == row {
			result.addTree(tree)
		}
	}
	return result
}

func (treeGroup *TreeGroup) getSurroundingGroups(tree Tree) []TreeGroup {
	col := treeGroup.getCol(tree.col)
	top := col[0:tree.row]
	bottom := col[tree.row+1:]

	row := treeGroup.getRow(tree.row)
	left := row[0:tree.col]
	right := row[tree.col+1:]

	return []TreeGroup{
		top,
		bottom,
		left,
		right,
	}
}

func (treeGroup *TreeGroup) isLowerThan(tree Tree) bool {
	for _, coverTree := range *treeGroup {
		if tree.height <= coverTree.height {
			return false
		}
	}
	return true
}

func (treeGroup *TreeGroup) isTreeVisible(tree Tree) bool {
	surroundingGroups := treeGroup.getSurroundingGroups(tree)
	for _, group := range surroundingGroups {
		if group.isLowerThan(tree) {
			return true
		}
	}
	return false
}

func (treeGroup *TreeGroup) countVisibleTrees() int {
	count := 0
	for _, tree := range *treeGroup {
		if treeGroup.isTreeVisible(tree) {
			count++
		}
	}
	return count
}

func parseData(data string) TreeGroup {
	treeGroup := TreeGroup{}
	rows := strings.Split(data, "\n")

	for rowIndex, row := range rows {
		trees := strings.Split(row, "")
		for colIndex, rawHeight := range trees {
			height := utils.SafeStringToInt(rawHeight)
			treeGroup.addTree(Tree{
				row:    rowIndex,
				col:    colIndex,
				height: height,
			})
		}
	}

	return treeGroup
}

func main() {
	data := utils.LoadData()
	treeGrid := parseData(data)
	visibleTreeCount := treeGrid.countVisibleTrees()

	println(visibleTreeCount)
}
