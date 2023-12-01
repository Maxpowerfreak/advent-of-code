package main

import (
	"bufio"
	"fmt"
	"regexp"
	"slices"
	"strconv"

	helpers "github.com/maxpowerfreak/advent-of-code"
)

var digitMap = map[string]string{
	"1": "1",
	"2": "2",
	"3": "3",
	"4": "4",
	"5": "5",
	"6": "6",
	"7": "7",
	"8": "8",
	"9": "9",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"eno":   "1",
	"owt":   "2",
	"eerht": "3",
	"ruof":  "4",
	"evif":  "5",
	"xis":   "6",
	"neves": "7",
	"thgie": "8",
	"enin":  "9",
}

func main() {
	body, err := helpers.GetInputResponseBody(2023, 1)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	pageScanner := bufio.NewScanner(body)
	pageScanner.Split(bufio.ScanLines)

	r, _ := regexp.Compile("([1-9]|one|two|three|four|five|six|seven|eight|nine)")
	rrev, _ := regexp.Compile("([1-9]|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)")

	var nums []int
	for pageScanner.Scan() {
		line := pageScanner.Text()

		fmt.Println(line)
		first := r.FindString(line)

		// Reverse line because regexp only reads from left to right and doesn't allow for overlaps
		lineSlice := []rune(line)
		slices.Reverse(lineSlice)
		revLine := string(lineSlice)

		second := rrev.FindString(revLine)

		numStr := digitMap[first] + digitMap[second]
		fmt.Println(numStr)
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
	}

	var total int
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}
