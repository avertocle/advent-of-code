package day17

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

var gInput []byte

func SolveP1() string {
	//simTest()
	hmax := 5000
	rmax := 2022
	grid := bytez.Init2D(hmax, 9, '.')
	xbounds := []int{0, 8}
	ybounds := []int{0, len(grid) - 1}
	markBounds(grid, xbounds, ybounds)
	var sh shape
	var didGoDown bool
	var shCtr int
	for i := 0; ; i++ {
		if shCtr >= rmax {
			break
		}
		sh = getNewShape(shCtr, grid)
		shCtr++
		didGoDown = true
		for n := 0; didGoDown; n++ {
			didGoDown = fallOneStep(grid, sh, gInput[i%(len(gInput))])
			i++
			//didGoDown = fallOneStepAndDisplay(grid, sh, gInput[i%(len(gInput))])
		}
		sh.markOnGrid(grid, '#')
		fmt.Printf("%v\n", hmax-getTowerHeight(grid))
		i--
	}
	h := getTowerHeight(grid)
	//fmt.Printf("rock(%v), height(%v)\n", shCtr, h)
	//grid[h][5] = '*'
	printGrid(grid)
	ans := hmax - h - 1
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func getNewShape(n int, grid [][]byte) shape {
	left, floor := 3, getTowerHeight(grid)
	var sh shape
	switch n % 5 {
	case 0:
		sh = NewSDash(left, floor)
	case 1:
		sh = NewSPlus(left, floor)
	case 2:
		sh = NewSChair(left, floor)
	case 3:
		sh = NewSLine(left, floor)
	case 4:
		sh = NewSbox(left, floor)
	}
	//fmt.Printf("init shape at : %v,%v\n", left, floor)
	//sh.print()
	//sh.markOnGrid(grid, '#')
	//printGrid(grid)
	return sh
}

// brute the fuck out of it
func getTowerHeight(grid [][]byte) int {
	h := len(grid) - 1
	for y := 0; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[0])-1; x++ {
			if grid[y][x] != cellEmpty {
				return y
			}
		}
	}
	return h
}

func fallOneStep(grid [][]byte, s shape, dir byte) bool {
	if dir == '<' {
		s.tryMoveLeft()
		lockOrIgnore(grid, s, "left")
	} else {
		s.tryMoveRight()
		lockOrIgnore(grid, s, "right")
	}
	s.tryMoveDown()
	didGoDwn := lockOrIgnore(grid, s, "down")
	return didGoDwn
}

func lockOrIgnore(grid [][]byte, s shape, msg string) bool {
	if s.willCollide(grid) {
		//fmt.Println(msg + " ignore")
		s.ignrMove()
		return false
	} else {
		//fmt.Println(msg + " lock")
		s.lockMove()
		return true
	}
}

func markBounds(grid [][]byte, xbounds, ybounds []int) {
	for y, row := range grid {
		for x, _ := range row {
			if x == xbounds[0] || x == xbounds[1] {
				grid[y][x] = '|'
			}
			if y == ybounds[1] {
				grid[y][x] = '-'
			}
			if (x == xbounds[0] || x == xbounds[1]) && y == ybounds[1] {
				grid[y][x] = '+'
			}
		}
	}
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = []byte(lines[0])
}

func fallOneStepAndDisplay(grid [][]byte, sh shape, dir byte) bool {
	sh.deleteFromGrid(grid)
	didGoDown := fallOneStep(grid, sh, dir)
	sh.markOnGrid(grid, '#')
	printGrid(grid)
	return didGoDown
}

func printGrid(grid [][]byte) {
	bytez.PPrint2D(grid)
	fmt.Println()
}
