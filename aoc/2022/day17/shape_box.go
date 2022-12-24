package day17

import . "github.com/avertocle/contests/io/geom"

type sbox struct {
	genShape
}

func NewSbox(left, floor int) *sbox {
	sh := &sbox{}
	sh.next = []*Coord2d{
		NewCoord2d(left, floor-5),
		NewCoord2d(left+1, floor-5),
		NewCoord2d(left, floor-4),
		NewCoord2d(left+1, floor-4),
	}
	sh.curr = make([]*Coord2d, len(sh.next))
	sh.lockMove()
	return sh
}
