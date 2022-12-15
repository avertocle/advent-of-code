package day18

import (
	"fmt"
	. "github.com/avertocle/contests/io/ds/bintree"
	"github.com/avertocle/contests/io/iutils"
	"strconv"

	"log"
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
		safeIncVal(tnLSib, tn.L.V)
		safeIncVal(tnRSib, tn.R.V)
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

func safeIncVal(tn *TNode, inc int) {
	if tn != nil {
		tn.V += inc
	}
}

func isLeafPair(tn *TNode) bool {
	return tn != nil && !isLeaf(tn) &&
		tn.L != nil && isLeaf(tn.L) &&
		tn.R != nil && isLeaf(tn.R)
}

func isLeaf(root *TNode) bool {
	return root.L == nil && root.R == nil
}

func magnitude(root *TNode) int {
	if root == nil {
		return 0
	}
	if isLeaf(root) {
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
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
}

func parsePair(parent *TNode, arr []byte) (int, *TNode, *TNode) {
	if x, err := strconv.Atoi(string(arr)); err == nil {
		return x, nil, nil
	} else if arr[0] == '[' {
		s, e, m := findSplit(arr)
		l := NewTNode(-1, parent)
		l.V, l.L, l.R = parsePair(l, arr[s+1:m])
		r := NewTNode(-1, parent)
		r.V, r.L, r.R = parsePair(r, arr[m+1:e])
		return -1, l, r
	} else {
		fmt.Printf("error : parsePair %v\n", string(arr))
		return -1, nil, nil
	}
}

// return index of "[" & "," & "]" in pattern [<>,<>]
func findSplit(arr []byte) (int, int, int) {
	s, e, sp := 0, math.MaxInt, math.MaxInt
	e = findMatchingEndBrace(arr)
	if arr[1] == '[' {
		sp = findMatchingEndBrace(arr[1:]) + 2
	} else {
		sp = findFirstComma(arr)
	}
	if s+e+sp+1 <= 0 { // maxint will rotate
		fmt.Errorf("error : findSplit %v\n", string(arr))
	}
	return s, e, sp
}

func findFirstComma(arr []byte) int {
	for i, c := range arr {
		if c == ',' {
			return i
		}
	}
	return math.MaxInt
}

// return index of matching ']' for arr[0] (which must be '[)
func findMatchingEndBrace(arr []byte) int {
	ctr := 0
	for i, b := range arr {
		if b == '[' {
			ctr++
		} else if b == ']' {
			ctr--
		}
		if ctr == 0 {
			return i
		}
	}
	return math.MaxInt
}

func parse(arr []byte) *TNode {
	root := NewTNode(-1, nil)
	root.V, root.L, root.R = parsePair(root, arr)
	return root
}
