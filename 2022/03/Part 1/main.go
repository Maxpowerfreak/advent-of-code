package main

import (
	"bufio"
	"fmt"

	helpers "github.com/maxpowerfreak/advent-of-code"
)

func main() {
	var alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	body, err := helpers.GetInputResponseBody(2022, 3)
	if err != nil {
		panic(err)
	}
	defer body.Close()

	pageScanner := bufio.NewScanner(body)
	pageScanner.Split(bufio.ScanLines)

	var totalPriority int
	for pageScanner.Scan() {
		rucksack := []rune(pageScanner.Text())
		size := len(rucksack)
		compartment1, compartment2 := rucksack[:size/2], rucksack[size/2:]

		comp1Unique, comp2Unique := unique(compartment1), unique(compartment2)

		val := intersect(comp1Unique, comp2Unique)
		for i, letter := range alphabet {
			if val == letter {
				totalPriority += i + 1
			}
		}
	}

	fmt.Println(totalPriority)
}

func unique(comp []rune) []rune {
	keys := make(map[rune]bool)

	var list []rune
	for _, char := range comp {
		if _, hit := keys[char]; !hit {
			keys[char] = true
			list = append(list, char)
		}
	}

	return list
}

// we know there is only a single intersect
func intersect(a []rune, b []rune) rune {
	hash := make(map[rune]struct{})

	for _, char := range a {
		hash[char] = struct{}{}
	}

	for _, char := range b {
		if _, ok := hash[char]; ok {
			return char
		}
	}

	return ' '
}
