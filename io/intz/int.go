package intz

import (
	"fmt"
	"math"
)

func Max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func Min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func Abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -1 * x
	}
}

func IsBounded(x, s, e int) bool {
	return x >= s && x <= e
}

func Trim(c int, x []int) int {
	if IsBounded(c, x[0], x[1]) {
		return c
	} else if c < x[0] {
		return x[0]
	} else if c > x[1] {
		return x[1]
	}
	fmt.Printf("trim : error %v in %v", c, x)
	return math.MaxInt
}

func IncBounded(x, inc, max int) int {
	inc %= max
	x += inc
	if x > max {
		x = x - max
	}
	return x
}
