package day10

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

var gInput [][]byte

const DirPath = "../2023/day10"

func SolveP1() string {
	start := bytez.Find2D(gInput, 'S')[0]
	pi, pj := start[0], start[1]
	ci, cj := start[0]+1, start[1] // choose manually by eyeballing / diff for input
	pathLen := 0
	for ci != start[0] || cj != start[1] {
		shape := gInput[ci][cj]
		//fmt.Printf("%v.", string(shape))
		ni, nj := getNextPos(pi, pj, ci, cj, shape)
		pi, pj = ci, cj
		ci, cj = ni, nj
		pathLen++
	}
	ans := (pathLen + 1) / 2
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	visualMap := bytez.Copy2D(gInput)
	start := bytez.Find2D(gInput, 'S')[0]
	pi, pj := start[0], start[1]
	ci, cj := start[0]+1, start[1] // choose manually by eyeballing / diff for input
	pathLen := 0
	for ci != start[0] || cj != start[1] {
		visualMap[ci][cj] = 'x'
		shape := gInput[ci][cj]
		//fmt.Printf("%v.", string(shape))
		ni, nj := getNextPos(pi, pj, ci, cj, shape)
		pi, pj = ci, cj
		ci, cj = ni, nj
		pathLen++
	}
	ans := (pathLen + 1) / 2
	for i, row := range visualMap {
		for j, cell := range row {
			if cell != 'x' && cell != '.' {
				visualMap[i][j] = ' '
			}
		}
	}
	bytez.PPrint2D(visualMap)
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

func getNextPos(pi, pj, ci, cj int, shape byte) (int, int) {
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
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
}
