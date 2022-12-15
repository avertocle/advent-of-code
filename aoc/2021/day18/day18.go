package day18

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"log"
	"math"
	"strings"
)

var gInput [][]byte

func SolveP1() string {
	var r1, r2, rs *tnode
	r1 = parse(gInput[0])
	//printTree(parse(gInput[2]), 0)
	for i := 1; i < len(gInput); i++ {
		fmt.Printf("adding %v and %v\n", i-1, i)
		r2 = parse(gInput[i])
		rs = sum(r1, r2)
		//printTree(rs, 0)
		for reduce(rs, 1) {
			fmt.Printf("reducing %v and %v\n", i-1, i)
		}
		//printTree(rs, 0)
		r1 = rs
		//break
	}
	printTree(r1, 0)
	ans := magnitude(r1)
	return fmt.Sprintf("%v", ans)
}

func magnitude(root *tnode) int {
	if root.v != -1 {
		return root.v
	}
	return 3*magnitude(root.l) + 2*magnitude(root.r)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

func parse(arr []byte) *tnode {
	root := newNode(-1, nil)
	root.v, root.l, root.r = parsePair(root, arr)
	return root
}

func sum(r1, r2 *tnode) *tnode {
	rs := newNode(-1, nil)
	rs.l = r1
	rs.r = r2
	return rs
}

func isLeafPair(root *tnode) bool {
	if root.v != -1 {
		return false
	}
	if root.l == nil || root.r == nil {
		log.Fatalf("error : isLeafPair : invalid prog state root(%+v)", root)
	}
	return root.l.v > -1 && root.r.v > -1
}

func reduce(root *tnode, depth int) bool {
	if root == nil {
		return false
	}
	if depth > 4 && isLeafPair(root) {
		explode(root)
		return true
	}
	if root.v > 9 {
		split(root)
		return false
	}

	if reduce(root.l, depth+1) {
		return true
	}
	if reduce(root.r, depth+1) {
		return true
	}
	return false
}

func explode(r *tnode) {
	incLeft(r.l, r.v)
	incRight(r.r, r.v)
	r.v = 0
	r.l = nil
	r.r = nil
}

func incLeft(r *tnode, inc int) {
	if r == nil {
		return
	} else if r.v != -1 {
		r.v += inc
		return
	} else if r.p == nil {
		return
	} else {
		incLeft(r.p.l, inc)
	}
}

func incRight(r *tnode, inc int) {
	if r == nil {
		return
	} else if r.v != -1 {
		r.v += inc
		return
	} else if r.p == nil {
		return
	} else {
		incRight(r.p.r, inc)
	}
}

func split(root *tnode) bool {
	if root.v <= 9 {
		return false
	}
	root.l = newNode(root.v/2, root)
	root.r = newNode(root.v-root.v/2, root)
	return true

}

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	gInput = iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
}

func parsePair(parent *tnode, arr []byte) (int, *tnode, *tnode) {
	//fmt.Printf("===> parsing %v\n", string(arr))
	if x := stringz.AtoiQ(string(arr), -1); x > -1 {
		return x, nil, nil
	} else if arr[0] == '[' {
		s, e, sp := findSplit(arr)
		//fmt.Printf("===> split %v, %v, %v\n\n", s, e, sp)
		l := newNode(-1, parent)
		l.v, l.l, l.r = parsePair(l, arr[s+1:sp])
		r := newNode(-1, parent)
		r.v, r.l, r.r = parsePair(r, arr[sp+1:e])
		return -1, l, r
	} else {
		fmt.Printf("error : parsePair %v\n", string(arr))
		return -1, nil, nil
	}
}

// return index of "[" & "," & "]" in pattern [<>,<>]
func findSplit(arr []byte) (int, int, int) {
	s, e, sp := math.MaxInt, math.MaxInt, math.MaxInt
	s = 0
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
	fmt.Printf("error : findFirstComma arr (%v)\n", string(arr))
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
	fmt.Printf("error : findMatchingEndBrace arr (%v)\n", string(arr))
	return math.MaxInt
}

/***** Structs *****/

type tnode struct {
	v int
	l *tnode
	r *tnode
	p *tnode
}

func newNode(v int, p *tnode) *tnode {
	return &tnode{
		v: v,
		l: nil,
		r: nil,
		p: p,
	}
}

func printTree(r *tnode, depth int) {
	if r == nil {
		return
	}
	printWithDepth(fmt.Sprintf("(%v)", r.v), depth)
	printTree(r.l, depth+1)
	printTree(r.r, depth+1)
}

func printWithDepth(s string, d int) {
	fmt.Printf("%v- %v\n", strings.Repeat(" ", 2*d), s)
}
