package day17

import . "github.com/avertocle/contests/io/geom"

type sline struct {
	genShape
}

func NewSLine(left, floor int) *sline {
	sh := &sline{}
	sh.next = []*Coord2d{
		NewCoord2d(left, floor-7),
		NewCoord2d(left, floor-6),
		NewCoord2d(left, floor-5),
		NewCoord2d(left, floor-4),
	}
	sh.curr = make([]*Coord2d, len(sh.next))
	sh.lockMove()
	return sh
}
