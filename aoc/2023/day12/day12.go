package day12

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"strings"
)

var gInputPatterns []string
var gInputCounts [][]int

const DirPath = "../2023/day12"

func SolveP1() string {
	ans := "0"
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
	gInputPatterns = make([]string, len(lines))
	gInputCounts = make([][]int, len(lines))
	for i, line := range lines {
		tokens := strings.Fields(line)
		gInputPatterns[i] = tokens[0]
		gInputCounts[i] = iutils.ExtractInt1DFromString0D(tokens[1], ",", -1)
	}
	fmt.Println(gInputPatterns)
	fmt.Println(gInputCounts)
}
