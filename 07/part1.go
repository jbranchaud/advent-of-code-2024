package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var lineMatcher = regexp.MustCompile(`^(\d+): (.*)\z$`)

func part1(debug bool) {
	// puzzleToSolve := "sample"
	puzzleToSolve := "real"

	input, err := getInput("part1", puzzleToSolve)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	var testValueSum int64
	testValueSum = 0

	for _, line := range input {
		matches := lineMatcher.FindStringSubmatch(line)
		testValue, err := strconv.ParseInt(matches[1], 10, 64)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fields := strings.Fields(matches[2])
		operands := []int64{}
		for _, field := range fields {
			operand, err := strconv.ParseInt(field, 10, 64)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			operands = append(operands, operand)
		}

		operatorsNeeded := len(operands) - 1
		operatorOrderings := allPermutations([]int{}, operatorsNeeded, []int{0, 1})

		for _, ordering := range operatorOrderings {
			runningTotal := operands[0]
			debugStr := strings.Builder{}
			debugStr.WriteString(fmt.Sprintf("%d", operands[0]))
			for j, operatorIndex := range ordering {
				switch operatorIndex {
				case 0:
					debugStr.WriteString(" + ")
					runningTotal = add(runningTotal, operands[j+1])
				case 1:
					debugStr.WriteString(" * ")
					runningTotal = multiply(runningTotal, operands[j+1])
				}
				debugStr.WriteString(fmt.Sprintf("%d", operands[j+1]))
			}

			if testValue == runningTotal {
				if debug {
					fmt.Println(">", debugStr.String(), "=>", testValue)
				}
				testValueSum += testValue
				if debug {
					fmt.Printf("- %s\n", formatWithUnderscores(testValueSum))
				}
				break
			}
		}
	}

	fmt.Println("Part 1 solution -->", testValueSum)
}

func add(op1 int64, op2 int64) int64 {
	return op1 + op2
}

func multiply(op1 int64, op2 int64) int64 {
	return op1 * op2
}

func allPermutations(accum []int, length int, vals []int) [][]int {
	var perms [][]int

	if length > 0 {
		for _, val := range vals {
			result := allPermutations(append(accum, val), length-1, vals)
			perms = slices.Concat(perms, result)
		}

		return perms
	} else {
		return append(perms, accum)
	}
}

func formatWithUnderscores(n int64) string {
	// Convert number to string
	str := strconv.FormatInt(n, 10)

	// Handle numbers less than 1000
	if len(str) <= 3 {
		return str
	}

	// Work from right to left, adding underscores
	var result strings.Builder
	for i := 0; i < len(str); i++ {
		// Add underscore before every 3rd digit from the right, but not at the start
		if i > 0 && (len(str)-i)%3 == 0 {
			result.WriteByte('_')
		}
		result.WriteByte(str[i])
	}

	return result.String()
}
