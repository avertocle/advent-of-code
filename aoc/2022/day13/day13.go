package day13

import (
	"fmt"
	"github.com/avertocle/contests/io/boolz"
	"github.com/avertocle/contests/io/bytez"
	. "github.com/avertocle/contests/io/ds/ntree"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/outils"
	"strconv"
)

var gInput [][]string
var gPairCount int

func SolveP1() string {
	var r1, r2 *TNode
	areOrdered := boolz.Init1D(gPairCount, false)
	for i := 0; i < gPairCount; i++ {
		r1 = parse([]byte(gInput[0][i]))
		prettyPrint(r1, fmt.Sprintf("r1 - %v", i+1))
		r1 = parse([]byte(gInput[1][i]))
		prettyPrint(r2, fmt.Sprintf("r2 - %v", i+1))
		if checkOrder(r1, r2) {
			areOrdered[i] = true
		}
	}
	ans := boolz.Count1D(areOrdered, true)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

func checkOrder(r1, r2 *TNode) bool {
	return true
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
	outils.PrettyArray2DString(gInput)
}

func parse(arr []byte) *TNode {
	fmt.Println(string(arr))
	r := NewTNode(-1, nil)
	r.AddChildren(parseList(r, arr))
	return r
}

func parseList(parent *TNode, arr []byte) []*TNode {
	tn := NewTNode(-1, parent)
	if x, err := strconv.Atoi(string(arr)); err == nil {
		tn.V = x
		return []*TNode{tn}
	}
	if len(arr) == 0 {
		return []*TNode{}
	}

	s := 0
	e := bytez.FindNestedMatch(arr, ']')
	if e == -1 {
		fmt.Printf("==> 0. %v\n", string(arr))
	}
	sepIdxs := make([]int, 0)
	ptr := s
	sepIdx := 0
	for ptr < e {
		if arr[ptr] == '[' {
			sepIdx = bytez.FindNestedMatch(arr[ptr+1:], ']')
			if sepIdx == -1 {
				fmt.Printf("==> 1. %v\n", string(arr[ptr:]))
			}
		} else {
			sepIdx = bytez.FindFirst(arr[ptr:], ',')
			if sepIdx == -1 {
				fmt.Printf("==> 2. %v\n", string(arr[ptr:]))
			}
		}
		sepIdxs = append(sepIdxs, sepIdx)
		ptr = sepIdx + 1
		break
	}
	sepIdxs = append(sepIdxs, e)
	s = 1
	fmt.Printf("==> sepIdxs (%+v)\n", sepIdxs)
	for _, si := range sepIdxs {
		tn.AddChildren(parseList(tn, arr[s:si]))
		s = si + 1
	}
	return []*TNode{}
}

func prettyPrint(r *TNode, msg string) {
	fmt.Println(msg)
	PrintFlattenedLeafOnly(r)
	fmt.Println()
}
