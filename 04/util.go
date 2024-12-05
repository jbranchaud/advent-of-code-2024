package main

import "strings"

func rotateLines(lines []string) []string {
	var rotatedLines []string

	outerLoopCount := len(lines[0])
	innerLoopCount := len(lines)

	for i := range outerLoopCount {
		var builder strings.Builder
		builder.Grow(innerLoopCount)

		for j := range innerLoopCount {
			currChar := lines[j][i]
			builder.WriteByte(currChar)
		}

		rotatedLines = append(rotatedLines, builder.String())
	}

	return rotatedLines
}

// func reverseLines(lines []string) []string {
// 	var reversedLines []string

// 	for _, line := range lines {
// 		reversedLines = append(reversedLines, reverseString(line))
// 	}

// 	return reversedLines
// }

func reverseString(str string) string {
	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func shiftLinesForDiagonals(direction string, lines []string) []string {
	var diagonalLines []string

	maxShift := len(lines) - 1
	for i, line := range lines {
		var leftPad int
		var rightPad int

		switch direction {
		case "left":
			leftPad = maxShift - i
			rightPad = i
		case "right":
			leftPad = i
			rightPad = maxShift - i
		default:
			panic("left or right not specified for diagonal builder")
		}

		var builder strings.Builder
		if leftPad > 0 {
			builder.WriteString(strings.Repeat("*", leftPad))
		}
		builder.WriteString(line)
		if rightPad > 0 {
			builder.WriteString(strings.Repeat("*", rightPad))
		}

		diagonalLines = append(diagonalLines, builder.String())
	}

	return rotateLines(diagonalLines)
}

//     ..X...
//    .SAMX.
//   .A..A.
//  XMAS.S
// .X....
