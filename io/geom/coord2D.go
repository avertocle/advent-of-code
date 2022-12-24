package geom

import "fmt"

type Coord2d struct {
	X int
	Y int
}

func NewCoord2d(x, y int) *Coord2d {
	return &Coord2d{X: x, Y: y}
}

func (o *Coord2d) Clone() *Coord2d {
	return &Coord2d{
		X: o.X,
		Y: o.Y,
	}
}

// todo : works only for Q4, fix this
// make generic for [][]array, tl, br
func (o *Coord2d) IsInside(boundTL, boundBR *Coord2d) bool {
	return o.X >= boundTL.X &&
		o.X <= boundBR.X &&
		o.Y <= boundTL.Y &&
		o.Y >= boundBR.Y
}

func IsValidCoord2D(x, y, rows, cols int) bool {
	return !(x < 0 || y < 0 || x >= rows || y >= cols)
}

func ApplyToAdjacent(g [][]int, x, y, rows, cols int, diag bool, f func(int) int) {
	ApplyIfValid(g, x+1, y, rows, cols, f)
	ApplyIfValid(g, x-1, y, rows, cols, f)
	ApplyIfValid(g, x, y+1, rows, cols, f)
	ApplyIfValid(g, x, y-1, rows, cols, f)
	if diag {
		ApplyIfValid(g, x-1, y-1, rows, cols, f)
		ApplyIfValid(g, x-1, y+1, rows, cols, f)
		ApplyIfValid(g, x+1, y-1, rows, cols, f)
		ApplyIfValid(g, x+1, y+1, rows, cols, f)
	}
}

func ApplyIfValid(g [][]int, x, y, rows, cols int, f func(int) int) {
	if IsValidCoord2D(x, y, rows, cols) {
		g[x][y] = f(g[x][y])
	}
}

func Unique1DIntIn2DInt(arr [][]int) int {
	m := make(map[string]bool)
	for _, row := range arr {
		m[fmt.Sprintf("%v", row)] = true
	}
	return len(m)
}
