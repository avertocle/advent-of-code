package day17

import . "github.com/avertocle/contests/io/geom"

type splus struct {
	genShape
}

func NewSPlus(left, floor int) *splus {
	sh := &splus{}
	sh.next = []*Coord2d{
		NewCoord2d(left+1, floor-6),
		NewCoord2d(left, floor-5),
		NewCoord2d(left+1, floor-5),
		NewCoord2d(left+2, floor-5),
		NewCoord2d(left+1, floor-4),
	}
	sh.curr = make([]*Coord2d, len(sh.next))
	sh.lockMove()
	return sh
}
