package day23

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/tpz"
)

var DirPath = "../2023/day23"
var gInput [][]byte
var gStart, gEnd *arrz.Idx2D

//2320 too-high
// 2320 too-high

func SolveP1() string {
	ans := 0
	player := gStart
	maxPathLen := new(int)
	hikeSim(player, make(tpz.StringSet), 0, maxPathLen)
	ans = *maxPathLen - 1
	return fmt.Sprintf("%v", ans)
}

func hikeSim(c *arrz.Idx2D, visited tpz.StringSet, pathLen int, maxPathLen *int) {
	visited[c.ToKey()] = true
	pathLen++
	if c.IsEqual(gEnd) {
		if pathLen > *maxPathLen {
			*maxPathLen = pathLen
			//grid := bytez.Copy2D(gInput)
			fmt.Printf("%v.%v.%v\n", c.Str(), pathLen, *maxPathLen)
			//for k, v := range visited {
			//	if v {
			//		ij := arrz.NewIdx2DFromKey(k)
			//		grid[ij.I][ij.J] = '*'
			//	}
			//}
			//arrz.PPrint2D(grid)
		}
	} else {
		nextIdxs := findNextVisitableIndexesP2(c, visited)
		for _, idx := range nextIdxs {
			hikeSim(idx, visited, pathLen, maxPathLen)
		}
	}
	pathLen--
	visited[c.ToKey()] = false
}

func findNextVisitableIndexes(cur *arrz.Idx2D, visited tpz.StringSet) []*arrz.Idx2D {
	nextPos := make([]*arrz.Idx2D, 0)
	if hasSlide(cur) {
		p := cur.Clone()
		if gInput[p.I][p.J] == '>' {
			p.J += 1
		} else if gInput[p.I][p.J] == '<' {
			p.J -= 1
		} else if gInput[p.I][p.J] == '^' {
			p.I -= 1
		} else if gInput[p.I][p.J] == 'v' {
			p.I += 1
		} else {
			errz.HardAssert(false, "invalid slide")
		}
		nextPos = append(nextPos, p)
		return nextPos
	} else {
		p := arrz.NewIdx2D(cur.I+1, cur.J)
		if isVisitable(p) && !visited[p.ToKey()] && gInput[p.I][p.J] != '^' {
			nextPos = append(nextPos, p)
		}
		p = arrz.NewIdx2D(cur.I-1, cur.J)
		if isVisitable(p) && !visited[p.ToKey()] && gInput[p.I][p.J] != 'v' {
			nextPos = append(nextPos, p)
		}
		p = arrz.NewIdx2D(cur.I, cur.J+1)
		if isVisitable(p) && !visited[p.ToKey()] && gInput[p.I][p.J] != '<' {
			nextPos = append(nextPos, p)
		}
		p = arrz.NewIdx2D(cur.I, cur.J-1)
		if isVisitable(p) && !visited[p.ToKey()] && gInput[p.I][p.J] != '>' {
			nextPos = append(nextPos, p)
		}
		return nextPos
	}
}

func findNextVisitableIndexesP2(cur *arrz.Idx2D, visited tpz.StringSet) []*arrz.Idx2D {
	nextPos := make([]*arrz.Idx2D, 0)
	nbrs := cur.Neighbours(false)
	for _, p := range nbrs {
		if isVisitable(p) && !visited[p.ToKey()] {
			nextPos = append(nextPos, p)
		}
	}
	return nextPos
}

func hasSlide(idx *arrz.Idx2D) bool {
	return gInput[idx.I][idx.J] == '>' || gInput[idx.I][idx.J] == '<' || gInput[idx.I][idx.J] == '^' || gInput[idx.I][idx.J] == 'v'
}

func isVisitable(idx *arrz.Idx2D) bool {
	//return !(gInput[idx.I][idx.J] == '#')
	return !(idx.I < 0 || idx.I >= len(gInput) || idx.J < 0 || idx.J >= len(gInput[0]) || gInput[idx.I][idx.J] == '#')
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
	gStart = arrz.NewIdx2D(0, bytez.FindSubseq1D(gInput[0], []byte{'.'})[0])
	gEnd = arrz.NewIdx2D(len(gInput)-1, bytez.FindSubseq1D(gInput[len(gInput)-1], []byte{'.'})[0])
	fmt.Printf("gStart : %v | gEnd : %v\n", gStart.Str(), gEnd.Str())
}
