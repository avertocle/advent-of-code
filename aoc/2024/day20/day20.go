package day20

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/mapz"
	"slices"
)

const DirPath = "../2024/day20"

var gInput [][]byte

type idx = arrz.Idx2D[int]

const (
	space = byte('.')
	wall  = byte('#')
)

// 861156 too low

func SolveP1() string {
	ans := 0
	grid, start, end := initStuff()
	basePath := findBasePath(grid, start, end)
	fmt.Println("base path len", len(basePath))
	basePathDistMap := makeBasePathDistMap(basePath)
	savingsToCheatStartCountMap := make(map[int]map[string]bool)
	var prev *idx
	for i, p := range basePath {
		dests := findPossibleDestAfterCheatP1(grid, p, prev)
		for d, v := range dests {
			pathLen := i + 1 + 1 + basePathDistMap[d]
			pathDiff := len(basePath) - pathLen
			if pathDiff > 0 {
				if _, ok := savingsToCheatStartCountMap[pathDiff]; !ok {
					savingsToCheatStartCountMap[pathDiff] = make(map[string]bool)
				}
				fkey := fmt.Sprintf("%v-%v", p.ToKey(), v.ToKey())
				savingsToCheatStartCountMap[pathDiff][fkey] = true
			}
		}
		prev = p
	}
	fmt.Println()
	keys := mapz.Keys(savingsToCheatStartCountMap)
	slices.Sort(keys)
	for _, k := range keys {
		fmt.Println(k, len(savingsToCheatStartCountMap[k]))
		if k >= 100 {
			ans += len(savingsToCheatStartCountMap[k])
		}
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	grid, start, end := initStuff()
	basePath := findBasePath(grid, start, end)
	fmt.Println("base path len", len(basePath))
	basePathDistMap := makeBasePathDistMap(basePath)
	savingsToCheatStartCountMap := make(map[int]map[string]bool)
	for i, p := range basePath {
		dests := findPossibleDestAfterCheatP2(grid, p, 21)
		for cKey, c := range dests {
			pathLen := i + c.dist + basePathDistMap[c.e.ToKey()]
			pathDiff := len(basePath) - pathLen
			if pathDiff > 0 {
				if _, ok := savingsToCheatStartCountMap[pathDiff]; !ok {
					savingsToCheatStartCountMap[pathDiff] = make(map[string]bool)
				}
				savingsToCheatStartCountMap[pathDiff][cKey] = true
			}
		}
	}
	ans := 0
	keys := mapz.Keys(savingsToCheatStartCountMap)
	slices.Sort(keys)
	for _, k := range keys {
		if k >= 100 {
			fmt.Println(k, len(savingsToCheatStartCountMap[k]))
			ans += len(savingsToCheatStartCountMap[k])
		}
	}
	return fmt.Sprintf("%v", ans)
}

type cheat struct {
	s, e *idx
	dist int
}

func (c *cheat) ToKey() string {
	return fmt.Sprintf("%v-%v", c.s.ToKey(), c.e.ToKey())
}

func findPossibleDestAfterCheatP2(grid [][]byte, cheatPoint *idx, simCount int) map[string]*cheat {
	dests := make(map[string]*cheat)
	for i := 0; i < simCount; i++ {
		for j := 0; j < simCount-i; j++ {
			//fmt.Println("adding for md ", i+j)
			p1, p2, p3, p4 := cheatPoint.Clone(), cheatPoint.Clone(), cheatPoint.Clone(), cheatPoint.Clone()
			p1.MoveBy(-i, -j)
			p2.MoveBy(-i, j)
			p3.MoveBy(i, -j)
			p4.MoveBy(i, j)
			addToDestsIfValid(grid, cheatPoint, p1, i+j, dests)
			addToDestsIfValid(grid, cheatPoint, p2, i+j, dests)
			addToDestsIfValid(grid, cheatPoint, p3, i+j, dests)
			addToDestsIfValid(grid, cheatPoint, p4, i+j, dests)
		}
	}
	return dests
}

func addToDestsIfValid(grid [][]byte, s, e *idx, dist int, dests map[string]*cheat) {
	if e.IsInBounds(len(grid), len(grid[0])) && grid[e.I][e.J] != wall {
		c := &cheat{s: s, e: e, dist: dist}
		if _, ok := dests[c.ToKey()]; !ok {
			dests[c.ToKey()] = c
		}
		if dests[c.ToKey()].dist > c.dist {
			dests[c.ToKey()].dist = c.dist
		}
	}
}

func findPossibleDestAfterCheatP1(grid [][]byte, cheatPoint, prev *idx) map[string]*idx {
	dests := make(map[string]*idx)
	nbrsL1 := findValidNbrs(grid, cheatPoint, prev, true)
	for _, n1 := range nbrsL1 {
		nbrsL2 := findValidNbrs(grid, n1, cheatPoint, false)
		for _, n2 := range nbrsL2 {
			dests[n2.ToKey()] = n2
		}
	}

	return dests
}

func makeBasePathDistMap(path []*idx) map[string]int {
	distMap := make(map[string]int)
	for i, p := range path {
		distMap[p.ToKey()] = len(path) - i
	}
	return distMap

}

func findBasePath(grid [][]byte, s, e *idx) []*idx {
	path := make([]*idx, 0)
	var curr, prev *idx
	for curr = s.Clone(); !curr.IsEqual(e); {
		path = append(path, curr)
		nbrs := findValidNbrs(grid, curr, prev, false)
		errz.HardAssert(len(nbrs) > 0, "no valid nbrs found for %v", curr)
		prev = curr
		curr = nbrs[0]
	}
	return path
}

func findValidNbrs(grid [][]byte, curr, prev *idx, includeWalls bool) []*idx {
	nbrs := make([]*idx, 0)
	allNbrs := curr.Neighbours(false)
	for _, n := range allNbrs {
		if n.IsInBounds(len(grid), len(grid[0])) && !n.IsEqual(prev) {
			if includeWalls || grid[n.I][n.J] != wall {
				nbrs = append(nbrs, n)
			}
		}
	}
	return nbrs
}

func initStuff() ([][]byte, *idx, *idx) {
	grid := arrz.Copy2D(gInput)
	s := arrz.NewIdx2D(arrz.Find2D(grid, 'S')[0]...)
	e := arrz.NewIdx2D(arrz.Find2D(grid, 'E')[0]...)
	return grid, s, e
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
}
