package day21

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/errz"
	"strings"
)

const (
	A1     = 10
	Empty1 = 11

	up     = 0
	down   = 1
	left   = 2
	right  = 3
	A2     = 4
	Empty2 = 5
)

/*
+---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+
	| 0 | A |
	+---+---+

    +---+---+
    | ^ | A |
+---+---+---+
| < | v | > |
+---+---+---+
*/

var numArr = [][]int{{7, 8, 9}, {4, 5, 6}, {1, 2, 3}, {Empty1, 0, A1}}
var dirArr = [][]int{{Empty2, up, A2}, {left, down, right}}

func makeManDisMaps() (map[string][]string, map[string][]string) {
	manDisNum := make(map[string][]string)
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			a1 := arrz.Find2D(numArr, i)[0]
			a2 := arrz.Find2D(numArr, j)[0]
			//fmt.Println(i, j, a1, a2, numz.Abs[int](a2[0]-a1[0])+numz.Abs[int](a2[1]-a1[1]))
			dx, dy := a2[0]-a1[0], a2[1]-a1[1]
			seq := ""
			if dy > 0 {
				seq += strings.Repeat(">", dy)
			}
			if dx < 0 {
				seq += strings.Repeat("^", -dx)
			}
			if dx > 0 {
				seq += strings.Repeat("v", dx)
			}
			if dy < 0 {
				seq += strings.Repeat("<", -dy)
			}
			ii, jj := i+'0', j+'0'
			if i == 10 {
				ii = 'A'
			}
			if j == 10 {
				jj = 'A'
			}
			allSeq := filterDuplicates(generateAllPerms(seq))
			allSeq = filterAllPermsByEmptyPosTouch(allSeq, i, j, numArr)
			manDisNum[toKey(byte(ii), byte(jj))] = allSeq
		}
	}
	manDisDir := make(map[string][]string)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			a1 := arrz.Find2D(dirArr, i)[0]
			a2 := arrz.Find2D(dirArr, j)[0]
			//fmt.Println(i, j, a1, a2, numz.Abs[int](a2[0]-a1[0])+numz.Abs[int](a2[1]-a1[1]))
			dx, dy := a2[0]-a1[0], a2[1]-a1[1]
			seq := ""
			if dy > 0 {
				seq += strings.Repeat(">", dy)
			}
			if dx > 0 {
				seq += strings.Repeat("v", dx)
			}
			if dy < 0 {
				seq += strings.Repeat("<", -dy)
			}
			if dx < 0 {
				seq += strings.Repeat("^", -dx)
			}
			allSeq := filterDuplicates(generateAllPerms(seq))
			allSeq = filterAllPermsByEmptyPosTouch(allSeq, i, j, dirArr)
			m := map[int]byte{up: '^', down: 'v', left: '<', right: '>', A2: 'A', Empty2: 'E'}
			manDisDir[toKey(m[i], m[j])] = allSeq

		}
	}
	return manDisNum, manDisDir
}

func filterDuplicates(arr []string) []string {
	m := make(map[string]bool)
	ans := make([]string, 0)
	for _, a := range arr {
		if _, ok := m[a]; !ok {
			m[a] = true
			ans = append(ans, a)
		}
	}
	return ans
}

func generateAllPerms(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}
	perms := make([]string, 0)
	for i, c := range str {
		subPerms := generateAllPerms(str[0:i] + str[i+1:])
		for _, p := range subPerms {
			perms = append(perms, string(c)+p)
		}
	}
	return perms
}

func filterAllPermsByEmptyPosTouch(perms []string, sChar, eChar int, arr [][]int) []string {
	ans := make([]string, 0)
	for _, p := range perms {
		if filterOnePermByEmptyPosTouch(p, sChar, eChar, arr) {
			ans = append(ans, p)
		}
	}
	return ans
}

func filterOnePermByEmptyPosTouch(str string, sChar, eChar int, arr [][]int) bool {
	s := arrz.Find2D(arr, sChar)[0]
	for _, c := range str {
		if (len(arr) == 4 && arr[s[0]][s[1]] == Empty1) || (len(arr) == 2 && arr[s[0]][s[1]] == Empty2) {
			fmt.Println("skipped : ", str, sChar, eChar, arr, s)
			return false
		}
		switch c {
		case '^':
			s = []int{s[0] - 1, s[1]}
		case 'v':
			s = []int{s[0] + 1, s[1]}
		case '<':
			s = []int{s[0], s[1] - 1}
		case '>':
			s = []int{s[0], s[1] + 1}
		case A1, A2, 'A':
			errz.SoftAssert(arr[s[0]][s[1]] == eChar, "end char not reached on A : char(%v), idx(%v), echar(%v)", c, s, eChar)
		}
	}
	return true
}
