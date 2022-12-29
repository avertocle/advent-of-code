package day20

import (
	"fmt"
	"github.com/avertocle/contests/io/ds/ll/cdll"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"math"
)

var gInput []int

func SolveP1() string {
	sn := cdll.FromArray(gInput)
	idxs := makeIndices(sn)
	mixAllNTimes(idxs, 1)
	keyIdxs := []int{1000, 2000, 3000}
	ans := calcCoordSum(idxs, keyIdxs)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	sn := cdll.FromArray(gInput)
	idxs := makeIndices(sn)
	applyDecKey(idxs, 811589153)
	mixAllNTimes(idxs, 10)
	keyIdxs := []int{1000, 2000, 3000}
	ans := calcCoordSum(idxs, keyIdxs)
	return fmt.Sprintf("%v", ans)
}

func mixAllNTimes(idxs []*cdll.Node, mixN int) {
	for n := 0; n < mixN; n++ {
		for i := 0; i < len(gInput); i++ {
			mixOnePoint(i, idxs)
		}
	}
}

func mixOnePoint(idx int, inpIdxs []*cdll.Node) {
	cn := inpIdxs[idx]
	dis := intz.Abs(cn.V) % (len(gInput) - 1)
	if dis == 0 {
		return
	}
	var fn *cdll.Node
	if cn.V > 0 {
		fn = cdll.NavFwd(cn, dis)
	} else {
		fn = cdll.NavRev(cn, dis+1)
	}
	cdll.DelMe(cn)
	cdll.AddAfterMe(fn, cn)
}

func calcCoordSum(idxs []*cdll.Node, keyIdxs []int) int {
	zeros := cdll.FindNodesByVal(idxs[0], 0)
	errz.HardAssert(len(zeros) == 1, "error : %v zeros in mixed, exp 1", len(zeros))
	zn := zeros[0]
	disRot, coordSum := -1, 0
	coords := make([]int, len(keyIdxs))
	for i, idx := range keyIdxs {
		disRot = idx % len(gInput)
		coords[i] = cdll.NavFwd(zn, disRot).V
		coordSum += coords[i]
	}
	//fmt.Printf("res : zero(%v) coords(%v) sum(%v)\n", zn.V, coords, coordSum)
	return coordSum
}

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractInt1DFromString1D(lines, ",", 0, math.MinInt)
}

func makeIndices(sn *cdll.Node) []*cdll.Node {
	idxs := make([]*cdll.Node, len(gInput))
	i := 0
	t := sn
	for {
		idxs[i] = t
		i++
		t = t.N
		if t.Id == sn.Id {
			break
		}
	}
	return idxs
}

func applyDecKey(idxs []*cdll.Node, key int) {
	for _, tn := range idxs {
		tn.V *= key
	}
}
