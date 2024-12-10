package main

import (
	"fmt"
)

func part2(debug bool) {
	puzzleToSolve := "sample"
	// puzzleToSolve := "real"

	input, err := getInput("part2", puzzleToSolve)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	if debug {
		fmt.Println(input)
	}

	fmt.Println("Part 2 incoming...")
}
