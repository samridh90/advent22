package main

import (
	"fmt"
	"sort"

	"github.com/samridh90/advent22/shared"
)

type point struct {
	x int
	y int
}

type pointWithSteps struct {
	pt    point
	steps int
}

type adjacencyList map[point][]point

func buildAdjList(input []string) (adjacencyList, point, point, []point) {
	adjList := adjacencyList{}
	allStarts := []point{}
	rows := len(input)
	cols := len(input[0])
	var start point
	var end point

	graph := [][]rune{}
	for row, line := range input {
		chars := []rune{}
		for col, char := range line {
			chars = append(chars, char)
			charStr := string(char)
			if charStr == "a" || charStr == "S" {
				allStarts = append(allStarts, point{row, col})
			}
		}
		graph = append(graph, chars)
	}

	for row, line := range input {
		for col, char := range line {
			charStr := string(char)
			curPoint := point{row, col}
			adjList[curPoint] = []point{}
			if charStr == "S" {
				start = curPoint
				char = 97
			} else if charStr == "E" {
				end = curPoint
				char = 122
			}
			neighbors := []point{{row + 1, col}, {row - 1, col}, {row, col + 1}, {row, col - 1}}
			for _, neighbor := range neighbors {
				if neighbor.x < 0 || neighbor.x >= rows || neighbor.y < 0 || neighbor.y >= cols {
					continue
				}

				neighborChar := graph[neighbor.x][neighbor.y]
				neighborStr := string(neighborChar)
				if neighborStr == "S" {
					neighborChar = 97
				} else if neighborStr == "E" {
					neighborChar = 122
				}
				if neighborChar-char <= 1 {
					adjList[curPoint] = append(adjList[curPoint], neighbor)
				}
			}
		}
	}

	return adjList, start, end, allStarts
}

func shortestPath(graph adjacencyList, start point, end point) int {
	visited := map[point]int{start: 0}
	queue := []pointWithSteps{{start, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentPoint := current.pt
		nextSteps := current.steps + 1
		for _, neighbor := range graph[currentPoint] {
			neighborSteps, ok := visited[neighbor]
			if !ok || nextSteps < neighborSteps {
				visited[neighbor] = nextSteps
				if neighbor != end {
					queue = append(queue, pointWithSteps{pt: neighbor, steps: nextSteps})
				}
			}
		}
	}
	_, ok := visited[end]
	if ok {
		return visited[end]
	}
	return int(^uint(0) >> 1)
}

func main() {
	files := []string{"./12/test.txt", "./12/input.txt"}
	for _, file := range files {
		fmt.Println(file)
		input, _ := shared.ReadFile(file)
		graph, start, end, allStarts := buildAdjList(input)
		fmt.Println(shortestPath(graph, start, end))
		allStartShortestPaths := []int{}
		for _, aStart := range allStarts {
			allStartShortestPaths = append(allStartShortestPaths, shortestPath(graph, aStart, end))
		}
		sort.Slice(allStartShortestPaths, func(i, j int) bool {
			return allStartShortestPaths[i] < allStartShortestPaths[j]
		})
		fmt.Println(allStartShortestPaths[0])
	}
}
