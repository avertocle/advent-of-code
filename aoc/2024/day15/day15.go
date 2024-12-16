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

const Box = 'O'
const Wall = '#'
const Space = '.'
const Robot = '@'

func SolveP1() string {
	ans := 0
	grid := arrz.Copy2D(gInput)
	curr := arrz.NewIdx2D[int](bytez.Find2D(grid, '@')[0]...)
	//bytez.PPrint2D(grid)
	//fmt.Println()
	for i := 0; i < len(gPath); i++ {
		grid[curr.I][curr.J] = Space
		moveRobot(grid, curr, gPath[i])
		grid[curr.I][curr.J] = Robot
		//fmt.Println(string(gPath[i]))
		//bytez.PPrint2D(grid)
		//fmt.Println()
	}
	ans = calcGpsSum(grid)
	return fmt.Sprintf("%v", ans)
}

func calcGpsSum(grid [][]byte) int {
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == Box {
				sum += i*100 + j
			}
		}
	}
	return sum
}

func moveRobot(grid [][]byte, curr *arrz.Idx2D[int], dir byte) {
	switch dir {
	case 'v':
		moveDown(grid, curr)
	case '^':
		moveUp(grid, curr)
	case '<':
		moveLeft(grid, curr)
	case '>':
		moveRight(grid, curr)
	}
}

func moveRight(grid [][]byte, c *arrz.Idx2D[int]) {
	if grid[c.I][c.J+1] == Wall {
		return
	} else if grid[c.I][c.J+1] == Space {
		c.MoveBy(0, 1)
		return
	}

	nextSpace := c.J + 1
	for ; nextSpace < len(grid[0]) && grid[c.I][nextSpace] == Box; nextSpace++ {
	}
	fmt.Println(c, nextSpace, string(grid[c.I][c.J]), string(grid[c.I][nextSpace]))
	if grid[c.I][nextSpace] == Space {
		grid[c.I][c.J] = Space
		grid[c.I][nextSpace] = Box
		c.MoveBy(0, 1)
	}
}

func moveLeft(grid [][]byte, c *arrz.Idx2D[int]) {
	if grid[c.I][c.J-1] == Wall {
		return
	} else if grid[c.I][c.J-1] == Space {
		c.MoveBy(0, -1)
		return
	}

	nextSpace := c.J - 1
	for ; nextSpace >= 0 && grid[c.I][nextSpace] == Box; nextSpace-- {
	}
	fmt.Println(c, nextSpace, string(grid[c.I][c.J]), string(grid[c.I][nextSpace]))
	if grid[c.I][nextSpace] == Space {
		grid[c.I][c.J] = Space
		grid[c.I][nextSpace] = Box
		c.MoveBy(0, -1)
	}
}

func moveUp(grid [][]byte, c *arrz.Idx2D[int]) {
	if grid[c.I-1][c.J] == Wall {
		return
	} else if grid[c.I-1][c.J] == Space {
		c.MoveBy(-1, 0)
		return
	}

	nextSpace := c.I - 1
	for ; nextSpace >= 0 && grid[nextSpace][c.J] == Box; nextSpace-- {
	}
	fmt.Println(c, nextSpace, string(grid[c.I][c.J]), string(grid[nextSpace][c.J]))
	if grid[nextSpace][c.J] == Space {
		grid[c.I][c.J] = Space
		grid[nextSpace][c.J] = Box
		c.MoveBy(-1, 0)
	}
}

func moveDown(grid [][]byte, c *arrz.Idx2D[int]) {
	if grid[c.I+1][c.J] == Wall {
		return
	} else if grid[c.I+1][c.J] == Space {
		c.MoveBy(1, 0)
		return
	}

	nextSpace := c.I + 1
	for ; nextSpace < len(grid) && grid[nextSpace][c.J] == Box; nextSpace++ {
	}
	fmt.Println(c, nextSpace, string(grid[c.I][c.J]), string(grid[nextSpace][c.J]))
	if grid[nextSpace][c.J] == Space {
		grid[c.I][c.J] = Space
		grid[nextSpace][c.J] = Box
		c.MoveBy(1, 0)
	}
}

//func moveDown(grid [][]byte, c *arrz.Idx2D[int], moveCount int) *arrz.Idx2D[int] {
//	if grid[c.I+1][c.J] == 'O' {
//		c.MoveBy(1, 0)
//		return c
//	}
//	nextWall := -1
//	blocksFound := 0
//	for i := c.I; i < len(grid); i++ {
//		if grid[i][c.J] == 'O' {
//			blocksFound++
//		} else if grid[i][c.J] == '#' {
//			nextWall = i
//			break
//		}
//	}
//	errz.HardAssert(blocksFound > 0, "at lease one  block should be there")
//	for i := c.I; i >= nextWall; i++ {
//		if i < nextWall-blocksFound {
//			grid[i][c.J] = '.'
//		} else {
//			grid[i][c.J] = 'O'
//		}
//	}
//	c.MoveBy(blocksFound, 0)
//	return c
//}

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
	parts := iutils.BreakByEmptyLineString1D(lines)
	gInput = iutils.ExtractByte2DFromString1D(parts[0], "", nil, 0)
	gPath = []byte(strings.Join(parts[1], ""))
}
