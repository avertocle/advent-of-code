package day18

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"github.com/avertocle/contests/io/tpz"
	"math"
)

const DirPath = "../2024/day18"

var gDims *arrz.Idx2D[int]
var gInput []*arrz.Idx2D[int]
var gMaxBlockCount int

const (
	safe   = byte('.')
	unsafe = byte('#')
)

func SolveP1() string {
	grid, start, end, costGrid, unvisited := initStuff(gMaxBlockCount)
	findShortestPath(grid, start, end, costGrid, unvisited)
	ans := costGrid[end.I][end.J]
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := ""
	// binary search
	s, e, m, a := gMaxBlockCount, len(gInput), -1, -1
	for s < e {
		m = (e + s) / 2
		grid, start, end, costGrid, unvisited := initStuff(m)
		findShortestPath(grid, start, end, costGrid, unvisited)
		cost := costGrid[end.I][end.J]
		if cost == math.MaxInt/2 {
			a, e = m, m
		} else {
			s = m + 1
		}
	}
	ans = fmt.Sprintf("%v,%v", gInput[a].J, gInput[a].I)
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

func initStuff(maxBlockCount int) ([][]byte, *arrz.Idx2D[int], *arrz.Idx2D[int], [][]int, tpz.StringSet) {
	grid := arrz.MarkOnNewGrid2D(gInput[0:maxBlockCount+1], gDims, safe, unsafe, false)
	s, e := arrz.NewIdx2D(0, 0), arrz.NewIdx2D(gDims.I-1, gDims.J-1)
	costGrid := arrz.Init2D(gDims.I, gDims.J, math.MaxInt/2)
	costGrid[s.I][s.J] = 0
	unvisited := initUnvisited(grid)
	return grid, s, e, costGrid, unvisited
}

func findShortestPath(grid [][]byte, curr, end *arrz.Idx2D[int], costs [][]int, unvisited tpz.StringSet) {
	delete(unvisited, curr.ToKey())
	if curr.IsEqual(end) {
		return
	}
	nbrs := curr.Neighbours(false)
	for _, n := range nbrs {
		if n.IsInBounds(len(grid), len(grid[0])) && grid[n.I][n.J] == safe {
			if costs[n.I][n.J] > costs[curr.I][curr.J]+1 {
				costs[n.I][n.J] = costs[curr.I][curr.J] + 1
			}
		}
	}
	if nextToVisit := findNextToVisit(costs, unvisited); nextToVisit != nil {
		findShortestPath(grid, nextToVisit, end, costs, unvisited)
	}
}

func initUnvisited(grid [][]byte) tpz.StringSet {
	unvisited := make(tpz.StringSet)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == safe {
				unvisited[arrz.NewIdx2D(i, j).ToKey()] = true
			}
		}
	}
	return unvisited
}

func findNextToVisit(costGrid [][]int, unvisited tpz.StringSet) *arrz.Idx2D[int] {
	minCost := math.MaxInt
	var minCostNode *arrz.Idx2D[int]
	for k, _ := range unvisited {
		node := arrz.NewIdx2DFromKey[int](k)
		if minCost > costGrid[node.I][node.J] {
			minCost = costGrid[node.I][node.J]
			minCostNode = node
		}
	}
	return minCostNode
}

func PrintCostGrid(costGrid [][]int) {
	for _, row := range costGrid {
		for _, cell := range row {
			if cell == math.MaxInt/2 {
				fmt.Printf("-- ")
			} else {
				fmt.Printf("%02d ", cell)
			}
		}
		fmt.Println()
	}
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = make([]*arrz.Idx2D[int], 0)
	sections := iutils.BreakByEmptyLineString1D(lines)
	gMaxBlockCount = stringz.AtoI(sections[0][1], -1)
	// switching I,J while taking input to make it more intuitive
	t := iutils.ExtractInt1DFromString0D(sections[0][0], ",", -1)
	gDims = arrz.NewIdx2D[int](t[1], t[0])
	for _, line := range sections[1] {
		t := iutils.ExtractInt1DFromString0D(line, ",", -1)
		gInput = append(gInput, arrz.NewIdx2D[int](t[1], t[0]))
	}
}
