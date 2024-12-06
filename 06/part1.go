package main

import (
	"fmt"
)

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

type Point struct {
	X int
	Y int
}

type Guard struct {
	Loc Point
	Dir Direction
}

func (g *Guard) Move(steps int) {
	switch g.Dir {
	case UP:
		g.Loc.Y -= steps
	case RIGHT:
		g.Loc.X += steps
	case DOWN:
		g.Loc.Y += steps
	case LEFT:
		g.Loc.X -= steps
	}
}

func (g *Guard) MoveTo(position Point) {
	currPositions := fmt.Sprintf("g.Loc: (%d, %d), pos: (%d,%d)", g.Loc.X, g.Loc.Y, position.X, position.Y)
	switch g.Dir {
	case UP:
		assert(g.Loc.X == position.X, currPositions)
		steps := g.Loc.Y - position.Y
		g.Move(steps)
	case RIGHT:
		assert(g.Loc.Y == position.Y, currPositions)
		steps := position.X - g.Loc.X
		g.Move(steps)
	case DOWN:
		assert(g.Loc.X == position.X, currPositions)
		steps := position.Y - g.Loc.Y
		g.Move(steps)
	case LEFT:
		assert(g.Loc.Y == position.Y, currPositions)
		steps := g.Loc.X - position.X
		g.Move(steps)
	}
}

func (g *Guard) TurnRight() {
	g.Dir = (g.Dir + 1) % 4
}

func (g *Guard) GetLocation() Point {
	return g.Loc
}

func (g *Guard) SetLocation(p Point) {
	g.Loc = p
}

type Event int

const (
	ENCOUNTERED_OBSTACLE Event = iota
	REACHED_EDGE
)

func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func upsertPoint(pointMap map[int][]Point, key int, p Point) {
	points := pointMap[key]
	if len(points) == 0 {
		points = []Point{p}
	} else {
		points = append(points, p)
	}

	pointMap[key] = points
}

func listPositionsBetween(p1 Point, p2 Point) []Point {
	positions := []Point{}
	if p1.X == p2.X {
		// grab positions along Y axis
		if p1.Y > p2.Y {
			for i := p2.Y; i <= p1.Y; i++ {
				positions = append(positions, Point{X: p1.X, Y: i})
			}
		} else if p1.Y < p2.Y {
			for i := p1.Y; i <= p2.Y; i++ {
				positions = append(positions, Point{X: p1.X, Y: i})
			}
		} else {
			return []Point{p1}
		}

		return positions
	}
	if p1.Y == p2.Y {
		// grab positions along X axis
		if p1.X > p2.X {
			for i := p2.X; i <= p1.X; i++ {
				positions = append(positions, Point{X: i, Y: p1.Y})
			}
		} else if p1.X < p2.X {
			for i := p1.X; i <= p2.X; i++ {
				positions = append(positions, Point{X: i, Y: p1.Y})
			}
		} else {
			return []Point{p1}
		}

		return positions
	}

	panic(fmt.Sprintf("These two points aren't compat: (%d,%d) <> (%d,%d)", p1.X, p1.Y, p2.X, p2.Y))
}

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
	// var obstacles []Point
	var obstaclesByX map[int][]Point
	var obstaclesByY map[int][]Point
	obstaclesByX = make(map[int][]Point)
	obstaclesByY = make(map[int][]Point)

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

	for guardOnMap {
		event, nextPosition := getNextPosition(guard, obstaclesByX, obstaclesByY, maxX, maxY)

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

	count := len(locationsVisited)

	if debug {
		fmt.Println(input)
	}

	fmt.Println("Part 1 locations visited:", count)
}

