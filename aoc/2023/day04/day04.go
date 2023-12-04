package day04

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"math"
	"strings"
)

var gInput [][][]int

func SolveP1() string {
	ans := 0
	for _, card := range gInput {
		cwNums := intz.Intersect1D(card[0], card[1])
		cwNums = intz.Filter1D(cwNums, func(arr []int, i int) bool {
			return arr[i] != -1
		})
		//fmt.Println(calcPointsP1(cwNums), cwNums)
		ans += calcPointsP1(cwNums)
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func calcPointsP1(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return int(math.Pow(2, float64(len(arr)-1)))
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = make([][][]int, len(lines))
	for i, line := range lines {
		b, c := parseLine(line)
		gInput[i] = [][]int{b, c}
		fmt.Println(b, c)
	}
}

func parseLine(line string) ([]int, []int) {
	tokens := strings.Split(strings.Split(line, ": ")[1], " | ")
	cNums := iutils.ExtractInt1DFromString0D(tokens[0], " ", -1)
	wNums := iutils.ExtractInt1DFromString0D(tokens[1], " ", -1)
	return cNums, wNums
}
