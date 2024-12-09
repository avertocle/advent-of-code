package arrz

import (
	"fmt"
	"github.com/avertocle/contests/io/clr"
	"github.com/avertocle/contests/io/cmz"
)

func Count2D[T cmz.PrimitivePlus](grid [][]T, v T) int {
	return CountIf2D(grid, func(b T, i int, j int) bool {
		return b == v
	})
}

func CountIf2D[T any](arr [][]T, f func(T, int, int) bool) int {
	count := 0
	for i, row := range arr {
		for j, cell := range row {
			if f(cell, i, j) {
				count++
			}
		}
	}
	return count
}

func Init2D[T any](rows, cols int, b T) [][]T {
	ans := make([][]T, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]T, cols)
		for j, _ := range ans[i] {
			ans[i][j] = b
		}
	}
	return ans
}

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

func PPrint2D[T cmz.PrimitivePlus](arr [][]T) {
	for _, row := range arr {
		for _, cell := range row {
			fmt.Printf("(%v) | ", clr.Gen(cell, clr.Cyan))
		}
		fmt.Println()
	}
	fmt.Println()
}
