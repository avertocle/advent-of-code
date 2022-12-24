package day15

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/rangez"
	"github.com/avertocle/contests/io/stringz"
	"math"
	"time"
)

var gInpSen [][]int
var gInpBea [][]int
var gBoundTL []int
var gBoundBR []int
var gOffset []int

func SolveP1() string {
	//rowIdx := 2000000 + gOffset[1]
	////rowIdx := 10 + gOffset[1]
	//finBea, found := getMarkedCount(rowIdx, 0)
	//getMarkedRow(gridRow, rowIdx)
	////fmt.Printf("%03d : %v\n", rowIdx, string(gridRow))
	//ans := bytez.Count1D(gridRow, '#')
	ans := 0
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	//max := 20
	max := 4000000
	t0sec, t1sec := time.Now().Unix(), time.Now().Unix()
	ctr0, ctr1 := gOffset[1], 0
	var finBea [][]int
	var fb []int
	var found bool
	for y := gOffset[1]; y < max+gOffset[1]; y++ {
		t1sec = time.Now().Unix()
		ctr1 = y
		if t1sec-t0sec >= 1 {
			fmt.Printf("%vs : %v\n", t1sec-t0sec, ctr1-ctr0)
			t0sec = t1sec
			ctr0 = ctr1
		}
		fb, found = getMarkedCount(y, max)
		if found {
			finBea = append(finBea, fb)
			//break
		}
	}
	intz.PPrint2D(finBea)
	a := int64(finBea[0][0] - gOffset[0])
	b := int64(finBea[0][1] - gOffset[1])
	fmt.Printf("fnbea : %v,%v\n", a, b)
	fmt.Printf("gCtr : %v,%v\n", gCtr1, gCtr2)
	ans := (a)*4000000 + b
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

var gCtr1, gCtr2 int

func getMarkedCount(rowIdx int, max int) ([]int, bool) {
	var mran [][]int
	var ran []int
	for i := 0; i < len(gInpSen); i++ {
		ran = getProjectedRange(rowIdx, gInpSen[i], gInpBea[i])
		if len(ran) > 0 {
			mran = rangez.Union1D(mran, ran)
		}
	}
	//l1 := len(mran)
	if len(mran) > 2 {
		//intz.PPrint2D(mran)
	}
	rowBeacons := getBeaconsInRow(rowIdx)
	for i := 0; i < len(rowBeacons); i++ {
		mran = rangez.Union1D(mran, rowBeacons[i])
	}
	//	l2 := len(mran)
	cutLeft := []int{gBoundTL[0] * 10, gOffset[0] - 1}
	cutRight := []int{gOffset[0] + max + 1, gBoundBR[0] * 10}
	mran = rangez.Union1D(mran, cutLeft)
	//	l3 := len(mran)
	mran = rangez.Union1D(mran, cutRight)
	//	l4 := len(mran)
	//if len(mran) > 2 {
	//	fmt.Printf("mran len history : %v %v %v %v\n", l1, l2, l3, l4)
	//	fmt.Printf("cuts : %v, %v\n", []int{gBoundTL[0], gOffset[0]}, []int{gOffset[0] + max, gBoundBR[0]})
	//	intz.PPrint2D(mran)
	//}
	//errz.HardAssert(len(mran) <= 2, "error merged-range-len = %v", len(mran))
	if len(mran) == 2 {
		//errz.HardAssert(mran[1][0]-mran[0][1] <= 2, "error big gap in merged-range")
		if mran[1][0]-mran[0][1] == 2 {
			fmt.Println()
			intz.PPrint2D(mran)
			fmt.Println()
			return []int{mran[0][1] + 1, rowIdx}, true
		} else {
			gCtr2++
		}
	} else if len(mran) > 2 {
		gCtr1++
	}
	return nil, false
}

func getBeaconsInRow(rowIdx int) [][]int {
	beacons := [][]int{}
	for _, bea := range gInpBea {
		if bea[1] == rowIdx {
			beacons = append(beacons, []int{bea[0], bea[0]})
		}
	}
	return beacons
}

func trimx(r []int, max int) []int {
	return []int{
		intz.Max(gOffset[0], r[0]),
		intz.Min(gOffset[0]+max, r[1]),
	}
}

func getProjectedRange(rowIdx int, sen, bea []int) []int {
	mdis := calcManDis(sen, bea)
	xdis := mdis - intz.Abs(sen[1]-rowIdx)
	if xdis < 0 {
		return []int{}
	}
	ks, ke := sen[0]-xdis, sen[0]+xdis
	errz.HardAssert(ks >= 0 && ke >= 0,
		"error : (%v, %v), mdis(%v), xdis(%v), rowIdx(%v), %v, %v",
		ks, ke, mdis, xdis, rowIdx, sen, bea)
	return []int{ks, ke}
}

//
//func markImpacted(gridRow []byte, rowIdx int, sen, bea []int) {
//	mdis := calcManDis(sen, bea)
//	xdis := mdis - intz.Abs(sen[1]-rowIdx)
//	if xdis < 0 {
//		return
//	}
//	ks, ke := sen[0]-xdis, sen[0]+xdis
//	errz.HardAssert(ks >= 0 && ke >= 0 && ks < len(gridRow) && ke < len(gridRow),
//		"error : (%v, %v), mdis(%v), xdis(%v), rowIdx(%v), %v, %v",
//		ks, ke, mdis, xdis, rowIdx, sen, bea)
//	for i := ks; i <= ke; i++ {
//		gridRow[i] = '#'
//	}
//}
//
//func getMarkedRow(gridRow []byte, rowIdx int) {
//	for i := 0; i < len(gInpSen); i++ {
//		markImpacted(gridRow, rowIdx, gInpSen[i], gInpBea[i])
//	}
//	markBeacons(gridRow, rowIdx)
//}
//
//func markBeacons(gridRow []byte, rowIdx int) {
//	for _, bea := range gInpBea {
//		if bea[1] == rowIdx {
//			gridRow[bea[0]] = 'B'
//		}
//	}
//}
//
//func markImpacted(gridRow []byte, rowIdx int, sen, bea []int) {
//	mdis := calcManDis(sen, bea)
//	xdis := mdis - intz.Abs(sen[1]-rowIdx)
//	if xdis < 0 {
//		return
//	}
//	ks, ke := sen[0]-xdis, sen[0]+xdis
//	errz.HardAssert(ks >= 0 && ke >= 0 && ks < len(gridRow) && ke < len(gridRow),
//		"error : (%v, %v), mdis(%v), xdis(%v), rowIdx(%v), %v, %v",
//		ks, ke, mdis, xdis, rowIdx, sen, bea)
//	for i := ks; i <= ke; i++ {
//		gridRow[i] = '#'
//	}
//}

/***** Input *****/

func calcManDis(c1, c2 []int) int {
	return intz.Abs(c1[0]-c2[0]) + intz.Abs(c1[1]-c2[1])
}

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)

	gInpSen = make([][]int, len(lines))
	gInpBea = make([][]int, len(lines))
	maxX, maxY := math.MinInt, math.MinInt
	minX, minY := math.MaxInt, math.MaxInt
	px, py := 0, 0
	var tokens []string
	for i := 0; i < len(lines); i++ {
		tokens = stringz.SplitMulti(lines[i], []string{" ", "=", ":", ","})
		px = stringz.AtoiQ(tokens[3], math.MinInt)
		py = stringz.AtoiQ(tokens[6], math.MinInt)
		maxX, maxY = intz.Max(maxX, px), intz.Max(maxY, py)
		minX, minY = intz.Min(minX, px), intz.Min(minY, py)
		gInpSen[i] = []int{px, py}

		px = stringz.AtoiQ(tokens[13], math.MinInt)
		py = stringz.AtoiQ(tokens[16], math.MinInt)
		maxX, maxY = intz.Max(maxX, px), intz.Max(maxY, py)
		minX, minY = intz.Min(minX, px), intz.Min(minY, py)
		gInpBea[i] = []int{px, py}
	}
	gBoundTL = []int{minX, minY}
	gBoundBR = []int{maxX, maxY}
	mdis := calcManDis(gBoundTL, gBoundBR) + 10
	gOffset = []int{mdis, mdis}
	offsetAll(gInpBea, gOffset[0], gOffset[1])
	offsetAll(gInpSen, gOffset[0], gOffset[1])
	offsetAll([][]int{gBoundTL}, gOffset[0], gOffset[1])
	offsetAll([][]int{gBoundBR}, gOffset[0], gOffset[1])
	//outils.PrettyArray2DInt(gInpSen)
	//outils.PrettyArray2DInt(gInpBea)
	fmt.Printf("bounds = tl%v , br%v\n", gBoundTL, gBoundBR)
	fmt.Printf("offset = %v\n", gOffset)

}

func offsetAll(points [][]int, ox, oy int) {
	for i, _ := range points {
		points[i][0] += ox
		points[i][1] += oy
	}
}
