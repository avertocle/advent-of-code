package day22

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

var gInpMap [][]byte
var gInpRows, gInpCols int
var gInpPath string

func SolveP1() string {
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
	i := 0
	cols := 0
	for _, l := range lines {
		if len(l) > cols {
			cols = len(l)
		}
		if len(l) == 0 {
			break
		}
		i++
	}
	gInpMap = iutils.ExtractByte2DFromString1D(lines[0:i], "", nil, '*')
	gInpRows = i
	gInpCols = cols
	gInpMap = bytez.Pad2D(gInpMap, gInpRows, gInpCols, 0, ' ')
	gInpPath = lines[i+1]
	//outils.PrettyArray2DByte(gInpMap)
	fmt.Println(gInpRows, gInpCols)
	//fmt.Println(gInpPath)
}
