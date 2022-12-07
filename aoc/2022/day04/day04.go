package day04

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
	"log"
)

// gInput : [][]int : each row contains start and end ranges of both elves {e1s,e1e,e2s,e2e}
var gInput [][]int

/***** PART 01 Functions *****/

func SolveP1() string {
	count := 0
	for _, arr := range gInput {
		if arr[0] >= arr[2] && arr[1] <= arr[3] ||
			arr[2] >= arr[0] && arr[3] <= arr[1] {
			count++
		}
	}
	return fmt.Sprintf("%v", count)
}

/***** PART 02 Functions *****/

func SolveP2() string {
	count := 0
	for _, arr := range gInput {
		if inRange(arr[0], arr[2], arr[3]) ||
			inRange(arr[1], arr[2], arr[3]) ||
			inRange(arr[2], arr[0], arr[1]) ||
			inRange(arr[3], arr[0], arr[1]) {
			count++
		}
	}
	return fmt.Sprintf("%v", count)
}

/***** Common Functions *****/

// todo : move to utils
func inRange(x, rStart, rEnd int) bool {
	return x >= rStart && x <= rEnd
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	linesSplit := iutils.ExtractString2DFromString1D(lines, ",", nil, "")
	gInput = make([][]int, len(lines))
	for i, ls := range linesSplit {
		gInput[i] = iutils.ExtractInt1DFromString0D(ls[0], "-", -1)
		gInput[i] = append(gInput[i], iutils.ExtractInt1DFromString0D(ls[1], "-", -1)...)
	}
}
