package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/samridh90/advent22/shared"
)

type monkey struct {
	items       []int64
	operation   func(int64) int64
	test        func(int64) bool
	testDivider int64
	testResult  map[bool]int64
}

func parseItems(input string) []int64 {
	fields := strings.Fields(strings.TrimSpace(input))
	result := []int64{}
	for _, field := range fields {
		field = strings.Trim(field, ",")
		value, _ := strconv.ParseInt(field, 0, 64)
		result = append(result, value)
	}
	return result
}

func parseOperation(input string) func(int64) int64 {
	fields := strings.Fields(strings.TrimSpace(input))
	operator := fields[3]
	operand := fields[4]
	operandValue, err := strconv.ParseInt(operand, 0, 64)
	return func(old int64) int64 {
		if err == nil {
			if operator == "+" {
				return old + operandValue
			}
			if operator == "*" {
				return old * operandValue
			}
		} else {
			if operator == "+" {
				return old + old
			}
			if operator == "*" {
				return old * old
			}
		}
		return 0
	}
}

func parseTest(input string) (func(int64) bool, int64) {
	fields := strings.Fields(strings.TrimSpace(input))
	divisibleBy, _ := strconv.ParseInt(fields[2], 0, 64)
	return func(new int64) bool {
		return new%divisibleBy == 0
	}, divisibleBy
}

func parseTestResult(input string) int64 {
	fields := strings.Fields(strings.TrimSpace(input))
	passToMonkey, _ := strconv.ParseInt(fields[3], 0, 64)
	return passToMonkey
}

func parseMonkeys(input []string) []monkey {
	result := []monkey{}

	ts := func(ipStr string) string {
		return strings.TrimLeft(ipStr, " ")
	}

	for i := 0; i < len(input); i += 7 {
		curMonkey := monkey{}
		curMonkey.items = parseItems(strings.TrimPrefix(ts(input[i+1]), "Starting items:"))
		curMonkey.operation = parseOperation(strings.TrimPrefix(ts(input[i+2]), "Operation:"))
		curMonkey.test, curMonkey.testDivider = parseTest(strings.TrimPrefix(ts(input[i+3]), "Test:"))
		curMonkey.testResult = map[bool]int64{}
		curMonkey.testResult[true] = parseTestResult(strings.TrimPrefix(ts(input[i+4]), "If true:"))
		curMonkey.testResult[false] = parseTestResult(strings.TrimPrefix(ts(input[i+5]), "If false:"))
		result = append(result, curMonkey)
	}

	return result
}

func debugMonkeys(monkeys []monkey) {
	for i := 0; i < len(monkeys); i++ {
		fmt.Println("Monkey", i, monkeys[i].items)
	}
	fmt.Println()
}

func runSimulation(monkeys []monkey, numRounds int) int {
	numMonkeys := len(monkeys)
	numInspect := make([]int, numMonkeys)
	for i := 0; i < numRounds; i++ {
		for m := 0; m < numMonkeys; m++ {
			curMonkey := monkeys[m]
			for j := 0; j < len(curMonkey.items); j++ {
				numInspect[m]++
				curItem := curMonkey.items[j]
				newItem := curMonkey.operation(curItem)
				newItem /= 3
				testResult := curMonkey.test(newItem)
				nextMonkey := curMonkey.testResult[testResult]
				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, newItem)
			}
			monkeys[m].items = []int64{}
		}
		// debugMonkeys(monkeys)
	}

	sort.Slice(numInspect, func(i, j int) bool {
		return numInspect[i] < numInspect[j]
	})

	return numInspect[numMonkeys-1] * numInspect[numMonkeys-2]
}

func runSimulation2(monkeys []monkey, numRounds int) int64 {
	modFactor := int64(1)
	for _, monkey := range monkeys {
		modFactor *= monkey.testDivider
	}

	numMonkeys := len(monkeys)
	numInspect := make([]int64, numMonkeys)
	for i := 0; i < numRounds; i++ {
		for m := 0; m < numMonkeys; m++ {
			curMonkey := monkeys[m]
			for j := 0; j < len(curMonkey.items); j++ {
				numInspect[m]++
				curItem := curMonkey.items[j]
				newItem := curMonkey.operation(curItem)
				newItem %= modFactor
				testResult := curMonkey.test(newItem)
				nextMonkey := curMonkey.testResult[testResult]
				monkeys[nextMonkey].items = append(monkeys[nextMonkey].items, newItem)
			}
			monkeys[m].items = []int64{}
		}
	}

	sort.Slice(numInspect, func(i, j int) bool {
		return numInspect[i] < numInspect[j]
	})

	return numInspect[numMonkeys-1] * numInspect[numMonkeys-2]
}

func main() {
	files := []string{"./11/test.txt", "./11/input.txt"}
	for _, file := range files {
		test, _ := shared.ReadFile(file)
		monkeys := parseMonkeys(test)
		fmt.Println(file)
		fmt.Println(runSimulation(monkeys, 20))
		monkeys = parseMonkeys(test)
		fmt.Println(runSimulation2(monkeys, 10000))
	}
}
