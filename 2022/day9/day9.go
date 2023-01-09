package main

import (
	"strings"
	"time"

	"aoc.io/utils"
)

const GRID_LENGTH = 50
const FRAME_TIME = time.Millisecond * 100

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

var MOVE_INDICATOR = map[Direction]string{
	UP:    "⬆️",
	DOWN:  "⬇️",
	LEFT:  "⬅️",
	RIGHT: "➡️",
}

type Move struct {
	direction Direction
	distance  int
}

type Head struct {
	x int
	y int
}

func (head *Head) move(direction Direction) {
	switch direction {
	case UP:
		head.y++
	case DOWN:
		head.y--
	case RIGHT:
		head.x++
	case LEFT:
		head.x--
	}
}

func (head *Head) renderState() {
	for y := 0; y < GRID_LENGTH; y++ {
		for x := 0; x < GRID_LENGTH; x++ {
			if x == head.x && y == head.y {
				print("H")
			} else {
				print(".")
			}
		}
		print("\n")
	}
}

func executeMoves(moves []Move) {
	head := Head{GRID_LENGTH / 2, GRID_LENGTH / 2}
	for moveIndex, move := range moves {
		for i := 0; i < move.distance; i++ {
			head.move(move.direction)
			head.renderState()
			println("move ", moveIndex+1, "/", len(moves), " | dir ", MOVE_INDICATOR[move.direction], " | distance ", move.distance)
			time.Sleep(FRAME_TIME)
		}
	}
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
	executeMoves(moves)

	println(moves)
}
