package main

import (
	"aoc_2024/util"
	"fmt"
	"regexp"
	"strconv"
)

func part1() {
	programs, err := util.ReadLines("input1.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	mulMatcher := regexp.MustCompile(pattern)

	var sum int

	for _, program := range programs {
		matches := mulMatcher.FindAllStringSubmatch(program, -1)

		for _, match := range matches {
			// ignore match[0], that's the whole match
			// match[1] is first capture group
			// match[2] is second capture group

			op1, err := strconv.Atoi(match[1])
			if err != nil {
				panic(fmt.Errorf("error parsing op1 int: %v", err))
			}
			op2, err := strconv.Atoi(match[2])
			if err != nil {
				panic(fmt.Errorf("error parsing op2 int: %v", err))
			}

			sum += (op1 * op2)
		}
	}

	fmt.Println("Total: ", sum)
}
