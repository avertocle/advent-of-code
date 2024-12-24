package day21

import (
	"github.com/avertocle/contests/io/arrz"
	"strings"
)

func makeManDisMaps() (map[string]string, map[string]string) {
	numArr := [][]int{{7, 8, 9}, {4, 5, 6}, {1, 2, 3}, {Empty1, 0, A1}}
	manDisNum := make(map[string]string)
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
			manDisNum[toKey(byte(ii), byte(jj))] = seq
		}
	}
	dirArr := [][]int{{Empty2, up, A2}, {left, down, right}}
	manDisDir := make(map[string]string)
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
			m := map[int]byte{up: '^', down: 'v', left: '<', right: '>', A2: 'A', Empty2: 'E'}
			manDisDir[toKey(m[i], m[j])] = seq
		}
	}
	return manDisNum, manDisDir
}
