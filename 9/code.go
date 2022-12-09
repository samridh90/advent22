package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/samridh90/advent22/shared"
)

type Pos struct {
	x, y int
}

func parseMoveStr(moveStr string) (string, int) {
	fields := strings.Fields(moveStr)
	numMoves, _ := strconv.ParseInt(fields[1], 0, 64)
	return fields[0], int(numMoves)
}

func isHeadTooFar(head, tail Pos) bool {
	xDiff := int(math.Abs(float64(head.x - tail.x)))
	yDiff := int(math.Abs(float64(head.y - tail.y)))
	return xDiff > 1 || yDiff > 1
}

func moveInDirection(cur Pos, direction string) Pos {
	if direction == "R" {
		return Pos{cur.x + 1, cur.y}
	} else if direction == "U" {
		return Pos{cur.x, cur.y + 1}
	} else if direction == "L" {
		return Pos{cur.x - 1, cur.y}
	} else if direction == "D" {
		return Pos{cur.x, cur.y - 1}
	}
	return cur
}

func getNewTailPosition(head, tail Pos) Pos {
	xDiff := head.x - tail.x
	yDiff := head.y - tail.y
	absXDiff := int(math.Abs(float64(xDiff)))
	absYDiff := int(math.Abs(float64(yDiff)))

	if absXDiff == 2 && yDiff == 0 {
		return Pos{tail.x + (xDiff / absXDiff), tail.y}
	}
	if absYDiff == 2 && xDiff == 0 {
		return Pos{tail.x, tail.y + (yDiff / absYDiff)}
	}
	if (absXDiff == 1 && absYDiff == 2) ||
		(absXDiff == 2 && absYDiff == 1) ||
		(absXDiff == 2 && absYDiff == 2) {
		return Pos{tail.x + (xDiff / absXDiff), tail.y + (yDiff / absYDiff)}
	}
	return tail
}

func getTotalTailPositions(input []string, ropeLength int) int {
	visited := map[Pos]bool{}
	rope := make([]Pos, ropeLength)
	visited[rope[ropeLength-1]] = true

	for _, moveStr := range input {
		direction, numMoves := parseMoveStr(moveStr)
		for i := 0; i < numMoves; i++ {
			rope[0] = moveInDirection(rope[0], direction)
			for j := 1; j < ropeLength; j++ {
				if isHeadTooFar(rope[j-1], rope[j]) {
					rope[j] = getNewTailPosition(rope[j-1], rope[j])
				}
			}
			visited[rope[ropeLength-1]] = true
		}
	}
	return len(visited)
}

func main() {
	test, _ := shared.ReadFile("./9/test.txt")
	test1, _ := shared.ReadFile(("./9/test1.txt"))
	input, _ := shared.ReadFile("./9/input.txt")
	fmt.Println("Part1")
	fmt.Println(getTotalTailPositions(test, 2))
	fmt.Println(getTotalTailPositions(test1, 2))
	fmt.Println(getTotalTailPositions(input, 2))
	fmt.Println("Part2")
	fmt.Println(getTotalTailPositions(test, 10))
	fmt.Println(getTotalTailPositions(test1, 10))
	fmt.Println(getTotalTailPositions(input, 10))
}
