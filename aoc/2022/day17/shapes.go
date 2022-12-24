package day17

import (
	"fmt"
	. "github.com/avertocle/contests/io/geom"
)

const cellEmpty = '.'

// assumes positive coordinates only
type shape interface {
	markOnGrid([][]byte, byte)
	deleteFromGrid([][]byte)
	tryMoveLeft()
	tryMoveRight()
	tryMoveDown()
	willCollide([][]byte) bool
	getTop() []int
	lockMove()
	ignrMove()
	print()
}

type genShape struct {
	curr []*Coord2d
	next []*Coord2d
}

func (o *genShape) deleteFromGrid(grid [][]byte) {
	o.markOnGrid(grid, cellEmpty)
}

func (o *genShape) markOnGrid(grid [][]byte, b byte) {
	for _, c := range o.curr {
		grid[c.Y][c.X] = b
	}
}

func (o *genShape) getTop() []int {
	return []int{o.curr[0].X, o.curr[0].Y}
}

func (o *genShape) tryMoveLeft() {
	for _, c := range o.next {
		c.X -= 1
	}
}

func (o *genShape) tryMoveRight() {
	for _, c := range o.next {
		c.X += 1
	}
}

func (o *genShape) tryMoveDown() {
	for _, c := range o.next {
		c.Y += 1
	}
}

func (o *genShape) willCollide(grid [][]byte) bool {
	for _, c := range o.next {
		if grid[c.Y][c.X] != cellEmpty {
			//fmt.Printf("%v", string(grid[c.Y][c.X]))
			return true
		}
	}
	return false
}

func (o *genShape) lockMove() {
	for i, c := range o.next {
		o.curr[i] = c.Clone()
	}
}

func (o *genShape) ignrMove() {
	for i, c := range o.curr {
		o.next[i] = c.Clone()
	}
}

func (o *genShape) print() {
	fmt.Printf("curr = ")
	for _, c := range o.curr {
		fmt.Printf("(%v,%v)  ", c.X, c.Y)
	}
	fmt.Println()
	fmt.Printf("next = ")
	for _, c := range o.next {
		fmt.Printf("(%v,%v)  ", c.X, c.Y)
	}
	fmt.Println()
}
