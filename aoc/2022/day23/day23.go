package day23

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/mapz"
)

var gInput [][]byte
var gInpLen int
var gDirLoops map[int][]tDirLoop

const gPad = 5

type tDirLoop func(*elf, [][]byte) (bool, int, int, int)

const cellEmpty = '.'
const cellElf = '#'
const (
	north = 0
	south = 1
	west  = 2
	east  = 3
)

func SolveP1() string {
	elves, ground := parseElfPos()
	for r := 0; r < 10; r++ {
		runOneRound(elves, ground)
	}
	ans := calcEmptySlotsP1(ground)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	elves, ground := parseElfPos()
	r, movedCount := 0, 1
	for ; movedCount > 0; r++ {
		movedCount = runOneRound(elves, ground)
	}
	ans := r
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func runOneRound(elves []*elf, ground [][]byte) int {
	for _, e := range elves {
		e.updateProposal(ground)
	}
	collTracker := makeCollisionTracker(elves)
	movedCount := 0
	for _, e := range elves {
		if e.willMove {
			movedCount++
		}
	}
	for _, e := range elves {
		e.executeMove(ground, collTracker)
	}
	return movedCount
}

func isColliding(collTracker *mapz.MMIntInt, pos []int) bool {
	v, ok := collTracker.Get(pos[0], pos[1])
	errz.HardAssert(ok, "error : isColliding : not found %v", pos)
	errz.HardAssert(v > 0, "error : isColliding : len = %v, %v", v, pos)
	return v > 1
}

func areEmpty(ground [][]byte, pos [][]int) bool {
	for _, p := range pos {
		if ground[p[0]][p[1]] != cellEmpty {
			return false
		}
	}
	return true
}

func makeDirLoops() {
	gDirLoops = make(map[int][]tDirLoop)
	gDirLoops[north] = []tDirLoop{(*elf).canMoveSouth, (*elf).canMoveWest, (*elf).canMoveEast, (*elf).canMoveNorth}
	gDirLoops[south] = []tDirLoop{(*elf).canMoveWest, (*elf).canMoveEast, (*elf).canMoveNorth, (*elf).canMoveSouth}
	gDirLoops[west] = []tDirLoop{(*elf).canMoveEast, (*elf).canMoveNorth, (*elf).canMoveSouth, (*elf).canMoveWest}
	gDirLoops[east] = []tDirLoop{(*elf).canMoveNorth, (*elf).canMoveSouth, (*elf).canMoveWest, (*elf).canMoveEast}
}

func makeCollisionTracker(elves []*elf) *mapz.MMIntInt {
	collTracker := mapz.NewMMIntInt()
	for _, elf := range elves {
		if elf.willMove {
			collTracker.AddTo(elf.nextPos[0], elf.nextPos[1], 1)
		}
	}
	return collTracker
}

/***** P1 Functions *****/

func calcEmptySlotsP1(ground [][]byte) int {
	boundsTl, boundsBr := bytez.FindBounds2D(ground, cellEmpty)
	ans := bytez.CountInSection2D(ground, boundsTl, boundsBr, cellEmpty)
	return ans

}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
	gInpLen = len(gInput)
	gInput = bytez.Pad2D(gInput, gInpLen, gInpLen, gInpLen*gPad, '.')
	gInpLen = len(gInput)
	makeDirLoops()
}

func parseElfPos() ([]*elf, [][]byte) {
	elves := make([]*elf, 0)
	ground := bytez.Init2D(gInpLen, gInpLen, '.')
	for i := 0; i < gInpLen; i++ {
		for j := 0; j < gInpLen; j++ {
			if gInput[i][j] == cellElf {
				ground[i][j] = cellElf
				elves = append(elves, newElf(i, j))
			}
		}
	}
	return elves, ground
}

func debug_markBoundsAndShow(ground [][]byte) {
	boundsTl, boundsBr := bytez.FindBounds2D(ground, cellEmpty)
	//fmt.Printf("%v %v \n", boundsTl, boundsBr)
	//ground[boundsTl[0]][boundsTl[1]] = '*'
	//ground[boundsTl[0]][boundsBr[1]] = '*'
	//ground[boundsBr[0]][boundsTl[1]] = '*'
	//ground[boundsBr[0]][boundsBr[1]] = '*'
	bytez.PPrint2D(bytez.Extract2D(ground, boundsTl, boundsBr, '$'))
}

/***** DS *****/

type elf struct {
	i        int
	j        int
	d        int
	nextPos  []int
	willMove bool
}

func (e *elf) str() string {
	return fmt.Sprintf("[%v,%v], %v, %v, %v", e.i, e.j, e.d, e.nextPos, e.willMove)
}

func newElf(i, j int) *elf {
	return &elf{
		i: i,
		j: j,
		d: east,
	}
}

func (e *elf) updateProposal(ground [][]byte) {
	if e.isAllAlone(ground) {
		e.d = (e.d + 1) % 4
		e.willMove = false
		return
	}
	dirLoop := gDirLoops[e.d]
	e.d = (e.d + 1) % 4
	for _, f := range dirLoop {
		if ok, i, j, _ := f(e, ground); ok {
			e.nextPos = []int{i, j}
			e.willMove = true
			return
		}
	}
}

func (e *elf) executeMove(ground [][]byte, collTracker *mapz.MMIntInt) {
	if !e.willMove {
		return
	}
	if !isColliding(collTracker, e.nextPos) {
		ground[e.i][e.j] = '.'
		e.i, e.j = e.nextPos[0], e.nextPos[1]
		ground[e.i][e.j] = '#'
	}
	e.nextPos = nil
	e.willMove = false
}

func (e *elf) canMoveNorth(ground [][]byte) (bool, int, int, int) {
	pos := [][]int{{e.i - 1, e.j}, {e.i - 1, e.j + 1}, {e.i - 1, e.j - 1}}
	return areEmpty(ground, pos), e.i - 1, e.j, north
}

func (e *elf) canMoveSouth(ground [][]byte) (bool, int, int, int) {
	pos := [][]int{{e.i + 1, e.j}, {e.i + 1, e.j + 1}, {e.i + 1, e.j - 1}}
	return areEmpty(ground, pos), e.i + 1, e.j, south
}

func (e *elf) canMoveWest(ground [][]byte) (bool, int, int, int) {
	pos := [][]int{{e.i, e.j - 1}, {e.i - 1, e.j - 1}, {e.i + 1, e.j - 1}}
	return areEmpty(ground, pos), e.i, e.j - 1, west
}

func (e *elf) canMoveEast(ground [][]byte) (bool, int, int, int) {
	pos := [][]int{{e.i, e.j + 1}, {e.i - 1, e.j + 1}, {e.i + 1, e.j + 1}}
	return areEmpty(ground, pos), e.i, e.j + 1, east
}

func (e *elf) isAllAlone(ground [][]byte) bool {
	pos := [][]int{
		{e.i - 1, e.j - 1}, {e.i - 1, e.j}, {e.i - 1, e.j + 1},
		{e.i, e.j - 1}, {e.i, e.j + 1},
		{e.i + 1, e.j - 1}, {e.i + 1, e.j}, {e.i + 1, e.j + 1},
	}
	return areEmpty(ground, pos)
}
