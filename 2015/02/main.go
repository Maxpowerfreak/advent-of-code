package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"

	helpers "github.com/maxpowerfreak/advent-of-code"
)

func main() {
	body, err := helpers.GetInputResponseBody(2015, 2)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	pageScanner := bufio.NewScanner(body)
	pageScanner.Split(bufio.ScanLines)

	var total, totalRibbon int
	for pageScanner.Scan() {
		line := pageScanner.Text()

		// l * w * h
		fmt.Println(line)

		l, w, h := inputToMeasurements(line)

		// area calc
		fmt.Printf("2*l*w=%d 2*w*h=%d 2*h*l=%d min=%d\n", 2*l*w, 2*w*h, 2*h*l, min(2*l*w, 2*w*h, 2*h*l))
		total += (2 * l * w) + (2 * w * h) + (2 * h * l) + (min(l*w, w*h, h*l))

		// ribbon calc
		ribSlice := []int{l, w, h}
		slices.Sort(ribSlice)

		totalRibbon += 2*ribSlice[0] + 2*ribSlice[1] + l*w*h
		fmt.Printf("Ribbon: %d\n\n", 2*ribSlice[0]+2*ribSlice[1]+l*w*h)
	}

	fmt.Println(total)
	fmt.Println(totalRibbon)
}

func inputToMeasurements(line string) (l, w, h int) {
	strSlice := strings.Split(line, "x")
	l, _ = strconv.Atoi(strSlice[0])
	w, _ = strconv.Atoi(strSlice[1])
	h, _ = strconv.Atoi(strSlice[2])

	return
}
