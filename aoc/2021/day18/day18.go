package day18

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	. "github.com/avertocle/contests/io/ds/bintree"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"strconv"

	"math"
)

var gInput [][]byte

func SolveP1() string {
	var r1, r2, rs *TNode
	r1 = parse(gInput[0])
	for i := 1; i < len(gInput); i++ {
		r2 = parse(gInput[i])
		rs = sumAndReduce(r1, r2)
		r1 = rs
	}
	//prettyPrint(rs, "final : ")
	mag := magnitude(rs)
	return fmt.Sprintf("%v", mag)
}

func SolveP2() string {
	var r1, r2, rs *TNode
	var mag, magMax int
	for i := 0; i < len(gInput); i++ {
		for j := 0; j < len(gInput); j++ {
			if i != j {
				r1 = parse(gInput[i])
				r2 = parse(gInput[j])
				rs = sumAndReduce(r1, r2)
				mag = magnitude(rs)
				if mag > magMax {
					magMax = mag
				}
			}
		}
	}
	return fmt.Sprintf("%v", magMax)
}

func sumAndReduce(r1, r2 *TNode) *TNode {
	rs := NewTNode(-1, nil)
	rs.L, rs.R = r1, r2
	r1.P, r2.P = rs, rs
	for reduce(rs) {
	}
	return rs
}

func reduce(root *TNode) bool {
	didReduce := false
	for explodeAll(root, root, 1) {
		didReduce = true
		//prettyPrint(root, "explode : ")
	}
	if didReduce = split(root); didReduce {
		//prettyPrint(root, "split : ")
	}
	return didReduce
}

func explodeAll(root, tn *TNode, depth int) bool {
	if tn == nil {
		return false
	}
	if depth > 4 && isLeafPair(tn) {
		tnLSib, tnRSib := findSiblings(root, tn)
		incIfNotNil(tnLSib, tn.L.V)
		incIfNotNil(tnRSib, tn.R.V)
		tn.L, tn.R, tn.V = nil, nil, 0
		return true
	}
	if explodeAll(root, tn.L, depth+1) {
		return true
	} else if explodeAll(root, tn.R, depth+1) {
		return true
	} else {
		return false
	}
}

func split(tn *TNode) bool {
	if tn == nil {
		return false
	}
	if tn.V > 9 {
		tn.L = NewTNode(tn.V/2, tn)
		tn.R = NewTNode(tn.V-tn.V/2, tn)
		tn.V = -1
		return true
	}
	if split(tn.L) {
		return true
	} else if split(tn.R) {
		return true
	} else {
		return false
	}
}

func findSiblings(root, tn *TNode) (*TNode, *TNode) {
	flatTree := FlattenLeafOnly(root)
	lIdx, rIdx := 0, 0
	for i := 0; i < len(flatTree); i++ {
		if tn.L.CompareTo(flatTree[i]) {
			lIdx = i
		}
		if tn.R.CompareTo(flatTree[i]) {
			rIdx = i
		}
	}
	var tnLSib, tnRSib *TNode
	if lIdx > 0 && lIdx < len(flatTree) {
		tnLSib = flatTree[lIdx-1]
	}
	if rIdx >= 0 && rIdx < len(flatTree)-1 {
		tnRSib = flatTree[rIdx+1]
	}
	return tnLSib, tnRSib
}

/***** Common Functions *****/

func incIfNotNil(tn *TNode, inc int) {
	if tn != nil {
		tn.V += inc
	}
}

func isLeafPair(tn *TNode) bool {
	return tn != nil && !tn.IsLeaf() &&
		tn.L != nil && tn.L.IsLeaf() &&
		tn.R != nil && tn.R.IsLeaf()
}

func magnitude(root *TNode) int {
	if root == nil {
		return 0
	}
	if root.IsLeaf() {
		return root.V
	}
	return 3*magnitude(root.L) + 2*magnitude(root.R)
}

func prettyPrint(tn *TNode, msg string) {
	fmt.Printf(msg)
	PrintInorderLeafOnly(tn)
	fmt.Println()
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "parseInput error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
}

func parsePair(parent *TNode, arr []byte) *TNode {
	tn := NewTNode(-1, parent)
	if x, err := strconv.Atoi(string(arr)); err == nil {
		tn.V = x
		return tn
	}
	s, e, m := findSplit(arr)
	tn.L = parsePair(tn, arr[s+1:m])
	tn.R = parsePair(tn, arr[m+1:e])
	return tn
}

// return index of "[" & "," & "]" in pattern [<>,<>]
func findSplit(arr []byte) (int, int, int) {
	s, e, sp := 0, math.MinInt, math.MinInt
	e = bytez.FindNestedMatch(arr, ']')
	if arr[1] == '[' {
		sp = bytez.FindNestedMatch(arr[1:], ']') + 2
	} else {
		sp = bytez.FindFirst(arr, ',')
	}
	errz.HardAssert(e >= 0 && sp >= 0, "error : findSplit %v\n", string(arr))
	return s, e, sp
}

func parse(arr []byte) *TNode {
	root := parsePair(nil, arr)
	return root
}
