package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "turns on debug mode, extra logging")
	flag.Parse()

	positionalArgs := flag.Args()

	if len(positionalArgs) < 1 {
		fmt.Println("Please specify which part to run: 1 or 2")
		os.Exit(1)
	}

	switch positionalArgs[0] {
	case "1":
		part1(debug)
	case "2":
		part2(debug)
	default:
		fmt.Printf("Invalid part specified: %s. Please use 1 or 2\n", positionalArgs[0])
		os.Exit(1)
	}
}
