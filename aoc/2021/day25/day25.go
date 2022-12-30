package day25

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

var gInput *input

func SolveP1() string {
	movesDone := 0
	for i := 0; true; i++ {
		movesDone = 0
		movesDone += runSim(moveHori)
		movesDone += runSim(moveVert)
		if movesDone == 0 {
			movesDone = i + 1
			break
		}
	}
	ans := movesDone
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func runSim(move func(int, int) (bool, int, int)) int {
	gridNew := bytez.Init2D(gInput.rows, gInput.cols, 0)
	movesDone := 0
	for i := 0; i < gInput.rows; i++ {
		for j := 0; j < gInput.cols; j++ {
			if ok, i1, j1 := move(i, j); ok {
				gridNew[i1][j1] = gInput.grid[i][j]
				gridNew[i][j] = '.'
				movesDone++
			} else if gridNew[i][j] == 0 {
				gridNew[i][j] = gInput.grid[i][j]
			}
		}
	}
	gInput.grid = gridNew
	return movesDone
}

func moveHori(i, j int) (bool, int, int) {
	if gInput.grid[i][j] == '>' {
		if j+1 < gInput.cols && gInput.grid[i][j+1] == '.' {
			return true, i, j + 1
		} else if j+1 == gInput.cols && gInput.grid[i][0] == '.' {
			return true, i, 0
		} else {
			return false, -1, -1
		}
	} else {
		return false, -1, -1
	}
}

func moveVert(i, j int) (bool, int, int) {
	if gInput.grid[i][j] == 'v' {
		if i+1 < gInput.rows && gInput.grid[i+1][j] == '.' {
			return true, i + 1, j
		} else if i+1 == gInput.rows && gInput.grid[0][j] == '.' {
			return true, 0, j
		} else {
			return false, -1, -1
		}
	} else {
		return false, -1, -1
	}
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	grid := iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
	gInput = &input{
		grid: grid,
		rows: len(grid),
		cols: len(grid[0]),
	}
}

type input struct {
	grid [][]byte
	rows int
	cols int
}
