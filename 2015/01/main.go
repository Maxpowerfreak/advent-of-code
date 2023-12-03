package main

import (
	"bufio"
	"fmt"

	helpers "github.com/maxpowerfreak/advent-of-code"
)

func main() {
	body, err := helpers.GetInputResponseBody(2015, 1)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	pageScanner := bufio.NewScanner(body)
	pageScanner.Split(bufio.ScanLines)

	var floor int
	for pageScanner.Scan() {
		line := pageScanner.Text()

		var enteredBasement bool
		for i, char := range line {
			if char == '(' {
				floor++
			} else {
				floor--
			}

			if !enteredBasement && floor < 0 {
				fmt.Println(i + 1)
				enteredBasement = true
			}
		}
	}

	fmt.Println(floor)
}
