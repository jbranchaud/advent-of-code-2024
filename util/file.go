package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadPairs(filename string) ([]int, []int, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	lineNumber := 0

	var firstList []int
	var secondList []int

	// Read each line
	for scanner.Scan() {
		lineNumber += 1

		// Split the line into fields
		fields := strings.Fields(scanner.Text())

		// Check if we have exactly two values
		if len(fields) != 2 {
			return nil, nil, fmt.Errorf("line %d: expected 2 fields, got %d", lineNumber, len(fields))
		}

		locationIdOne, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, nil, fmt.Errorf("line %d: expected field 1 to be int, got %s", lineNumber, fields[0])
		}
		locationIdTwo, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, nil, fmt.Errorf("line %d: expected field 2 to be int, got %s", lineNumber, fields[1])
		}

		firstList = append(firstList, locationIdOne)
		secondList = append(secondList, locationIdTwo)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	return firstList, secondList, nil
}
