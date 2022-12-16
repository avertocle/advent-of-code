package bytez

import (
	"strings"
)

func FindSubseq1D(arr, pat []byte) []int {
	lenA := len(arr)
	lenP := len(pat)
	if lenA == 0 || lenP == 0 || lenA < lenP {
		return []int{}
	}
	if lenA == lenP && strings.Compare(string(arr), string(pat)) == 0 {
		return []int{0}
	}

	ans := make([]int, lenA)
	match := false
	k := 0
	for i := 0; i < lenA; i++ {
		if arr[i] == pat[0] && lenA-i >= lenP {
			match = true
			for j := 0; j < lenP; j++ {
				if arr[i+j] != pat[j] {
					match = false
				}
			}
			if match {
				ans[k] = i
				k++
			}
		}
	}
	return ans[0:k]
}

func GroupUniq1D(arr []byte) map[byte]int {
	m := make(map[byte]int)
	for _, b := range arr {
		if v, ok := m[b]; ok {
			m[b] = v + 1
		} else {
			m[b] = 1
		}
	}
	return m
}

func FindFirst(arr []byte, b byte) int {
	for i, c := range arr {
		if c == b {
			return i
		}
	}
	return -1
}

/*
FindNestedMatch
arr[0] must be the start char
return index of closing char in a seq of nested chars
*/
func FindNestedMatch(arr []byte, endByte byte) int {
	ctr := 0
	begByte := arr[0]
	for i, b := range arr {
		if b == begByte {
			ctr++
		} else if b == endByte {
			ctr--
		}
		if ctr == 0 {
			return i
		}
	}
	return -1
}
