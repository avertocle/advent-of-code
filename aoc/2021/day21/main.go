package main

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/metrics"
	"log"
)

const inputFilePath = "input.txt"

func main() {
	metrics.ProgStart()
	input := getInputOrDie()

	ans01 := solvePartOne(input)
	fmt.Printf("answer part-01 = %v\n", ans01)

	ans02 := solvePartTwo(input)
	fmt.Printf("answer part-02 = %v\n", ans02)
}

func solvePartOne(input [][]byte) int {
	return 0
}

func solvePartTwo(input [][]byte) int {
	return 0
}

/***** Common Functions *****/

/***** PART 01 Functions *****/

/***** PART 02 Functions *****/

/***** Input *****/

// input : [][]byte : each row contains a move pair
func getInputOrDie() [][]byte {
	_, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	return nil
}
