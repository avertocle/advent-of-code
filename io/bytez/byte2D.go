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
