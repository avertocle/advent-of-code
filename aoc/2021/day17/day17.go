package day17

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/geom"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/numz"
	"github.com/avertocle/contests/io/stringz"
	"math"
)

var gInput [][]int

// todo : hit and trial hack, should check the actual way to find this one
const gVyMax = 1000

func SolveP1() string {
	bounds := parseBounds()
	ans, _ := runSim(bounds)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	bounds := parseBounds()
	_, ans := runSim(bounds)
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func runSim(bounds []*geom.Coord2d) (int, int) {
	vxMin, vxMax := 1, bounds[1].X
	vyMin, vyMax := bounds[1].Y, gVyMax
	didHit := false
	h, hmax, vCtr := 0, math.MinInt, 0
	for vx := vxMin; vx <= vxMax; vx++ {
		for vy := vyMin; vy <= vyMax; vy++ {
			didHit, h = runSimForOneVector(vx, vy, bounds)
			//fmt.Printf("%v,%v = %v hmax@%v\n", vx, vy, didHit, h)
			hmax = numz.Max(h, hmax)
			if didHit {
				vCtr++
			}
		}
	}
	return hmax, vCtr
}

// return did-hit-target, max-height
func runSimForOneVector(vx, vy int, bounds []*geom.Coord2d) (bool, int) {
	c := geom.NewCoord2d(0, 0)
	hmax := math.MinInt
	for c.X <= bounds[1].X && c.Y >= bounds[1].Y {
		//fmt.Printf("path = pos(%v,%v) v(%v,%v)\n", c.X, c.Y, vx, vy)
		c.X += vx
		c.Y += vy
		vx = numz.Max(0, vx-1)
		vy -= 1
		hmax = numz.Max(c.Y, hmax)
		if c.IsInside(bounds[0], bounds[1]) {
			return true, hmax
		}
	}
	return false, math.MinInt
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	tokens := stringz.SplitMulti(lines[0], []string{":", "=", "..", ",", " "})
	gInput = [][]int{
		{stringz.AtoI(tokens[4], math.MinInt), stringz.AtoI(tokens[5], math.MinInt)},
		{stringz.AtoI(tokens[8], math.MinInt), stringz.AtoI(tokens[9], math.MinInt)},
	}
	//fmt.Printf("%v\n", strings.Join(tokens, "|"))
	//outils.PrettyArray2DInt(gInput)
}

func parseBounds() []*geom.Coord2d {
	b := []*geom.Coord2d{
		geom.NewCoord2d(gInput[0][0], gInput[1][1]),
		geom.NewCoord2d(gInput[0][1], gInput[1][0]),
	}
	//fmt.Printf("bounds : tl(%v,%v) br(%v,%v)\n", b[0].X, b[0].Y, b[1].X, b[1].Y)
	return b
}
