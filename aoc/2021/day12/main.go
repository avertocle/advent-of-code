package main

import (
	"fmt"
	"github.com/avertocle/contests/io/ds"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/outils"
	"github.com/avertocle/contests/metrics"
	"log"
	"strings"
)

const inputFilePath = "input_final.txt"

type visitCriteriaFunc func(sv string, graph *ds.Graph, visMap map[string]int) bool

func main() {
	metrics.ProgStart()
	input := getInputOrDie()
	input.PrintAdList()

	ans01 := solvePartOne(input)
	fmt.Printf("answer part-01 = %v\n", ans01)

	ans02 := solvePartTwo(input)
	fmt.Printf("answer part-02 = %v\n", ans02)
}

func solvePartOne(iGraph *ds.Graph) int {
	vis := make(map[string]int)
	pathCount := new(int)
	*pathCount = 0
	traverse("start", iGraph, vis, pathCount, canVisitP1)
	return *pathCount
}

func solvePartTwo(iGraph *ds.Graph) int {
	vis := make(map[string]int)
	pathCount := new(int)
	*pathCount = 0
	visSmallTwice := new(bool)
	*visSmallTwice = false
	traverse("start", iGraph, vis, pathCount, canVisitP2)
	return *pathCount
}

func traverse(sv string, graph *ds.Graph, visMap map[string]int,
	pathCount *int, canVisit visitCriteriaFunc) {
	if isEnd(sv) {
		*pathCount++
		return
	}
	if !canVisit(sv, graph, visMap) {
		return
	}
	visit(sv, visMap)
	adj, _ := graph.AdList[sv]
	for v, _ := range adj {
		traverse(v, graph, visMap, pathCount, canVisit)
	}
	unVisit(sv, visMap)
}

/***** PART 01 Functions *****/

func canVisitP1(v string, graph *ds.Graph, visMap map[string]int) bool {
	if !isVisited(v, visMap) {
		return true
	} else if isBigCave(v) {
		return true
	} else {
		return false
	}
}

/***** PART 02 Functions *****/

func canVisitP2(v string, graph *ds.Graph, visMap map[string]int) bool {
	if !isVisited(v, visMap) {
		return true
	} else if isBigCave(v) {
		return true
	} else if isSmallCave(v) && !hasAnySmallCaveVisitedTwice(visMap) {
		return true
	} else {
		return false
	}
}

func hasAnySmallCaveVisitedTwice(visMap map[string]int) bool {
	for v, count := range visMap {
		if isSmallCave(v) && count > 1 {
			return true
		}
	}
	return false
}

/***** Common Functions *****/

func isStart(v string) bool {
	return strings.Compare(v, "start") == 0
}

func isEnd(v string) bool {
	return strings.Compare(v, "end") == 0
}

func isVisited(v string, visMap map[string]int) bool {
	_, ok := visMap[v]
	return ok
}

func isBigCave(v string) bool {
	b := []byte(v)[0]
	return b >= 'A' && b <= 'Z'
}

func isSmallCave(v string) bool {
	return !isBigCave(v) && !isStart(v) && !isEnd(v)
}

func visit(v string, visMap map[string]int) {
	if c, ok := visMap[v]; !ok {
		visMap[v] = 1
	} else {
		visMap[v] = c + 1
	}
}

func unVisit(v string, visMap map[string]int) {
	c, ok := visMap[v]
	if !ok || c < 1 {
		fmt.Errorf("invalid viscount %v %+v", v, visMap)
		return
	} else if c == 1 {
		delete(visMap, v)
	} else {
		visMap[v] = c - 1
	}
}

/***** Input *****/

// input : [][]int : each row contains start and end ranges of both elves {e1s,e1e,e2s,e2e}
func getInputOrDie() *ds.Graph {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	linesSplit := iutils.ExtractString2DFromString1D(lines, "-", nil, "")
	outils.PrettyArray2DString(linesSplit)
	input := ds.NewGraph()
	for _, ls := range linesSplit {
		input.AddVertex(ls[0], 1, arrToMap(ls[1]))
		input.AddVertex(ls[1], 1, arrToMap(ls[0]))
	}
	return input
}

func arrToMap(arr ...string) map[string]int {
	m := make(map[string]int)
	for _, a := range arr {
		m[a] = 1
	}
	return m
}
