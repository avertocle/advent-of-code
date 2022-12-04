package main

import (
	"fmt"
	"github.com/avertocle/contests/io"
	"github.com/avertocle/contests/io/ds"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/metrics"
	"log"
)

const inputFilePath = "input_small.txt"

func main() {
	metrics.ProgStart()
	input := getInputOrDie()

	input.PrintAdList()

	input.PrintAdMat()

	ans01 := solvePartOne(input)
	fmt.Printf("answer part-01 = %v\n", ans01)

	ans02 := solvePartTwo(input)
	fmt.Printf("answer part-02 = %v\n", ans02)
}

func solvePartOne(input *ds.Graph) int {
	return 0
}

func solvePartTwo(input *ds.Graph) int {
	return 0
}

/***** Common Functions *****/

/***** Input *****/

// input : [][]int : each row contains start and end ranges of both elves {e1s,e1e,e2s,e2e}
func getInputOrDie() *ds.Graph {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	linesSplit := iutils.ExtractString2DFromString1D(lines, "-", nil, "")
	io.PrettyArray2DString(linesSplit)
	input := ds.NewGraph()
	for _, ls := range linesSplit {
		input.AddVertex(ls[0], arrToMap(ls[1]))
		input.AddVertex(ls[1], arrToMap(ls[0]))
	}
	return input
}

func arrToMap(arr ...string) map[string]int {
	m := make(map[string]int)
	for _, a := range arr {
		m[a] = 1
	}
	return m
}

/***** PART 01 Functions *****/

/***** PART 02 Functions *****/
