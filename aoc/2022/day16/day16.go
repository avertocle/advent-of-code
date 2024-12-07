package day16

import (
	"fmt"
	"github.com/avertocle/contests/io/ds"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"math"
	"strings"
)

var gInput *ds.Graph

func SolveP1() string {
	spaths := findShortestPaths()
	nonZeroV := findNonZeroValves()
	opened := make(map[string]int)
	ans := traverse("AA", 30, nonZeroV, opened, spaths, 0)
	return fmt.Sprintf("%v", ans)
}

func traverse(node string, timeRem int, nonZeroV, opened, spaths map[string]int, depth int) int {
	if timeRem <= 0 {
		return 0
	}
	var maxP, p, currP int
	var wasOpened bool
	if canOpen(node, opened) {
		timeRem-- // valve open time cost
		currP = gInput.VMap[node] * (timeRem)
		open(node, opened)
		wasOpened = true
	}
	nbrs := findNextUnopened(nonZeroV, opened)
	plen := 0
	for _, nbr := range nbrs {
		plen = getSPath(node, nbr, spaths)
		p = traverse(nbr, timeRem-plen, nonZeroV, opened, spaths, depth+1) // valve travel time cost
		if p > maxP {
			maxP = p
		}
	}
	if wasOpened {
		close(node, opened)
	}
	//fmt.Printf("%v vis %v, tr %v, p %v\n", strings.Repeat(" ", 2*depth), node, timeRem, currP)
	return currP + maxP
}

func findNextUnopened(nonZeroV, opened map[string]int) []string {
	next := make([]string, 0)
	for v, _ := range nonZeroV {
		if _, ok := opened[v]; !ok {
			next = append(next, v)
		}
	}
	return next
}

func findNonZeroValves() map[string]int {
	nzv := make(map[string]int)
	for v, flow := range gInput.VMap {
		if flow > 0 {
			nzv[v] = flow
		}
	}
	return nzv
}

func sp(snode, enode string, vis map[string]int) int {
	if strings.Compare(snode, enode) == 0 {
		return 0
	}
	p, minP := 0, math.MaxInt/2
	vis[snode] = 1
	nbrs := gInput.AdList[snode]
	for nbr, _ := range nbrs {
		if _, ok := vis[nbr]; ok {
			continue
		}
		p = sp(nbr, enode, vis)
		if p < minP {
			minP = p
		}
	}
	delete(vis, snode)
	return minP + 1
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func open(node string, vis map[string]int) {
	vis[node] = 1
}

func close(node string, vis map[string]int) {
	delete(vis, node)
}

func canOpen(node string, opened map[string]int) bool {
	return gInput.VMap[node] > 0 && !isOpen(node, opened)
}

func isOpen(node string, opened map[string]int) bool {
	_, ok := opened[node]
	return ok
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)

	gInput = ds.NewGraph()
	var tokens []string
	var rate int
	var v string
	var adjMap map[string]int
	for i := 0; i < len(lines); i++ {
		tokens = stringz.SplitMulti(lines[i], []string{" ", "=", ";", ","})
		rate = stringz.AtoI(tokens[5], math.MinInt)
		v = tokens[1]
		adjMap = make(map[string]int)
		for j := 11; j < len(tokens); j += 2 {
			adjMap[tokens[j]] = 1
		}
		gInput.AddVertex(v, rate, adjMap)
	}
}

func findShortestPaths() map[string]int {
	m := make(map[string]int)
	vlist := gInput.VList()
	n2 := ""
	p := math.MaxInt
	for i := 0; i < len(vlist); i++ {
		n1 := vlist[i]
		for j := 0; j < len(vlist); j++ {
			if i == j {
				continue
			}
			n2 = vlist[j]
			p = sp(n1, n2, make(map[string]int))
			k1, k2 := keys(n1, n2)
			m[k1] = p
			m[k2] = p
		}
	}
	return m
}

func keys(n1, n2 string) (string, string) {
	return fmt.Sprintf("%v-%v", n1, n2), fmt.Sprintf("%v-%v", n2, n1)
}

func getSPath(n1, n2 string, spaths map[string]int) int {
	k, _ := keys(n1, n2)
	return spaths[k]
}
