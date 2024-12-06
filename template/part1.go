package main

import (
	"fmt"
)

func part1(debug bool) {
	puzzleToSolve := "sample"
	// puzzleToSolve := "real"

	input, err := getInput("part1", puzzleToSolve)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	if debug {
		fmt.Println(input)
	}

	fmt.Println("Part 1 incoming...")
}
