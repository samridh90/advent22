package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/samridh90/advent22/shared"
)

func computeMaxCalories(input []string, top int) int64 {
	allMax := []int64{}
	sumSoFar := int64(0)

	for _, calStr := range input {
		if calStr == "" {
			allMax = append(allMax, sumSoFar)
			sumSoFar = 0
		} else {
			cal, err := strconv.ParseInt(calStr, 10, 64)
			if err != nil {
				break
			}
			sumSoFar += cal
		}

	}
	sort.Slice(allMax, func(i, j int) bool {
		return allMax[i] > allMax[j]
	})
	sum := int64(0)

	for _, v := range allMax[:top] {
		sum += v
	}
	return sum
}

func main() {
	test, _ := shared.ReadFile("./1/test.txt")
	test = append(test, "")
	input, _ := shared.ReadFile("./1/input.txt")
	input = append(input, "")
	fmt.Println("Part1")
	fmt.Println(computeMaxCalories(test, 1))
	fmt.Println(computeMaxCalories(input, 1))
	fmt.Println("Part2")
	fmt.Println(computeMaxCalories(test, 3))
	fmt.Println(computeMaxCalories(input, 3))
}
