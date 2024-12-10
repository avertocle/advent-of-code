package day10

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/cmz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

const DirPath = "../2024/day10"

var gInput [][]byte
var gStarts []*arrz.Idx2D
var gEnds []*arrz.Idx2D

func SolveP1() string {
	ans := 0
	trailCount := new(int)
	for _, s := range gStarts {
		calcTrailCountP1(s, make(cmz.MapVisited), trailCount)
	}
	ans = *trailCount
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	trailCount := new(int)
	for _, s := range gStarts {
		for _, e := range gEnds {
			calcTrailCountP2(s, e, make(cmz.MapVisited), trailCount)
		}
	}
	ans = *trailCount
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func calcTrailCountP1(start *arrz.Idx2D, visited cmz.MapVisited, trailCount *int) {
	visited[start.ToKey()] = true
	if gInput[start.I][start.J] == '9' {
		*trailCount++
		return
	}
	visitableNbrs := getVisitableNbrs(start, visited)
	for _, n := range visitableNbrs {
		calcTrailCountP1(n, visited, trailCount)
	}
	visited[start.ToKey()] = false
}

/***** P2 Functions *****/

func calcTrailCountP2(start *arrz.Idx2D, end *arrz.Idx2D, visited cmz.MapVisited, trailCount *int) {
	visited[start.ToKey()] = true
	if start.IsEqual(end) {
		*trailCount++
		visited[start.ToKey()] = false
		return
	}
	visitableNbrs := getVisitableNbrs(start, visited)
	for _, n := range visitableNbrs {
		calcTrailCountP2(n, end, visited, trailCount)
	}
	visited[start.ToKey()] = false
}

/***** Common Functions *****/

func getVisitableNbrs(p *arrz.Idx2D, visited cmz.MapVisited) []*arrz.Idx2D {
	allNbrs := p.Neighbours(false)
	visitableNbrs := make([]*arrz.Idx2D, 0)
	for _, n := range allNbrs {
		if n.IsInBounds(len(gInput), len(gInput[0])) && !visited[n.ToKey()] && gInput[n.I][n.J] == gInput[p.I][p.J]+1 {
			visitableNbrs = append(visitableNbrs, n)
		}
	}
	return visitableNbrs
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
	points := arrz.Find2D(gInput, '0')
	gStarts = make([]*arrz.Idx2D, 0)
	gEnds = make([]*arrz.Idx2D, 0)
	for _, p := range points {
		gStarts = append(gStarts, arrz.NewIdx2D(p[0], p[1]))
	}
	points = arrz.Find2D(gInput, '9')
	for _, p := range points {
		gEnds = append(gEnds, arrz.NewIdx2D(p[0], p[1]))
	}
}
