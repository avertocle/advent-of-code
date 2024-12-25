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

func SolveP1() string {
	ans := solve(2, 100)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := solve(20, 100)
	return fmt.Sprintf("%v", ans)
}

func solve(cheatTime, minSavings int) int {
	grid := arrz.Copy2D(gInput)
	start := arrz.NewIdx2D(arrz.Find2D(grid, 'S')[0]...)
	end := arrz.NewIdx2D(arrz.Find2D(grid, 'E')[0]...)
	basePath, basePathDistMap := findBasePath(grid, start, end)
	savingToCheatCountMap := make(map[int]map[string]bool)
	for i, p := range basePath {
		dests := findAllDestsAfterCheat(grid, p, cheatTime+1)
		for cKey, c := range dests {
			pathLen := i + c.dist + basePathDistMap[c.e.ToKey()]
			pathDiff := len(basePath) - pathLen
			if pathDiff > 0 {
				if _, ok := savingToCheatCountMap[pathDiff]; !ok {
					savingToCheatCountMap[pathDiff] = make(map[string]bool)
				}
				savingToCheatCountMap[pathDiff][cKey] = true
			}
		}
	}
	ans := 0
	for k, v := range savingToCheatCountMap {
		if k >= minSavings {
			ans += len(v)
		}
	}
	return ans
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

type cheat struct {
	s, e *idx
	dist int
}

func (c *cheat) ToKey() string {
	return fmt.Sprintf("%v-%v", c.s.ToKey(), c.e.ToKey())
}

func findAllDestsAfterCheat(grid [][]byte, cheatPoint *idx, simCount int) map[string]*cheat {
	// finds all points at simCount (man-dis) away from cheatPoint
	dests := make(map[string]*cheat)
	for i := 0; i < simCount; i++ {
		for j := 0; j < simCount-i; j++ {
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

func findBasePath(grid [][]byte, s, e *idx) ([]*idx, map[string]int) {
	path := make([]*idx, 0)
	var curr, prev *idx
	var nbrs []*idx
	for curr = s.Clone(); !curr.IsEqual(e); {
		path = append(path, curr)
		nbrs = make([]*idx, 0)
		for _, n := range curr.Neighbours(false) {
			if n.IsInBounds(len(grid), len(grid[0])) &&
				!n.IsEqual(prev) && grid[n.I][n.J] != wall {
				nbrs = append(nbrs, n)
			}
		}
		errz.HardAssert(len(nbrs) > 0, "no valid nbrs found for %v", curr)
		prev, curr = curr, nbrs[0]
	}
	distMap := make(map[string]int)
	for i, p := range path {
		distMap[p.ToKey()] = len(path) - i
	}
	return path, distMap
}

func PrintSortedSavingsMap(m map[int]map[string]bool) {
	keys := mapz.Keys(m)
	slices.Sort(keys)
	for _, k := range keys {
		fmt.Println(k, len(m[k]))
	}
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
}
