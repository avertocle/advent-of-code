package day22

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/geom"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/typez"
	"strings"
)

var gInput []*brick

// 32933
// 11079
var gSpace [][][]int

var DirPath = "../2023/day22"

func SolveP1() string {
	ans := 0
	initStuff()
	allBrickSupportMap := makeAllBrickSupportMap()
	for id, _ := range gInput {
		countBricksOnlySupportedBy := calcCountBricksOnlySupportedBy(id, allBrickSupportMap)
		if countBricksOnlySupportedBy == 0 {
			ans++
		}
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	initStuff()
	allBrickSupportMap := makeAllBrickSupportMap()
	for id, _ := range gInput {
		crm := calcRemovedBricksCascading(id, allBrickSupportMap)
		errz.HardAssert(len(crm) > 0 && crm[id], "always contains itself %v %v", id, len(crm))
		ans += len(crm) - 1
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func calcCountBricksOnlySupportedBy(id int, allBrickSupportMap typez.MapIIB) int {
	count := 0
	for id2, oneBrickSupportMap := range allBrickSupportMap {
		if id2 == id {
			continue
		}
		if len(oneBrickSupportMap) == 1 && oneBrickSupportMap[id] {
			count++
		}
	}
	return count

}

/***** P2 Functions *****/

func calcRemovedBricksCascading(id int, allBrickSupportMap typez.MapIIB) typez.MapIB {
	removedBricks := make(typez.MapIB)
	removedBricks[id] = true
	for {
		nextRemovedBricks := getNextRemovedBricks(removedBricks, allBrickSupportMap)
		for id2, _ := range nextRemovedBricks {
			removedBricks[id2] = true
		}
		if len(nextRemovedBricks) == 0 {
			break
		}
	}
	return removedBricks
}

func getNextRemovedBricks(removedBricks typez.MapIB, brickSupportMap typez.MapIIB) typez.MapIB {
	nextRemoved := make(typez.MapIB)
	for id, _ := range gInput {
		allSupportingBricksRemoved := checkAllSupportingBricksRemoved(id, removedBricks, brickSupportMap)
		if allSupportingBricksRemoved && !removedBricks[id] {
			nextRemoved[id] = true
		}
	}
	return nextRemoved
}

func checkAllSupportingBricksRemoved(id int, removedBricks typez.MapIB, brickSupportMap typez.MapIIB) bool {
	dependencies := brickSupportMap[id]
	if len(dependencies) == 0 {
		return false
	}
	for id2, _ := range dependencies {
		if _, ok := removedBricks[id2]; !ok {
			return false
		}
	}
	return true
}

/***** Common Functions *****/

func initStuff() {
	gSpace = intz.Init3D(500, 500, 500, -1)
	placeAllBricksInSpace()
	_ = moveBricksDown()
}

func makeAllBrickSupportMap() typez.MapIIB {
	allBrickSupportMap := make(typez.MapIIB)
	for id, b := range gInput {
		allBrickSupportMap[id], _ = getSupportingBricks(b)
	}
	return allBrickSupportMap
}

// returns supporting bricks and if the brick is supported at all
func getSupportingBricks(b *brick) (typez.MapIB, bool) {
	supportingBricks := make(typez.MapIB)
	//fmt.Printf("getSupportingBricks : checking brick %v\n", b.str())
	if b.s.Z == 1 || b.e.Z == 1 {
		return supportingBricks, true
	}

	for i := b.s.X; i <= b.e.X; i++ {
		if gSpace[i][b.s.Y][b.s.Z-1] != -1 {
			supportingBricks[gSpace[i][b.s.Y][b.s.Z-1]] = true
		}
	}
	for i := b.s.Y; i <= b.e.Y; i++ {
		if gSpace[b.s.X][i][b.s.Z-1] != -1 {
			supportingBricks[gSpace[b.s.X][i][b.s.Z-1]] = true
		}
	}
	if gSpace[b.s.X][b.s.Y][b.s.Z-1] != -1 && gSpace[b.s.X][b.s.Y][b.e.Z-1] != -1 {
		supportingBricks[gSpace[b.s.X][b.s.Y][b.s.Z-1]] = true
	}

	return supportingBricks, len(supportingBricks) > 0
}

func removeBrickFromSpace(b *brick) {
	for i := b.s.X; i <= b.e.X; i++ {
		gSpace[i][b.s.Y][b.s.Z] = -1
	}
	for i := b.s.Y; i <= b.e.Y; i++ {
		gSpace[b.s.X][i][b.s.Z] = -1
	}
	for i := b.s.Z; i <= b.e.Z; i++ {
		gSpace[b.s.X][b.s.Y][i] = -1
	}

}

func placeBrickInSpace(b *brick, id int) {
	for i := b.s.X; i <= b.e.X; i++ {
		gSpace[i][b.s.Y][b.s.Z] = id
	}
	for i := b.s.Y; i <= b.e.Y; i++ {
		gSpace[b.s.X][i][b.s.Z] = id
	}
	for i := b.s.Z; i <= b.e.Z; i++ {
		gSpace[b.s.X][b.s.Y][i] = id
	}
}

func placeAllBricksInSpace() {
	for id, b := range gInput {
		placeBrickInSpace(b, id)
	}
}

func moveBricksDown() typez.MapIB {
	noBricksMoved := false
	bricksMoved := make(typez.MapIB)
	for noBricksMoved == false {
		bricksMovedCount := 0
		for id, b := range gInput {
			for {
				if _, isSupported := getSupportingBricks(b); isSupported {
					break
				}
				removeBrickFromSpace(b)
				b.moveDown()
				placeBrickInSpace(b, id)
				bricksMoved[id] = true
				bricksMovedCount++
			}
		}
		if bricksMovedCount == 0 {
			noBricksMoved = true
		}
	}
	return bricksMoved
}

func printAllBrickSupportMap(brickSupportMap typez.MapIIB) {
	for id, b := range gInput {
		fmt.Printf("allBrickSupportMap : brick-%02d : %v : %v\n", id, b.str(), brickSupportMap[id])
	}
}

/***** Structs *****/

type brick struct {
	s    *geom.Coord3d
	e    *geom.Coord3d
	oldS *geom.Coord3d
	oldE *geom.Coord3d
}

func (b *brick) moveDown() {
	b.s.MoveBy([]int{0, 0, -1})
	b.e.MoveBy([]int{0, 0, -1})
}

func (b *brick) str() string {
	return fmt.Sprintf("%v ~ %v", b.s.Str(), b.e.Str())
}

func newBrick(str string) *brick {
	t := strings.Split(str, "~")
	return &brick{
		s: geom.NewCoord3dFromVec(iutils.ExtractInt1DFromString0D(t[0], ",", -1)),
		e: geom.NewCoord3dFromVec(iutils.ExtractInt1DFromString0D(t[1], ",", -1)),
	}
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = make([]*brick, 0)
	for _, line := range lines {
		gInput = append(gInput, newBrick(line))
	}
}
