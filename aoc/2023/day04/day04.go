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

const DirPath = "../2023/day04"

func SolveP1() string {
	ans := 0
	for i, _ := range gInput {
		cwNumsCount := findCWNumCount(i)
		ans += calcPointsP1(cwNumsCount)
	}

	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	cwNumsCounts := intz.Init1D(len(gInput), 0)
	for i, _ := range gInput {
		cwNumsCounts[i] = findCWNumCount(i)
	}
	cardArr := intz.Init1D(len(gInput), 0)
	for i, _ := range gInput {
		processP2(i, cardArr, cwNumsCounts)
	}
	ans := intz.Sum1D(cardArr)
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func findCWNumCount(cardIdx int) int {
	cwNums := intz.Intersect1D(gInput[cardIdx][0], gInput[cardIdx][1])
	cwNumCount := intz.Reduce1D(cwNums, 0, func(ans int, arr []int, i int) int {
		if arr[i] != -1 {
			return ans + 1
		} else {
			return ans
		}
	})
	return cwNumCount
}

/***** P1 Functions *****/

func calcPointsP1(cwNumsCount int) int {
	if cwNumsCount == 0 {
		return 0
	}
	return int(math.Pow(2, float64(cwNumsCount-1)))
}

/***** P2 Functions *****/

// 1 2 4 8 14 1
// 1 4 8 16 0 0
func processP2(idx int, ans []int, cwNumsCounts []int) {
	ans[idx] += 1
	for i := 1; i <= cwNumsCounts[idx]; i++ {
		processP2(idx+i, ans, cwNumsCounts)
	}
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = make([][][]int, len(lines))
	for i, line := range lines {
		b, c := parseLine(line)
		gInput[i] = [][]int{b, c}
	}
}

func parseLine(line string) ([]int, []int) {
	tokens := strings.Split(strings.Split(line, ": ")[1], " | ")
	cNums := iutils.ExtractInt1DFromString0D(tokens[0], " ", -1)
	wNums := iutils.ExtractInt1DFromString0D(tokens[1], " ", -1)
	return cNums, wNums
}
