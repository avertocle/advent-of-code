package day18

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"math"
	"strconv"
)

var gInputDirec []byte
var gInputDist []int
var gInputColor []string

const DirPath = "../2023/day18"

func SolveP1() string {
	//ans := 0
	//ground := digHole()
	//floodFill(ground, []int{0, 0})
	//ans += bytez.Count2D(ground, '#')
	//ans += bytez.Count2D(ground, '.')
	//return fmt.Sprintf("%v", ans)
	ans := int64(0)
	holeVertices := digHoleV2()
	fmt.Println(holeVertices)
	fmt.Println(len(holeVertices))
	area := findPolygonArea(holeVertices)
	circ := int64(findPolygonCircumference(holeVertices))
	verts := int64(len(holeVertices))
	fmt.Println(area, circ, verts)
	ans = area + circ/2 + 1
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := int64(0)
	// 0 means R, 1 means D, 2 means L, and 3 means U.
	dirInts := []byte{'R', 'D', 'L', 'U'}
	for i, c := range gInputColor {
		x, _ := strconv.ParseInt(string([]byte(c)[2:7]), 16, 32)
		gInputDist[i] = int(x)
		gInputDirec[i] = dirInts[[]byte(c)[7]-'0']
		//fmt.Printf("%v %v %v\n", string([]byte(c)[2:7]), gInputDist[i], string(gInputDirec[i]))
	}
	holeVertices := digHoleV2()
	fmt.Println(holeVertices)
	fmt.Println(len(holeVertices))
	area := findPolygonArea(holeVertices)
	circ := int64(findPolygonCircumference(holeVertices))
	verts := int64(len(holeVertices))
	fmt.Println(area, circ, verts)
	ans = area + circ/2 + 1
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func digHoleV2() [][]int {
	spos := []int{0, 0}
	holeVertices := [][]int{spos}
	l, r, u, d := math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for i := 0; i < len(gInputDirec); i++ {
		spos = tracePathV2(spos, gInputDist[i], gInputDirec[i])
		holeVertices = append(holeVertices, spos)
		l, r, u, d = intz.Min(l, spos[1]), intz.Max(r, spos[1]), intz.Min(u, spos[0]), intz.Max(d, spos[0])
	}
	for i := 0; i < len(holeVertices); i++ {
		holeVertices[i][0] -= u
		holeVertices[i][1] -= l
	}
	return holeVertices
}

func digHole() [][]byte {
	l, r, u, d := calcGridDimensions()
	ground := bytez.Init2D(u+d+1, l+r+1, '.')
	spos := []int{u, l}
	l, r, u, d = math.MaxInt, math.MinInt, math.MaxInt, math.MinInt
	for i := 0; i < len(gInputDirec); i++ {
		spos = tracePath(ground, spos, gInputDist[i], gInputDirec[i])
		l, r, u, d = intz.Min(l, spos[1]), intz.Max(r, spos[1]), intz.Min(u, spos[0]), intz.Max(d, spos[0])
	}
	ground = bytez.Extract2D(ground, []int{u, l}, []int{d, r}, 'o')
	ground = bytez.Pad2D(ground, len(ground), len(ground[0]), 1, '.')
	return ground
}
func findPolygonArea(vertices [][]int) int64 {
	//Let 'vertices' be an array of N pairs (x,y), indexed from 0
	//Let 'area' = 0.0
	//for i = 0 to N-1, do
	//Let j = (i+1) mod N
	//Let area = area + vertices[i].x * vertices[j].y
	//Let area = area - vertices[i].y * vertices[j].x
	//end for
	//	Return 'area'
	area := int64(0)
	for i := 0; i < len(vertices)-1; i++ {
		j := i + 1
		area += int64(vertices[i][1])*int64(vertices[j][0]) - int64(vertices[i][0])*int64(vertices[j][1])
	}
	return area / 2
}

func findPolygonCircumference(vertices [][]int) int {
	circumference := 0
	for i := 0; i < len(vertices)-1; i++ {
		j := (i + 1) % (len(vertices) - 1)
		if vertices[i][0] == vertices[j][0] {
			circumference += intz.Abs(vertices[i][1] - vertices[j][1])
		} else {
			circumference += intz.Abs(vertices[i][0] - vertices[j][0])
		}
	}
	return circumference
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

//
//func floodFillV2(holeVertices [][]int, cpos []int) {
//	ground[cpos[0]][cpos[1]] = 'x'
//	npos := [][]int{
//		{cpos[0] - 1, cpos[1]},
//		{cpos[0] + 1, cpos[1]},
//		{cpos[0], cpos[1] - 1},
//		{cpos[0], cpos[1] + 1},
//	}
//	for _, np := range npos {
//		if bytez.IsValidIndex(ground, np[0], np[1]) {
//			if ground[np[0]][np[1]] == '.' {
//				floodFill(ground, np)
//			}
//		}
//	}
//}

func tracePathV2(spos []int, dist int, direc byte) []int {
	switch direc {
	case 'L':
		return []int{spos[0], spos[1] - dist}
	case 'R':
		return []int{spos[0], spos[1] + dist}
	case 'U':
		return []int{spos[0] - dist, spos[1]}
	case 'D':
		return []int{spos[0] + dist, spos[1]}
	}
	errz.HardAssert(false, "Invalid direction | %v,%v,%v", spos, dist, direc)
	return []int{-1, -1}
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
	gInputDirec = iutils.ExtractByte1DFromString1D(lines, " ", 0, 0)
	gInputDist = iutils.ExtractInt1DFromString1D(lines, " ", 1, -1)
	gInputColor = iutils.ExtractString1DFromString1D(lines, " ", 2, "")
}
