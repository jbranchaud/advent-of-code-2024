package main

import (
	"fmt"
	"regexp"
)

func part1(debug bool) {
	// puzzleToSolve := "small"
	// puzzleToSolve := "medium"
	puzzleToSolve := "real"

	input, err := getInput("part1", puzzleToSolve)
	if err != nil {
		fmt.Printf("Error getting input: %v\n", err)
		return
	}
	occurrences := findXmas(input, debug)

	fmt.Printf("Solution to %s --> %d", puzzleToSolve, occurrences)
}

func findXmas(lines []string, debug bool) int {
	horizontalLines := lines
	verticalLines := rotateLines(lines)

	// for left diagonal, push from the top
	leftDiagonalLines := shiftLinesForDiagonals("left", lines)

	// for right diagonal, push from bottom
	rightDiagonalLines := shiftLinesForDiagonals("right", lines)

	horizontalCount := getXmasCount(horizontalLines)
	if debug {
		fmt.Println("Horizontal count:", horizontalCount)
	}
	verticalCount := getXmasCount(verticalLines)
	if debug {
		fmt.Println("Vertical count:", verticalCount)
	}
	leftDiagonalCount := getXmasCount(leftDiagonalLines)
	if debug {
		fmt.Println("Left Diagonal count:", leftDiagonalCount)
	}
	rightDiagonalCount := getXmasCount(rightDiagonalLines)
	if debug {
		fmt.Println("Right Diagonal count:", rightDiagonalCount)
	}

	return horizontalCount + verticalCount + leftDiagonalCount + rightDiagonalCount
}

var xmasMatcher = regexp.MustCompile(`XMAS`)

func getXmasCount(lines []string) int {
	var count int

	for _, line := range lines {
		count += getXmasCountForLine(line)
	}

	return count
}

func getXmasCountForLine(line string) int {
	reversedLine := reverseString(line)

	forwardMatches := xmasMatcher.FindAllString(line, -1)
	backwardMatches := xmasMatcher.FindAllString(reversedLine, -1)

	return len(forwardMatches) + len(backwardMatches)
}
