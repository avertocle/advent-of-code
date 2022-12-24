package day12

import (
	"fmt"
	"github.com/avertocle/contests/io/boolz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"math"
)

var gInput [][]byte
var gInpRows int
var gInpCols int
var gStart []int
var gEnd []int

func SolveP1() string {
	starts := [][]int{gStart}
	ans := findSPathFrom(starts)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	starts := findVerticesByValue('a')
	ans := findSPathFrom(starts)
	return fmt.Sprintf("%v", ans)
}

func findSPathFrom(starts [][]int) int {
	var spaths [][]int
	var visited [][]bool
	fmt.Printf("total starting points : %v\n", len(starts))
	p, minP := 0, math.MaxInt/2
	for _, sv := range starts {
		spaths = intz.Init2D(gInpRows, gInpCols, math.MaxInt/2)
		visited = boolz.Init2D(gInpRows, gInpCols, false)
		spaths[sv[0]][sv[1]] = 0
		findSPath(sv, gEnd, spaths, visited)
		p = spaths[gEnd[0]][gEnd[1]]
		minP = intz.Min(minP, p)
		//fmt.Printf("%03d : s = (%v,%v) path = %v\n", i, sv[0], sv[1], p)
	}
	return minP
}

/***** Common Functions *****/

func findSPath(sv, ev []int, spaths [][]int, visited [][]bool) {
	if isEqual(sv, ev) {
		return
	}
	visited[sv[0]][sv[1]] = true
	nbrs := getVisitableNbrs(sv)
	for _, n := range nbrs {
		if !visited[n[0]][n[1]] {
			spaths[n[0]][n[1]] = intz.Min(spaths[n[0]][n[1]], 1+spaths[sv[0]][sv[1]])
		}
	}
	sv = findClosestVertex(spaths, visited)
	findSPath(sv, ev, spaths, visited)
}

func findClosestVertex(spaths [][]int, visited [][]bool) []int {
	var r, c int
	minP := math.MaxInt
	for i := 0; i < gInpRows; i++ {
		for j := 0; j < gInpCols; j++ {
			if visited[i][j] {
				continue
			}
			if spaths[i][j] < minP {
				r, c, minP = i, j, spaths[i][j]
			}
		}
	}
	return []int{r, c}
}

func getVisitableNbrs(v []int) [][]int {
	r, c := v[0], v[1]
	allNbrs := [][]int{{r + 1, c}, {r - 1, c}, {r, c + 1}, {r, c - 1}}
	visNbrs := make([][]int, 0)
	for _, nbr := range allNbrs {
		if inBounds(nbr) && isVisitable(v, nbr) {
			visNbrs = append(visNbrs, nbr)
		}
	}
	return visNbrs
}

func inBounds(v []int) bool {
	r, c := v[0], v[1]
	return r >= 0 && r < gInpRows && c >= 0 && c < gInpCols
}

func isEqual(v1, v2 []int) bool {
	return v1[0] == v2[0] && v1[1] == v2[1]
}

func isVisitable(vsrc, vdes []int) bool {
	return gInput[vdes[0]][vdes[1]] <= gInput[vsrc[0]][vsrc[1]]+1
}

func findVerticesByValue(b byte) [][]int {
	ans := make([][]int, 0)
	for i, row := range gInput {
		for j, cell := range row {
			if cell == b {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
	gInpRows = len(gInput)
	gInpCols = len(gInput[0])
	gStart = findVerticesByValue('S')[0]
	gEnd = findVerticesByValue('E')[0]
	gInput[gStart[0]][gStart[1]] = 'a'
	gInput[gEnd[0]][gEnd[1]] = 'z'
}
