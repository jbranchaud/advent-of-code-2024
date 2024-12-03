package main

import (
	"aoc_2024/util"
	"fmt"
)

func part2() {
	reports, err := util.ReadLinesOfInts("input1.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	var safeReports [][]int

	for _, report := range reports {
		if isMostlyValidReport(report, 0) {
			safeReports = append(safeReports, report)
		}
	}

	fmt.Printf("Number of safe reports: %d\n", len(safeReports))
}

func isMostlyValidReport(report []int, depth int) bool {
	if len(report) == 0 {
		return true
	}

	previousLevel, remainingReport := util.Split(report)

	direction := "UNKNOWN"
	reportStatus := "VALID"
	for _, level := range remainingReport {
		if direction == "UNKNOWN" {
			if previousLevel < level && level-previousLevel <= 3 {
				direction = "INCREASING"
			} else if previousLevel > level && previousLevel-level <= 3 {
				direction = "DECREASING"
			} else {
				reportStatus = "INVALID"
				break
			}
		} else if direction == "INCREASING" {
			if previousLevel < level && level-previousLevel <= 3 {
				// no-op
			} else {
				reportStatus = "INVALID"
				break
			}
		} else if direction == "DECREASING" {
			if previousLevel > level && previousLevel-level <= 3 {
				// no-op
			} else {
				reportStatus = "INVALID"
				break
			}
		}

		previousLevel = level
	}

	if reportStatus == "VALID" {
		return true
	} else {
		if depth > 0 {
			return false
		}

		anyValidPartialReport := false

		fmt.Printf("%d\n", report)
		for i := range len(report) {
			var partialReport []int

			for j, level := range report {
				if i != j {
					partialReport = append(partialReport, level)
				}
			}
			fmt.Printf("- %d: %d\n", i, partialReport)

			anyValidPartialReport = anyValidPartialReport || isMostlyValidReport(partialReport, 1)
		}

		return anyValidPartialReport
	}
}
