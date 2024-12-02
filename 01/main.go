package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify which part to run: 1 or 2")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "1":
		part1()
	case "2":
		part2()
	default:
		fmt.Printf("Invalid part specified: %s. Please use 1 or 2\n", os.Args[1])
		os.Exit(1)
	}
}
