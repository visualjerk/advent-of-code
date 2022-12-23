package main

import (
	"fmt"
	"os"
	"strings"
)

func loadData() string {
	file := "input"

	if os.Getenv("TEST") != "" {
		file = "test_input"
	}

	data, error := os.ReadFile(file)

	if error != nil {
		return ""
	}

	return string(data)
}

type Shape int

const (
	ROCK Shape = iota
	PAPER
	SCISSORS
)

var OPPONENT_SHAPES = map[string]Shape{
	"A": ROCK,
	"B": PAPER,
	"C": SCISSORS,
}

var MY_SHAPES = map[string]Shape{
	"X": ROCK,
	"Y": PAPER,
	"Z": SCISSORS,
}

var SHAPES_SCORE = map[Shape]int{
	ROCK: 1,
	PAPER: 2,
	SCISSORS: 3,
}

type RoundOutcome int

const (
	LOSS RoundOutcome = iota
	DRAW
	WIN
)

var OUTCOME_SCORE = map[RoundOutcome]int{
	LOSS: 0,
	DRAW: 3,
	WIN: 6,
}


func getRounds(data string) [][2]string {
	lines := strings.Split(data, "\n")
	rounds := [][2]string{}

	for i := 0; i < len(lines); i++ {
		shapes := strings.Split(lines[i], " ")
		shapePairs := [2]string{shapes[0], shapes[1]}
		rounds = append(rounds, shapePairs)
	}

	return rounds
}

func getOutcome(opponentShape Shape, myShape Shape) RoundOutcome {
	if (opponentShape == myShape) {
		return DRAW
	}

	if (opponentShape == ROCK && myShape == PAPER) {
		return WIN
	}

	if (opponentShape == PAPER && myShape == SCISSORS) {
		return WIN
	}

	if (opponentShape == SCISSORS && myShape == ROCK) {
		return WIN
	}

	return LOSS
}

func getRoundScore(round [2]string) int {
	opponentShape := OPPONENT_SHAPES[round[0]]
	myShape := MY_SHAPES[round[1]]

	shapeScore := SHAPES_SCORE[myShape]

	outcome := getOutcome(opponentShape, myShape)
	outcomeScore := OUTCOME_SCORE[outcome]

	return shapeScore + outcomeScore
}

func getTotalScore(rounds [][2]string) int {
	score := 0

	for i := 0; i < len(rounds); i++ {
		score += getRoundScore(rounds[i])
	}

	return score
}

func main() {
	data := loadData()
	rounds := getRounds(data)
	totalScore := getTotalScore(rounds)

	fmt.Println(totalScore)
}
