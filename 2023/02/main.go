package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	helpers "github.com/maxpowerfreak/advent-of-code"
)

const (
	maxRed   = 12
	maxGreen = 13
	maxBlue  = 14
	redStr   = "red"
	greenStr = "green"
	blueStr  = "blue"
)

type Color struct {
	Count int
}

// Game represents a game and it's highest number of seen cubes of a given color
type Game struct {
	Id       int
	Red      Color
	Green    Color
	Blue     Color
	MinRed   Color
	MinGreen Color
	MinBlue  Color
}

func (g *Game) Compare(red, green, blue Color) {
	g.Red.Count = max(g.Red.Count, red.Count)
	g.Green.Count = max(g.Green.Count, green.Count)
	g.Blue.Count = max(g.Blue.Count, blue.Count)
}

func (g *Game) IsPossible() bool {
	return g.Red.Count <= maxRed &&
		g.Green.Count <= maxGreen &&
		g.Blue.Count <= maxBlue
}

func newGame(id int) Game {
	return Game{
		Id: id,
	}
}

func main() {
	body, err := helpers.GetInputResponseBody(2023, 2)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	pageScanner := bufio.NewScanner(body)
	pageScanner.Split(bufio.ScanLines)

	var maxTotal, minTotal int
	for pageScanner.Scan() {
		line := pageScanner.Text()
		fmt.Println(line)

		splitLine := strings.Split(line, ":")
		gameStr, showingsStr := splitLine[0], splitLine[1]

		gameIdStr := strings.TrimLeft(gameStr, "Game ")
		gameId, _ := strconv.Atoi(gameIdStr)

		curGame := newGame(gameId)

		showingsSlice := strings.Split(showingsStr, ";")
		for _, showing := range showingsSlice {
			colorMap := showingStrToMap(showing)
			curGame.Compare(colorMap[redStr], colorMap[greenStr], colorMap[blueStr])
		}

		fmt.Println(curGame)
		if curGame.IsPossible() {
			maxTotal += curGame.Id
		}

		minTotal += (curGame.Red.Count * curGame.Green.Count * curGame.Blue.Count)
	}

	fmt.Printf("\nPart 1: %d\nPart 2: %d", maxTotal, minTotal)
}

func showingStrToMap(str string) map[string]Color {
	strSlice := strings.Split(str, ",")

	colorMap := make(map[string]Color, 0)
	for _, hand := range strSlice {
		handSlice := strings.Split(strings.TrimLeft(hand, " "), " ")

		countStr, color := handSlice[0], handSlice[1]
		count, _ := strconv.Atoi(countStr)

		colorMap[color] = Color{
			Count: count,
		}
	}

	return colorMap
}
