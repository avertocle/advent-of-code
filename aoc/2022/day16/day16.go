package day16

import (
	"fmt"
	"github.com/avertocle/contests/io/ds"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"math"
)

var gInput *ds.Graph

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

	gInput = ds.NewGraph()
	var tokens []string
	var rate int
	var v string
	var adjMap map[string]int
	for i := 0; i < len(lines); i++ {
		tokens = stringz.SplitMulti(lines[i], []string{" ", "=", ";", ","})
		rate = stringz.AtoiQ(tokens[5], math.MinInt)
		v = tokens[1]
		adjMap = make(map[string]int)
		for j := 11; j < len(tokens); j += 2 {
			adjMap[tokens[j]] = 1
		}
		gInput.AddVertex(v, rate, adjMap)
	}
	gInput.PrintAdList()
}
