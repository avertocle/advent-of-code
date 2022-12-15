package day22

import (
	"fmt"
	"github.com/avertocle/contests/io/geom"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"log"
	"math"
)

var gInput []string

func SolveP1() string {
	offset := []int{50, 50, 50}
	bounds := [][]int{{-50, 50}, {-50, 50}, {-50, 50}}
	cube := intz.Init3D(101, 101, 101, 0)
	begs, ends, vals := parseInstructions()
	for i, _ := range gInput {
		if begs[i].InBounds(bounds) && ends[i].InBounds(bounds) {
			changeCubeState(cube, offset, bounds, begs[i], ends[i], vals[i])
		}
	}
	ans := intz.Count3d(cube, 1)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func changeCubeState(cube [][][]int, offset []int, bounds [][]int, beg, end *geom.Coord3d, val int) {
	beg.Trim(bounds).Move(offset)
	end.Trim(bounds).Move(offset)
	_ = intz.SetSub3D(cube,
		[]int{beg.X, end.X},
		[]int{beg.Y, end.Y},
		[]int{beg.Z, end.Z}, val)
	//fmt.Printf("set %v values from (%v) to (%v) to %v\n", count, beg.Str(), end.Str(), val)
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	gInput = lines
}

func parseInstructions() ([]*geom.Coord3d, []*geom.Coord3d, []int) {
	begs := make([]*geom.Coord3d, len(gInput))
	ends := make([]*geom.Coord3d, len(gInput))
	vals := make([]int, len(gInput))
	for i, l := range gInput {
		begs[i], ends[i], vals[i] = parseOneInstruction(l)
	}
	return begs, ends, vals
}

func parseOneInstruction(l string) (*geom.Coord3d, *geom.Coord3d, int) {
	tokens := stringz.SplitMulti(l, []string{" ", "=", ",", ".."})
	beg := geom.NewCoord3d(stringz.AtoiQ(tokens[2], math.MinInt),
		stringz.AtoiQ(tokens[5], math.MinInt),
		stringz.AtoiQ(tokens[8], math.MinInt))
	end := geom.NewCoord3d(stringz.AtoiQ(tokens[3], math.MinInt),
		stringz.AtoiQ(tokens[6], math.MinInt),
		stringz.AtoiQ(tokens[9], math.MinInt))
	val := (len(tokens[0]) - 1) % 2 // lols
	//fmt.Printf("input : (%v) to (%v) to (%v)\n", beg.Str(), end.Str(), val)
	return beg, end, val
}
