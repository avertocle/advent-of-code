package intz

import "sort"

func Init1D(size, val int) []int {
	ans := make([]int, size)
	if val == 0 {
		return ans
	}
	Map1D(ans, func(arr []int, i int) int {
		return val
	})
	return ans
}

func FindByVal1D(arr []int, v int) []int {
	ans := make([]int, 0)
	for i, x := range arr {
		if x == v {
			ans = append(ans, i)
		}
	}
	return ans
}

func FindMax1D(arr []int) (int, int) {
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

func FindMin1D(arr []int) (int, int) {
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

func FindMid1D(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)/2]
}

func Sum1D(arr []int) int {
	return Reduce1D(arr, func(ans int, i1d []int, i int) int {
		return ans + i1d[i]
	})
}

func Mul1D(arr []int) int {
	return Reduce1D(arr, func(ans int, arr []int, i int) int {
		return ans * arr[i]
	})
}

// f : (ans, arr, i) => ans
func Reduce1D(arr []int, f func(int, []int, int) int) int {
	if len(arr) == 0 {
		return 0
	}
	ans := arr[0]
	for i := 1; i < len(arr); i++ {
		ans = f(ans, arr, i)
	}
	return ans
}

// f : (ans, arr, i) => ans
func Map1D(arr []int, f func([]int, int) int) {
	for i := 0; i < len(arr); i++ {
		arr[0] = f(arr, i)
	}
}
