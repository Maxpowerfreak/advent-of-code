package main

import (
	"bufio"
	"fmt"
	"strconv"

	helpers "github.com/maxpowerfreak/advent-of-code"
)

// Solution filled with copy/paste which could've been turned to functions
// But I got lazy and it was late
func main() {
	body, err := helpers.GetInputResponseBody(2023, 3)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	pageScanner := bufio.NewScanner(body)
	pageScanner.Split(bufio.ScanLines)

	var curLine, nextLine []rune
	moreLines := true

	// Retrieve first and second line as current and next
	pageScanner.Scan()
	curLine = []rune(pageScanner.Text())
	pageScanner.Scan()
	nextLine = []rune(pageScanner.Text())

	var total int

	// First check happens outside the loop as we don't have a complete set of 3 scans
	// That way I can avoid having an if statement for the first and last scan or over engineering even further
	fmt.Println(string(curLine))
	var firstHighIndex int
	for i, char := range curLine {
		if i < firstHighIndex {
			// Skip as we are still within an already scanned number
			continue
		}

		// Look for digit
		if isDigit(char) {
			var hasSymbol bool
			// Digit detected, begin gathering idx range and filling number
			// Find end of number
			firstHighIndex = getLastIndexForNumber(i, curLine)

			// Get search range for symbol (accounting for array bounds)
			minSearch, maxSearch := max(0, i-1), min(len(curLine), firstHighIndex+1)

			// Check for symbol on current line
			if !hasSymbol && sublineHasSymbol(curLine[minSearch:maxSearch]) {
				hasSymbol = true
			}

			// Check for symbol on last line
			if !hasSymbol && sublineHasSymbol(nextLine[minSearch:maxSearch]) {
				hasSymbol = true
			}

			// if symbol is found, add to total
			if hasSymbol {
				numSlice := curLine[i:firstHighIndex]
				num, err := strconv.Atoi(string(numSlice))

				if err != nil {
					panic(err)
				}

				total += num
			} else {
			}

		}
	}

	// rotate
	prevLine := make([]rune, len(curLine))
	copy(prevLine, curLine)
	copy(curLine, nextLine)
	pageScanner.Scan()

	// Normal lines
	for moreLines {
		nextLine = []rune(pageScanner.Text())
		fmt.Println(string(curLine))
		var highIdx int

		for i, char := range curLine {
			if i < highIdx {
				// Skip as we are still within an already scanned number
				continue
			}

			// Look for digit
			if isDigit(char) {
				var hasSymbol bool
				// Digit detected, begin gathering idx range and filling number
				// Find end of number
				highIdx = getLastIndexForNumber(i, curLine)

				// Get search range for symbol (accounting for array bounds)
				minSearch, maxSearch := max(0, i-1), min(len(curLine), highIdx+1)

				// Check symbol on previous line
				if sublineHasSymbol(prevLine[minSearch:maxSearch]) {
					hasSymbol = true
				}

				// Check for symbol on current line
				if !hasSymbol && sublineHasSymbol(curLine[minSearch:maxSearch]) {
					hasSymbol = true
				}

				// Check for symbol on last line
				if !hasSymbol && sublineHasSymbol(nextLine[minSearch:maxSearch]) {
					hasSymbol = true
				}

				// if symbol is found, add to total
				if hasSymbol {
					numSlice := curLine[i:highIdx]
					num, err := strconv.Atoi(string(numSlice))

					if err != nil {
						panic(err)
					}

					total += num
				}
			}

		}

		copy(prevLine, curLine)
		copy(curLine, nextLine)
		moreLines = pageScanner.Scan()
	}

	// Check the last case, which is only a prevLine and curLine case
	fmt.Println(string(curLine))
	var lastHighIndex int
	for i, char := range curLine {
		if i < lastHighIndex {
			// Skip as we are still within an already scanned number
			continue
		}

		// Look for digit
		if isDigit(char) {
			var hasSymbol bool
			// Digit detected, begin gathering idx range and filling number
			// Find end of number
			lastHighIndex = getLastIndexForNumber(i, curLine)

			// Get search range for symbol (accounting for array bounds)
			minSearch, maxSearch := max(0, i-1), min(len(curLine), lastHighIndex+1)

			// Check for symbol on previous line
			if !hasSymbol && sublineHasSymbol(prevLine[minSearch:maxSearch]) {
				hasSymbol = true
			}

			// Check for symbol on current line
			if !hasSymbol && sublineHasSymbol(curLine[minSearch:maxSearch]) {
				hasSymbol = true
			}

			// if symbol is found, add to total
			if hasSymbol {
				numSlice := curLine[i:lastHighIndex]
				num, err := strconv.Atoi(string(numSlice))

				if err != nil {
					panic(err)
				}

				total += num
			}

		}
	}

	fmt.Println(total)
	fmt.Printf("Real answer : %d\n", 536576)
}

func getLastIndexForNumber(curIndex int, curLine []rune) int {
	for ; curIndex < len(curLine); curIndex++ {
		if !isDigit(curLine[curIndex]) {
			break
		}
	}

	return curIndex
}

func isDigit(char rune) bool {
	// '0' = 48 and '9' = 57
	return char >= 48 && char <= 57
}

func sublineHasSymbol(subline []rune) bool {
	for _, char := range subline {
		if !isDigit(char) && char != '.' {
			return true
		}
	}

	return false
}
