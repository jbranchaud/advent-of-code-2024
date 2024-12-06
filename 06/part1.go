package main

import (
	"fmt"
)

func part1(debug bool) {
	// puzzleToSolve := "sample"
	puzzleToSolve := "real"

	input, err := getInput("part1", puzzleToSolve)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	maxX := len(input[0])
	maxY := len(input)

	var guard Guard
	var obstaclesByX map[int][]Point
	var obstaclesByY map[int][]Point
	obstaclesByX = make(map[int][]Point)
	obstaclesByY = make(map[int][]Point)

	for i, line := range input {
		for j, cell := range line {
			x := j
			y := i
			p := Point{X: x, Y: y}
			switch cell {
			case '#':
				upsertPoint(obstaclesByX, x, p)
				upsertPoint(obstaclesByY, y, p)
			case '^':
				guard = Guard{Dir: UP, Loc: Point{X: x, Y: y}}
			default:
				// no-op
			}
		}
	}

	guardOnMap := true
	locationsVisited := make(map[Point]int)

	for guardOnMap {
		event, nextPosition := getNextPosition(guard, obstaclesByX, obstaclesByY, maxX, maxY)

		positions := listPositionsBetween(guard.Loc, nextPosition)
		for _, position := range positions {
			locationsVisited[position]++
		}

		if event == REACHED_EDGE {
			// reached the edge of the map
			guardOnMap = false
		} else if event == ENCOUNTERED_OBSTACLE {
			// walk guard to the next obstacle, turn, and continue
			guard.SetLocation(nextPosition)
			guard.TurnRight()
		} else {
			panic(fmt.Sprintf("Unknown event %d", event))
		}
	}

	count := len(locationsVisited)

	if debug {
		fmt.Println(input)
	}

	fmt.Println("Part 1 locations visited:", count)
}
