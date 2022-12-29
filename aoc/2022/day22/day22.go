package day22

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
)

var gInpMap [][]byte
var gInpRows, gInpCols int
var gInpPath [][]int
var gInpRowBounds [][]int
var gInpColBounds [][]int
var debugMap [][]byte

const cellEmpty = '.'
const cellWall = '#'
const cellGap = ' '

const (
	right = 0
	down  = 1
	left  = 2
	up    = 3
)

func SolveP1() string {
	debugMap = bytez.Extract2D(gInpMap, []int{0, 0}, []int{gInpRows - 1, gInpCols - 1}, '*')
	ci, cj, cf := 1, 9, right
	for _, step := range gInpPath {
		ci, cj, cf = walkOneStep(ci, cj, cf, step[0], step[1])
	}
	//bytez.PPrint2D(debugMap)
	ans := 1000*ci + 4*cj + cf
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func walkOneStep(ci, cj, cf, dis, turn int) (int, int, int) {
	//fmt.Printf("step (%v,%v) @ %v dis(%v) turn(%v) => ", ci, cj, cf, dis, turn)
	switch cf {
	case right:
		for d := 0; d < dis; d++ {
			ci, cj = walkOneStepRight(ci, cj)
			debugMap[ci][cj] = ftoc(cf)
		}
		cf = calcNewFacing(cf, turn)
		debugMap[ci][cj] = ftoc(cf)
	case down:
		for d := 0; d < dis; d++ {
			ci, cj = walkOneStepDown(ci, cj)
			debugMap[ci][cj] = ftoc(cf)
		}
		cf = calcNewFacing(cf, turn)
		debugMap[ci][cj] = ftoc(cf)
	case left:
		for d := 0; d < dis; d++ {
			ci, cj = walkOneStepLeft(ci, cj)
			debugMap[ci][cj] = ftoc(cf)
		}
		cf = calcNewFacing(cf, turn)
		debugMap[ci][cj] = ftoc(cf)
	case up:
		for d := 0; d < dis; d++ {
			ci, cj = walkOneStepUp(ci, cj)
			debugMap[ci][cj] = ftoc(cf)
		}
		cf = calcNewFacing(cf, turn)
		debugMap[ci][cj] = ftoc(cf)
	default:
		errz.HardAssert(false, "walkOneStep : invalid curr-facing %v", cf)
		break
	}
	//fmt.Printf("step (%v,%v) @ %v\n", ci, cj, cf)
	return ci, cj, cf
}

func walkOneStepLeft(ci, cj int) (int, int) {
	cin, cjn := ci, cj-1
	if gInpMap[cin][cjn] == cellWall {
		return ci, cj
	} else if gInpMap[cin][cjn] == cellGap {
		cjn = gInpRowBounds[ci][1]
		if gInpMap[cin][cjn] == cellWall {
			return ci, cj
		} else {
			return cin, cjn
		}
	} else {
		return cin, cjn
	}
}

func walkOneStepRight(ci, cj int) (int, int) {
	cin, cjn := ci, cj+1
	if gInpMap[cin][cjn] == cellWall {
		return ci, cj
	} else if gInpMap[cin][cjn] == cellGap {
		cjn = gInpRowBounds[ci][0]
		if gInpMap[cin][cjn] == cellWall {
			return ci, cj
		} else {
			return cin, cjn
		}
	} else {
		return cin, cjn
	}
}

func walkOneStepDown(ci, cj int) (int, int) {
	cin, cjn := ci+1, cj
	if gInpMap[cin][cjn] == cellWall {
		return ci, cj
	} else if gInpMap[cin][cjn] == cellGap {
		cin = gInpColBounds[cj][0]
		if gInpMap[cin][cjn] == cellWall {
			return ci, cj
		} else {
			return cin, cjn
		}
	} else {
		return cin, cjn
	}
}

func walkOneStepUp(ci, cj int) (int, int) {
	cin, cjn := ci-1, cj
	if gInpMap[cin][cjn] == cellWall {
		return ci, cj
	} else if gInpMap[cin][cjn] == cellGap {
		cin = gInpColBounds[cj][1]
		if gInpMap[cin][cjn] == cellWall {
			return ci, cj
		} else {
			return cin, cjn
		}
	} else {
		return cin, cjn
	}
}

func calcNewFacing(cf, turn int) int {
	if turn == left {
		cf--
		if cf == -1 {
			cf = 3
		}
	} else if turn == right {
		cf++
		if cf == 4 {
			cf = 0
		}
	} else {
		fmt.Println("wierd turn", turn)
		return cf
	}
	return cf
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	var pathIdx int
	gInpRows, gInpCols, pathIdx = parseDims(lines)
	gInpRows += 2
	gInpCols += 2
	gInpMap = bytez.Init2D(gInpRows, gInpCols, cellGap)
	for i := 1; i < gInpRows-1; i++ {
		for j := 1; j < gInpCols-1; j++ {
			if j <= len(lines[i-1]) {
				gInpMap[i][j] = lines[i-1][j-1]
			}
		}
	}

	parseBounds()
	parsePath(lines[pathIdx])
	//fmt.Println(gInpRows, gInpCols)
	//bytez.PPrint2D(gInpMap)
	//intz.PPrint2D(gInpRowBounds)
	//intz.PPrint2D(gInpColBounds)
}

func parseBounds() {
	gInpRowBounds = make([][]int, gInpRows)
	gInpColBounds = make([][]int, gInpCols)
	for i := 0; i < gInpRows; i++ {
		j := 0
		for ; j < gInpCols; j++ {
			if gInpMap[i][j] != cellGap {
				gInpRowBounds[i] = []int{j, -1}
				break
			}
		}
		for j++; j < gInpCols; j++ {
			if gInpMap[i][j] == cellGap {
				gInpRowBounds[i][1] = j - 1
				break
			}
		}
	}

	for j := 0; j < gInpCols; j++ {
		i := 0
		for ; i < gInpRows; i++ {
			if gInpMap[i][j] != cellGap {
				gInpColBounds[j] = []int{i, -1}
				break
			}
		}
		for i++; i < gInpRows; i++ {
			if gInpMap[i][j] == cellGap {
				gInpColBounds[j][1] = i - 1
				break
			}
		}
	}
}

func parseDims(lines []string) (int, int, int) {
	rows := 0
	cols := 0
	for _, l := range lines {
		if len(l) > cols {
			cols = len(l)
		}
		if len(l) == 0 {
			break
		}
		rows++
	}
	return rows, cols, rows + 1
}

func parsePath(line string) {
	gInpPath = make([][]int, 0)
	var step []int
	path := []byte(line)
	for i := 0; i < len(path); {
		step = []int{-1, -1}
		j := i
		for ; j < len(path) && path[j] != 'L' && path[j] != 'R'; j++ {
		}
		step[0] = stringz.AtoiQ(string(path[i:j]), -1)
		if j < len(path) {
			if path[j] == 'L' {
				step[1] = 2
			} else if path[j] == 'R' {
				step[1] = 0
			}
		}
		gInpPath = append(gInpPath, step)
		i = j + 1
	}
}

func ftoc(f int) byte {
	switch f {
	case right:
		return '>'
	case down:
		return 'v'
	case left:
		return '<'
	case up:
		return '^'
	}
	return '$'
}
