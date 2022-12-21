package day15

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/outils"
	"github.com/avertocle/contests/io/stringz"
	"math"
)

var gInpSen [][]int
var gInpBea [][]int
var gBoundTL []int
var gBoundBR []int

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

	gInpSen = make([][]int, len(lines))
	gInpBea = make([][]int, len(lines))
	maxX, maxY := math.MinInt, math.MinInt
	minX, minY := math.MaxInt, math.MaxInt
	px, py := 0, 0
	var tokens []string
	for i := 0; i < len(lines); i++ {
		tokens = stringz.SplitMulti(lines[i], []string{" ", "=", ":", ","})
		px = stringz.AtoiQ(tokens[3], math.MinInt)
		py = stringz.AtoiQ(tokens[6], math.MinInt)
		maxX, maxY = intz.Max(maxX, px), intz.Max(maxY, py)
		minX, minY = intz.Min(minX, px), intz.Min(minY, py)
		gInpSen[i] = []int{px, py}

		px = stringz.AtoiQ(tokens[13], math.MinInt)
		py = stringz.AtoiQ(tokens[16], math.MinInt)
		maxX, maxY = intz.Max(maxX, px), intz.Max(maxY, py)
		minX, minY = intz.Min(minX, px), intz.Min(minY, py)
		gInpBea[i] = []int{px, py}
	}
	gBoundTL = []int{minX, 0} // sand dropping from 500,0
	gBoundBR = []int{maxX, maxY}

	outils.PrettyArray2DInt(gInpSen)
	outils.PrettyArray2DInt(gInpBea)
	fmt.Printf("\n\n bounds : tl(%v,%v) br(%v,%v) \n\n", minX, minY, maxX, maxY)
}
