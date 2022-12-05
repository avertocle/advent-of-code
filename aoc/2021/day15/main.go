package main

import (
	"fmt"
	"github.com/avertocle/contests/io/geom"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/outils"
	"log"
	"math"
	"sort"

	"github.com/avertocle/contests/metrics"
)

const inputFilePath = "input.txt"

func main() {
	metrics.ProgStart()
	input := getInputOrDie()
	metrics.InputLen(len(input))

	ans := problem1(input)
	fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

var pathsComp int
var pathsDisc int
var pathCache [][]int

func problem1(graph [][]byte) int {
	rows := len(graph)
	cols := len(graph[0])
	pathCache = intz.Init2D(-1, rows, cols)
	pathCache[rows-1][cols-1] = int(graph[rows-1][cols-1] - '0')
	fmt.Printf("%v,%v\n", rows, cols)
	findMinPathCost(graph, 0, 0, rows, cols)
	outils.PrettyArray2DInt(pathCache)
	return pathCache[0][0]
}

func findMinPathCost(graph [][]byte, x, y, rows, cols int) int {
	if !geom.IsValidCoord2D(x, y, rows, cols) {
		return math.MaxInt32
	} else if pathCache[x][y] >= 0 {
		return pathCache[x][y]
	} else if x == rows-1 && y == cols-1 {
		return int(graph[x][y] - '0')
	} else {
		x0 := findMinPathCost(graph, x+1, y, rows, cols)
		x1 := findMinPathCost(graph, x, y+1, rows, cols)
		pathCache[x][y] = int(graph[x][y]-'0') + minInt(x0, x1)
		return pathCache[x][y]
	}
}

func problem2() {

}

func minInt(x ...int) int {
	sort.Ints(x)
	return x[0]
}

func getInputOrDie() [][]byte {
	lines, err := iutils.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}

	input := iutils.ExtractByte2DFromString1D(lines, "", nil, 0)
	return input
}
