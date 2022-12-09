package main

import (
	"fmt"
	"unicode"

	"github.com/samridh90/advent22/shared"
)

func computePriority(item rune) int {
	if unicode.IsUpper(item) {
		return (26 + (int(item) - 64))
	}
	return (int(item) - 96)
}

func computeRucksackPriorities(input []string) int {
	totalPriority := 0

	for _, allItems := range input {
		itemSet := make(map[rune]bool)

		length := len(allItems)
		compartmentLength := length / 2
		first := allItems[:compartmentLength]
		second := allItems[compartmentLength:]
		for _, item := range first {
			itemSet[item] = true
		}
		for _, item := range second {
			if _, ok := itemSet[item]; ok {
				totalPriority += computePriority(item)
				break
			}
		}
	}

	return totalPriority
}

func computeRucksackPriorities2(input []string) int {
	totalPriority := 0

	for i := 0; i < len(input); i += 3 {
		first := input[i]
		second := input[i+1]
		third := input[i+2]

		itemSet := make(map[rune]bool)
		for _, item := range first {
			itemSet[item] = true
		}

		candidateSet := make(map[rune]bool)
		for _, item := range second {
			if _, ok := itemSet[item]; ok {
				candidateSet[item] = true
			}
		}

		for _, item := range third {
			if _, ok := candidateSet[item]; ok {
				totalPriority += computePriority(item)
				break
			}
		}
	}
	return totalPriority
}

func main() {
	test, _ := shared.ReadFile("./3/test.txt")
	input, _ := shared.ReadFile("./3/input.txt")
	fmt.Println("Part1")
	fmt.Println(computeRucksackPriorities(test))
	fmt.Println(computeRucksackPriorities(input))
	fmt.Println("Part2")
	fmt.Println(computeRucksackPriorities2(test))
	fmt.Println(computeRucksackPriorities2(input))
}
