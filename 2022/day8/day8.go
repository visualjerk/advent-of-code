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

func (treeGroup *TreeGroup) reverse() TreeGroup {
	result := TreeGroup{}
	for i := len(*treeGroup) - 1; i >= 0; i-- {
		tree := (*treeGroup)[i]
		result = append(result, tree)
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
		top.reverse(),
		bottom,
		left.reverse(),
		right,
	}
}

func (treeGroup *TreeGroup) getViewingDistance(tree Tree) int {
	distance := 0
	for _, viewTree := range *treeGroup {
		distance++
		if tree.height <= viewTree.height {
			return distance
		}
	}
	return distance
}

func (treeGroup *TreeGroup) getScenicScore(tree Tree) int {
	score := 1
	surroundingGroups := treeGroup.getSurroundingGroups(tree)
	for _, group := range surroundingGroups {
		score = score * group.getViewingDistance(tree)
	}
	return score
}

func (treeGroup *TreeGroup) getBestScenicScore() int {
	bestScore := 0
	for _, tree := range *treeGroup {
		treeScore := treeGroup.getScenicScore(tree)
		if bestScore < treeScore {
			bestScore = treeScore
		}
	}
	return bestScore
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
	visibleTreeCount := treeGrid.getBestScenicScore()

	println(visibleTreeCount)
}
