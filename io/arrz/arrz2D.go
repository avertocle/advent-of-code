package arrz

import (
	"fmt"
	//"github.com/avertocle/contests/io/clr"
	"github.com/avertocle/contests/io/cmz"
)

func Copy2D[T any](source [][]T) [][]T {
	var def T
	ans := Init2D(len(source), len(source[0]), def)
	for i := 0; i < len(source); i++ {
		for j := 0; j < len(source[0]); j++ {
			ans[i][j] = source[i][j]
		}
	}
	return ans
}

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

func Reduce2d[T any, U int | int64](arr [][]T, init U, redFunc func([][]T, U, U) U) U {
	ans := init
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			ans += redFunc(arr, U(i), U(j))
		}
	}
	return ans
}

//func Reduce2d[T any, U any](grid [][]T, f func(U, T) U, init U) U {
//}

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

func GetElementAt2D[T cmz.Primitive](grid [][]T, index []int, isInfinite bool) T {
	i, j := index[0], index[1]
	if isInfinite {
		i = index[0] % len(grid)
		j = index[1] % len(grid[0])
	}
	return grid[i][j]
}

func MarkOnNewGrid2D[T any](points []*Idx2D[int], dims *Idx2D[int], defaultVal, marker T, print bool) [][]T {
	grid := Init2D(dims.I, dims.J, defaultVal)
	MarkOnGrid2D(grid, points, marker)
	if print {
		PPrint2D(grid)
	}
	return grid
}

func MarkOnGrid2D[T any](grid [][]T, points []*Idx2D[int], marker T) {
	for _, x := range points {
		grid[x.I][x.J] = marker
	}
}

func PPrint2D[T any](arr [][]T) {
	for _, row := range arr {
		for _, cell := range row {
			val := any(cell)
			if _, ok := val.(byte); ok {
				val = fmt.Sprintf("%c", val)
			}
			fmt.Printf("%v", val)
			//			fmt.Printf("%v", clr.Gen(val, clr.Cyan))
		}
		fmt.Println()
	}
	fmt.Println()
}
