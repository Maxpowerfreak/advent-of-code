package main

import (
	"bufio"
	"fmt"

	helpers "github.com/maxpowerfreak/advent-of-code"
)

func main() {
	body, err := helpers.GetInputResponseBody(2022, 3)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	pageScanner := bufio.NewScanner(body)
	pageScanner.Split(bufio.ScanLines)

	// build groups of rucksack
	var groups [][][]rune
	var group [][]rune
	for pageScanner.Scan() {
		if len(group) == 3 {
			groups = append(groups, group)
			group = nil
		}
		group = append(group, []rune(pageScanner.Text()))
	}
	groups = append(groups, group)

	var totalPriority int
	for _, group := range groups {
		intersection := intersect(group[0], group[1])

		// escape early
		if len(intersection) == 1 {
			fmt.Println(string(intersection[0]), calculatePriority(intersection[0]))
			totalPriority += calculatePriority(intersection[0])
			continue
		}

		intersection = intersect(intersection, group[2])
		fmt.Println(string(intersection[0]), calculatePriority(intersection[0]))
		totalPriority += calculatePriority(intersection[0])
	}

	fmt.Println(totalPriority)
}

// we know there is only a single intersect
func intersect(a []rune, b []rune) []rune {
	hash := make(map[rune]struct{})
	var intersection []rune

	for _, char := range a {
		hash[char] = struct{}{}
	}

	for _, char := range b {
		if _, ok := hash[char]; ok {
			intersection = append(intersection, char)
		}
	}

	return intersection
}

func calculatePriority(val rune) int {
	var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i, letter := range alphabet {
		if val == letter {
			return i + 1
		}
	}

	return 0
}
