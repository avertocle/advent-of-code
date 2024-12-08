package geom

import (
	"fmt"
	"github.com/avertocle/contests/io/cmz"
	"math"
)

type Coord2D[T cmz.Number] struct {
	X T
	Y T
}

func (o *Coord2D[T]) IsInside(boundTL, boundBR *Coord2D[T]) bool {
	return o.X >= boundTL.X &&
		o.X <= boundBR.X &&
		o.Y <= boundTL.Y &&
		o.Y >= boundBR.Y
}

func (o *Coord2D[T]) MoveBy(dx, dy T) *Coord2D[T] {
	o.X += dx
	o.Y += dy
	return o
}

func (o *Coord2D[T]) Str() string {
	return fmt.Sprintf("%v,%v", o.X, o.Y)
}

func (o *Coord2D[T]) IsEqual(o1 *Coord2D[T]) bool {
	return o.X == o1.X && o.Y == o1.Y

}

func PPrintCoord2D[T cmz.Number](coords []*Coord2D[T]) {
	for _, c := range coords {
		fmt.Println(c.Str())
	}
}

func NewCoord2D[T cmz.Number](x, y T) *Coord2D[T] {
	return &Coord2D[T]{X: x, Y: y}
}

type Line2D[T cmz.Number] struct {
	m T
	c T
}

func NewLine2D[T cmz.Number](cd, vel *Coord2D[T]) *Line2D[T] {
	m := (vel.Y) / (vel.X)
	c := (cd.Y) - m*(cd.X)
	//fmt.Println(m, c)
	return &Line2D[T]{m: m, c: c}
}

func LineIntersect2D[T cmz.Number](l1, l2 *Line2D[T]) *Coord2D[T] {
	x := (l2.c - l1.c) / (l1.m - l2.m)
	y := l1.m*x + l1.c
	return NewCoord2D[T](x, y)
}

func Dist2D[T cmz.Number](c1, c2 *Coord2D[T]) float64 {
	d2 := (c1.X-c2.X)*(c1.X-c2.X) + (c1.Y-c2.Y)*(c1.Y-c2.Y)
	return math.Sqrt(float64(d2))
}
