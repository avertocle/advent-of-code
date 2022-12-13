package bytez

func Init2D(rows, cols int, b byte) [][]byte {
	ans := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]byte, cols)
		for j, _ := range ans[i] {
			ans[i][j] = b
		}
	}
	return ans
}

func Count2D(grid [][]byte, v byte) int {
	return CountIf2D(grid, func(b byte, i int, i2 int) bool {
		return b == v
	})
}

func CountIf2D(grid [][]byte, f func(byte, int, int) bool) int {
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

func ExtractSq2D(source [][]byte, center []int, size int, padding byte) [][]byte {
	ans := Init2D(size, size, padding)
	si, sj := center[0]-(size/2), center[1]-(size/2)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if si+i >= 0 && sj+j >= 0 && si+i < len(source) && sj+j < len(source[0]) {
				ans[i][j] = source[si+i][sj+j]
			}
		}
	}
	return ans
}

func Pad2D(arr [][]byte, rows, cols, padSize int, padVal byte) [][]byte {
	ans := Init2D(rows+2*padSize, cols+2*padSize, padVal)
	for i, row := range arr {
		for j, _ := range row {
			ans[i+padSize][j+padSize] = arr[i][j]
		}
	}
	return ans
}
