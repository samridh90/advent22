package main

import (
	"fmt"
	"strconv"

	"github.com/samridh90/advent22/shared"
)

func parseMatrix(input []string) [][]int {
	result := [][]int{}
	for _, line := range input {
		tmp := []int{}
		for _, char := range line {
			num, _ := strconv.ParseInt(string(char), 0, 64)
			tmp = append(tmp, int(num))
		}
		result = append(result, tmp)
	}

	return result
}

func isTreeVisible(row, col int, matrix [][]int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	treeHeight := matrix[row][col]
	topVisible := true
	leftVisible := true
	rightVisible := true
	bottomVisible := true
	for i := 0; i < row; i++ {
		if matrix[i][col] >= treeHeight {
			topVisible = false
			break
		}
	}
	if topVisible {
		return true
	}
	for i := row + 1; i < rows; i++ {
		if matrix[i][col] >= treeHeight {
			bottomVisible = false
			break
		}
	}
	if bottomVisible {
		return true
	}
	for j := 0; j < col; j++ {
		if matrix[row][j] >= treeHeight {
			leftVisible = false
			break
		}
	}
	if leftVisible {
		return true
	}
	for j := col + 1; j < cols; j++ {
		if matrix[row][j] >= treeHeight {
			rightVisible = false
			break
		}
	}
	return rightVisible
}

func findVisibleTrees(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	count := 2 * (rows + cols - 2)

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if isTreeVisible(i, j, matrix) {
				count++
			}
		}
	}
	return count
}

func viewingScore(row, col int, matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	topScore := 0
	bottomScore := 0
	leftScore := 0
	rightScore := 0
	treeHeight := matrix[row][col]
	for i := row - 1; i >= 0; i-- {
		topScore++
		if matrix[i][col] >= treeHeight {
			break
		}
	}
	for i := row + 1; i < rows; i++ {
		bottomScore++
		if matrix[i][col] >= treeHeight {
			break
		}
	}
	for j := col - 1; j >= 0; j-- {
		leftScore++
		if matrix[row][j] >= treeHeight {
			break
		}
	}
	for j := col + 1; j < cols; j++ {
		rightScore++
		if matrix[row][j] >= treeHeight {
			break
		}
	}
	return topScore * bottomScore * leftScore * rightScore
}

func findMaxViewingScore(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])
	maxScore := 0

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			score := viewingScore(i, j, matrix)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

func main() {
	test, _ := shared.ReadFile("./8/test.txt")
	testMatrix := parseMatrix(test)
	input, _ := shared.ReadFile("./8/input.txt")
	inputMatrix := parseMatrix(input)
	fmt.Println("Part1")
	fmt.Println(findVisibleTrees(testMatrix))
	fmt.Println(findVisibleTrees(inputMatrix))
	fmt.Println("Part2")
	fmt.Println(findMaxViewingScore(testMatrix))
	fmt.Println(findMaxViewingScore(inputMatrix))
}
