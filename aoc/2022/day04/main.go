package main

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/outils"
	"github.com/avertocle/contests/metrics"
	"log"
)

const inputFilePath = "input_small.txt"

func main() {
	metrics.ProgStart()
	input := getInputOrDie()

	outils.PrettyArray2DInt(input)

	ans01 := solvePartOne(input)
	fmt.Printf("answer part-01 = %v\n", ans01)

	ans02 := solvePartTwo(input)
	fmt.Printf("answer part-02 = %v\n", ans02)
}

func solvePartOne(input [][]int) int {
	count := 0
	for _, arr := range input {
		if arr[0] >= arr[2] && arr[1] <= arr[3] ||
			arr[2] >= arr[0] && arr[3] <= arr[1] {
			count++
		}
	}
	return count
}

func solvePartTwo(input [][]int) int {
	count := 0
	for _, arr := range input {
		if inRange(arr[0], arr[2], arr[3]) ||
			inRange(arr[1], arr[2], arr[3]) ||
			inRange(arr[2], arr[0], arr[1]) ||
			inRange(arr[3], arr[0], arr[1]) {
			count++
		}
	}
	return count
}

func inRange(x, rStart, rEnd int) bool {
	return x >= rStart && x <= rEnd
}

/***** Common Functions *****/

/***** Input *****/

// input : [][]int : each row contains start and end ranges of both elves {e1s,e1e,e2s,e2e}
func getInputOrDie() [][]int {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	linesSplit := iutils.ExtractString2DFromString1D(lines, ",", nil, "")
	input := make([][]int, len(lines))
	for i, ls := range linesSplit {
		input[i] = iutils.ExtractInt1DFromString0D(ls[0], "-", -1)
		input[i] = append(input[i], iutils.ExtractInt1DFromString0D(ls[1], "-", -1)...)
	}
	return input
}

/***** PART 01 Functions *****/

/***** PART 02 Functions *****/
