package main

import (
	"strings"

	"aoc.io/utils"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

var RAW_DIRECTION_MAP = map[string]Direction{
	"U": UP,
	"D": DOWN,
	"L": LEFT,
	"R": RIGHT,
}

type Move struct {
	direction Direction
	distance  int
}

func parseRawMove(data string) Move {
	parts := strings.Split(data, " ")

	return Move{
		direction: RAW_DIRECTION_MAP[parts[0]],
		distance:  utils.SafeStringToInt(parts[1]),
	}
}

func parseData(data string) (moves []Move) {
	rawMoves := strings.Split(data, "\n")

	for _, rawMove := range rawMoves {
		moves = append(moves, parseRawMove(rawMove))
	}

	return moves
}

func main() {
	data := utils.LoadData()
	moves := parseData(data)

	println(moves)
}
