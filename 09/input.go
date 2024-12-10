package main

import (
	"aoc_2024/util"
	"fmt"
	"path/filepath"
)

func getInput(part string, name string) ([]string, error) {
	var filename string

	switch name {
	case "sample":
		filename = "sample.txt"
	case "real":
		filename = "input1.txt"
	default:
		return []string{}, fmt.Errorf("getInput was given an unrecognized name: %s", name)
	}

	path := filepath.Join("inputs", filename)
	input, err := util.ReadLines(path)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return []string{}, err
	}

	return input, nil
}
