package main

import (
	"bufio"
	"fmt"
	helpers "github.com/maxpowerfreak/advent-of-code"
	"strconv"
	"strings"
)

func main() {
	body, err := helpers.GetInputResponseBody(2022, 5)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	pageScanner := bufio.NewScanner(body)
	pageScanner.Split(bufio.ScanLines)

	// No way do I want to make a parser for those stacks
	for i := 0; i < 10; i++ {
		// skip the stack setup
		pageScanner.Scan()
	}

	// write it all by hand
	stacks := buildStacks()

	for pageScanner.Scan() {
		line := pageScanner.Text()
		quantity, fromStackIndex, toStackIndex := parseLine(line)

		// let's try without checking to see if fromStack has enough for now
		var pop []rune
		pop = stacks[fromStackIndex][startIndexHandler(stacks[fromStackIndex], quantity):]
		stacks[fromStackIndex] = stacks[fromStackIndex][:startIndexHandler(stacks[fromStackIndex], quantity)]
		stacks[toStackIndex] = append(stacks[toStackIndex], pop...)
	}

	for _, stack := range stacks {
		fmt.Print(string(stack[len(stack)-1]))
	}
}

func buildStacks() [][]rune {
	var stacks [][]rune

	stacks = append(stacks, []rune("JHGMZNTF"))
	stacks = append(stacks, []rune("VWJ"))
	stacks = append(stacks, []rune("GVLJBTH"))
	stacks = append(stacks, []rune("BPJNCDVL"))
	stacks = append(stacks, []rune("FWSMPRG"))
	stacks = append(stacks, []rune("GHCFBNVM"))
	stacks = append(stacks, []rune("DHGMR"))
	stacks = append(stacks, []rune("HNMVZD"))
	stacks = append(stacks, []rune("GNFH"))

	return stacks
}

func parseLine(line string) (int, int, int) {
	lineSplit := strings.Split(line, " ")

	// 0 : move
	// 1 : 6
	// 2 : from
	// 3 : 4
	// 4 : to
	// 5 : 3
	return atoi(lineSplit[1]), atoi(lineSplit[3]) - 1, atoi(lineSplit[5]) - 1
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)

	return i
}

func startIndexHandler(slice []rune, quantity int) int {
	sliceSize := len(slice)
	if sliceSize <= quantity {
		return 0
	}
	return sliceSize - quantity
}
