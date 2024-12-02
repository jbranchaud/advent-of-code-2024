package main

import (
	"aoc_2024/util"
	"fmt"
)

// Part 1 --> 660
func part1() {
	reports, err := util.ReadLinesOfInts("input1.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	var safeReports [][]int

	for _, report := range reports {
		firstLevel, remainingReport := util.Split(report)
		if isValidReport(firstLevel, remainingReport, "UNKNOWN", 0) {
			safeReports = append(safeReports, report)
		}
	}

	fmt.Printf("Number of safe reports: %d\n", len(safeReports))
}

func isValidReport(currentLevel int, report []int, direction string, depth int) bool {
	if len(report) == 0 {
		return true
	}

	changeMagnitude := "STABLE"
	nextDirection := "UNKNOWN"
	if currentLevel == report[0] {
		nextDirection = "FLAT"
		changeMagnitude = "NONE"
	} else if currentLevel < report[0] {
		nextDirection = "INCREASING"
		if report[0]-currentLevel > 3 {
			changeMagnitude = "UNSTABLE"
		}
	} else {
		nextDirection = "DECREASING"
		if currentLevel-report[0] > 3 {
			changeMagnitude = "UNSTABLE"
		}
	}

	nextLevel, remainingReport := util.Split(report)

	switch {
	case nextDirection == "FLAT":
		return false
	case changeMagnitude == "STABLE" && direction == "UNKNOWN":
		return isValidReport(nextLevel, remainingReport, nextDirection, depth)
	case changeMagnitude == "STABLE" && direction == nextDirection:
		return isValidReport(nextLevel, remainingReport, nextDirection, depth)
	default:
		return false
	}
}
