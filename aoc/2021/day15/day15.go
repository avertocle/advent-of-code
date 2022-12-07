package day15

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"log"
	"math"
)

// Nbr = neighbour

var gInput [][]byte
var gInpRows, gInpCols int

func SolveP1() string {
	vis := intz.Init2D(0, gInpRows, gInpCols)
	shortPaths := intz.Init2D(math.MaxInt, gInpRows, gInpCols)
	ans := traverse(shortPaths, vis)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	gInput, gInpRows, gInpCols = getBloatedInputP2()
	vis := intz.Init2D(0, gInpRows, gInpCols)
	shortPaths := intz.Init2D(math.MaxInt, gInpRows, gInpCols)
	ans := traverse(shortPaths, vis)
	return fmt.Sprintf("%v", ans)
}

func getBloatedInputP2() ([][]byte, int, int) {
	ans := bytez.Init2D(gInpRows*5, gInpCols*5, 0)
	x := 0
	for i := 0; i < gInpRows; i++ {
		for j := 0; j < gInpCols; j++ {
			x = riskAt(i, j)
			for k := 0; k < 5; k++ {
				for l := 0; l < 5; l++ {
					ans[i+k*gInpRows][j+l*gInpCols] = byte(foldInc(x, (k+l)*gInpCols)) + '0'
				}
			}
		}
	}
	return ans, gInpRows * 5, gInpCols * 5
}

func foldInc(v, folds int) int {
	for i := 1; i <= folds; i++ {
		v++
		if v == 10 {
			v = 1
		}
	}
	return v
}

func traverse(shortPaths [][]int, vis [][]int) int {
	curr := []int{0, 0}
	shortPaths[0][0] = 0
	for curr != nil {
		//fmt.Printf("visiting (%v,%v)\n", curr[0], curr[1])
		x, y := curr[0], curr[1]
		nbrs := getUnvisitedNbrs(x, y)
		for _, n := range nbrs {
			nx, ny := n[0], n[1]
			shortPaths[nx][ny] = intz.Min(shortPaths[nx][ny], shortPaths[x][y]+riskAt(nx, ny))
		}
		vis[x][y] = 1
		curr = findClosestUnvisitedVertex(shortPaths, vis)
	}
	ans := shortPaths[gInpRows-1][gInpCols-1]
	return ans
}

func updateNbrPaths(curr []int, shortPaths [][]int, vis [][]int) {
}

// N^2 hack to avoid implementing priority queue
func findClosestUnvisitedVertex(shortPaths [][]int, vis [][]int) []int {
	min := math.MaxInt
	var ans []int
	for i, row := range shortPaths {
		for j, cell := range row {
			if !isVisited(i, j, vis) && cell <= min {
				min = cell
				ans = []int{i, j}
			}
		}
	}
	return ans
}

func getUnvisitedNbrs(x, y int) [][]int {
	next := make([][]int, 0)
	for _, n := range [][]int{{x + 1, y}, {x, y + 1}, {x - 1, y}, {x, y - 1}} {
		if inBounds(n[0], n[1]) {
			next = append(next, n)
		}
	}
	return next
}

/***** Common Functions *****/

func isVisited(x, y int, vis [][]int) bool {
	return vis[x][y] > 0
}

func inBounds(x, y int) bool {
	if x >= 0 && x < gInpRows &&
		y >= 0 && y < gInpCols {
		return true
	}
	return false
}

func riskAt(x, y int) int {
	return int(gInput[x][y]-'1') + 1
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
	gInpRows = len(gInput)
	gInpCols = len(gInput[0])
	//outils.PrettyArray2DByte(gInput)
}
