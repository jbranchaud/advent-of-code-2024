package main

import (
	"aoc_2024/util"
	"fmt"
	"strings"
)

func getInput(part string, name string) ([]string, error) {
	var small []string
	if part == "part1" {
		// Part 1: should have 5 occurrences
		small = strings.Split(`..X...
.SAMX.
.A..A.
XMAS.S
.X....`, "\n")
	} else if part == "part2" {
		small = strings.Split(`M.S
.A.
M.S`, "\n")
	} else {
		return []string{}, fmt.Errorf("No input for Part %s, Name %s", part, name)
	}

	var medium []string
	if part == "part1" {
		// Part 1: should have 18 occurrences
		medium = strings.Split(`MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`, "\n")
	} else if part == "part2" {
		medium = strings.Split(`.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`, "\n")
	} else {
		return []string{}, fmt.Errorf("No input for Part %s, Name %s", part, name)
	}

	switch name {
	case "small":
		return small, nil
	case "medium":
		return medium, nil
	case "real":
		input, err := util.ReadLines("input1.txt")
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			return []string{}, err
		}

		return input, nil
	default:
		return []string{}, fmt.Errorf("getInput was given an unrecognized name: %s", name)
	}
}
