package main

import (
	"bufio"
	"fmt"
	helpers "github.com/maxpowerfreak/advent-of-code"
	"strings"
)

const rock = "rock"
const paper = "paper"
const scissor = "scissor"
const lose = "lose"
const draw = "draw"
const win = "win"

func main() {
	body, err := helpers.GetInputResponseBody(2022, 2)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	pageScanner := bufio.NewScanner(body)
	pageScanner.Split(bufio.ScanLines)

	var totalScore int
	for pageScanner.Scan() {
		letters := strings.Split(pageScanner.Text(), " ")

		theirs, mine := letterToHand(letters[0]), chooseMyHand(letterToHand(letters[0]), letterToStatus(letters[1]))

		totalScore += handToScore(mine) + score(mine, theirs)
	}

	fmt.Println(totalScore)
}

func letterToHand(letter string) string {
	switch letter {
	case "A":
		return rock
	case "B":
		return paper
	case "C":
		return scissor
	default:
		return ""
	}
}

func letterToStatus(letter string) string {
	switch letter {
	case "Y":
		return draw
	case "Z":
		return win
	default:
		return lose
	}
}

func chooseMyHand(otherHand, endResult string) string {
	if endResult == draw {
		return otherHand
	}

	if otherHand == rock {
		if endResult == win {
			return paper
		}

		return scissor
	} else if otherHand == paper {
		if endResult == win {
			return scissor
		}

		return rock
	}

	// otherHand == scissor
	if endResult == win {
		return rock
	}

	return paper
}

func handToScore(hand string) int {
	switch hand {
	case rock:
		return 1
	case paper:
		return 2
	case scissor:
		return 3
	default:
		return 0
	}
}

func score(mine, theirs string) int {
	if mine == theirs {
		return 3
	}

	switch mine {
	case rock:
		if theirs == paper {
			return 0
		}
	case paper:
		if theirs == scissor {
			return 0
		}
	case scissor:
		if theirs == rock {
			return 0
		}
	default:
		return 0
	}

	return 6
}
