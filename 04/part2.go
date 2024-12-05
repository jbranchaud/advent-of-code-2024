package main

import (
	"fmt"
	"regexp"
)

func part2(debug bool) {
	// puzzleToSolve := "small"
	// puzzleToSolve := "medium"
	puzzleToSolve := "real"

	input, err := getInput("part2", puzzleToSolve)
	if err != nil {
		fmt.Printf("Error getting input: %v\n", err)
		return
	}

	leftDiagonalLines := shiftLinesForDiagonals("left", input)
	rightDiagonalLines := shiftLinesForDiagonals("right", input)

	occurrences := countCrossingMases(leftDiagonalLines, rightDiagonalLines, debug)

	fmt.Printf("Solution to %s --> %d", puzzleToSolve, occurrences)
}

func countCrossingMases(leftDiagonalLines []string, rightDiagonalLines []string, debug bool) int {
	if debug {
		fmt.Println("-- left --")
	}
	leftMasCoordinates := getMasCoordinates(leftDiagonalLines, debug)
	if debug {
		fmt.Println("-- right --")
	}
	rightMasCoordinates := getMasCoordinates(rightDiagonalLines, debug)

	var count int

	for _, leftCoord := range leftMasCoordinates {
		for _, rightCoord := range rightMasCoordinates {
			match, err := CompareCoordinates(leftCoord, rightCoord)
			if err != nil {
				panic(err)
			}

			if match {
				count++
			}
		}
	}

	return count
}

var masMatcher = regexp.MustCompile("MAS")
var samMatcher = regexp.MustCompile("SAM")

func getMasCoordinates(lines []string, debug bool) [][]int {
	var coords [][]int

	for i, line := range lines {
		masMatches := masMatcher.FindAllStringIndex(line, -1)
		samMatches := samMatcher.FindAllStringIndex(line, -1)

		matches := append(masMatches, samMatches...)

		for _, match := range matches {
			xCoord := match[0] + 1 // 'A' on the x axis
			yCoord := 0

			for n := 0; n < i; n++ {
				if lines[n][xCoord] != '*' {
					yCoord++
				}
			}

			if debug {
				fmt.Printf("X: %d, Y: %d\n", xCoord, yCoord)
			}

			coords = append(coords, []int{xCoord, yCoord})
		}
	}

	return coords
}

func CompareCoordinates(coord1, coord2 []int) (bool, error) {
	if len(coord1) != 2 || len(coord2) != 2 {
		return false, fmt.Errorf("Invalid coordinates: %v -- %v", coord1, coord2)
	}

	return (coord1[0] == coord2[0] && coord1[1] == coord2[1]), nil
}
