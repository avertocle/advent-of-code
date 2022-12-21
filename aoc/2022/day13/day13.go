package day13

import (
	"fmt"
	"github.com/avertocle/contests/io/boolz"
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
	areOrdered := boolz.Init1D(gPairCount, false)
	ans := 0
	for i := 0; i < gPairCount; i++ {
		r1 = parse(nil, []byte(gInput[0][i]))
		//prettyPrint(r1, fmt.Sprintf("%v : r1 : raw = %v : parsed = ", i+1, string([]byte(gInput[0][i]))))
		r2 = parse(nil, []byte(gInput[1][i]))
		//prettyPrint(r2, fmt.Sprintf("%v : r2 : raw = %v : parsed = ", i+1, string([]byte(gInput[1][i]))))
		res := compare(r1, r2)
		if res == 1 {
			areOrdered[i] = true
			ans += i + 1
		} else if res == 0 {
			fmt.Printf("%v : res(%v) : %v vs %v\n", i+1, res,
				string([]byte(gInput[0][i])), string([]byte(gInput[1][i])))
		}

	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

func compare(r1, r2 *TNode) int {

	res := 0

	if r1.IsLeaf() && r2.IsLeaf() {
		return compareVals(r1, r2)
	}
	if r1.IsLeaf() && !r2.IsLeaf() {
		r1.AddC(NewTNode(r1.V, r1))
		r1.V = -1
	}
	if !r1.IsLeaf() && r2.IsLeaf() {
		r2.AddC(NewTNode(r2.V, r2))
		r2.V = -1
	}
	r1c := r1.C
	r2c := r2.C
	for i := 0; i < len(r1c) && i < len(r2c); i++ {
		res = compare(r1c[i], r2c[i])
		if res != 0 {
			return res
		}
	}
	if len(r1c) < len(r2c) {
		return 1
	} else if len(r1c) > len(r2c) {
		return -1
	} else {
		return 0
	}
}

//func splitAndCompare(r1, r2 *TNode) int {
//	r :=
//}

func compareVals(r1, r2 *TNode) int {
	if r1.V < r2.V {
		return 1
	} else if r1.V > r2.V {
		return -1
	} else {
		return 0
	}
}

/***** Common Functions *****/

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
	//outils.PrettyArray2DString(gInput)
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
	//fmt.Printf("findChildExpressions = (%v) len(%v)\n", string(arr), len(arr))
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
		//fmt.Printf("findChildExpressions : %v to %v = ", s, t)
		//fmt.Printf("%v\n", string(arr[s:t]))
		ans = append(ans, arr[s:t])
		s = t + 1
	}
	return ans
}

func prettyPrint(r *TNode, msg string) {
	fmt.Printf("%v %v\n", msg, GetFlatStringLeafOnly(r))
	GetFlatStringLeafOnly(r)
}
