package main

import (
	"aoc_2024/util"
	"cmp"
	"fmt"
	"maps"
	"os"
	"slices"
)

type Block struct {
	Size     int
	Contents *int
	Parts    *[]int
}

func part1(debug bool) {
	puzzleToSolve := "sample"
	// puzzleToSolve := "real"

	input, err := getInput("part1", puzzleToSolve)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	diskMap, err := util.StringToInts(input[0], "")

	if debug {
		fmt.Printf("Disk Map: %v\n", diskMap)
	}

	// work through both ends of the list at the same time?
	// have a map to track gaps of a certain size, where the key is the gap size and the value is a slice of indexes
	file := make(map[int]Block)

	fileMapIndex := 0
	fileId := 0
	for i, val := range diskMap {
		if i%2 == 0 {
			// even indexes are file blocks
			fileIdParts := IntToParts(fileId)
			blockSize := len(fileIdParts)

			if debug {
				fmt.Printf("Adding content block: Num %d, File %d\n", val, fileId)
			}

			// for each of the blocks
			for range val {
				file[fileMapIndex] = Block{Size: blockSize, Contents: &fileId, Parts: &fileIdParts}
				newBlock := file[fileMapIndex]
				if debug {
					fmt.Printf("  Block contains: %d, %d, %v\n", newBlock.Size, *newBlock.Contents, *newBlock.Parts)
				}
				fileMapIndex += blockSize
			}

			// next item in the diskmap will need an incremented fileId
			fileId++
		} else {
			// odd indexes are free space / gaps
			file[fileMapIndex] = Block{Size: val, Contents: nil, Parts: nil}
			fileMapIndex += val
		}
	}

	if debug {
		printFile(file)
	}

	for true { // while still fragmented
		encounteredGaps := false
		blockIndexes := []int{}
		lastBlockIndex := 0
		contentBlocks := []Block{}
		gapBlocks := []Block{}
		for currIndex, block := range file {
			blockIndexes = append(blockIndexes, currIndex)

			if block.Contents != nil {
				contentBlocks = append(contentBlocks, block)

				if currIndex > lastBlockIndex {
					lastBlockIndex = currIndex
				}
			} else {
				gapBlocks = append(gapBlocks, block)

				encounteredGaps = true
			}
		}

		// slices.Sort(blockIndexes)
		SortItems(blockIndexes, ASC)

		if !encounteredGaps {
			break
		}

		// otherwise, place the last block at the earliest point we can in the
		// file where it will fit.
		blockToMove := file[lastBlockIndex]
		newIndex := -1
		newGapSize := 0
		for _, index := range blockIndexes {
			block := file[index]

			if block.Contents != nil {
				continue
			}

			emptyBlock := block

			remainingGap := emptyBlock.Size - blockToMove.Size
			if remainingGap >= 0 {
				newIndex = index
				newGapSize = remainingGap
				break
			}
		}

		if newIndex < 0 {
			fmt.Printf("Block doesn't fit: index %d, size %d, contents: %d", lastBlockIndex, blockToMove.Size, *blockToMove.Contents)
			os.Exit(1)
		}

		// remove the existing empty block
		// maps.DeleteFunc(file, func(index int, _ Block) bool {
		// 	return index == newIndex
		// })

		// remove empty block and add the new content block with overwrite
		// file[newIndex] = Block{Size: blockToMove.Size, Contents: &*blockToMove.Contents, Parts: &*blockToMove.Parts}
		file[newIndex] = blockToMove

		// add a new gap block if a gap remains
		if newGapSize > 0 {
			file[newIndex+blockToMove.Size] = Block{Size: newGapSize}
		}

		// remove the moved block from the end of the list
		maps.DeleteFunc(file, func(index int, _ Block) bool {
			return index == lastBlockIndex
		})
	}

	if debug {
		fmt.Println(input)
	}

	checksum := 0

	for key := range maps.Keys(file) {
		block := file[key]

		checksum += *block.Contents * key
	}

	fmt.Println("Part 1 solution ->", checksum)
}

func IntToParts(val int) []int {
	// return early for 0 because that doesn't work with the `divisor > val` check
	if val == 0 {
		return []int{0}
	}

	parts := []int{}

	// val := 3457
	// 1) 3457 / 1     % 10 => 7, parts: [7]
	// 2) 3457 / 10    % 10 => 5, parts: [5,7]
	// 3) 3457 / 100   % 10 => 4, parts: [4,5,7]
	// 4) 3457 / 1000  % 10 => 3, parts: [3,4,5,7]
	// 5) 10000 > 3457 => break
	divisor := 1
	for true {
		if divisor > val {
			break
		}

		part := val / divisor % 10
		// preprend each part
		parts = slices.Concat([]int{part}, parts)
		divisor *= 10
	}

	return parts
}

func printFile(file map[int]Block) {
	sortedKeys := []int{}
	for key := range maps.Keys(file) {
		sortedKeys = append(sortedKeys, key)
	}
	SortItems(sortedKeys, ASC)
	// slices.Sort(sortedKeys)

	fmt.Println("Sorted keys:", sortedKeys)
	firstBlock := file[sortedKeys[0]]
	fmt.Printf("[0]: Size: %d, Contents: %d, Parts: %v\n", firstBlock.Size, *firstBlock.Contents, *&firstBlock.Parts)

	for _, key := range sortedKeys {
		block := file[key]

		if block.Contents == nil {
			for range block.Size {
				fmt.Print(".")
			}
		} else {
			content := *block.Contents
			fmt.Printf("%d", content)
		}
	}

	fmt.Print("\n")
}

type Direction int

const (
	ASC Direction = iota
	DESC
)

func SortItems[T cmp.Ordered](items []T, dir Direction) {
	slices.SortFunc(items, func(i, j T) int {
		if dir == ASC {
			return cmp.Compare(i, j)
		} else if dir == DESC {
			return cmp.Compare(j, i)
		} else {
			panic(fmt.Sprintf("Unrecognized sort direction: %d", dir))
		}
	})
}
