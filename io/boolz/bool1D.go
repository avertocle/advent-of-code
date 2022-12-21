package boolz

func Init1D(size int, val bool) []bool {
	ans := make([]bool, size)
	if val == false {
		return ans
	}
	Map1D(ans, func(arr []bool, i int) bool {
		return val
	})
	return ans
}

func Init2D(rows, cols int, val bool) [][]bool {
	ans := make([][]bool, rows)
	for i, _ := range ans {
		ans[i] = Init1D(cols, val)
	}
	return ans
}

func Or1D(arr []bool) bool {
	return Reduce1D(arr, func(ans bool, b1d []bool, i int) bool {
		return ans || b1d[i]
	})
}

func And1D(arr []bool) bool {
	return Reduce1D(arr, func(ans bool, b1d []bool, i int) bool {
		return ans && b1d[i]
	})
}

func Count1D(arr []bool, val bool) int {
	count := 0
	for _, b := range arr {
		if b == val {
			count++
		}
	}
	return count
}

// f : (ans, arr, i) => ans
func Reduce1D(arr []bool, f func(bool, []bool, int) bool) bool {
	if len(arr) == 0 {
		return false
	}
	ans := arr[0]
	for i := 1; i < len(arr); i++ {
		ans = f(ans, arr, i)
	}
	return ans
}

// f : (ans, arr, i) => ans
func Map1D(arr []bool, f func([]bool, int) bool) {
	for i := 0; i < len(arr); i++ {
		arr[0] = f(arr, i)
	}
}
