package day18

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"math"
)

var gInput [][]int

func SolveP1() string {
	openSides := 0
	for i, _ := range gInput {
		openSides += calcOpenSides(i)
	}
	ans := openSides
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func calcOpenSides(idx int) int {
	currCube := gInput[idx]
	coveredCtr := 0
	for i, tempCube := range gInput {
		if i == idx {
			continue
		}
		if isTouching(currCube, tempCube) {
			coveredCtr++
		}
	}
	fmt.Printf("%v : %v = %v\n", idx, currCube, coveredCtr)
	return 6 - coveredCtr
}

func isTouching(c1, c2 []int) bool {
	if c1[0] == c2[0] && c1[1] == c2[1] && intz.Abs(c1[2]-c2[2]) == 1 {
		return true
	} else if c1[1] == c2[1] && c1[2] == c2[2] && intz.Abs(c1[0]-c2[0]) == 1 {
		return true
	} else if c1[2] == c2[2] && c1[0] == c2[0] && intz.Abs(c1[1]-c2[1]) == 1 {
		return true
	} else {
		return false
	}
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractInt2DFromString1D(lines, ",", nil, math.MinInt)
}
