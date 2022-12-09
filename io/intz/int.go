package intz

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
