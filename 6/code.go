package main

import (
	"bufio"
	"fmt"
	"os"
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

func isDistinctWindow(slice string) bool {
	windowSet := make(map[string]bool)
	for i := 0; i < len(slice); i++ {
		if _, ok := windowSet[string(slice[i])]; ok {
			return false
		}
		windowSet[string(slice[i])] = true
	}
	return true
}

func firstMarkerForSeq(input []string) []int {
	result := []int{}

	for _, inputStr := range input {
		for i := 4; i < len(inputStr); i++ {
			if isDistinctWindow(inputStr[i-4 : i]) {
				result = append(result, i)
				break
			}
		}
	}

	return result
}

func firstMarkerForMessage(input []string) []int {
	result := []int{}

	for _, inputStr := range input {
		for i := 14; i < len(inputStr); i++ {
			if isDistinctWindow(inputStr[i-14 : i]) {
				result = append(result, i)
				break
			}
		}
	}

	return result
}

func main() {
	test, _ := readFile("test.txt")
	input, _ := readFile("input.txt")
	fmt.Println("Part1")
	fmt.Println(firstMarkerForSeq(test))
	fmt.Println(firstMarkerForSeq(input))
	fmt.Println("Part1")
	fmt.Println(firstMarkerForMessage(test))
	fmt.Println(firstMarkerForMessage(input))
}
