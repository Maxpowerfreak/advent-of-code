package main

import (
	"bufio"
	"fmt"
	"unicode/utf8"

	helpers "github.com/maxpowerfreak/advent-of-code"
)

func main() {
	body, err := helpers.GetInputResponseBody(2022, 6)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	pageScanner := bufio.NewScanner(body)
	pageScanner.Split(bufio.ScanRunes)

	signalSlice := make([]rune, 4)
	// Initial 4 letters
	for i := 0; i < 4; i++ {
		pageScanner.Scan()
		char, _ := utf8.DecodeRune(pageScanner.Bytes())

		signalSlice[i] = char
	}

	if hasSignal(signalSlice) {
		fmt.Println(4)
		return
	}

	var numOfCharsScanned = 4
	for pageScanner.Scan() {
		char, _ := utf8.DecodeRune(pageScanner.Bytes())

		signalSlice = forward(signalSlice, char)
		numOfCharsScanned++

		if hasSignal(signalSlice) {
			fmt.Println(numOfCharsScanned)
			return
		}
	}
}

func forward(signalSlice []rune, char rune) []rune {
	return []rune{signalSlice[1], signalSlice[2], signalSlice[3], char}
}

func hasSignal(signalSlice []rune) bool {
	hashmap := make(map[rune]bool)

	for _, char := range signalSlice {
		if ok := hashmap[char]; ok {
			return false
		}

		hashmap[char] = true
	}

	return true
}
