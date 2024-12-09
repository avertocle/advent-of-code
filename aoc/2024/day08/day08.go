package day08

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/geom"
	"github.com/avertocle/contests/io/iutils"
)

const DirPath = "../2024/day08"

var gInput [][]byte

type Ant = geom.Coord2D[int]
type AntLocMap = map[byte][]*Ant

func SolveP1() string {
	ans := 0
	antMap := makeAntMap()
	markerGrid := arrz.Init2D(len(gInput), len(gInput[0]), byte('.'))
	for _, ants := range antMap {
		markAntiNodesForOneAntType(markerGrid, ants, findAntiNodesP1)
	}
	ans = arrz.Count2D(markerGrid, '#')
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	antMap := makeAntMap()
	markerGrid := arrz.Init2D(len(gInput), len(gInput[0]), byte('.'))
	for _, ants := range antMap {
		markAntiNodesForOneAntType(markerGrid, ants, findAntiNodesP2)
	}
	ans = arrz.Count2D(markerGrid, '#')
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func findAntiNodesP1(c1, c2 *Ant) []*Ant {
	dx, dy := find2dManDist(c1, c2)
	anLocs := []*geom.Coord2D[int]{
		geom.NewCoord2D(c1.X-dx, c1.Y-dy),
		geom.NewCoord2D(c2.X-dx, c2.Y-dy),
		geom.NewCoord2D(c1.X+dx, c1.Y+dy),
		geom.NewCoord2D(c2.X+dx, c2.Y+dy),
	}
	anLocsOnMap := make([]*Ant, 0)
	for _, c := range anLocs {
		if isInBounds(c) && !c1.IsEqual(c) && !c2.IsEqual(c) {
			anLocsOnMap = append(anLocsOnMap, c)
		}
	}
	return anLocsOnMap
}

/***** P2 Functions *****/

func findAntiNodesP2(c1, c2 *Ant) []*Ant {
	dx, dy := find2dManDist(c1, c2)
	// todo : ideally, should reduce dx,dy by dividing by gcd but input had dx,dy as co-primes so fuck it for now
	anLocs := make([]*Ant, 0)
	for i := 0; ; i++ {
		an := geom.NewCoord2D(c1.X+(i*dx), c1.Y+(i*dy))
		if isInBounds(an) {
			anLocs = append(anLocs, an)
		} else {
			break
		}
	}
	for i := 0; ; i++ {
		an := geom.NewCoord2D(c1.X-(i*dx), c1.Y-(i*dy))
		if isInBounds(an) {
			anLocs = append(anLocs, an)
		} else {
			break
		}
	}
	return anLocs
}

/***** Common Functions *****/

func markAntiNodesForOneAntType(markerGrid [][]byte, ants []*Ant, finderFunc func(*Ant, *Ant) []*Ant) {
	for i := 0; i < len(ants)-1; i++ {
		for j := i + 1; j < len(ants); j++ {
			anLoc := finderFunc(ants[i], ants[j])
			for _, c := range anLoc {
				markerGrid[c.X][c.Y] = '#'
			}
		}
	}
}

func isInBounds(c *Ant) bool {
	return c.X >= 0 && c.X < len(gInput) && c.Y >= 0 && c.Y < len(gInput[0])
}

func find2dManDist(c1, c2 *Ant) (int, int) {
	return c2.X - c1.X, c2.Y - c1.Y
}

func makeAntMap() AntLocMap {
	antMap := make(AntLocMap)
	for i, row := range gInput {
		for j, c := range row {
			if c != '.' {
				antMap[c] = append(antMap[c], geom.NewCoord2D(i, j))
			}
		}
	}
	return antMap
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
}
