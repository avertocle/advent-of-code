package geom

import (
	"fmt"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/numz"
)

/*
returns ref for easy chaining
*/

type Coord3d struct {
	X int
	Y int
	Z int
}

func NewCoord3d(x, y, z int) *Coord3d {
	return &Coord3d{X: x, Y: y, Z: z}
}

func NewCoord3dFromVec(v []int) *Coord3d {
	return &Coord3d{X: v[0], Y: v[1], Z: v[2]}
}

func (this *Coord3d) MoveBy(vec []int) *Coord3d {
	this.X += vec[0]
	this.Y += vec[1]
	this.Z += vec[2]
	return this
}

func (this *Coord3d) Trim(bounds [][]int) *Coord3d {
	this.X = numz.Trim(this.X, bounds[0])
	this.Y = numz.Trim(this.Y, bounds[1])
	this.Z = numz.Trim(this.Z, bounds[2])
	return this
}

func (this *Coord3d) InBounds(bounds [][]int) bool {
	return intz.InBounds3D(this.Arr(), bounds)
}

func (this *Coord3d) Arr() []int {
	return []int{this.X, this.Y, this.Z}
}

func (this *Coord3d) Str() string {
	return fmt.Sprintf("%v,%v,%v", this.X, this.Y, this.Z)
}

func C3DToBounds(cb, ce *Coord3d) [][]int {
	return [][]int{
		{cb.X, ce.X},
		{cb.Y, ce.Y},
		{cb.Z, ce.Z},
	}
}
