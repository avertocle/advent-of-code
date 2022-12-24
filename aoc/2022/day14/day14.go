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

const cellEmpty = '.'
const cellSand = 'o'
const cellRock = '#'

func SolveP1() string {
	grid := makeGrid(600)
	ans := 0
	for !simOneSandP1(grid) {
		ans++
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	grid := makeGrid(1000)
	ans := 0
	for !simOneSandP2(grid) {
		ans++
	}
	ans++ // to account for last rock on the origin
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

// returns sim-end-reached i.e inAbyss
func simOneSandP1(grid [][]byte) bool {
	sx, sy := 500, 0
	floor := gBoundBR[1]
	isSet := false
	for {
		isSet, sx, sy = move(grid, sx, sy)
		if sy == floor {
			return true
		}
		if isSet {
			grid[sx][sy] = cellSand
			return false
		}
	}
}

/***** P2 Functions *****/

// returns sim-end-reached i.e top reached
func simOneSandP2(grid [][]byte) bool {
	sx, sy := 500, 0
	floor := gBoundBR[1] + 1
	isSet := false
	for {
		isSet, sx, sy = move(grid, sx, sy)
		if sy == floor {
			grid[sx][sy] = cellSand
			return false
		}
		if isSet {
			grid[sx][sy] = cellSand
			if sx == 500 && sy == 0 {
				return true
			}
			return false
		}
	}
}

/***** Common Functions *****/

func move(grid [][]byte, x, y int) (bool, int, int) {
	if movePossible(grid, x, y+1) {
		return false, x, y + 1
	} else if movePossible(grid, x-1, y+1) {
		return false, x - 1, y + 1
	} else if movePossible(grid, x+1, y+1) {
		return false, x + 1, y + 1
	} else {
		return true, x, y
	}
}

func movePossible(grid [][]byte, x, y int) bool {
	return grid[x][y] == cellEmpty
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

func makeGrid(size int) [][]byte {
	grid := bytez.Init2D(size, size, '.')
	for _, rocks := range gInput {
		addRocksToGrid(grid, rocks)
	}
	return grid
}

func addRocksToGrid(grid [][]byte, rocks [][]int) {
	var xs, xe, ys, ye int
	for i := 0; i < len(rocks)-1; i++ {
		xs = intz.Min(rocks[i][0], rocks[i+1][0])
		xe = intz.Max(rocks[i][0], rocks[i+1][0])
		ys = intz.Min(rocks[i][1], rocks[i+1][1])
		ye = intz.Max(rocks[i][1], rocks[i+1][1])
		for x := xs; x <= xe; x++ {
			for y := ys; y <= ye; y++ {
				grid[x][y] = cellRock
			}
		}
	}
}

func printGrid(grid [][]byte) {
	fmt.Println()
	viewTL := []int{gBoundTL[0] - 20, gBoundTL[1]}
	viewBR := []int{gBoundBR[0] + 20, gBoundBR[1] + 3}
	outils.PrettyArray2DByte(bytez.Transpose2D(bytez.Extract2D(grid, viewTL, viewBR, '.')))
}
