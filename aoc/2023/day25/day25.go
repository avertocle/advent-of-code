package day25

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
)

var gInput [][]string

const DirPath = "../2023/day25"

func SolveP1() string {
	ans := 0
	compMap := make(map[string]int)
	for _, comps := range gInput {
		for _, comp := range comps {
			if _, ok := compMap[comp]; !ok {
				compMap[comp] = 0
			}
			compMap[comp]++
		}
	}
	adjMat := make([][]int, len(compMap))
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	for _, line := range lines {
		gInput = append(gInput, stringz.SplitMultiTrimSpace(line, []string{" ", ":"}))
	}
	arrz.PPrint2D(gInput)
}
