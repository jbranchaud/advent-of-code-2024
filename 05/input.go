package main

import (
	"aoc_2024/util"
	"fmt"
	"path/filepath"
)

func getInput(part string, name string) ([]string, []string, error) {
	var filename string

	switch name {
	case "sample":
		filename = "sample.txt"
	case "real":
		filename = "input1.txt"
	default:
		return []string{}, []string{}, fmt.Errorf("getInput was given an unrecognized name: %s", name)
	}

	path := filepath.Join("inputs", filename)
	input, err := util.ReadLines(path)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return []string{}, []string{}, err
	}

	var rules []string
	var updates []string

	typeToCollect := "rules"

	for i, line := range input {
		switch {
		case line == "":
			typeToCollect = "updates"
		case typeToCollect == "rules":
			rules = append(rules, line)
		case typeToCollect == "updates":
			updates = append(updates, line)
		default:
			msg := fmt.Sprintf("Something went wrong collating the input, line %d", i)
			panic(msg)
		}
	}

	return rules, updates, nil
}
