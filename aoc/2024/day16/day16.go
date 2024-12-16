package day16

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/cmz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"math"
	"slices"
)

const DirPath = "../2024/day16"

// assumes input grid is padded with a '#' layer, if not, pad first while parsing input
var gInput [][]byte
var gStart *arrz.Idx2DD[int]
var gEnd *arrz.Idx2DD[int]

func SolveP1() string {
	ans := 0
	fmt.Println(gStart, gEnd)
	costGrid := intz.Init3D(len(gInput), len(gInput[0]), 4, math.MaxInt/2)
	costGrid[gStart.I][gStart.J][gStart.D] = 0
	unvisited := makeUnvisited(gInput)
	findShortestPath(gInput, gStart, gEnd, costGrid, make(cmz.MapVisited), unvisited)
	//fmt.Println(costGrid[gEnd.I][gEnd.J])
	ans = slices.Min(costGrid[gEnd.I][gEnd.J])
	var ans2 int
	tilesVisited := make(cmz.MapVisited)
	findTileCountInAllShortestPaths(gInput, gStart, gEnd, costGrid, tilesVisited, make(cmz.MapVisited), &ans2)
	ans2 = len(tilesVisited)
	return fmt.Sprintf("%v, %v", ans, ans2)
}

func uniqueTiles(visited cmz.MapVisited) int {
	unique := make(cmz.MapVisited)
	for k, ok := range visited {
		if ok {
			x := arrz.NewIdx2DDFromKey[int](k)
			unique[arrz.NewIdx2D(x.I, x.J).ToKey()] = true
		}
	}
	return len(unique)
}

func findTileCountInAllShortestPaths(grid [][]byte, curr, end *arrz.Idx2DD[int], costGrid [][][]int, tilesVisited, visited cmz.MapVisited, count *int) {
	visited[curr.ToKey()] = true
	if curr.IsEqual(end, true) {
		for k, v := range visited {
			if v {
				x := arrz.NewIdx2DDFromKey[int](k)
				tilesVisited[arrz.NewIdx2D(x.I, x.J).ToKey()] = true
			}
		}
		visited[curr.ToKey()] = false
		return
	}
	nbrs, _ := findVisitableNbrs(grid, curr, visited)
	for _, n := range nbrs {

		if costGrid[n.I][n.J][n.D] == costGrid[curr.I][curr.J][curr.D]+1 || costGrid[n.I][n.J][n.D] == costGrid[curr.I][curr.J][curr.D]+1000 {
			findTileCountInAllShortestPaths(grid, n, end, costGrid, tilesVisited, visited, count)
		}
	}
	visited[curr.ToKey()] = false
}

func findShortestPath(grid [][]byte, curr, end *arrz.Idx2DD[int], costGrid [][][]int, visited, unvisited cmz.MapVisited) {
	//printDebug(grid, costGrid, visited, curr, end)
	if len(visited)%1000 == 0 {
		fmt.Printf("%v/%v -> ", len(visited), len(unvisited))
	}
	if curr.IsEqual(end, true) {
		return
	}
	visited[curr.ToKey()] = true
	delete(unvisited, curr.ToKey())
	nbrs, stepCosts := findVisitableNbrs(grid, curr, visited)
	for i, n := range nbrs {
		oldCost := costGrid[n.I][n.J][n.D]
		newCost := costGrid[curr.I][curr.J][curr.D] + stepCosts[i]
		if newCost < oldCost {
			costGrid[n.I][n.J][n.D] = newCost
		}
	}
	next := findNextTobeVisitedNode(costGrid, visited, unvisited)
	findShortestPath(grid, next, end, costGrid, visited, unvisited)
}

func makeUnvisited(grid [][]byte) cmz.MapVisited {
	unVisited := make(cmz.MapVisited)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			for d := 0; d < 4; d++ {
				if grid[i][j] != '#' {
					unVisited[arrz.NewIdx2DD(i, j, d).ToKey()] = true
				}
			}
		}
	}
	return unVisited
}

func findNextTobeVisitedNode(costGrid [][][]int, visited, unvisited cmz.MapVisited) *arrz.Idx2DD[int] {
	minCost := math.MaxInt
	var next *arrz.Idx2DD[int]
	for k, _ := range unvisited {
		idx := arrz.NewIdx2DDFromKey[int](k)
		if costGrid[idx.I][idx.J][idx.D] < minCost {
			minCost = costGrid[idx.I][idx.J][idx.D]
			next = idx
		}
	}
	return next
}

func findVisitableNbrs(grid [][]byte, curr *arrz.Idx2DD[int], visited cmz.MapVisited) ([]*arrz.Idx2DD[int], []int) {
	nextStates := curr.NextStates()
	nbrs, stepCosts := make([]*arrz.Idx2DD[int], 0), make([]int, 0)
	for _, ns := range nextStates {
		if ns.IsInBounds(len(grid), len(grid[0])) && grid[ns.I][ns.J] != '#' && !visited[ns.ToKey()] {
			stepCost := 1
			if ns.D != curr.D {
				stepCost = 1000
			}
			nbrs = append(nbrs, ns)
			stepCosts = append(stepCosts, stepCost)
		}
	}
	return nbrs, stepCosts
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
	start := arrz.Find2D(gInput, byte('S'))[0]
	end := arrz.Find2D(gInput, byte('E'))[0]
	gStart = arrz.NewIdx2DD(start[0], start[1], arrz.Right)
	gEnd = arrz.NewIdx2DD(end[0], end[1], arrz.Right)
}

func printDebug(grid [][]byte, costGrid [][][]int, visited cmz.MapVisited, curr, end *arrz.Idx2DD[int]) {
	fmt.Println(curr.Str(), end.Str())
	for k, _ := range visited {
		fmt.Printf("%v ", k)
	}
	fmt.Println()
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			minCost := slices.Min(costGrid[i][j])
			if minCost < math.MaxInt/2 {
				fmt.Printf("%v ", minCost)
			} else {
				fmt.Printf("%v ", string(grid[i][j]))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
