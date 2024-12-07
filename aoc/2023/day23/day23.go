package day23

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/cmz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
)

var DirPath = "../2023/day23"
var gInput [][]byte
var gStart, gEnd *arrz.Idx2D

//2320 too-high

func SolveP1() string {
	ans := 0
	player := gStart
	dpGrid := intz.Init2D(len(gInput), len(gInput[0]), 0)
	hikeSim(player, make(cmz.MapVisited), dpGrid)
	ans = dpGrid[gEnd.I][gEnd.J]
	//intz.PPrint2D(dpGrid)
	//bytez.PPrint2D(gInput)
	return fmt.Sprintf("%v", ans)
}

func hikeSim(c *arrz.Idx2D, visited cmz.MapVisited, dpGrid [][]int) {
	//markVisited(c, visited)
	//nextIdxs := findNextVisitableIndexes(c, visited)
	////fmt.Printf("hikeSim : c %v %v | %2d | nextIdxs : %v\n", c.Str(), string(gInput[c.I][c.J]), dpGrid[gEnd.I][gEnd.J], arrz.Idx2DListToStr(nextIdxs))
	//for _, idx := range nextIdxs {
	//	//if (dpGrid[c.I][c.J]+1, dpGrid[idx.I][idx.J]){
	//	//
	//	//}
	//	//dpGrid[idx.I][idx.J] = intz.Max(dpGrid[c.I][c.J]+1, dpGrid[idx.I][idx.J])
	//}
	//for _, idx := range nextIdxs {
	//	hikeSim(idx, visited, dpGrid)
	//}
	//
}

func findNextVisitableIndexes(cur *arrz.Idx2D, visited cmz.MapVisited) []*arrz.Idx2D {
	nextPos := make([]*arrz.Idx2D, 0)
	if cur.IsEqual(gEnd) {
		return nextPos
	}
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
		}
		nextPos = append(nextPos, p)
		return nextPos
	}

	nbrs := cur.Neighbours(true)
	for _, nbr := range nbrs {
		if isVisitable(nbr) && !visited[nbr.ToKey()] {
			nextPos = append(nextPos, nbr)
		}
	}
	return nextPos
}

func hasSlide(idx *arrz.Idx2D) bool {
	return gInput[idx.I][idx.J] == '>' || gInput[idx.I][idx.J] == '<' || gInput[idx.I][idx.J] == '^' || gInput[idx.I][idx.J] == 'v'
}

func markVisited(idx *arrz.Idx2D, visited cmz.MapVisited) {
	visited[idx.ToKey()] = true
}

func isVisitable(idx *arrz.Idx2D) bool {
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
	gStart = arrz.NewIdxP2D(0, bytez.FindSubseq1D(gInput[0], []byte{'.'})[0])
	gEnd = arrz.NewIdxP2D(len(gInput)-1, bytez.FindSubseq1D(gInput[len(gInput)-1], []byte{'.'})[0])
	fmt.Printf("gStart : %v | gEnd : %v\n", gStart.Str(), gEnd.Str())
}
