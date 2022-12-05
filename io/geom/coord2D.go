package geom

import "fmt"

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
