package numz

import (
	"fmt"
	"github.com/avertocle/contests/io/cmz"
	"math"
)

func Max[T cmz.Number](x, y T) T {
	if x >= y {
		return x
	} else {
		return y
	}
}

func Min[T cmz.Number](x, y T) T {
	if x <= y {
		return x
	} else {
		return y
	}
}

func Abs[T cmz.Number](x T) T {
	if x >= 0 {
		return x
	} else {
		return -1 * x
	}
}

func IsBounded[T cmz.Number](x, s, e T) bool {
	return x >= s && x <= e
}

func Trim[T cmz.Number](c T, x []T) T {
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

func Pow[T cmz.Number](x, n T) int64 {
	return int64(math.Pow(float64(x), float64(n)))
}

/*
IncBounded
returns x incremented and rotated (if) to be between 1 & max
*/
func IncBounded[T int | int64](x, inc, max T) T {
	inc %= max
	x += inc
	if x > max {
		x = x - max
	}
	return x
}
