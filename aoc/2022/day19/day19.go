package day19

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
)

var gInput [][][]int

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
	gInput = make([][][]int, len(lines))
	var t []string
	for i, l := range lines {
		t = stringz.SplitMulti(l, []string{" ", ":", "."})
		gInput[i] = [][]int{
			{stringz.AtoiQ(t[7], -1), 0, 0, 0},
			{stringz.AtoiQ(t[14], -1), 0, 0, 0},
			{stringz.AtoiQ(t[21], -1), stringz.AtoiQ(t[24], -1), 0, 0},
			{stringz.AtoiQ(t[31], -1), 0, stringz.AtoiQ(t[34], -1), 0},
		}
	}
	intz.PPrint3D(gInput)
}

// ore, clay, obsidian, geode
