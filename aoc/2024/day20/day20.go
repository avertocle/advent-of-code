package day20

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/cmz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"math"
)

const DirPath = "../2024/day20"

var gInput [][]byte

const (
	space = byte('.')
	wall  = byte('#')
)

func SolveP1() string {
	grid, start, end, costGrid, unvisited := initStuff()
	findShortestPath(grid, start, end, costGrid, unvisited)
	ans := costGrid[end.I][end.J]
	return fmt.Sprintf("%v", ans)
}

func xxx() {
	for i := 0; i < len(gInput); i++ {
		for j := 0; j < len(gInput[0]); j++ {
		}
	}
}

func initStuff() ([][]byte, *arrz.Idx2D[int], *arrz.Idx2D[int], [][]int, cmz.MapVisited) {
	grid := arrz.Copy2D(gInput)
	s := arrz.NewIdx2D(arrz.Find2D(grid, 'S')[0]...)
	e := arrz.NewIdx2D(arrz.Find2D(grid, 'E')[0]...)
	costGrid := arrz.Init2D(len(grid), len(grid[0]), math.MaxInt)
	costGrid[s.I][s.J] = 0
	unvisited := initUnvisited(grid)
	return grid, s, e, costGrid, unvisited
}

func findShortestPath(grid [][]byte, curr *arrz.Idx2D[int], end *arrz.Idx2D[int], costs [][]int, unvisited cmz.MapVisited) {
	delete(unvisited, curr.ToKey())
	if curr.IsEqual(end) {
		return
	}
	nbrs := findVisitableNbrs(grid, curr)
	for _, n := range nbrs {
		if costs[n.I][n.J] > costs[curr.I][curr.J]+1 {
			costs[n.I][n.J] = costs[curr.I][curr.J] + 1
		}
	}
	if nextToVisit := findNextToVisit(costs, unvisited); nextToVisit != nil {
		findShortestPath(grid, nextToVisit, end, costs, unvisited)
	}
}

func initUnvisited(grid [][]byte) cmz.MapVisited {
	unvisited := make(cmz.MapVisited)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == space {
				unvisited[arrz.NewIdx2D(i, j).ToKey()] = true
			}
		}
	}
	return unvisited
}

func findNextToVisit(costGrid [][]int, unvisited cmz.MapVisited) *arrz.Idx2D[int] {
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

func findVisitableNbrs(grid [][]byte, curr *arrz.Idx2D[int]) []*arrz.Idx2D[int] {
	nbrs := make([]*arrz.Idx2D[int], 0)
	for _, n := range curr.Neighbours(false) {
		if n.IsInBounds(len(grid), len(grid[0])) && grid[n.I][n.J] != wall {
			nbrs = append(nbrs, n)
		}
	}
	return nbrs
}

func SolveP2() string {
	ans := 0
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
}