func getNextPosition(
	guard Guard,
	obstaclesByX map[int][]Point,
	obstaclesByY map[int][]Point,
	maxX int,
	maxY int,
) (Event, Point) {
	// if the current position is on the edge and we're about to move "out of bounds"
	// So, we start with nextPosition being the current location of the guard before
	// advancing it.
	nextPosition := guard.Loc

	var stableAxis string
	var stableAxisVal int
	var terminalPoint Point
	var obstaclesOnAxis []Point
	// var closestObstacleCheck func(guardPos Point, obstaclePos Point, closestObstaclePos Point) bool
	var closestObstacleCheck = func(stableAxis string, guardPos Point, obstaclePos Point, closestObstaclePos Point) bool {
		if stableAxis == "x" {
			if guard.Dir == UP {
				return obstaclePos.Y < guardPos.Y && obstaclePos.Y > closestObstaclePos.Y
			}
			if guard.Dir == DOWN {
				return obstaclePos.Y > guardPos.Y && obstaclePos.Y < closestObstaclePos.Y
			}

			msg := fmt.Sprintf("Direction %d doesn't go with stableAxis x", guard.Dir)
			panic(msg)
		}

		if stableAxis == "y" {
			if guard.Dir == RIGHT {
				return obstaclePos.X > guardPos.X && obstaclePos.X < closestObstaclePos.X
			}
			if guard.Dir == LEFT {
				return obstaclePos.X < guardPos.X && obstaclePos.X > closestObstaclePos.X
			}

			msg := fmt.Sprintf("Direction %d doesn't go with stableAxis y", guard.Dir)
			panic(msg)
		}

		panic("Invalid stable axis")
	}

	if guard.Dir == UP || guard.Dir == DOWN {
		stableAxis = "x"
		stableAxisVal = nextPosition.X
		if guard.Dir == UP {
			terminalPoint = Point{X: stableAxisVal, Y: -1}
		} else {
			terminalPoint = Point{X: stableAxisVal, Y: maxY}
		}
		obstaclesOnAxis = obstaclesByX[stableAxisVal]
	} else if guard.Dir == RIGHT || guard.Dir == LEFT {
		stableAxis = "y"
		stableAxisVal = nextPosition.Y
		if guard.Dir == RIGHT {
			terminalPoint = Point{X: maxX, Y: stableAxisVal}
		} else {
			terminalPoint = Point{X: -1, Y: stableAxisVal}
		}
		obstaclesOnAxis = obstaclesByY[stableAxisVal]
	}

	closestObstaclePos := terminalPoint
	for _, obstaclePos := range obstaclesOnAxis {
		if closestObstacleCheck(stableAxis, guard.Loc, obstaclePos, closestObstaclePos) {
			closestObstaclePos = obstaclePos
		}
	}

	reachedEdge := false
	switch guard.Dir {
	case UP:
		nextPosition = Point{X: stableAxisVal, Y: closestObstaclePos.Y + 1}
		reachedEdge = closestObstaclePos.Y == -1
	case RIGHT:
		nextPosition = Point{X: closestObstaclePos.X - 1, Y: stableAxisVal}
		reachedEdge = closestObstaclePos.X == maxX
	case DOWN:
		nextPosition = Point{X: stableAxisVal, Y: closestObstaclePos.Y - 1}
		reachedEdge = closestObstaclePos.Y == maxY
	case LEFT:
		nextPosition = Point{X: closestObstaclePos.X + 1, Y: stableAxisVal}
		reachedEdge = closestObstaclePos.X == -1
	}

	if reachedEdge {
		return REACHED_EDGE, nextPosition
	} else {
		return ENCOUNTERED_OBSTACLE, nextPosition
	}

	// switch guard.Dir {
	// case UP:
	// 	stableAxis = "x"
	// x := nextPosition.X

	// if nextPosition.Y == 0 {
	// 	return REACHED_EDGE, nextPosition
	// } else {
	// 	// otherwise, we want to advance the position until we run into something
	// 	obstaclesOnColumn := obstaclesByX[x]

	// 	// find the largest Y on the column, that's what the guard will run into first
	// 	largestObstaclePosition := Point{X: x, Y: -1}
	// 	for _, obstaclePos := range obstaclesOnColumn {
	// 		if closestObstacleCheck(guard.Loc, obstaclePos, largestObstaclePosition) {
	// 			largestObstaclePosition = obstaclePos
	// 		}
	// 		// if obstaclePos.Y < guard.Loc.Y && obstaclePos.Y > largestObstaclePosition.Y {
	// 		// 	largestObstaclePosition = obstaclePos
	// 		// }
	// 	}

	// 	nextPosition = Point{X: x, Y: largestObstaclePosition.Y + 1}
	// 	if largestObstaclePosition.Y == -1 {
	// 		return REACHED_EDGE, nextPosition
	// 	} else {
	// 		return ENCOUNTERED_OBSTACLE, nextPosition
	// 	}
	// }
	// case RIGHT:
	// 	stableAxis = "y"
	// case DOWN:
	// 	stableAxis = "x"
	// case LEFT:
	// 	stableAxis = "y"
	// }
}
