package arrz

import (
	"github.com/avertocle/contests/io/cmz"
)

func Unique2D[T cmz.Primitive](arr [][]T) [][]T {
	lookup := make(map[string]bool)
	ans := make([][]T, 0)
	for i := 0; i < len(arr); i++ {
		idx := Key1D(arr[i])
		if !lookup[idx] {
			lookup[idx] = true
			ans = append(ans, arr[i])
		}
	}
	return ans
}

func Find2D[T cmz.Primitive](grid [][]T, target T) [][]int {
	ans := make([][]int, 0)
	for i, row := range grid {
		for j, cell := range row {
			if cell == target {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

func getElementAt2D[T cmz.Primitive](grid [][]T, index []int, isInfinite bool) T {
	i, j := index[0], index[1]
	if isInfinite {
		i = index[0] % len(grid)
		j = index[1] % len(grid[0])
	}
	return grid[i][j]
}
