package day12

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

const DirPath = "../2024/day12"

// 96179 too low

var gInput [][]byte
var gTempMarker, gPermMarker byte = '#', '.'

func SolveP1() string {
	ans := 0
	grid := arrz.Copy2D(gInput)
	var ur *arrz.Idx2D[int]
	for {
		if ur = findNextUnmappedRegion(grid, gPermMarker); ur == nil {
			break
		}
		urVal, area, parameter := grid[ur.I][ur.J], 0, 0
		markTemporarily(grid, ur, urVal, &area, &parameter)
		markPermanently(grid)
		ans += area * parameter
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	grid := arrz.Copy2D(gInput)
	var r *arrz.Idx2D[int]
	for {
		if r = findNextUnmappedRegion(grid, gPermMarker); r == nil {
			break
		}
		rName, area, corners := grid[r.I][r.J], 0, 0
		markTemporarily(grid, r, rName, &area, new(int))
		corners = countCorners(grid)
		markPermanently(grid)
		ans += area * corners
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func findNextUnmappedRegion(grid [][]byte, marker byte) *arrz.Idx2D[int] {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != marker {
				return arrz.NewIdx2D(i, j)
			}
		}
	}
	return nil
}

func markTemporarily(grid [][]byte, s *arrz.Idx2D[int], pVal byte, area, parameter *int) {
	if grid[s.I][s.J] == gTempMarker {
		return
	}
	grid[s.I][s.J] = gTempMarker
	toVisitNbrs, param := getNbrSplit(grid, s, pVal)
	*area++
	*parameter += param
	for _, n := range toVisitNbrs {
		markTemporarily(grid, n, pVal, area, parameter)
	}
}

func getNbrSplit(grid [][]byte, p *arrz.Idx2D[int], pVal byte) ([]*arrz.Idx2D[int], int) {
	allNbrs := p.Neighbours(false)
	toVisitNbrs := make([]*arrz.Idx2D[int], 0)
	param := 0
	for _, n := range allNbrs {
		if n.IsInBounds(len(grid), len(grid[0])) && grid[n.I][n.J] == pVal {
			toVisitNbrs = append(toVisitNbrs, n)
		}
		if !n.IsInBounds(len(grid), len(grid[0])) || (grid[n.I][n.J] != gTempMarker && grid[n.I][n.J] != pVal) {
			param++
		}
	}
	return toVisitNbrs, param
}

func markPermanently(grid [][]byte) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == gTempMarker {
				grid[i][j] = gPermMarker
			}
		}
	}
}

/***** P2 Functions *****/

func countCorners(gridOrig [][]byte) int {
	corners := 0
	grid := arrz.Copy2D(gridOrig)
	grid = bytez.Pad2D(grid, len(grid), len(grid[0]), 1, gPermMarker)
	points := bytez.Find2D(grid, gTempMarker)
	for i, row := range grid {
		for j, cell := range row {
			if cell != gTempMarker && cell != gPermMarker {
				grid[i][j] = gPermMarker
			}
		}
	}
	//arrz.PPrint2D(grid)
	//fmt.Println("points", points)
	for _, p := range points {
		i, j := p[0], p[1]
		if grid[i-1][j] == gPermMarker && grid[i][j-1] == gPermMarker {
			corners++ // convex corner
		}
		if grid[i-1][j] == gPermMarker && grid[i][j+1] == gPermMarker {
			corners++ // convex corner
		}
		if grid[i+1][j] == gPermMarker && grid[i][j-1] == gPermMarker {
			corners++ // convex corner
		}
		if grid[i+1][j] == gPermMarker && grid[i][j+1] == gPermMarker {
			corners++ // convex corner
		}
		if grid[i-1][j] == gTempMarker && grid[i][j-1] == gTempMarker && grid[i-1][j-1] != gTempMarker {
			corners++ // concave corner
		}
		if grid[i-1][j] == gTempMarker && grid[i][j+1] == gTempMarker && grid[i-1][j+1] != gTempMarker {
			corners++ // concave corner
		}
		if grid[i+1][j] == gTempMarker && grid[i][j-1] == gTempMarker && grid[i+1][j-1] != gTempMarker {
			corners++ // concave corner
		}
		if grid[i+1][j] == gTempMarker && grid[i][j+1] == gTempMarker && grid[i+1][j+1] != gTempMarker {
			corners++ // concave corner
		}
	}
	return corners
}

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
}
