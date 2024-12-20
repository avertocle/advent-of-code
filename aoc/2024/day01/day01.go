package day01

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/numz"
	"slices"
)

const DirPath = "../2024/day01"

var gInput [][]int

func SolveP1() string {
	ans := 0
	sorted1 := make([]int, len(gInput))
	sorted2 := make([]int, len(gInput))
	for i := 0; i < len(gInput); i++ {
		sorted1[i] = gInput[i][0]
		sorted2[i] = gInput[i][1]
	}
	slices.Sort(sorted1)
	slices.Sort(sorted2)
	for i := 0; i < len(sorted1); i++ {
		ans += numz.Abs(sorted1[i] - sorted2[i])
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	count := 0
	for i := 0; i < len(gInput); i++ {
		for j := 0; j < len(gInput); j++ {
			if gInput[i][0] == gInput[j][1] {
				count += 1
			}
			ans += gInput[i][0] * count
			count = 0
		}
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractInt2DFromString1D(lines, "  ", nil, -1)
	//intz.PPrint2D(gInput)
}
