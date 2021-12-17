package io

import "sort"

func Init2DByte(rows, cols int) [][]byte {
	ans := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]byte, cols)
	}
	return ans
}

func Init2DInt(val, rows, cols int) [][]int {
	ans := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			ans[i][j] = val
		}
	}
	return ans
}

func Init2DString(rows, cols int) [][]string {
	ans := make([][]string, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]string, cols)
	}
	return ans
}

func MaxInt(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func MinInt(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func CountIn2dByteIf(grid [][]byte, f func(byte, int, int) bool) int {
	count := 0
	for i, row := range grid {
		for j, cell := range row {
			if f(cell, i, j) {
				count++
			}
		}
	}
	return count
}

func FindMax1DInt(arr []int) (int, int) {
	max := arr[0]
	pos := 0
	for i, x := range arr {
		if x > max {
			max = x
			pos = i
		}
	}
	return max, pos
}

func FindMin1DInt(arr []int) (int, int) {
	max := arr[0]
	pos := 0
	for i, x := range arr {
		if x < max {
			max = x
			pos = i
		}
	}
	return max, pos
}

func FindMiddleElem1DInt(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)/2]
}
