package main

import (
	"fmt"
	"sort"

	"aoc_2024/util"
)

func part1() {
	listOne, listTwo, err := util.ReadPairs("./input1.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	sort.Ints(listOne)
	sort.Ints(listTwo)

	var differences []int
	differenceTotal := 0

	for i := 0; i < len(listOne); i++ {
		var diff int
		if listOne[i] > listTwo[i] {
			diff = listOne[i] - listTwo[i]
		} else {
			diff = listTwo[i] - listOne[i]
		}

		differenceTotal += diff
		differences = append(differences, diff)
	}

	// util.PrintFirstAndLast("One", listOne)
	// util.PrintFirstAndLast("Two", listTwo)

	fmt.Println(differenceTotal)
}
