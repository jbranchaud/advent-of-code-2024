package main

import (
	"aoc_2024/util"
	"fmt"
	"regexp"
	"strconv"
)

// https://adventofcode.com/2024/day/3
// part 2 solution => 102467299
func part2(debug bool) {
	programs, err := util.ReadLines("input1.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	mulMatcher := compileMatcher(`mul\((\d{1,3}),(\d{1,3})\)`)
	doMatcher := compileMatcher(`do\(\)`)
	dontMatcher := compileMatcher(`don't\(\)`)

	var sum int

	defaultStatus := "DO"
	// status can be "DO" or "DONT"
	status := defaultStatus

	for _, program := range programs {
		// Find indexes of all `mul(N,M)` in `program` string
		mulIndexes := mulMatcher.FindAllStringSubmatchIndex(program, -1)
		// Find indexes of all `do()` in `program` string
		doIndexes := doMatcher.FindAllStringSubmatchIndex(program, -1)
		// Find indexes of all `don't()` in `program` string
		dontIndexes := dontMatcher.FindAllStringSubmatchIndex(program, -1)

		// cursors for tracking place in each _Indexes slice
		mulCursor, doCursor, dontCursor := 0, 0, 0

		// step through ever single position in the `program` string
		// if the position matches a recognized syntax index, then parse it
		// otherwise no-op and proceed
		for i := range len(program) {
			switch {
			case len(mulIndexes) > mulCursor && mulIndexes[mulCursor][0] == i:
				sum += parseMultiplicationMatch(program, mulIndexes, &mulCursor, status, debug)
			case len(doIndexes) > doCursor && doIndexes[doCursor][0] == i:
				parseDoMatch(&doCursor, &status, debug)
			case len(dontIndexes) > dontCursor && dontIndexes[dontCursor][0] == i:
				parseDontMatch(&dontCursor, &status, debug)
			}
		}
	}

	fmt.Println("Total: ", sum)
}

func compileMatcher(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func parseMultiplicationMatch(program string, mulIndexes [][]int, mulCursor *int, status string, debug bool) int {
	match := mulIndexes[*mulCursor]
	*mulCursor++
	if status == "DONT" {
		if debug {
			fmt.Println("x", program[match[0]:match[1]])
		}

		return 0
	}

	op1, err := strconv.Atoi(program[match[2]:match[3]])
	if err != nil {
		panic(fmt.Errorf("error parsing op1 int: %v", err))
	}
	op2, err := strconv.Atoi(program[match[4]:match[5]])
	if err != nil {
		panic(fmt.Errorf("error parsing op2 int: %v", err))
	}

	mulResult := op1 * op2
	if debug {
		fmt.Println(program[match[0]:match[1]], "=>", mulResult)
	}
	return mulResult
}

func parseDoMatch(doCursor *int, status *string, debug bool) {
	*doCursor++
	*status = "DO"
	if debug {
		fmt.Println("- DO() -")
	}
}

func parseDontMatch(dontCursor *int, status *string, debug bool) {
	*dontCursor++
	*status = "DONT"
	if debug {
		fmt.Println("- DON'T() -")
	}
}
