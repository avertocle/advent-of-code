package intz

import "sort"

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
