package day18

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"math"
	"strings"
)

var gInputDirec []byte
var gInputDist []int
var gInputColor []string

const DirPath = "../2023/day18"

func SolveP1() string {
	ans := 0
	ground := digHole()
	fmt.Println(len(ground), len(ground[0]))
	floodFill(ground, []int{0, 0})
	//floodFillV2(ground)
	bytez.PPrint2D(ground)
	ans += bytez.Count2D(ground, '#')
	ans += bytez.Count2D(ground, '.')
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func digHole() [][]byte {
	l, r, u, d := calcGridDimensions()
	ground := bytez.Init2D(u+d+1, l+r+1, '.')
	spos := []int{u, l}
	l, r, u, d = math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for i := 0; i < len(gInputDirec); i++ {
		spos = tracePath(ground, spos, gInputDist[i], gInputDirec[i])
		l, r, u, d = intz.Min(l, spos[1]), intz.Max(r, spos[1]), intz.Min(u, spos[0]), intz.Max(d, spos[0])
	}
	fmt.Println(l, r, u, d)
	ground = bytez.Extract2D(ground, []int{u, l}, []int{d, r}, 'o')
	ground = bytez.Pad2D(ground, len(ground), len(ground[0]), 1, '.')
	return ground
}

func floodFill(ground [][]byte, cpos []int) {
	ground[cpos[0]][cpos[1]] = 'x'
	npos := [][]int{
		{cpos[0] - 1, cpos[1]},
		{cpos[0] + 1, cpos[1]},
		{cpos[0], cpos[1] - 1},
		{cpos[0], cpos[1] + 1},
	}
	for _, np := range npos {
		if bytez.IsValidIndex(ground, np[0], np[1]) {
			if ground[np[0]][np[1]] == '.' {
				floodFill(ground, np)
			}
		}
	}
}

func floodFillV2(ground [][]byte) {
	for i := 0; i < len(ground); i++ {
		for j := 0; j < len(ground[0]) && ground[i][j] == '.'; j++ {
			ground[i][j] = 'x'
		}
		for j := len(ground[0]) - 1; j >= 0 && ground[i][j] == '.'; j-- {
			ground[i][j] = 'x'
		}
	}
	for jj := 0; jj < len(ground[0]); jj++ {
		for ii := 0; ii < len(ground) && (ground[ii][jj] == '.' || ground[ii][jj] == 'x'); ii++ {
			ground[ii][jj] = 'x'
		}
		for ii := len(ground) - 1; ii >= 0 && (ground[ii][jj] == '.' || ground[ii][jj] == 'x'); ii-- {
			ground[ii][jj] = 'x'
		}
	}
}

func tracePath(ground [][]byte, spos []int, dist int, direc byte) []int {
	switch direc {
	case 'L':
		for d := 0; d < dist; d++ {
			ground[spos[0]][spos[1]-d] = '#'
		}
		return []int{spos[0], spos[1] - dist}

	case 'R':
		for d := 0; d < dist; d++ {
			ground[spos[0]][spos[1]+d] = '#'
		}
		return []int{spos[0], spos[1] + dist}

	case 'U':
		for d := 0; d < dist; d++ {
			ground[spos[0]-d][spos[1]] = '#'
		}
		return []int{spos[0] - dist, spos[1]}

	case 'D':
		for d := 0; d < dist; d++ {
			ground[spos[0]+d][spos[1]] = '#'
		}
		return []int{spos[0] + dist, spos[1]}
	}
	errz.HardAssert(false, "Invalid direction | %v,%v,%v", spos, dist, direc)
	return []int{-1, -1}
}

func getNextPos(pos []int, dist int, direc byte) []int {
	switch direc {
	case 'L':
		return []int{pos[0], pos[1] - dist}
	case 'R':
		return []int{pos[0], pos[1] + dist}
	case 'U':
		return []int{pos[0] - dist, pos[1]}
	case 'D':
		return []int{pos[0] + dist, pos[1]}
	}
	errz.HardAssert(false, "Invalid direction | %v,%v,%v", pos, dist, direc)
	return []int{-1, -1}
}

func calcGridDimensions() (int, int, int, int) {
	l, r, u, d := 0, 0, 0, 0
	for i := 0; i < len(gInputDirec); i++ {
		switch gInputDirec[i] {
		case 'L':
			l += gInputDist[i]
		case 'R':
			r += gInputDist[i]
		case 'U':
			u += gInputDist[i]
		case 'D':
			d += gInputDist[i]
		}
	}
	return l, r, u, d
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInputDirec, gInputDist, gInputColor = make([]byte, 0), make([]int, 0), make([]string, 0)
	for _, line := range lines {
		t := strings.Fields(line)
		gInputDirec = append(gInputDirec, []byte(t[0])[0])
		gInputDist = append(gInputDist, int([]byte(t[1])[0]-'0'))
		gInputColor = append(gInputColor, t[2])
	}
	fmt.Println(gInputDist)
}
