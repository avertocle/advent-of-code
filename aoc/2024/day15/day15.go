package day15

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/tpz"
	"strings"
)

const DirPath = "../2024/day15"

// assumes input grid is padded with a '#' layer, if not, pad first while parsing input
var gInput [][]byte
var gPath []byte

const Box = byte('O')
const Wall = byte('#')
const Space = byte('.')
const Robot = byte('@')
const Box2S = byte('[')
const Box2E = byte(']')

type idx = arrz.Idx2D[int]

func SolveP1() string {
	ans := 0
	grid := arrz.Copy2D(gInput)
	curr := findRobot(grid)
	for i := 0; i < len(gPath); i++ {
		moveRobot(grid, curr, gPath[i], false)
		curr = findRobot(grid)
	}
	ans = arrz.Reduce2d(grid, 0, calcGpsOneBox)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	grid := makeGridP2()
	curr := findRobot(grid)
	for i := 0; i < len(gPath); i++ {
		curr = findRobot(grid)
		moveRobot(grid, curr, gPath[i], true)
	}
	ans = arrz.Reduce2d(grid, 0, calcGpsOneBox)
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func moveRobot(grid [][]byte, curr *arrz.Idx2D[int], dir byte, bigWarehouse bool) {
	switch dir {
	case 'v':
		if bigWarehouse {
			moveVerticalP2(grid, []*arrz.Idx2D[int]{curr}, findNbrsBottom)
		} else {
			moveGeneric(grid, curr, func(x *idx) *idx { return arrz.NewIdx2D[int](x.I+1, x.J) })
		}
	case '^':
		if bigWarehouse {
			moveVerticalP2(grid, []*arrz.Idx2D[int]{curr}, findNbrsTop)
		} else {
			moveGeneric(grid, curr, func(x *idx) *idx { return arrz.NewIdx2D[int](x.I-1, x.J) })
		}
	case '<':
		moveGeneric(grid, curr, func(x *idx) *idx { return arrz.NewIdx2D[int](x.I, x.J-1) })
	case '>':
		moveGeneric(grid, curr, func(x *idx) *idx { return arrz.NewIdx2D[int](x.I, x.J+1) })
	}
}

func moveVerticalP2(grid [][]byte, cs []*idx, getNext func([][]byte, *idx) []*idx) bool {
	nbrs := findNeighbours(grid, cs, getNext)
	if hasWall(grid, nbrs) {
		return false
	}
	if isAllSpaces(grid, nbrs) {
		for i := 0; i < len(nbrs); i++ {
			grid[nbrs[i].I][nbrs[i].J] = grid[cs[i].I][cs[i].J]
			grid[cs[i].I][cs[i].J] = Space
		}
		return true
	}

	nbrs = filterOnlyBoxNbrs(grid, nbrs)
	if moveVerticalP2(grid, nbrs, getNext) {
		moveVerticalP2(grid, cs, getNext)
		return true
	}
	return false
}

func findNeighbours(grid [][]byte, cs []*idx, getNext func([][]byte, *idx) []*idx) []*idx {
	// find all neighbours of all cs and return unique ones in order to not disrupt the brackets []
	nbrsAll := make([]*idx, 0)
	nbrsSet := make(tpz.StringSet)
	for _, c := range cs {
		nbrsOne := getNext(grid, c)
		for _, n := range nbrsOne {
			if _, ok := nbrsSet[n.ToKey()]; !ok {
				nbrsSet[n.ToKey()] = true
				nbrsAll = append(nbrsAll, n)
			}
		}
	}
	return nbrsAll
}

func findNbrsTop(grid [][]byte, c *idx) []*idx {
	if grid[c.I-1][c.J] == Box2S {
		return []*idx{arrz.NewIdx2D[int](c.I-1, c.J), arrz.NewIdx2D[int](c.I-1, c.J+1)}
	} else if grid[c.I-1][c.J] == Box2E {
		return []*idx{arrz.NewIdx2D[int](c.I-1, c.J-1), arrz.NewIdx2D[int](c.I-1, c.J)}
	} else {
		return []*idx{arrz.NewIdx2D[int](c.I-1, c.J)}
	}
}

func filterOnlyBoxNbrs(grid [][]byte, cs []*idx) []*idx {
	cs2 := make([]*idx, 0)
	for _, c := range cs {
		if isPartOfBox(grid, c) {
			cs2 = append(cs2, c)
		}
	}
	return cs2
}

func findNbrsBottom(grid [][]byte, c *idx) []*idx {
	if grid[c.I+1][c.J] == Box2S {
		return []*idx{arrz.NewIdx2D[int](c.I+1, c.J), arrz.NewIdx2D[int](c.I+1, c.J+1)}
	} else if grid[c.I+1][c.J] == Box2E {
		return []*idx{arrz.NewIdx2D[int](c.I+1, c.J-1), arrz.NewIdx2D[int](c.I+1, c.J)}
	} else {
		return []*idx{arrz.NewIdx2D[int](c.I+1, c.J)}
	}
}

func hasWall(grid [][]byte, c []*idx) bool {
	for _, n := range c {
		if grid[n.I][n.J] == Wall {
			return true
		}
	}
	return false
}

func isAllSpaces(grid [][]byte, c []*idx) bool {
	for _, n := range c {
		if grid[n.I][n.J] != Space {
			return false
		}
	}
	return true
}

func isPartOfBox(grid [][]byte, c *idx) bool {
	return grid[c.I][c.J] == Box2S || grid[c.I][c.J] == Box2E
}

func makeGridP2() [][]byte {
	r, c := len(gInput), len(gInput[0])*2
	grid := arrz.Init2D(r, c, Space)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j += 2 {
			grid[i][j] = gInput[i][j/2]
			if grid[i][j] == Wall || grid[i][j] == Space {
				grid[i][j+1] = grid[i][j]
			} else if grid[i][j] == Robot {
				grid[i][j+1] = Space
			} else {
				grid[i][j] = Box2S
				grid[i][j+1] = Box2E
			}
		}
	}
	return grid
}

/***** Common Functions *****/

func moveGeneric(grid [][]byte, c *idx, getNext func(*idx) *idx) bool {
	c1 := getNext(c)
	canMove := false
	if grid[c1.I][c1.J] == Wall {
		canMove = false
	} else if grid[c1.I][c1.J] == Space {
		canMove = true
	} else {
		canMove = moveGeneric(grid, c1, getNext)
	}

	if canMove {
		grid[c1.I][c1.J] = grid[c.I][c.J]
		grid[c.I][c.J] = Space
	}
	return canMove
}

func findRobot(grid [][]byte) *idx {
	return arrz.NewIdx2D[int](bytez.Find2D(grid, '@')[0]...)
}

func calcGpsOneBox(arr [][]byte, i, j int) int {
	if arr[i][j] == Box || arr[i][j] == Box2S {
		return i*100 + j
	}
	return 0
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	parts := iutils.BreakByEmptyLineString1D(lines)
	gInput = iutils.ExtractByte2DFromString1D(parts[0], "", nil, 0)
	gPath = []byte(strings.Join(parts[1], ""))
}
