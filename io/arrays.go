package io

func Init2DByte(rows, cols int) [][]byte {
	ans := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]byte, cols)
	}
	return ans
}

func Init2DInt(rows, cols int) [][]int {
	ans := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]int, cols)
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
