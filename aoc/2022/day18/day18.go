package day18

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/numz"
	"math"
)

var gInput [][]int
var gInpMaxCoord int

const (
	empty   = 0
	cube    = 1
	steamed = 2
)

func SolveP1() string {
	openSides := 0
	for i, _ := range gInput {
		openSides += 6 - sidesTouchingCubes(i)
	}
	ans := openSides
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	gridSize := gInpMaxCoord + 1 // enclose in a tight large cube
	grid3d := intz.Init3D(gridSize, gridSize, gridSize, empty)
	markCubes(grid3d)
	markSteamed(grid3d, gridSize, []int{0, 0, 0})
	pockets := findPockets(grid3d, gridSize)
	openSides := 0
	for i, _ := range gInput {
		openSides += 6 - sidesTouchingCubes(i) - sidesTouchingPockets(i, pockets)
	}
	ans := openSides
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func areTouching(c1, c2 []int) bool {
	if c1[0] == c2[0] && c1[1] == c2[1] && numz.Abs(c1[2]-c2[2]) == 1 {
		return true
	} else if c1[1] == c2[1] && c1[2] == c2[2] && numz.Abs(c1[0]-c2[0]) == 1 {
		return true
	} else if c1[2] == c2[2] && c1[0] == c2[0] && numz.Abs(c1[1]-c2[1]) == 1 {
		return true
	} else {
		return false
	}
}

func sidesTouchingCubes(idx int) int {
	currCube := gInput[idx]
	touchCtr := 0
	for i, tempCube := range gInput {
		if i == idx {
			continue
		}
		if areTouching(currCube, tempCube) {
			touchCtr++
		}
	}
	return touchCtr
}

/***** P2 Functions *****/

func markCubes(grid3d [][][]int) {
	for _, c := range gInput {
		grid3d[c[0]][c[1]][c[2]] = cube
	}
}

func markSteamed(grid3d [][][]int, gridSize int, c []int) {
	if !isInBounds(c, gridSize) {
		return
	}
	v := grid3d[c[0]][c[1]][c[2]]
	if v == steamed || v == cube {
		return
	}
	grid3d[c[0]][c[1]][c[2]] = steamed
	markSteamed(grid3d, gridSize, []int{c[0] + 1, c[1], c[2]})
	markSteamed(grid3d, gridSize, []int{c[0] - 1, c[1], c[2]})
	markSteamed(grid3d, gridSize, []int{c[0], c[1] + 1, c[2]})
	markSteamed(grid3d, gridSize, []int{c[0], c[1] - 1, c[2]})
	markSteamed(grid3d, gridSize, []int{c[0], c[1], c[2] + 1})
	markSteamed(grid3d, gridSize, []int{c[0], c[1], c[2] - 1})
}

func findPockets(grid3d [][][]int, gridSize int) [][]int {
	pockets := make([][]int, 0)
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			for k := 0; k < gridSize; k++ {
				if grid3d[i][j][k] == empty {
					pockets = append(pockets, []int{i, j, k})
				}
			}
		}
	}
	return pockets
}

func sidesTouchingPockets(idx int, pockets [][]int) int {
	currCube := gInput[idx]
	touchCtr := 0
	for _, tempCube := range pockets {
		if areTouching(currCube, tempCube) {
			touchCtr++
		}
	}
	return touchCtr
}

func isInBounds(c []int, gridSize int) bool {
	return c[0] >= 0 && c[0] < gridSize &&
		c[1] >= 0 && c[1] < gridSize &&
		c[2] >= 0 && c[2] < gridSize
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractInt2DFromString1D(lines, ",", nil, math.MinInt)
	gInpMaxCoord = -1
	for _, c := range gInput {
		gInpMaxCoord, _ = intz.FindMax1D([]int{c[0], c[1], c[2], gInpMaxCoord})
	}
	//fmt.Println("max coord : ", gInpMaxCoord)
}
