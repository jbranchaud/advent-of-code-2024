package main

import (
	"fmt"
)

type Segment struct {
	Dir   Direction
	Start Point
	End   Point
}

func part2(debug bool) {
	puzzleToSolve := "sample"
	// puzzleToSolve := "real"

	input, err := getInput("part2", puzzleToSolve)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	maxX := len(input[0])
	maxY := len(input)

	var guard Guard
	// var obstacles []Point
	var obstaclesByX map[int][]Point
	var obstaclesByY map[int][]Point
	obstaclesByX = make(map[int][]Point)
	obstaclesByY = make(map[int][]Point)

	var pathSegments []Segment

	for i, line := range input {
		for j, cell := range line {
			// TODO: Make sure I didn't get X/Y backwards
			x := j
			y := i
			p := Point{X: x, Y: y}
			switch cell {
			case '#':
				// obstacles = append(obstacles, p)
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
	newObstacleOptions := make(map[Point]int)

	for guardOnMap {
		event, nextPosition, potentialObstaclePositions, updatedPathSegments :=
			searchAlongNextSegment(
				guard,
				pathSegments,
				obstaclesByX,
				obstaclesByY,
				maxX,
				maxY,
				debug,
			)

		pathSegments = updatedPathSegments

		for _, obstaclePosition := range potentialObstaclePositions {
			if debug {
				fmt.Printf("New Obstacle Pos: (%d,%d)\n", obstaclePosition.X, obstaclePosition.Y)
			}
			newObstacleOptions[obstaclePosition]++
		}

		positions := listPositionsBetween(guard.Loc, nextPosition)
		for _, position := range positions {
			locationsVisited[position]++
		}

		if event == REACHED_EDGE {
			guardOnMap = false
			// add all positions between curr guard pos and nextPosition
		} else if event == ENCOUNTERED_OBSTACLE {
			// add all positions between curr guard pos and nextPosition
			guard.SetLocation(nextPosition)
			guard.TurnRight()
		} else {
			panic(fmt.Sprintf("Unknown event %d", event))
		}
	}

	count := len(newObstacleOptions)

	if debug {
		fmt.Println(input)
	}

	fmt.Println("Part 2 potential obstacles:", count)
}
