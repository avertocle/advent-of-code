package day17

import . "github.com/avertocle/contests/io/geom"

type sdash struct {
	genShape
}

func NewSDash(left, bottom int) *sdash {
	sh := &sdash{}
	sh.next = []*Coord2d{
		NewCoord2d(left, bottom-4),
		NewCoord2d(left+1, bottom-4),
		NewCoord2d(left+2, bottom-4),
		NewCoord2d(left+3, bottom-4),
	}
	sh.curr = make([]*Coord2d, len(sh.next))
	sh.lockMove()
	return sh
}
