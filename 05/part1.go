package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(debug bool) {
	// puzzleToSolve := "sample"
	puzzleToSolve := "real"

	ruleLines, updateLines, err := getInput("part1", puzzleToSolve)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	var sum int

	// build `rules` map
	rulesMap := buildRulesMap(ruleLines, debug)

	if debug {
		fmt.Println("Rules Map:")
		for key, values := range rulesMap {
			fmt.Printf(" . %s: %v\n", key, values)
		}
	}

	// parse each update line
	var updates [][]string
	for _, updateLine := range updateLines {
		updates = append(updates, strings.Split(updateLine, ","))
	}

	// walk through updates
	for i, update := range updates {
		valid := true

		for j := len(update) - 1; j >= 0; j-- {
			pageToCheck := update[j]
			rules := rulesMap[pageToCheck]

			for k := range j {
				if slices.Contains(rules, update[k]) {
					valid = false
					break
				}
			}
		}

		if valid {
			// grab middle value and add to sum
			middle := update[len(update)/2 : len(update)/2+1]
			middleAsInt, err := strconv.Atoi(middle[0])
			if err != nil {
				fmt.Printf("%d: Unable to parse value to int: %s\n", i, middle)
				os.Exit(1)
			}

			sum += middleAsInt

			if debug {
				fmt.Printf("%d: Valid update: %s, %v\n", i, middle, update)
			}
		} else {
			if debug {
				fmt.Printf("%d: Invalid update: %v\n", i, update)
			}
		}
	}

	fmt.Println("Part 1 solution -->", sum)
}
