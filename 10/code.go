package main

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/samridh90/advent22/shared"
)

func parseInstruction(input string) (int, int) {
	if input == "noop" {
		return 1, 0
	}
	fields := strings.Fields(input)
	value, _ := strconv.ParseInt(fields[1], 0, 64)
	return 2, int(value)
}

func computeSignalStrengths(input []string) int {
	registerX := 1
	signalStrengths := []int{}
	cyclesOfInterest := []int{20, 60, 100, 140, 180, 220}
	cycle := 0

	for _, instruction := range input {
		numCycles, registerValue := parseInstruction(instruction)
		for i := 0; i < numCycles; i++ {
			cycle++
			if slices.Contains(cyclesOfInterest, cycle) {
				signalStrengths = append(signalStrengths, cycle*registerX)
			}
		}
		registerX += registerValue
	}

	sumOfSignalStrengths := 0
	for _, value := range signalStrengths {
		sumOfSignalStrengths += value
	}

	return sumOfSignalStrengths
}

func drawPixels(input []string) {
	registerX := 1
	cycle := 0
	crt := [][]string{
		make([]string, 40),
		make([]string, 40),
		make([]string, 40),
		make([]string, 40),
		make([]string, 40),
		make([]string, 40),
	}

	for _, instruction := range input {
		numCycles, registerValue := parseInstruction(instruction)
		spritePostion := []int{registerX - 1, registerX, registerX + 1}
		for i := 0; i < numCycles; i++ {
			row := cycle / 40
			col := cycle % 40

			if slices.Contains(spritePostion, col) {
				crt[row][col] = "#"
			} else {
				crt[row][col] = "."
			}

			cycle++
		}
		registerX += registerValue
	}

	for i := 0; i < len(crt); i++ {
		for j := 0; j < len(crt[0]); j++ {
			fmt.Print(crt[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	test, _ := shared.ReadFile("./10/test.txt")
	input, _ := shared.ReadFile("./10/input.txt")
	fmt.Println("Part1")
	fmt.Println(computeSignalStrengths(test))
	fmt.Println(computeSignalStrengths(input))
	fmt.Println("Part2")
	drawPixels(test)
	drawPixels(input)
}
