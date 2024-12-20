package day06

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"strings"
)

var gInputTimeP1 []int
var gInputDistP1 []int
var gInputTimeP2 int
var gInputDistP2 int

const DirPath = "../2023/day06"

func SolveP1() string {
	ans := 1
	for i := 0; i < len(gInputTimeP1); i++ {
		winCtr := countWinScenarios(gInputTimeP1[i], gInputDistP1[i])
		ans *= winCtr
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := countWinScenarios(gInputTimeP2, gInputDistP2)
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

/***** P1 Functions *****/

func countWinScenarios(tf, df int) int {
	winCtr := 0
	for t := 1; t <= tf; t++ {
		if canWinRace(t, tf, df) {
			winCtr++
		}
	}
	return winCtr
}

func canWinRace(t0, tf, df int) bool {
	// d = t0 * (tf -t0)
	return t0*(tf-t0) > df
}

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInputTimeP1 = iutils.ExtractInt1DFromString1D(
		stringz.SplitMultiTrimSpace(lines[0], []string{":", " "})[1:], "", -1, -1)
	gInputDistP1 = iutils.ExtractInt1DFromString1D(
		stringz.SplitMultiTrimSpace(lines[1], []string{":", " "})[1:], "", -1, -1)

	gInputTimeP2 = stringz.AtoI(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""), -1)
	gInputDistP2 = stringz.AtoI(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""), -1)

}
