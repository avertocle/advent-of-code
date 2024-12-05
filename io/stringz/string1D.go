package stringz

import (
	"fmt"
	"strings"
)

func PrintWithIndex1D(arr []string) {
	for i, s := range arr {
		fmt.Printf("%2d => %v\n", i, s)
	}
}

func Has1D(arr []string, str string) bool {
	return len(Find1D(arr, str)) > 0
}

func Find1D(arr []string, str string) []int {
	ans := make([]int, 0)
	for i, s := range arr {
		if strings.Compare(s, str) == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

func FindEmpty1D(arr []string) []int {
	return Find1D(arr, "")
}
