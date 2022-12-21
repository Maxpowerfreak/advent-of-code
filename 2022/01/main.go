package main

import (
	"bufio"
	"fmt"
	"strconv"

	helpers "github.com/maxpowerfreak/advent-of-code"
)

func main() {
	body, err := helpers.GetInputResponseBody(2022, 1)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	pageScanner := bufio.NewScanner(body)
	pageScanner.Split(bufio.ScanLines)

	var currentCount, highestCount, num int
	for pageScanner.Scan() {
		line := pageScanner.Text()
		if len(line) == 0 {
			if currentCount > highestCount {
				highestCount = currentCount
			}

			currentCount, num = 0, 0
		} else {
			num, err = strconv.Atoi(line)
			if err != nil {
				panic(fmt.Errorf("failed to convert [%s] to a number", line))
			}
		}

		currentCount += num
	}

	if currentCount > highestCount {
		highestCount = currentCount
	}

	fmt.Println(highestCount)
}
