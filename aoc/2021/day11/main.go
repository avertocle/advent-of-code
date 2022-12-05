package main

import (
	"fmt"
	"github.com/avertocle/contests/io/intz"
	input2 "github.com/avertocle/contests/io/iutils"
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
		log.Fatalf("iutils error | %v", err)
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
	flashCount := 0
	for i := 0; i < simCount; i++ {
		flashCount += iterate()
		fmt.Printf("%v.", i)
	}
	return flashCount
}

func problem2() int {
	for i := 0; ; i++ {
		fmt.Printf("%v.", i)
		iterate()
		if AreAllZero() {
			return i + 1
		}
	}
}

func AreAllZero() bool {
	sum := 0
	for _, row := range in.grid {
		for _, ele := range row {
			sum += ele
		}
	}
	return sum == 0
}

func iterate() int {
	step1()
	flashCount := step2()
	//iutils.PrettyArray2DInt(in.grid)
	for x := flashCount; x > 0; {
		x = step2()
		//iutils.PrettyArray2DInt(in.grid)
		flashCount += x
	}
	step3()
	//iutils.PrettyArray2DInt(in.grid)
	return flashCount
}

func step1() {
	for i, row := range in.grid {
		for j, _ := range row {
			in.grid[i][j]++
		}
	}
}

func step2() int {
	flashCount := 0
	for i, row := range in.grid {
		for j, ele := range row {
			if ele > 9 {
				flashCount++
				in.grid[i][j] = -1 * (in.grid[i][j])
				intz.ApplyToAdjacent(in.grid, i, j, in.rows, in.cols, true, func(b int) int { return b + 1 })
			}
		}
	}
	return flashCount
}

func step3() {
	for i, row := range in.grid {
		for j, ele := range row {
			if ele > 9 || ele < 0 {
				in.grid[i][j] = 0
			}
		}
	}
}
