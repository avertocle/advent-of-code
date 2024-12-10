package day10

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/cmz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

var gInput [][]byte
var gStartShape byte
var sStartIdx *arrz.Idx2D
var sStartNextIdx *arrz.Idx2D

const DirPath = "../2023/day10"

func SolveP1() string {
	start, startNext := sStartIdx, sStartNextIdx
	loop := findLoop(start, startNext, gInput)
	ans := (len(loop) + 1) / 2
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	loopMarker, expandMarker, fillMarker := byte('#'), byte('!'), byte(' ')
	grid, start := makeExpandedGrid(expandMarker)
	startNext := arrz.NewIdx2D(start.I+1, start.J) // set manually or fix getNextPos to handle no-prev case
	loop := findLoop(start, startNext, grid)
	floodFill(grid, arrz.NewIdx2D(0, 0), loop, fillMarker)
	markLoopOnGrid(grid, loop, loopMarker)
	for i := 0; i < len(grid); i += 2 {
		for j := 0; j < len(grid[0]); j += 2 {
			if grid[i][j] != loopMarker && grid[i][j] != fillMarker && grid[i][j] != expandMarker {
				ans++
			}
		}
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

func makeExpandedGrid(marker byte) ([][]byte, *arrz.Idx2D) {
	// hor : - F 7
	// ver : | L J
	grid := make([][]byte, len(gInput)*2)
	for i := 0; i < len(gInput); i++ {
		grid[2*i] = make([]byte, len(gInput[0])*2)
		for j := 0; j < len(gInput[0]); j++ {
			grid[2*i][2*j] = gInput[i][j]
			if gInput[i][j] == '-' || gInput[i][j] == 'F' || gInput[i][j] == 'L' {
				grid[2*i][2*j+1] = '-'
			} else {
				grid[2*i][2*j+1] = marker
			}
		}
		grid[2*i+1] = make([]byte, len(gInput[0])*2)
		for j := 0; j < len(gInput[0]); j++ {
			if gInput[i][j] == '|' || gInput[i][j] == 'F' || gInput[i][j] == '7' {
				grid[2*i+1][2*j] = '|'
			} else {
				grid[2*i+1][2*j] = marker
			}
			grid[2*i+1][2*j+1] = marker
		}
	}
	grid = bytez.Pad2D(grid, len(grid), len(grid[0]), 2, marker)
	newStart := arrz.NewIdx2D(sStartIdx.I*2+2, sStartIdx.J*2+2)

	return grid, newStart
}

func floodFill(grid [][]byte, curr *arrz.Idx2D, loop cmz.MapVisited, marker byte) {
	grid[curr.I][curr.J] = marker
	neighbours := curr.Neighbours(false)
	for _, n := range neighbours {
		if n.IsInBounds(len(grid), len(grid[0])) && !loop[n.ToKey()] && grid[n.I][n.J] != marker {
			floodFill(grid, n, loop, marker)
		}
	}
}

func markLoopOnGrid(grid [][]byte, loop cmz.MapVisited, marker byte) {
	for p, _ := range loop {
		pp := arrz.NewIdx2DFromKey(p)
		grid[pp.I][pp.J] = marker
	}
}

/***** Common Functions *****/

func findLoop(start, startNext *arrz.Idx2D, grid [][]byte) cmz.MapVisited {
	prev, curr := start.Clone(), startNext.Clone()
	loop := make(cmz.MapVisited)
	loop[prev.ToKey()] = true
	for !curr.IsEqual(start) {
		loop[curr.ToKey()] = true
		shape := grid[curr.I][curr.J]
		next := arrz.NewIdx2D(getNextPos(prev, curr, shape))
		prev = curr
		curr = next
	}
	return loop
}

func getNextPos(prev, next *arrz.Idx2D, shape byte) (int, int) {
	pi, pj := prev.I, prev.J
	ci, cj := next.I, next.J
	switch shape {
	case '|':
		if pi == ci-1 {
			return ci + 1, cj
		} else {
			return ci - 1, cj
		}
	case '-':
		if pj == cj-1 {
			return ci, cj + 1
		} else {
			return ci, cj - 1
		}
	case 'L':
		if pi == ci-1 {
			return ci, cj + 1
		} else {
			return ci - 1, cj
		}
	case 'J':
		if pi == ci-1 {
			return ci, cj - 1
		} else {
			return ci - 1, cj
		}
	case '7':
		if pi == ci+1 {
			return ci, cj - 1
		} else {
			return ci + 1, cj
		}
	case 'F':
		if pi == ci+1 {
			return ci, cj + 1
		} else {
			return ci + 1, cj
		}
	default:
		errz.HardAssert(false, "invalid shape | %v", string(shape))
	}
	errz.HardAssert(false, "invalid shape or coordinates | (%v,%v : %v, %v : %v)",
		pi, pj, ci, cj, shape)

	return -1, -1
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gStartShape = lines[0][0]
	temp := iutils.ExtractInt1DFromString0D(lines[1], ",", -1)
	sStartIdx = arrz.NewIdx2D(temp[0], temp[1])
	temp = iutils.ExtractInt1DFromString0D(lines[2], ",", -1)
	sStartNextIdx = arrz.NewIdx2D(temp[0], temp[1])
	gInput = iutils.ExtractByte2DFromString1D(lines[3:], "", nil, 0)
	gInput[sStartIdx.I][sStartIdx.J] = gStartShape
}
