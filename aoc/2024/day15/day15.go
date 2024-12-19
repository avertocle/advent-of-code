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

// old debug 334 next move ^
// old debug 4699 next move v
// 1475511 too high

func SolveP1() string {
	ans := 0
	//grid := arrz.Copy2D(gInput)
	//curr := findRobot(grid)
	////arrz.PPrint2D(grid)
	//for i := 0; i < len(gPath); i++ {
	//	curr = findRobot(grid)
	//	moveRobotP1(grid, curr, gPath[i])
	//}
	//ans = calcGpsSum(grid)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	grid := makeGridP2()
	//	grid := arrz.Copy2D(gInput)
	//fmt.Println("Grid: ", arrz.Count2D(grid, Wall), arrz.Count2D(grid, Space), arrz.Count2D(grid, Box), arrz.Count2D(grid, Box2S), arrz.Count2D(grid, Box2E), calcGpsSum(grid))
	curr := findRobot(grid)
	for i := 0; i < len(gPath); i++ {
		fmt.Println("\n--------------------\n")
		arrz.PPrint2D(grid)
		fmt.Println("Next move: ", string(gPath[i]), " | ", i)
		curr = findRobot(grid)
		moveRobotP2(grid, curr, gPath[i])
	}
	//fmt.Println("Grid: ", arrz.Count2D(grid, Wall), arrz.Count2D(grid, Space), arrz.Count2D(grid, Box), arrz.Count2D(grid, Box2S), arrz.Count2D(grid, Box2E))
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
		return []*idx{arrz.NewIdx2D[int](c.I-1, c.J-1), arrz.NewIdx2D[int](c.I-1, c.J)}
	} else {
		return []*idx{arrz.NewIdx2D[int](c.I-1, c.J)}
	}
}

func findBoxesBottom(grid [][]byte, c *idx) []*idx {
	if grid[c.I+1][c.J] == Box2S {
		return []*idx{arrz.NewIdx2D[int](c.I+1, c.J), arrz.NewIdx2D[int](c.I+1, c.J+1)}
	} else if grid[c.I+1][c.J] == Box2E {
		// make sure to return the same J point first, the ordering is used later
		return []*idx{arrz.NewIdx2D[int](c.I+1, c.J-1), arrz.NewIdx2D[int](c.I+1, c.J)}
	} else {
		return []*idx{arrz.NewIdx2D[int](c.I+1, c.J)}
	}
}

func isPartOfBox(grid [][]byte, c *idx) bool {
	return grid[c.I][c.J] == Box2S || grid[c.I][c.J] == Box2E
}

func pruneNbs(grid [][]byte, nbs []*idx) []*idx {
	nbsMap := make(map[string][]*idx)
	nbsOrder := make([]string, 0)
	for i := 0; i < len(nbs); i++ {
		n := nbs[i]
		if grid[n.I][n.J] == Box2S {
			key := nbs[i].ToKey() + "-" + nbs[i+1].ToKey()
			if _, ok := nbsMap[key]; !ok {
				nbsOrder = append(nbsOrder, key)
				nbsMap[key] = []*idx{nbs[i], nbs[i+1]}
			}
			i++
		} else {
			errz.HardAssert(grid[n.I][n.J] != Box2E, "invalid box end : %v : nbs = %v", n.Str(), arrz.Idx2DListToStr(nbs))
			nbsMap[n.ToKey()] = []*idx{n}
			nbsOrder = append(nbsOrder, n.ToKey())
		}
	}
	nbs = make([]*idx, 0)
	for _, key := range nbsOrder {
		nbs = append(nbs, nbsMap[key]...)
	}
	return nbs
}

func moveVerticalP2(grid [][]byte, cs []*idx, getNext func([][]byte, *idx) []*idx) bool {
	nbs := make([]*idx, 0)
	for _, c := range cs {
		nbs = append(nbs, getNext(grid, c)...)
	}
	nbs = pruneNbs(grid, nbs)
	//fmt.Println(" | cs = ", arrz.Idx2DListToStr(cs), "nbs = ", arrz.Idx2DListToStr(nbs))
	allNbrClear := true
	for _, n := range nbs {
		if grid[n.I][n.J] == Wall {
			return false
		} else if isPartOfBox(grid, n) {
			allNbrClear = false
			break
		}
	}
	if allNbrClear {
		for i := 0; i < len(nbs); i++ {
			grid[nbs[i].I][nbs[i].J] = grid[cs[i].I][cs[i].J]
			grid[cs[i].I][cs[i].J] = Space
		}
		return true
	} else {
		nbs2 := make([]*idx, 0)
		for _, n := range nbs {
			if isPartOfBox(grid, n) {
				nbs2 = append(nbs2, n)
			}
		}
		r := moveVerticalP2(grid, nbs2, getNext)
		if r {
			moveVerticalP2(grid, cs, getNext)
			return true
		}
	}
	return false
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
			} else if grid[i][j] == Box2S {
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
	fmt.Println("Input Parsed: ", len(gPath))
}
