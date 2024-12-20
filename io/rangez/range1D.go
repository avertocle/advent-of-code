package rangez

import (
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/numz"
)

/*
Intersec1D
r1, r2 : []int{startIdx, endIdx} both inclusive
empty range : []int{}
*/
func Intersec1D(r1, r2 []int) []int {
	if len(r1) == 0 || len(r2) == 0 {
		return []int{}
	}
	s, e := numz.Max(r1[0], r2[0]), numz.Min(r1[1], r2[1])
	if s > e {
		return []int{}
	} else {
		return []int{s, e}
	}
}

/*
Union1D
r1 is sorted and kept sorted always
// may be buggy cos of slice issues, be careful and test
*/
func Union1D(big [][]int, small []int) [][]int {
	if len(big) == 0 {
		return [][]int{small}
	} else if len(small) == 0 {
		return big
	}
	sp, so := GetPlace(big, small[0])
	ep, eo := GetPlace(big, small[1])
	errz.HardAssert(sp >= 0 && so > -2, "get-place-err for %v : %v,%v", small, sp, so, big)
	errz.HardAssert(ep >= 0 && eo > -2, "get-place-err for %v : %v,%v", small, ep, eo, big)

	//fmt.Printf("%v,%v,%v,%v\n", sp, so, ep, eo)
	// out-left
	if ep == 0 && eo == -1 {
		return append([][]int{small}, big...)
	}

	// out-right
	if sp == len(big)-1 && so == 1 {
		return append(big, small)
	}

	// contained in one
	if sp == ep && so == 0 && eo == 0 {
		return big
	}

	// contains all
	if sp == 0 && so == -1 && ep == len(big)-1 && eo == 1 {
		return [][]int{small}
	}

	// spans some partially
	if so == 0 && eo == 0 {
		z := make([][]int, 0)
		z = append(z, big[0:sp]...)
		z = append(z, []int{big[sp][0], big[ep][1]})
		z = append(z, big[ep+1:]...)
		return z
	}

	// contains some fully end
	if so == -1 && eo == -1 {
		z := make([][]int, 0)
		z = append(z, big[0:sp]...)
		z = append(z, small)
		z = append(z, big[ep:]...)
		return z
	}

	// contains some fully till end
	if so == -1 && ep == len(big)-1 && eo == 1 {
		z := append(big[0:sp], small)
		return z
	}

	// start in, end out
	if so == 0 && eo == -1 {
		z := make([][]int, 0)
		z = append(z, big[0:sp]...)
		z = append(z, []int{big[sp][0], small[1]})
		z = append(z, big[ep:]...)
		return z
	}

	// start in, end at end
	if so == 0 && eo == 1 {
		z := make([][]int, 0)
		z = append(z, big[0:sp]...)
		z = append(z, []int{big[sp][0], small[1]})
		return z
	}

	// start out, end in
	if so == -1 && eo == 0 {
		z := make([][]int, 0)
		z = append(z, big[0:sp]...)
		z = append(z, []int{small[0], big[ep][1]})
		z = append(z, big[ep+1:]...)
		return z
	}

	errz.HardAssert(false, "Union1D err %v, %v, %v, %v", sp, so, ep, eo)

	return [][]int{}
}

/*
GetPlace
returns tightest index, placement : -1, 0, 1 (left, in, out)
*/
func GetPlace(r [][]int, x int) (int, int) {
	for i := 0; i < len(r); i++ {
		if IsOutLeft(r[i], x) {
			return i, -1
		} else if IsInside(r[i], x) {
			return i, 0
		} else if i == len(r)-1 && IsOutRight(r[len(r)-1], x) {
			return i, 1
		}
	}
	intz.PPrint2D(r)
	errz.HardAssert(false, "%v\n", x)
	return -2, -2
}

/*
GetOrientation
returns 0,1,2,3,4,5 out-left, intersect-left, contained, intersect-right, out-right, contains
*/
func GetOrientation1(r []int, r1 []int) int {
	s, e := r1[0], r1[1]
	if IsOutLeft(r, e) { // out-left
		return 0
	} else if IsOutLeft(r, s) && IsInside(r, e) { // intersect-left
		return 1
	} else if IsInside(r, s) && IsInside(r, e) { // contained
		return 2
	} else if IsInside(r, s) && IsOutRight(r, e) { // intersect-right
		return 3
	} else if IsOutRight(r, s) { // out-left
		return 4
	} else if IsOutLeft(r, s) && IsOutRight(r, e) {
		return 5
	} else {
		return 6
	}
}

func IsInside(r []int, x int) bool {
	return x >= r[0] && x <= r[1]
}

func IsOutLeft(r []int, x int) bool {
	return x < r[0]
}

func IsOutRight(r []int, x int) bool {
	return x > r[1]
}
