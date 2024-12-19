package day15

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
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
type funcNext = func(*idx) *idx

func SolveP1() string {
	ans := 0
	grid := arrz.Copy2D(gInput)
	curr := findRobot(grid)
	//arrz.PPrint2D(grid)
	for i := 0; i < len(gPath); i++ {
		curr = findRobot(grid)
		moveRobotP1(grid, curr, gPath[i])
	}
	ans = calcGpsSum(grid)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	grid := makeGridP2()
	curr := findRobot(grid)
	for i := 0; i < len(gPath); i++ {
		fmt.Println("\n--------------------\n")
		arrz.PPrint2D(grid)
		fmt.Println("Next move: ", string(gPath[i]))
		curr = findRobot(grid)
		moveRobotP2(grid, curr, gPath[i])
	}
	arrz.PPrint2D(grid)
	ans = calcGpsSum(grid)
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func moveRobotP1(grid [][]byte, curr *arrz.Idx2D[int], dir byte) {
	switch dir {
	case 'v':
		moveGenericP1(grid, curr, func(x *idx) *idx { return arrz.NewIdx2D[int](x.I+1, x.J) })
	case '^':
		moveGenericP1(grid, curr, func(x *idx) *idx { return arrz.NewIdx2D[int](x.I-1, x.J) })
	case '<':
		moveGenericP1(grid, curr, func(x *idx) *idx { return arrz.NewIdx2D[int](x.I, x.J-1) })
	case '>':
		moveGenericP1(grid, curr, func(x *idx) *idx { return arrz.NewIdx2D[int](x.I, x.J+1) })
	}
}

func moveGenericP1(grid [][]byte, c *idx, getNext func(*idx) *idx) bool {
	canMove := false
	c1 := getNext(c)
	if grid[c1.I][c1.J] == Wall {
		canMove = false
	} else if grid[c1.I][c1.J] == Space {
		canMove = true
	} else {
		canMove = moveGenericP1(grid, c1, getNext)
	}

	if canMove {
		grid[c1.I][c1.J] = grid[c.I][c.J]
		grid[c.I][c.J] = Space
	}
	return canMove
}

/***** P2 Functions *****/

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

func moveRobotP2(grid [][]byte, curr *arrz.Idx2D[int], dir byte) {
	switch dir {
	case 'v':
		moveVerticalP2(grid, []*arrz.Idx2D[int]{curr}, findBoxesBottom)
	case '^':
		moveVerticalP2(grid, []*arrz.Idx2D[int]{curr}, findBoxesTop)
	case '<':
		moveGenericP1(grid, curr, func(x *idx) *idx { return arrz.NewIdx2D[int](x.I, x.J-1) })
	case '>':
		moveGenericP1(grid, curr, func(x *idx) *idx { return arrz.NewIdx2D[int](x.I, x.J+1) })
	}
}

func findBoxesTop(grid [][]byte, c *idx) []*idx {
	if grid[c.I-1][c.J] == Box2S {
		return []*idx{arrz.NewIdx2D[int](c.I-1, c.J), arrz.NewIdx2D[int](c.I-1, c.J+1)}
	} else if grid[c.I-1][c.J] == Box2E {
		return []*idx{arrz.NewIdx2D[int](c.I-1, c.J), arrz.NewIdx2D[int](c.I-1, c.J-1)}
	} else {
		return []*idx{arrz.NewIdx2D[int](c.I-1, c.J)}
	}
}

func findBoxesBottom(grid [][]byte, c *idx) []*idx {
	if grid[c.I+1][c.J] == Box2S {
		return []*idx{arrz.NewIdx2D[int](c.I+1, c.J), arrz.NewIdx2D[int](c.I+1, c.J+1)}
	} else if grid[c.I+1][c.J] == Box2E {
		// make sure to return the same J point first, the ordering is used later
		return []*idx{arrz.NewIdx2D[int](c.I+1, c.J), arrz.NewIdx2D[int](c.I+1, c.J-1)}
	} else {
		return []*idx{arrz.NewIdx2D[int](c.I+1, c.J)}
	}
}

func isPartOfBox(grid [][]byte, c *idx) bool {
	return grid[c.I][c.J] == Box2S || grid[c.I][c.J] == Box2E
}

func moveVerticalP2(grid [][]byte, cs []*idx, getNext func([][]byte, *idx) []*idx) bool {
	isBox := len(cs) == 2
	if !isBox {
		c := cs[0]
		nbs := getNext(grid, c)
		canMove, isBlockedByBox := checkMobility(grid, nbs)
		if canMove {
			grid[nbs[0].I][nbs[0].J] = grid[c.I][c.J]
			grid[c.I][c.J] = Space
			return true
		} else if isBlockedByBox {
			didNbrMove := moveVerticalP2(grid, nbs, getNext)
			if didNbrMove {
				moveVerticalP2(grid, cs, getNext)
				return true
			}
		}
	} else {
		c1, c2 := cs[0], cs[1]
		nbs1, nbs2 := getNext(grid, c1), getNext(grid, c2)
		canMove1, isBlockedByBox1 := checkMobility(grid, nbs1)
		canMove2, isBlockedByBox2 := checkMobility(grid, nbs2)
		if canMove1 && canMove2 {
			grid[nbs1[0].I][nbs1[0].J] = grid[c1.I][c1.J]
			grid[c1.I][c1.J] = Space
			grid[nbs2[0].I][nbs2[0].J] = grid[c2.I][c2.J]
			grid[c2.I][c2.J] = Space
			return true
		} else if isBlockedByBox1 && isBlockedByBox2 {
			didNbrMove1 := moveVerticalP2(grid, nbs1, getNext)
			didNbrMove2 := moveVerticalP2(grid, nbs2, getNext)
			if didNbrMove1 && didNbrMove2 {
				moveVerticalP2(grid, cs, getNext)
				return true
			}
		} else if isBlockedByBox1 {
			didNbrMove := moveVerticalP2(grid, nbs1, getNext)
			if didNbrMove {
				moveVerticalP2(grid, cs, getNext)
				return true
			}
		} else if isBlockedByBox2 {
			didNbrMove := moveVerticalP2(grid, nbs2, getNext)
			if didNbrMove {
				moveVerticalP2(grid, cs, getNext)
				return true
			}
		}
	}
	return false
}

func checkMobility(grid [][]byte, nbs []*idx) (bool, bool) {
	canMove, isBlockedByBox := false, false
	if len(nbs) == 1 {
		isBlockedByBox = false
		n := nbs[0]
		if grid[n.I][n.J] == Wall {
			canMove = false
		} else if grid[n.I][n.J] == Space {
			canMove = true
		}
	} else {
		isBlockedByBox = true
	}
	return canMove, isBlockedByBox
}

/***** Common Functions *****/

func findRobot(grid [][]byte) *idx {
	return arrz.NewIdx2D[int](bytez.Find2D(grid, '@')[0]...)
}

func calcGpsSum(grid [][]byte) int {
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == Box {
				sum += i*100 + j
			}
			if grid[i][j] == Box2S {
				sum += i*100 + j
			}
		}
	}
	return sum
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	parts := iutils.BreakByEmptyLineString1D(lines)
	gInput = iutils.ExtractByte2DFromString1D(parts[0], "", nil, 0)
	gPath = []byte(strings.Join(parts[1], ""))
}
