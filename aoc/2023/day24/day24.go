package day24

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/geom"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
)

const DirPath = "../2023/day24"

var gInputPos [][]int
var gInputVel [][]int

func SolveP1() string {
	ans := 0
	coords := make([]*geom.Coord2D[float64], len(gInputPos))
	vels := make([]*geom.Coord2D[float64], len(gInputVel))
	lines := make([]*geom.Line2D[float64], len(gInputPos))
	for i := range gInputPos {
		coords[i] = geom.NewCoord2D(float64(gInputPos[i][0]), float64(gInputPos[i][1]))
		vels[i] = geom.NewCoord2D(float64(gInputVel[i][0]), float64(gInputVel[i][1]))
		lines[i] = geom.NewLine2D[float64](coords[i], vels[i])
	}
	//boundTL := geom.NewCoord2D(7.0, 27.0)
	//boundBR := geom.NewCoord2D(27.0, 7.0)
	boundTL := geom.NewCoord2D(200000000000000.0, 400000000000000.0)
	boundBR := geom.NewCoord2D(400000000000000.0, 200000000000000.0)
	intersects := make([]*geom.Coord2D[float64], 0)
	intersectsIn := make([]*geom.Coord2D[float64], 0)
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			in := geom.LineIntersect2D(lines[i], lines[j])
			intersects = append(intersects, in)
			if in.IsInside(boundTL, boundBR) &&
				ifFuturePoint(coords[i], in, vels[i]) &&
				ifFuturePoint(coords[j], in, vels[j]) {
				intersectsIn = append(intersectsIn, in)
			}
		}
	}
	//geom.PPrintCoord2D(intersects)
	//fmt.Println()
	//geom.PPrintCoord2D(intersectsIn)

	ans = len(intersectsIn)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func ifFuturePoint(cc, ic, v *geom.Coord2D[float64]) bool {
	nc := geom.NewCoord2D(cc.X+v.X, cc.Y+v.Y)
	d1 := geom.Dist2D(cc, ic)
	d2 := geom.Dist2D(nc, ic)
	return d1 >= d2
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInputPos = make([][]int, len(lines))
	gInputVel = make([][]int, len(lines))
	for i, line := range lines {
		t := stringz.SplitMultiTrimSpace(line, []string{",", " ", "@"})
		gInputPos[i] = []int{stringz.AtoI(t[0], -1), stringz.AtoI(t[1], -1), stringz.AtoI(t[2], -1)}
		gInputVel[i] = []int{stringz.AtoI(t[3], -1), stringz.AtoI(t[4], -1), stringz.AtoI(t[5], -1)}
	}
}
