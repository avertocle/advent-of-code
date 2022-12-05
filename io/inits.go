package io

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
