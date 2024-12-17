package day14

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

var gInput [][]byte

const DirPath = "../2023/day14"

func SolveP1() string {
	ans := 0
	grid := bytez.Copy2D(gInput)
	doNorthTilt(grid)
	ans = calcLoad(grid)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	resetPoint := 0
	history := make([][][]byte, 0)
	history = append(history, gInput)
	historyIndex := -1
	resetLen := -1
	initialCut := -1
	grid := bytez.Copy2D(gInput)
	for i := 0; ; i++ {
		//if i%10000000 == 0 {
		//	fmt.Printf("%v.", i)
		//}
		doOneCycle(grid)
		//arrz.PPrint2D(grid)
		if historyIndex = matchInHistory(grid, history); historyIndex != -1 {
			resetPoint = i + 1
			resetLen = resetPoint - historyIndex
			initialCut = historyIndex
			arrz.PPrint2D(grid)
			break
		} else {
			history = append(history, bytez.Copy2D(grid))
		}
	}
	remainingCycles := (1000000000 - initialCut) % (resetLen)
	fmt.Println("repeated : %v - %v = %v => %v \n", resetPoint, resetLen, initialCut, remainingCycles)
	for i := 0; i < remainingCycles; i++ {
		doOneCycle(grid)
	}
	ans = calcLoad(grid)
	return fmt.Sprintf("%v", ans)
}

func doOneCycle(grid [][]byte) {
	doNorthTilt(grid)
	doWestTilt(grid)
	doSouthTilt(grid)
	doEastTilt(grid)
}

func matchInHistory(grid [][]byte, history [][][]byte) int {
	for i, h := range history {
		if bytez.Compare2D(grid, h) == 0 {
			return i
		}
	}
	return -1
}

/***** Common Functions *****/

func calcLoad(grid [][]byte) int {
	load := 0
	for i := 0; i < len(grid); i++ {
		//fmt.Println(i, bytez.Count1D(grid[i], 'O'))
		load += (len(grid) - i) * bytez.Count1D(grid[i], 'O')
	}
	return load
}

func doNorthTilt(grid [][]byte) {
	for j := 0; j < len(grid[0]); j++ {
		x := -1
		for i := 0; i < len(grid); i++ {
			if grid[i][j] == '#' {
				x = i
			} else if grid[i][j] == 'O' {
				x++
				if x != i {
					grid[x][j] = 'O'
					grid[i][j] = '.'
				}
			}
		}
	}
}

func doSouthTilt(grid [][]byte) {
	for j := 0; j < len(grid[0]); j++ {
		x := len(grid)
		for i := len(grid) - 1; i >= 0; i-- {
			if grid[i][j] == '#' {
				x = i
			} else if grid[i][j] == 'O' {
				x--
				if x != i {
					grid[x][j] = 'O'
					grid[i][j] = '.'
				}
			}
		}
	}
}

func doWestTilt(grid [][]byte) {
	for ii := 0; ii < len(grid); ii++ {
		x := -1
		for jj := 0; jj < len(grid[0]); jj++ {
			if grid[ii][jj] == '#' {
				x = jj
			} else if grid[ii][jj] == 'O' {
				x++
				if x != jj {
					grid[ii][x] = 'O'
					grid[ii][jj] = '.'
				}
			}
		}
	}
}

func doEastTilt(grid [][]byte) {
	for ii := 0; ii < len(grid); ii++ {
		x := len(grid)
		for jj := len(grid[0]) - 1; jj >= 0; jj-- {
			if grid[ii][jj] == '#' {
				x = jj
			} else if grid[ii][jj] == 'O' {
				x--
				if x != jj {
					grid[ii][x] = 'O'
					grid[ii][jj] = '.'
				}
			}
		}
	}
}

/***** P1 Functions *****/

func calcColumnLoad(grid [][]byte, col int) int {
	load := 0

	// two pointers, bitch !!
	x := -1
	for i := 0; i < len(grid); i++ {
		if grid[i][col] == '#' {
			x = i
		} else if grid[i][col] == 'O' {
			x++
			load += len(grid) - x
		} else if grid[i][col] == '.' {

		}
	}
	return load
}

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
}
