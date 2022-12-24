package day13

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	. "github.com/avertocle/contests/io/ds/ntree"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"strconv"
)

var gInput [][]string
var gPairCount int

func SolveP1() string {
	var r1, r2 *TNode
	var r1str, r2str string
	ans := 0
	for i := 0; i < gPairCount; i++ {
		r1str, r2str = gInput[0][i], gInput[1][i]
		r1 = parse(nil, []byte(r1str))
		r2 = parse(nil, []byte(r2str))
		res := compare(r1, r2)
		if res == 1 {
			ans += i + 1
		} else if res == 0 {
			fmt.Printf("%v : res(%v) : %v vs %v\n", i+1, res, r1str, r2str)
		}
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	//inpRows := len(gInput)
	rnlen := gPairCount*2 + 2
	rn := make([]*TNode, rnlen)
	for i, row := range gInput {
		for j, cell := range row {
			rn[i*gPairCount+j] = parse(nil, []byte(cell))
		}
	}
	rdiv1 := parse(nil, []byte("[[2]]"))
	rdiv2 := parse(nil, []byte("[[6]]"))
	rn[rnlen-2] = rdiv1
	rn[rnlen-1] = rdiv2
	var temp *TNode
	for i := len(rn) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if compare(rn[j], rn[j+1]) == -1 {
				temp = rn[j]
				rn[j] = rn[j+1]
				rn[j+1] = temp
			}
		}
	}
	ans := 1
	for i, r := range rn {
		if r.Id == rdiv1.Id || r.Id == rdiv2.Id {
			ans *= i + 1
		}
	}
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func compare(r1, r2 *TNode) int {
	res := 0
	if r1.IsLeaf() && r2.IsLeaf() {
		return compareVals(r1, r2)
	}
	if r1.IsLeaf() && !r2.IsLeaf() && r1.V > -1 {
		r1.AddC(NewTNode(r1.V, r1))
		r1.V = -1
	}
	if !r1.IsLeaf() && r2.IsLeaf() && r2.V > -1 {
		r2.AddC(NewTNode(r2.V, r2))
		r2.V = -1
	}
	for i := 0; i < len(r1.C) && i < len(r2.C); i++ {
		res = compare(r1.C[i], r2.C[i])
		if res != 0 {
			return res
		}
	}
	return compareLens(r1, r2)
}

func compareLens(r1, r2 *TNode) int {
	if len(r1.C) < len(r2.C) {
		return 1
	} else if len(r1.C) > len(r2.C) {
		return -1
	} else {
		return 0
	}
}

func compareVals(r1, r2 *TNode) int {
	if r1.V < r2.V {
		return 1
	} else if r1.V > r2.V {
		return -1
	} else {
		return 0
	}
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gPairCount = (len(lines) + 1) / 3
	gInput = [][]string{make([]string, gPairCount), make([]string, gPairCount)}
	for i := 0; i < len(lines); i += 3 {
		gInput[0][i/3] = lines[i]
		gInput[1][i/3] = lines[i+1]
	}
}

func parse(parent *TNode, arr []byte) *TNode {
	if len(arr) == 0 {
		return nil
	}

	tn := NewTNode(-1, parent)
	if x, err := strconv.Atoi(string(arr)); err == nil {
		tn.V = x
		return tn
	}

	e := bytez.FindNestedMatch(arr, ']')
	childExps := findChildExpressions(arr[1:e])
	for _, ce := range childExps {
		tn.AddC(parse(tn, ce))
	}
	return tn
}

func findChildExpressions(arr []byte) [][]byte {
	ans := make([][]byte, 0)
	if len(arr) == 0 {
		return ans
	}

	s, t := 0, 0
	for s < len(arr) {
		if arr[s] == '[' {
			t = s + bytez.FindNestedMatch(arr[s:], ']') + 1 // t at comma after [<>]
		} else {
			t = bytez.FindFirst(arr[s:], ',') // t at comma after V
			if t == -1 {                      // the end digit will not have a comma after it
				t = len(arr)
			} else {
				t = s + t
			}
		}
		ans = append(ans, arr[s:t])
		s = t + 1
	}
	return ans
}
