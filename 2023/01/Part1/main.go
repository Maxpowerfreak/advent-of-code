package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"

	helpers "github.com/maxpowerfreak/advent-of-code"
)

func main() {
	body, err := helpers.GetInputResponseBody(2023, 1)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	pageScanner := bufio.NewScanner(body)
	pageScanner.Split(bufio.ScanLines)

	var nums []int
	for pageScanner.Scan() {
		line := pageScanner.Text()
		fmt.Println(line)

		lineRunes := []rune(line)
		lineLen := len(lineRunes)

		var firstDigit, lastDigit rune
		// forward travel
		for i := 0; i < lineLen; i++ {
			if isDigit(lineRunes[i]) {
				firstDigit = lineRunes[i]
				break
			}
		}

		// reverse travel
		for i := lineLen - 1; i >= 0; i-- {
			if isDigit(lineRunes[i]) {
				lastDigit = lineRunes[i]
				break
			}
		}

		// only a single digit in this string
		str := string(firstDigit) + string(lastDigit)

		fmt.Println(str)

		num, _ := strconv.Atoi(str)

		nums = append(nums, num)
	}

	var total int
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func isDigit(r rune) bool {
	digits := []rune{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}

	return slices.Contains(digits, r)
}
