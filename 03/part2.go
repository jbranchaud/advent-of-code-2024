package main

import (
	"aoc_2024/util"
	"fmt"
	"regexp"
	"strconv"
)

// part 2 solution => 102467299
func part2(debug bool) {
	programs, err := util.ReadLines("input1.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	mulPattern := `mul\((\d{1,3}),(\d{1,3})\)`
	mulMatcher := regexp.MustCompile(mulPattern)

	doPattern := `do\(\)`
	doMatcher := regexp.MustCompile(doPattern)
	dontPattern := `don't\(\)`
	dontMatcher := regexp.MustCompile(dontPattern)

	var sum int

	defaultStatus := "DO"
	// status can be "DO" or "DONT"
	status := defaultStatus

	for _, program := range programs {
		mulIndexes := mulMatcher.FindAllStringSubmatchIndex(program, -1)
		doIndexes := doMatcher.FindAllStringSubmatchIndex(program, -1)
		dontIndexes := dontMatcher.FindAllStringSubmatchIndex(program, -1)
		// matches := mulMatcher.FindAllStringSubmatch(program, -1)

		mulCursor, doCursor, dontCursor := 0, 0, 0

		parseMultiplicationMatch := func() int {
			match := mulIndexes[mulCursor]
			mulCursor++
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

		parseDoMatch := func() {
			doCursor++
			status = "DO"
			if debug {
				fmt.Println("- DO() -")
			}
		}

		parseDontMatch := func() {
			dontCursor++
			status = "DONT"
			if debug {
				fmt.Println("- DON'T() -")
			}
		}

		for i := range len(program) {
			switch {
			case len(mulIndexes) > mulCursor && mulIndexes[mulCursor][0] == i:
				sum += parseMultiplicationMatch()
			case len(doIndexes) > doCursor && doIndexes[doCursor][0] == i:
				parseDoMatch()
			case len(dontIndexes) > dontCursor && dontIndexes[dontCursor][0] == i:
				parseDontMatch()
			}
		}
	}

	fmt.Println("Total: ", sum)
}
