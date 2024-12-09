package mathz

import (
	"github.com/avertocle/contests/io/numz"
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int) int {
	return int((int64(a) * int64(b)) / int64(GCD(a, b)))
}

func LCMArr(arr []int) int {
	ans := arr[0]
	for i := 1; i < len(arr); i++ {
		ans = LCM(ans, arr[i])
	}
	return ans
}

func GCDArr(arr []int) int {
	ans := 1
	for i := 1; i < len(arr); i++ {
		ans = numz.Min(ans, GCD(ans, arr[i]))
	}
	return ans
}
