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

var ROUND_OUTCOMES = map[string]RoundOutcome{
	"X": LOSS,
	"Y": DRAW,
	"Z": WIN,
}

var OUTCOME_SCORE = map[RoundOutcome]int{
	LOSS: 0,
	DRAW: 3,
	WIN: 6,
}

var SHAPE_TO_WIN = map[Shape]Shape{
	ROCK: PAPER,
	PAPER: SCISSORS,
	SCISSORS: ROCK,
}

var SHAPE_TO_LOSE = map[Shape]Shape{
	ROCK: SCISSORS,
	PAPER: ROCK,
	SCISSORS: PAPER,
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

func getMyShape(opponentShape Shape, outcome RoundOutcome) Shape {
	if (outcome == WIN) {
		return SHAPE_TO_WIN[opponentShape]
	}

	if (outcome == LOSS) {
		return SHAPE_TO_LOSE[opponentShape]
	}

	return opponentShape
}

func getRoundScore(round [2]string) int {
	opponentShape := OPPONENT_SHAPES[round[0]]
	outcome := ROUND_OUTCOMES[round[1]]
	outcomeScore := OUTCOME_SCORE[outcome]

	myShape := getMyShape(opponentShape, outcome)
	shapeScore := SHAPES_SCORE[myShape]

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
