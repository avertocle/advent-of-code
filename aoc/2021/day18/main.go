package main

import (
	"fmt"
	input2 "github.com/avertocle/contests/io/input"
	"log"

	"github.com/avertocle/contests/io"
	"github.com/avertocle/contests/metrics"
)

const inputFilePath = "input.txt"

var in *input

type input struct {
	grid [][]int
	rows int
	cols int
}

func main() {
	metrics.ProgStart()
	in = getInputOrDie()
	metrics.InputLen(in.rows)

	ans := problem1()
	//ans := problem2()
	fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func getInputOrDie() *input {
	lines, err := input2.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}
	grid := io.String1DToInt2D(lines, "")
	return &input{
		grid: grid,
		rows: len(grid),
		cols: len(grid[0]),
	}
}

/***** Logic Begins here *****/

const simCount = 100

func problem1() int {
	return 0
}
