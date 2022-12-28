package day24

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

var gInput [][]byte
var gInpRows int
var gInpCols int
var gInpS []int
var gInpE []int

func SolveP1() string {
	gInput[gInpS[0]][gInpS[1]] = 'S'
	gInput[gInpE[0]][gInpE[1]] = 'E'
	bytez.PPrint2D(gInput)
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
	gInpRows = len(gInput)
	gInpCols = len(gInput[0])
	gInpS = []int{0, 1}
	gInpE = []int{gInpRows - 1, gInpCols - 2}
}
