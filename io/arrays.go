package io

import (
	"sort"
	"strings"
)

func Init2DByte(rows, cols int, b byte) [][]byte {
	ans := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]byte, cols)
		for j, _ := range ans[i] {
			ans[i][j] = b
		}
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

func Find1DByteIn1DByte(arr, pat []byte) []int {
	lenA := len(arr)
	lenP := len(pat)
	if lenA == 0 || lenP == 0 || lenA < lenP {
		return []int{}
	}
	if lenA == lenP && strings.Compare(string(arr), string(pat)) == 0 {
		return []int{0}
	}

	ans := make([]int, lenA)
	match := false
	k := 0
	for i := 0; i < lenA; i++ {
		if arr[i] == pat[0] && lenA-i >= lenP {
			match = true
			for j := 0; j < lenP; j++ {
				if arr[i+j] != pat[j] {
					match = false
				}
			}
			if match {
				ans[k] = i
				k++
			}
		}
	}
	return ans[0:k]
}

func CountUniqByteIn1DByte(arr []byte) map[byte]int {
	m := make(map[byte]int)
	for _, b := range arr {
		if v, ok := m[b]; ok {
			m[b] = v + 1
		} else {
			m[b] = 1
		}
	}
	return m
}
