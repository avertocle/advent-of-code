package day08

import (
	"fmt"
	"github.com/avertocle/contests/io/boolz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/numz"
	"log"
)

var input [][]int // square input
var inpSize int

func SolveP1() string {
	vis := intz.Init2D(inpSize, inpSize, 0) // only to debug
	visCount := 0
	isVis := boolz.Init1D(4, true)
	for i, row := range input {
		for j, _ := range row {
			isVis[0], _ = checkLeft(i, j)
			isVis[1], _ = checkRight(i, j)
			isVis[2], _ = checkUp(i, j)
			isVis[3], _ = checkDown(i, j)
			if boolz.Or1D(isVis) {
				vis[i][j] = 1
				visCount++
			}
		}
	}
	return fmt.Sprintf("%v", visCount)
}

func SolveP2() string {
	maxScore := 0
	vdis := intz.Init1D(4, 0)
	score := 0
	for i, row := range input {
		for j, _ := range row {
			_, vdis[0] = checkLeft(i, j)
			_, vdis[1] = checkRight(i, j)
			_, vdis[2] = checkUp(i, j)
			_, vdis[3] = checkDown(i, j)
			score = intz.Mul1D(vdis)
			maxScore = numz.Max(score, maxScore)
		}
	}
	return fmt.Sprintf("%v", maxScore)
}

/***** Common Functions *****/

// returns isVis, vdis
func checkLeft(i, j int) (bool, int) {
	isVis := true
	vdis := 0
	for x := j - 1; x >= 0; x-- {
		vdis++
		if input[i][x] >= input[i][j] {
			isVis = false
			break
		}
	}
	return isVis, vdis
}

func checkRight(i, j int) (bool, int) {
	isVis := true
	vdis := 0
	for x := j + 1; x < inpSize; x++ {
		vdis++
		if input[i][x] >= input[i][j] {
			isVis = false
			break
		}
	}
	return isVis, vdis
}

func checkUp(i, j int) (bool, int) {
	isVis := true
	vdis := 0
	for x := i - 1; x >= 0; x-- {
		vdis++
		if input[x][j] >= input[i][j] {
			isVis = false
			break
		}
	}
	return isVis, vdis
}

func checkDown(i, j int) (bool, int) {
	isVis := true
	vdis := 0
	for x := i + 1; x < inpSize; x++ {
		vdis++
		if input[x][j] >= input[i][j] {
			isVis = false
			break
		}
	}
	return isVis, vdis
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	input = iutils.ExtractInt2DFromString1D(lines, "", nil, 0)
	inpSize = len(input)
}
