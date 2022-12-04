// package main

// import (
// 	"fmt"
// 	"log"
// 	"strings"

// 	"github.com/avertocle/contests/iutils"
// 	"github.com/avertocle/contests/metrics"
// )

// const inputFilePath = "input_small.txt"

// func main() {
// 	metrics.ProgStart()
// 	iutils := getInputOrDie()
// 	metrics.InputLen(iutils.size)

// 	metrics.ProgEnd()
// 	fmt.Printf("metrics : [%v]", metrics.ToString())
// }

// func problem1() {

// }

// func problem2() {

// }

// func getInputOrDie() *graph {
// 	lines, err := iutils.FromFile(inputFilePath, true)
// 	if err != nil {
// 		log.Fatalf("iutils error | %v", err)
// 	}

// 	tokens := make([]string, 2)
// 	nodeSet := make(map[string]bool, 0)
// 	paths := iutils.Init2DString(len(lines), 2)
// 	for i, l := range lines {
// 		tokens = strings.Split(l, "-")
// 		//fmt.Printf("%v,%v|", tokens[0], tokens[1])
// 		if _, ok := nodeSet[tokens[0]]; !ok {
// 			nodeSet[tokens[0]] = true
// 		}
// 		if _, ok := nodeSet[tokens[1]]; !ok {
// 			nodeSet[tokens[1]] = true
// 		}
// 		paths[i][0] = tokens[0]
// 		paths[i][1] = tokens[1]
// 	}

// 	iutils := newGraph(len(nodeSet))
// 	for _, l := range lines {
// 		tokens = strings.Split(l, "-")
// 		iutils.addPaths(paths)
// 	}

// 	iutils.PrettyArray2DByte(iutils.mat)
// 	fmt.Printf("%v\n", iutils.iToCap)

// 	return iutils
// }

// type graph struct {
// 	vToI   map[string]int
// 	iToV   map[int]string
// 	iToCap map[int]bool
// 	mat    [][]byte
// 	size   int
// 	nc     int // current node count
// }

// func newGraph(size int) *graph {
// 	g := new(graph)
// 	g.size = size
// 	g.mat = iutils.Init2DByte(size, size, 0)
// 	g.vToI = make(map[string]int)
// 	g.iToV = make(map[int]string)
// 	g.iToCap = make(map[int]bool)
// 	return g
// }

// func (g *graph) addPaths(paths [][]string) {
// 	for _, rows := range paths {
// 		g.addNode(rows[0])
// 		g.addNode(rows[1])
// 		g.mat[g.vToI[rows[0]]][g.vToI[rows[1]]] = 1
// 		g.mat[g.vToI[rows[1]]][g.vToI[rows[0]]] = 1
// 	}
// }

// func (g *graph) addNode(node string) {
// 	if _, ok := g.vToI[node]; ok {
// 		return
// 	} else {
// 		g.vToI[node] = g.nc
// 		g.iToV[g.nc] = node
// 		if strings.Compare(node, strings.ToUpper(node)) == 0 {
// 			g.iToCap[g.nc] = true
// 		}
// 		g.nc++
// 	}
// }
