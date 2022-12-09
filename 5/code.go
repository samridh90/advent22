package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samridh90/advent22/shared"
)

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	if l == 0 {
		return nil, ""
	}
	return s[:l-1], s[l-1]
}

func (s stack) Top() string {
	l := len(s)
	return s[l-1]
}

func makeStacks(input []string) ([]stack, []string) {
	result := []stack{}
	stacksStart := 0
	for i, str := range input {
		if len(str) == 0 {
			stacksStart = i - 1
			break
		}
	}

	stacksLine := input[stacksStart]
	stackIndices := strings.Fields(stacksLine)
	offsets := []int{}
	for _, stackIndex := range stackIndices {
		offsets = append(offsets, strings.Index(stacksLine, stackIndex))
	}

	for _, offset := range offsets {
		newStack := stack{}
		for i := stacksStart - 1; i >= 0; i-- {
			line := input[i]
			if len(line) < offset {
				break
			}
			value := strings.TrimSpace(string(line[offset]))
			if len(value) > 0 {
				newStack = newStack.Push(value)
			} else {
				break
			}
		}
		result = append(result, newStack)
	}

	return result, input[stacksStart+2:]
}

func parseMove(moveStr string) (int, int, int) {
	fields := strings.Fields(moveStr)
	numValues, _ := strconv.ParseInt(fields[1], 0, 64)
	fromStack, _ := strconv.ParseInt(fields[3], 0, 64)
	toStack, _ := strconv.ParseInt(fields[5], 0, 64)
	return int(numValues), int(fromStack - 1), int(toStack - 1)
}

func doMove(stacks []stack, numValues, fromStack, toStack int) {
	for i := 0; i < numValues; i++ {
		value := ""
		stacks[fromStack], value = stacks[fromStack].Pop()
		stacks[toStack] = stacks[toStack].Push(value)
	}
}

func makeMoves(input []string, stacks []stack) string {
	result := []string{}
	for _, moveStr := range input {
		numValues, fromStack, toStack := parseMove(moveStr)
		doMove(stacks, numValues, fromStack, toStack)
	}
	for _, stack := range stacks {
		result = append(result, stack.Top())
	}
	return strings.Join(result, "")
}

func doMove1(stacks []stack, numValues, fromStack, toStack int) {
	tmp := []string{}
	for i := 0; i < numValues; i++ {
		value := ""
		stacks[fromStack], value = stacks[fromStack].Pop()
		tmp = append(tmp, value)
	}
	for i := len(tmp) - 1; i >= 0; i-- {
		value := tmp[i]
		stacks[toStack] = stacks[toStack].Push(value)
	}
}

func makeMoves1(input []string, stacks []stack) string {
	result := []string{}
	for _, moveStr := range input {
		numValues, fromStack, toStack := parseMove(moveStr)
		doMove1(stacks, numValues, fromStack, toStack)
	}
	for _, stack := range stacks {
		result = append(result, stack.Top())
	}
	return strings.Join(result, "")
}

func main() {
	test, _ := shared.ReadFile("./5/test.txt")
	input, _ := shared.ReadFile("./5/input.txt")
	testStacks, test := makeStacks(test)
	inputStacks, input := makeStacks(input)
	fmt.Println("Part1")
	fmt.Println(makeMoves(test, testStacks))
	fmt.Println(makeMoves(input, inputStacks))

	test, _ = shared.ReadFile("./5/test.txt")
	input, _ = shared.ReadFile("./5/input.txt")
	testStacks, test = makeStacks(test)
	inputStacks, input = makeStacks(input)
	fmt.Println("Part2")
	fmt.Println(makeMoves1(test, testStacks))
	fmt.Println(makeMoves1(input, inputStacks))
}
