package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	helpers "github.com/maxpowerfreak/advent-of-code"
)

type pair struct {
	low  int
	high int
}

func main() {
	body, err := helpers.GetInputResponseBody(2022, 4)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	pageScanner := bufio.NewScanner(body)
	pageScanner.Split(bufio.ScanLines)

	var innieCount int
	for pageScanner.Scan() {
		line := pageScanner.Text()

		split := strings.Split(line, ",")
		pair1Str, pair2Str := strings.Split(split[0], "-"), strings.Split(split[1], "-")
		pair1 := pair{
			low:  atoiNoErr(pair1Str[0]),
			high: atoiNoErr(pair1Str[1]),
		}
		pair2 := pair{
			low:  atoiNoErr(pair2Str[0]),
			high: atoiNoErr(pair2Str[1]),
		}

		if overlaps(pair2, pair1) {
			innieCount++
		}
	}

	fmt.Println(innieCount)
}

func overlaps(pair1, pair2 pair) bool {
	return math.Max(float64(pair1.low), float64(pair2.low)) <= math.Min(float64(pair1.high), float64(pair2.high))
}

func atoiNoErr(s string) int {
	i, _ := strconv.Atoi(s)

	return i
}
