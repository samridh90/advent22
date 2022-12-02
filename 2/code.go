package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func computeScore(input []string) int {
	individualScore := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	score := map[string]int{
		"AX": 3,
		"AY": 6,
		"AZ": 0,
		"BX": 0,
		"BY": 3,
		"BZ": 6,
		"CX": 6,
		"CY": 0,
		"CZ": 3,
	}
	finalScore := 0
	for _, matchStr := range input {
		parts := strings.Split(matchStr, " ")
		finalScore += (individualScore[parts[1]] + score[strings.Join(parts, "")])
	}
	return finalScore

}

func computeScore2(input []string) int {
	resultScore := map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}
	score := map[string]int{
		"AX": 3,
		"AY": 1,
		"AZ": 2,
		"BX": 1,
		"BY": 2,
		"BZ": 3,
		"CX": 2,
		"CY": 3,
		"CZ": 1,
	}
	finalScore := 0
	for _, matchStr := range input {
		parts := strings.Split(matchStr, " ")
		finalScore += (resultScore[parts[1]] + score[strings.Join(parts, "")])
	}
	return finalScore

}

func main() {
	test, _ := readFile("test.txt")
	input, _ := readFile("input.txt")
	fmt.Println("Part1")
	fmt.Println(computeScore(test))
	fmt.Println(computeScore(input))
	fmt.Println("Part2")
	fmt.Println(computeScore2(test))
	fmt.Println(computeScore2(input))
}
