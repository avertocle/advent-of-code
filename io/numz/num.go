package numz

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/tpz"
	"math"
)

func Max[T tpz.Number](x, y T) T {
	if x >= y {
		return x
	} else {
		return y
	}
}

func Min[T tpz.Number](x, y T) T {
	if x <= y {
		return x
	} else {
		return y
	}
}

func Abs[T tpz.Number](x T) T {
	if x >= 0 {
		return x
	} else {
		return -1 * x
	}
}

func IsBounded[T tpz.Number](x, s, e T) bool {
	return x >= s && x <= e
}

func Trim[T tpz.Number](c T, x []T) T {
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

func Pow[T tpz.Number](x, n T) int64 {
	return int64(math.Pow(float64(x), float64(n)))
}

/*
IncBounded
@deprecated as it does not handle negative increments use IncBoundedV2 instead
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

func IncBoundedV2[T int | int64](x, inc, min, max T) T {
	errz.SoftAssert(x >= min && x <= max, "IncBoundedV2 : input outside bounds : %v + %v [%v to %v]", x, inc, min, max)
	xo := x
	bSize := max - min + 1
	inc = (Abs(inc) % bSize) * (inc / Abs(inc)) // reduce inc keeping the sign
	x += inc
	if x < min {
		x = max - (min - x) + 1
	}
	if x > max {
		x = min + (x - max) - 1
	}
	errz.HardAssert(x >= min && x <= max, "IncBoundedV2 : output outside bounds : %v + %v [%v to %v] = %v", xo, inc, min, max, x)
	//fmt.Printf("IncBoundedV2 : %v + %v [%v to %v] = %v\n", xo, inc, min, max, x)
	return x
}
