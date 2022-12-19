package day14

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/outils"
	"github.com/avertocle/contests/io/stringz"
	"math"
	"strings"
)

var gInput [][][]int
var gBoundTL []int
var gBoundBR []int

func SolveP1() string {
	grid := bytez.Init2D(600, 200, '.')
	for _, rocks := range gInput {
		addRocksToGrid(grid, rocks)
	}
	grid[500][0] = 'S'
	printGrid(grid)
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/
func addRocksToGrid(grid [][]byte, rocks [][]int) {
	var xs, xe, ys, ye int
	for i := 0; i < len(rocks)-1; i++ {
		xs = intz.Min(rocks[i][0], rocks[i+1][0])
		xe = intz.Max(rocks[i][0], rocks[i+1][0])
		ys = intz.Min(rocks[i][1], rocks[i+1][1])
		ye = intz.Max(rocks[i][1], rocks[i+1][1])
		for x := xs; x <= xe; x++ {
			for y := ys; y <= ye; y++ {
				grid[x][y] = '#'
			}
		}
	}
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)

	var tokens []string
	var points [][]int
	gInput = make([][][]int, len(lines))
	maxX, maxY := math.MinInt, math.MinInt
	minX, minY := math.MaxInt, math.MaxInt
	px, py := 0, 0
	for i, l := range lines {
		tokens = strings.Fields(l)
		points = make([][]int, 0)
		for j := 0; j < len(tokens); j += 2 {
			px = stringz.AtoiQ(strings.Split(tokens[j], ",")[0], -1)
			py = stringz.AtoiQ(strings.Split(tokens[j], ",")[1], -1)
			maxX, maxY = intz.Max(maxX, px), intz.Max(maxY, py)
			minX, minY = intz.Min(minX, px), intz.Min(minY, py)
			points = append(points, []int{px, py})
		}
		gInput[i] = points
	}
	gBoundTL = []int{minX, 0} // sand dropping from 500,0
	gBoundBR = []int{maxX, maxY}
	fmt.Printf("\n\n bounds : tl(%v,%v) br(%v,%v) \n\n", minX, minY, maxX, maxY)
}

func printGrid(grid [][]byte) {
	outils.PrettyArray2DByte(bytez.Transpose2D(bytez.Extract2D(grid, gBoundTL, gBoundBR, '.')))
}
