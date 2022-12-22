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

	signalSlice := make([]rune, 14)
	// Initial 4 letters
	for i := 0; i < 14; i++ {
		pageScanner.Scan()
		char, _ := utf8.DecodeRune(pageScanner.Bytes())

		signalSlice[i] = char
	}

	if hasSignal(signalSlice) {
		fmt.Println(14)
		return
	}

	var numOfCharsScanned = 14
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
	newSlice := make([]rune, 14)

	for i := 0; i < 13; i++ {
		newSlice[i] = signalSlice[i+1]
	}

	newSlice[13] = char

	return newSlice
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
