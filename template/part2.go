package main

import (
	"aoc_2024/util"
	"fmt"
)

func part2(debug bool) {
	input, err := util.ReadLines("input1.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	if debug {
		fmt.Println(input)
	}

	fmt.Println("Part 2 incoming...")
}
