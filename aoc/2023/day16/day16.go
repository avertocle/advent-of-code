package day16

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

var gInput [][]byte

const DirPath = "../2023/day16"

const (
	DirUp = iota
	DirRight
	DirDown
	DirLeft
)

func SolveP1() string {
	ans := calcEnergizedTileCount([]int{0, 0, DirRight})
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	etCount := 0
	for i := 0; i < len(gInput); i++ {
		etCount = calcEnergizedTileCount([]int{i, 0, DirRight})
		if etCount > ans {
			ans = etCount
		}
		etCount = calcEnergizedTileCount([]int{i, len(gInput[0]) - 1, DirLeft})
		if etCount > ans {
			ans = etCount
		}
	}
	for j := 0; j < len(gInput[0]); j++ {
		etCount = calcEnergizedTileCount([]int{0, j, DirDown})
		if etCount > ans {
			ans = etCount
		}
		etCount = calcEnergizedTileCount([]int{len(gInput) - 1, j, DirUp})
		if etCount > ans {
			ans = etCount
		}
	}
	return fmt.Sprintf("%v", ans)
}
func calcEnergizedTileCount(startLight []int) int {
	markedGrid := traceLightPathP1(startLight)
	//bytez.PPrint2D(markedGrid)
	return bytez.Count2D(markedGrid, 'X')
}

func traceLightPathP1(startLight []int) [][]byte {
	markedGrid := bytez.Init2D(len(gInput), len(gInput[0]), '.')
	currLights := [][]int{startLight}
	var lightHistory = make(map[string]bool)
	markedGrid[startLight[0]][startLight[1]] = 'X'
	for i := 0; i < 10000; i++ {
		nextLights := make([][]int, 0)
		for _, cl := range currLights {
			temp := getNextPos(cl)
			for _, tl := range temp {
				if isCoordValid(tl[0], tl[1]) {
					markedGrid[tl[0]][tl[1]] = 'X'
					if _, ok := lightHistory[fmt.Sprintf("%v", tl)]; !ok {
						nextLights = append(nextLights, tl)
						lightHistory[fmt.Sprintf("%v", tl)] = true
					}
				}
			}
		}
		currLights = nextLights
		//if i%100 == 0 {
		//	fmt.Printf("%v(%v) --> ", i, len(currLights))
		//}
	}
	fmt.Printf("%v-", len(currLights))
	return markedGrid
}

func isCoordValid(ci, cj int) bool {
	return ci >= 0 && ci < len(gInput) && cj >= 0 && cj < len(gInput[0])
}

func getNextPos(light []int) [][]int {
	ci, cj, dir := light[0], light[1], light[2]
	if ci < 0 || cj < 0 {
		fmt.Println(ci, cj)
	}
	mirror := gInput[ci][cj]
	switch mirror {
	case '/':
		switch dir {
		case DirUp:
			return [][]int{{ci, cj + 1, DirRight}}
		case DirRight:
			return [][]int{{ci - 1, cj, DirUp}}
		case DirDown:
			return [][]int{{ci, cj - 1, DirLeft}}
		case DirLeft:
			return [][]int{{ci + 1, cj, DirDown}}
		}
	case '\\':
		switch dir {
		case DirUp:
			return [][]int{{ci, cj - 1, DirLeft}}
		case DirRight:
			return [][]int{{ci + 1, cj, DirDown}}
		case DirDown:
			return [][]int{{ci, cj + 1, DirRight}}
		case DirLeft:
			return [][]int{{ci - 1, cj, DirUp}}
		}
	case '|':
		switch dir {
		case DirUp, DirDown:
			return [][]int{nextPosByNewDirec(ci, cj, dir)}
		case DirLeft, DirRight:
			return [][]int{nextPosByNewDirec(ci, cj, DirUp), nextPosByNewDirec(ci, cj, DirDown)}
		}
	case '-':
		switch dir {
		case DirUp, DirDown:
			return [][]int{nextPosByNewDirec(ci, cj, DirLeft), nextPosByNewDirec(ci, cj, DirRight)}
		case DirLeft, DirRight:
			return [][]int{nextPosByNewDirec(ci, cj, dir)}
		}
	case '.':
		return [][]int{nextPosByNewDirec(ci, cj, dir)}

	}
	errz.HardAssert(false, "getNextPos : invalid state | %v, %v,%v,%v", string(mirror), ci, cj, dir)
	return nil
}

func nextPosByNewDirec(ci, cj, dir int) []int {
	switch dir {
	case DirUp:
		return []int{ci - 1, cj, dir}
	case DirRight:
		return []int{ci, cj + 1, dir}
	case DirDown:
		return []int{ci + 1, cj, dir}
	case DirLeft:
		return []int{ci, cj - 1, dir}
	}
	errz.HardAssert(false, "nextPosByNewDirec : invalid state | %v,%v,%v", ci, cj, dir)
	return nil
}

/***** Common Functions *****/

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
}
