package main

import "fmt"

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

func (p Point) Clone() Point {
	return Point{
		X: p.X,
		Y: p.Y,
	}
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

func reachablePathSegmentToRight(dir Direction, pathSegments []Segment, currStep Point, obstaclesByX map[int][]Point, obstaclesByY map[int][]Point, debug bool) bool {
	anyValidSegment := false
	dirToRight := (dir + 1) % 4

	for _, segment := range pathSegments {
		if segment.Dir != dirToRight {
			continue
		}

		switch segment.Dir {
		case UP:
			stableAxisVal := currStep.X

			// check that segment is off to right, not off to left
			// or that we are ON the segment
			if segment.Start.X != stableAxisVal {
				continue // not reachable on current axis
			}
			if segment.End.Y > currStep.Y {
				continue // not to the right
			}

			blockingObstacle := false

			if debug {
				fmt.Printf(". >>> segment: %d, (%d,%d)-->(%d,%d), for: (%d,%d)\n", segment.Dir, segment.Start.X, segment.Start.Y, segment.End.X, segment.End.Y, currStep.X, currStep.Y)
			}

			for _, obstaclePos := range obstaclesByX[stableAxisVal] {
				if obstaclePos.Y < currStep.Y && obstaclePos.Y > segment.Start.Y {
					blockingObstacle = true
					break
				}
			}

			if blockingObstacle {
				continue
			}
		case RIGHT:
			stableAxisVal := currStep.Y

			// check that segment is off to right, not off to left
			// or that we are ON the segment
			if segment.Start.Y != stableAxisVal {
				continue // not reachable on current axis
			}
			if segment.End.X < currStep.X {
				continue // not to the right
			}

			blockingObstacle := false

			if debug {
				fmt.Printf(". >>> segment: %d, (%d,%d)-->(%d,%d), for: (%d,%d)\n", segment.Dir, segment.Start.X, segment.Start.Y, segment.End.X, segment.End.Y, currStep.X, currStep.Y)
			}

			for _, obstaclePos := range obstaclesByY[stableAxisVal] {
				if obstaclePos.X > currStep.X && obstaclePos.X < segment.Start.X {
					blockingObstacle = true
					break
				}
			}

			if blockingObstacle {
				continue
			}
		case DOWN:
			stableAxisVal := currStep.X

			if debug {
				fmt.Printf(". >>> segment: %d, (%d,%d)-->(%d,%d), for: (%d,%d)", segment.Dir, segment.Start.X, segment.Start.Y, segment.End.X, segment.End.Y, currStep.X, currStep.Y)
			}

			// check that segment is off to right, not off to left
			// or that we are ON the segment
			if segment.Start.X != stableAxisVal {
				fmt.Print(" &\n")
				continue // not reachable on current axis
			}
			if segment.End.Y < currStep.Y {
				fmt.Print(" $\n")
				continue // not to the right
			}

			blockingObstacle := false

			for _, obstaclePos := range obstaclesByX[stableAxisVal] {
				if obstaclePos.Y > currStep.Y && obstaclePos.Y < segment.Start.Y {
					blockingObstacle = true
					fmt.Print(" @\n")
					break
				}
			}

			if blockingObstacle {
				continue
			}
		case LEFT:
			stableAxisVal := currStep.Y

			// check that segment is off to right, not off to left
			// or that we are ON the segment
			if segment.Start.Y != stableAxisVal {
				continue // not reachable on current axis
			}
			if segment.End.X > currStep.X {
				continue // not to the right
			}

			blockingObstacle := false

			if debug {
				fmt.Printf(". >>> segment: %d, (%d,%d)-->(%d,%d), for: (%d,%d)\n", segment.Dir, segment.Start.X, segment.Start.Y, segment.End.X, segment.End.Y, currStep.X, currStep.Y)
			}

			for _, obstaclePos := range obstaclesByY[stableAxisVal] {
				if obstaclePos.X < currStep.X && obstaclePos.X > segment.Start.X {
					blockingObstacle = true
					break
				}
			}

			if blockingObstacle {
				continue
			}
		}

		if debug {
			fmt.Printf("* +++ segment: %d, (%d,%d)-->(%d,%d)\n", segment.Dir, segment.Start.X, segment.Start.Y, segment.End.X, segment.End.Y)
		}

		anyValidSegment = true
		break
	}

	return anyValidSegment
}

func searchAlongNextSegment(
	guard Guard,
	pathSegments []Segment,
	obstaclesByX map[int][]Point,
	obstaclesByY map[int][]Point,
	maxX int,
	maxY int,
	debug bool,
) (Event, Point, []Point, []Segment) {
	var potentialObstaclePositions []Point

	// we want to do much of the same as `getNextPosition`
	// but for each Point we are about to advance, we need to check if
	// it would be a good place to put an obstacle.
	// If it is, we add it to a list that gets returned.
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
		for nextY := guard.Loc.Y; nextY > closestObstaclePos.Y; nextY-- {
			currStep := Point{X: stableAxisVal, Y: nextY}
			nextStep := Point{X: stableAxisVal, Y: nextY - 1}
			nextStepUnobstructed := nextY-1 != closestObstaclePos.Y

			if nextStepUnobstructed && reachablePathSegmentToRight(guard.Dir, pathSegments, currStep, obstaclesByX, obstaclesByY, debug) {
				potentialObstaclePositions = append(potentialObstaclePositions, nextStep)
			}
		}
		nextPosition = Point{X: stableAxisVal, Y: closestObstaclePos.Y + 1}
		reachedEdge = closestObstaclePos.Y == -1
	case RIGHT:
		for nextX := guard.Loc.X; nextX < closestObstaclePos.X; nextX++ {
			currStep := Point{X: nextX, Y: stableAxisVal}
			nextStep := Point{X: nextX + 1, Y: stableAxisVal}
			nextStepUnobstructed := nextX+1 != closestObstaclePos.X

			if nextStepUnobstructed && reachablePathSegmentToRight(guard.Dir, pathSegments, currStep, obstaclesByX, obstaclesByY, debug) {
				potentialObstaclePositions = append(potentialObstaclePositions, nextStep)
			}
		}
		nextPosition = Point{X: closestObstaclePos.X - 1, Y: stableAxisVal}
		reachedEdge = closestObstaclePos.X == maxX
	case DOWN:
		for nextY := guard.Loc.Y; nextY < closestObstaclePos.Y; nextY++ {
			currStep := Point{X: stableAxisVal, Y: nextY}
			nextStep := Point{X: stableAxisVal, Y: nextY + 1}
			nextStepUnobstructed := nextY+1 != closestObstaclePos.Y

			if nextStepUnobstructed && reachablePathSegmentToRight(guard.Dir, pathSegments, currStep, obstaclesByX, obstaclesByY, debug) {
				potentialObstaclePositions = append(potentialObstaclePositions, nextStep)
			}
		}
		nextPosition = Point{X: stableAxisVal, Y: closestObstaclePos.Y - 1}
		reachedEdge = closestObstaclePos.Y == maxY
	case LEFT:
		for nextX := guard.Loc.X; nextX > closestObstaclePos.X; nextX-- {
			currStep := Point{X: nextX, Y: stableAxisVal}
			nextStep := Point{X: nextX - 1, Y: stableAxisVal}
			nextStepUnobstructed := nextX-1 != closestObstaclePos.X

			if nextStepUnobstructed && reachablePathSegmentToRight(guard.Dir, pathSegments, currStep, obstaclesByX, obstaclesByY, debug) {
				potentialObstaclePositions = append(potentialObstaclePositions, nextStep)
			}
		}
		nextPosition = Point{X: closestObstaclePos.X + 1, Y: stableAxisVal}
		reachedEdge = closestObstaclePos.X == -1
	}

	newSegment := Segment{Dir: guard.Dir, Start: guard.Loc.Clone(), End: nextPosition.Clone()}
	if debug {
		fmt.Printf("- new segment: %d, (%d,%d)-->(%d,%d)\n", newSegment.Dir, newSegment.Start.X, newSegment.Start.Y, newSegment.End.X, newSegment.End.Y)
	}
	updatedPathSegments := append(pathSegments, newSegment)

	if reachedEdge {
		return REACHED_EDGE, nextPosition, potentialObstaclePositions, updatedPathSegments
	} else {
		return ENCOUNTERED_OBSTACLE, nextPosition, potentialObstaclePositions, updatedPathSegments
	}
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
}
