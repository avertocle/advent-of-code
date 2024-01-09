package day21

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

var gInput [][]byte

const DirPath = "../2023/day21"

func SolveP1() string {
	ans := 0
	start := [][]int{arrz.Find2D(gInput, 'S')[0]}
	gInput[start[0][0]][start[0][1]] = '.'
	currState := newState()
	currState.maxStepsP1 = 64
	goBFS(gInput, start, currState)
	ans = currState.visitedNodeCtr
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func goBFS(ground [][]byte, toVisit [][]int, currState *state) {
	//fmt.Println("toVisit", toVisit)
	if len(toVisit) == 0 || currState.isEndStateP1() {
		return
	}
	nextToVisit := make([][]int, 0)
	for _, tv := range toVisit {
		nextToVisit = append(nextToVisit, getValidNeighbours(ground, tv)...)
	}
	nextToVisit = arrz.Unique2D(nextToVisit)
	currState.visitedNodeCtr = len(toVisit)
	currState.stepCtr++
	goBFS(ground, nextToVisit, currState)
}

func getValidNeighbours(ground [][]byte, index []int) [][]int {
	neighbours := arrz.Neighbours2D(index)
	criterion1 := arrz.IsValidIndexCriterion2D[byte]
	criterion2 := arrz.MakeValueCriterion2D(byte('.'))
	criteria := []arrz.CriterionFunc[byte]{criterion1, criterion2}
	return arrz.GenericSelect2D(ground, neighbours, criteria)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)

}

/***** Interfaces *****/

type state struct {
	visitedNodeCtr int
	stepCtr        int
	maxStepsP1     int
}

func (s state) isEndStateP1() bool {
	return s.stepCtr > s.maxStepsP1
}

func newState() *state {
	return &state{}
}
