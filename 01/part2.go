package main

import (
	"fmt"
	"sort"

	"aoc_2024/util"
)

func part2() {
	listOne, listTwo, err := util.ReadPairs("./input1.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	sort.Ints(listOne)
	sort.Ints(listTwo)

	similarityScore := 0

	for _, item := range listOne {
		appearanceCount := 0

		for _, currentItem := range listTwo {
			if item == currentItem {
				appearanceCount += 1
			}

			if item < currentItem {
				break
			}
		}

		similarityScore += (item * appearanceCount)
	}

	// util.PrintFirstAndLast("One", listOne)
	// util.PrintFirstAndLast("Two", listTwo)

	fmt.Println(similarityScore)
}
