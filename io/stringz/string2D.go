package stringz

func Init2D(rows, cols int) [][]string {
	ans := make([][]string, rows)
	for i, _ := range ans {
		ans[i] = make([]string, cols)
	}
	return ans
}

func Fill2D(s2d [][]string, val string) {
	for i, s1d := range s2d {
		for j, _ := range s1d {
			s2d[i][j] = val
		}
	}
}

func GetTranspose2D(s2d [][]string, rows, cols int) [][]string {
	ans := Init2D(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ans[i][j] = s2d[j][i]
		}
	}
	return ans
}
