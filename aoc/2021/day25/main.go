package main

import (
	"fmt"
	"log"

	"github.com/avertocle/contests/io"
	"github.com/avertocle/contests/metrics"
)

const inputFilePath = "input.txt"

var in *input

type input struct {
	grid [][]byte
	rows int
	cols int
}

func main() {
	metrics.ProgStart()
	in = getInputOrDie()
	metrics.InputLen(in.rows)

	ans := problem1()
	fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func getInputOrDie() *input {
	lines, err := io.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}
	fmt.Printf("input-len = %v\n", len(lines))
	grid := io.String1DToByte2D(lines)
	return &input{
		grid: grid,
		rows: len(grid),
		cols: len(grid[0]),
	}
}

/***** Logic Begins here *****/

const simCount = 40

func problem1() int {
	movesDone := 0
	for i := 0; true; i++ {
		movesDone = 0
		//io.PrettyArray2DByte(in.grid)
		movesDone += iterate1(moveHori1)
		movesDone += iterate1(moveVert1)
		if movesDone == 0 {
			return i + 1
		}
	}
	return 0
}

func iterate1(move func(int, int) (bool, int, int)) int {
	gridNew := io.Init2DByte(in.rows, in.cols, 0)
	movesDone := 0
	for i := 0; i < in.rows; i++ {
		for j := 0; j < in.cols; j++ {
			if ok, i1, j1 := move(i, j); ok {
				gridNew[i1][j1] = in.grid[i][j]
				gridNew[i][j] = '.'
				movesDone++
			} else if gridNew[i][j] == 0 {
				gridNew[i][j] = in.grid[i][j]
			}
		}
	}
	in.grid = gridNew
	return movesDone
}

func moveHori1(i, j int) (bool, int, int) {
	if in.grid[i][j] == '>' {
		if j+1 < in.cols && in.grid[i][j+1] == '.' {
			return true, i, j + 1
		} else if j+1 == in.cols && in.grid[i][0] == '.' {
			return true, i, 0
		} else {
			return false, -1, -1
		}
	} else {
		return false, -1, -1
	}
}

func moveVert1(i, j int) (bool, int, int) {
	if in.grid[i][j] == 'v' {
		if i+1 < in.rows && in.grid[i+1][j] == '.' {
			return true, i + 1, j
		} else if i+1 == in.rows && in.grid[0][j] == '.' {
			return true, 0, j
		} else {
			return false, -1, -1
		}
	} else {
		return false, -1, -1
	}
}

func problem2() int {
	return 0
}
