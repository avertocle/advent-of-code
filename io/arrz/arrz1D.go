package arrz

import (
	"fmt"
	"github.com/avertocle/contests/io/clr"
	"github.com/avertocle/contests/io/mapz"
	"github.com/avertocle/contests/io/tpz"
)

func Key1D[T tpz.Primitive](keys []T) string {
	return fmt.Sprintf("%v", keys)
}

func RemoveElement1D[T any](arr []T, index int) []T {
	return append(append([]T{}, arr[:index]...), arr[index+1:]...)
}

func RemoveElementByVal1D[T tpz.PrimitivePlus](arr []T, val T) []T {
	ans := make([]T, 0)
	for _, v := range arr {
		if v != val {
			ans = append(ans, v)
		}
	}
	return ans
}

func FindMid1D[T any](arr []T) T {
	return arr[len(arr)/2]
}

func PPrint1D[T tpz.PrimitivePlus](arr []T, withIndex bool) {
	for i, cell := range arr {
		if withIndex {
			fmt.Printf("(%v:%v) | ", clr.Int(i, clr.Cyan), clr.Gen(cell, clr.Cyan))
		} else {
			fmt.Printf("(%v) | ", clr.Gen(cell, clr.Cyan))
		}
	}
	fmt.Println()
}

func ToStr1D[T any](arr []T, sep string) string {
	ans := ""
	for _, cell := range arr {
		ans += fmt.Sprintf("%v%v", sep, cell)
	}
	return ans[1:]
}

func Join1D[T tpz.PrimitivePlus](arrays ...[]T) []T {
	ans := make([]T, 0)
	for _, arr := range arrays {
		ans = append(ans, arr...)
	}
	return ans
}

func Upscale1D[T interface{}](arr []T) [][]T {
	ans := make([][]T, len(arr))
	for i, x := range arr {
		ans[i] = []T{x}
	}
	return ans
}

func SwapRangesInPlace1D[T any](arr []T, srcRange, dstRange []int) {
	for i := 0; i <= (srcRange[1] - srcRange[0]); i++ {
		arr[srcRange[0]+i], arr[dstRange[0]+i] = arr[dstRange[0]+i], arr[srcRange[0]+i]
	}
}

// FindByVal1D : Find 'count' indices of a value in a 1D array,
// to find all, send count = len(arr),
// boundIndex = nil searches the entire array
func FindByVal1D[T tpz.PrimitivePlus](arr []T, v T, boundIndex []int, maxCount int) []int {
	ans := make([]int, 0)
	foundCtr := 0
	if boundIndex == nil {
		boundIndex = []int{0, len(arr)}
	}
	for i := boundIndex[0]; i < boundIndex[1]; i++ {
		if arr[i] == v && foundCtr < maxCount {
			ans = append(ans, i)
			foundCtr++
		}
	}
	return ans
}

// FindRepeatedByVal1D : Finds repeated chunk of an element in a 1D array,
// returns indices
// e.g ("aabbbaabbababbbbaaa, b) finds bbb, bb, b, bbbb upto count instances
// to find all, send count = len(arr),
// search happens within bounds, bounds = nil searches the entire array
func FindRepeatedByVal1D[T tpz.PrimitivePlus](arr []T, v T, bounds []int, minLength, maxCount int) [][]int {
	ans := make([][]int, 0)
	if bounds == nil {
		bounds = []int{0, len(arr)}
	}
	foundCtr := 0
	for i := bounds[0]; i < bounds[1] && foundCtr < maxCount; i++ {
		if arr[i] == v {
			start := i
			for ; i < bounds[1] && arr[i] == v; i++ {
			}
			if i-start >= minLength {
				ans = append(ans, []int{start, i - 1})
				foundCtr++
			}
		}
	}
	return ans
}

func PushToSet[T tpz.PrimitivePlus](set tpz.Set[T], arr []T) {
	for _, v := range arr {
		set[v] = true
	}
}

//func Union1D[T tpz.PrimitivePlus](arr1, arr2 []T) []T {
//	m := make(tpz.Set[T])
//	PushToSet(m, arr1)
//	PushToSet(m, arr2)
//	return mapz.Keys(m)
//}

func Union1D[T tpz.PrimitivePlus](arrs ...[]T) []T {
	m := make(tpz.Set[T])
	for _, arr := range arrs {
		PushToSet(m, arr)
	}
	return mapz.Keys(m)
}

func Intersection1D[T tpz.PrimitivePlus](arrs ...[]T) []T {
	m := make(tpz.Set[T])
	PushToSet(m, arrs[0])
	for i := 1; i < len(arrs); i++ {
		m2 := make(tpz.Set[T])
		for _, e := range arrs[i] {
			if _, ok := m[e]; ok {
				m2[e] = true
			}
		}
		m = m2
	}
	return mapz.Keys(m)
}
