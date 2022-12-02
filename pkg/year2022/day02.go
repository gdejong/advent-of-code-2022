package year2022

import (
	"strings"
)

type Day02 struct{}

func (p Day02) PartA(lines []string) any {
	totalScore := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) != 2 {
			panic("expected 2 fields")
		}

		opponentShape := letterToShape(fields[0])
		myShape := letterToShape(fields[1])

		totalScore += playRoundPart1(opponentShape, myShape)
	}

	return totalScore
}

func (p Day02) PartB(lines []string) any {
	totalScore := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) != 2 {
			panic("expected 2 fields")
		}

		opponentShape := letterToShape(fields[0])
		actionToPlay := fields[1]

		totalScore += playRoundPart2(opponentShape, actionToPlay)
	}

	return totalScore
}

func playRoundPart2(opponentShape shape, actionToPlay string) int {
	var myShape shape

	switch actionToPlay {
	case needToEndInADraw:
		myShape = opponentShape
	case needToLose:
		myShape = shapeToLose(opponentShape)
	case needToWin:
		myShape = shapeToWin(opponentShape)
	default:
		panic("impossible")
	}

	return playRoundPart1(opponentShape, myShape)
}

func shapeToLose(s shape) shape {
	if s == rock {
		return scissors
	}
	if s == paper {
		return rock
	}
	if s == scissors {
		return paper
	}

	panic("impossible")
}

func shapeToWin(s shape) shape {
	if s == rock {
		return paper
	}
	if s == paper {
		return scissors
	}
	if s == scissors {
		return rock
	}

	panic("impossible")
}

func playRoundPart1(opponentShape, myShape shape) int {
	// The score for a single round is the score for the shape you selected...
	score := myShape.score()

	// ... plus the score for the outcome of the round
	score += scoreForRoundPart1(opponentShape, myShape)

	return score
}

func scoreForRoundPart1(opponentShape, myShape shape) int {
	// draw?
	if opponentShape == myShape {
		return 3
	}

	// win?
	if (myShape == rock && opponentShape == scissors) || (myShape == scissors && opponentShape == paper) || (myShape == paper && opponentShape == rock) {
		return 6
	}

	// loose
	return 0
}

const opponentRock = "A"
const myRock = "X"

const opponentPaper = "B"
const myPaper = "Y"

const opponentScissors = "C"
const myScissors = "Z"

const needToLose = "X"
const needToEndInADraw = "Y"
const needToWin = "Z"

const (
	rock     = shape(iota)
	paper    = shape(iota)
	scissors = shape(iota)
)

type shape int

func (s shape) score() int {
	switch s {
	case rock:
		return 1
	case paper:
		return 2
	case scissors:
		return 3
	}

	panic("Unknown shape")
}

func letterToShape(input string) shape {
	if input == opponentRock || input == myRock {
		return rock
	}

	if input == opponentPaper || input == myPaper {
		return paper
	}

	if input == opponentScissors || input == myScissors {
		return scissors
	}

	panic("Unknown shape")
}
