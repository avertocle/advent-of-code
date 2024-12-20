package day17

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/numz"

	"math"
)

var gInput [][]int

const DirPath = "../2023/day17"

const (
	U = 0
	D = 1
	L = 2
	R = 3
)

func SolveP1() string {
	ans := 0
	//ps := newState(len(gInput), len(gInput[0]), 4, 4, 0, 3)
	//startNode := ps.getNode(0, 0, 0, 0)
	//ps.setCost(startNode, 0)
	//goDijkstra(gInput, startNode, ps)
	//ans = ps.getEndNodeMinCost()
	return fmt.Sprintf("%v", ans)

	//printHM(currState.heatLossMap)
	//currState.setCost(currState.getNode(0, 0, 1, 0), 0)
	//currState.setCost(currState.getNode(0, 0, 2, 0), 0)
	//currState.setCost(currState.getNode(0, 0, 3, 0), 0)

}

func SolveP2() string {
	ans := 0
	ps := newState(len(gInput), len(gInput[0]), 4, 11, 4, 10)
	startNode := ps.getNode(0, 0, 0, 0)
	ps.setCost(startNode, 0)
	goDijkstra(gInput, startNode, ps)
	ans = ps.getEndNodeMinCost()
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

/***** P1 Functions *****/

func goDijkstra(ground [][]int, cn *node, ps *state) {
	if cn == nil {
		return
	}
	nbrs := getValidNeighbours(ground, cn, ps)
	for _, nb := range nbrs {
		h := ground[nb.i][nb.j]
		costCn := ps.getCost(cn)
		costNb := ps.getCost(nb)
		c := 0
		if costCn == math.MaxInt {
			c = costNb
		} else {
			c = numz.Min(costCn+h, costNb)
		}
		ps.setCost(nb, c)
	}
	ps.visit(cn)
	nn := getNextNode(nbrs, ps)
	goDijkstra(ground, nn, ps)
}

func getNextNode(nbrs []*node, ps *state) *node {
	for _, n := range nbrs {
		if !ps.isVisited(n) {
			ps.unvisited[n.keyStr()] = n
		}
	}
	nn, nnCost := (*node)(nil), math.MaxInt
	for _, n := range ps.unvisited {
		c := ps.getCost(n)
		if c < nnCost {
			nn = n
			nnCost = c
		}
	}
	return nn
}

func getValidNeighbours(ground [][]int, cn *node, ps *state) []*node {
	switch cn.v {
	case U:
		return makeNodes(ground, [][]int{
			{cn.i - 1, cn.j, U, cn.d + 1, cn.d, 0},
			{cn.i, cn.j - 1, L, 1, cn.d, 1},
			{cn.i, cn.j + 1, R, 1, cn.d, 1},
		}, ps)
	case D:
		return makeNodes(ground, [][]int{
			{cn.i + 1, cn.j, D, cn.d + 1, cn.d, 0},
			{cn.i, cn.j - 1, L, 1, cn.d, 1},
			{cn.i, cn.j + 1, R, 1, cn.d, 1},
		}, ps)
	case L:
		return makeNodes(ground, [][]int{
			{cn.i, cn.j - 1, L, cn.d + 1, cn.d, 0},
			{cn.i - 1, cn.j, U, 1, cn.d, 1},
			{cn.i + 1, cn.j, D, 1, cn.d, 1},
		}, ps)
	case R:
		return makeNodes(ground, [][]int{
			{cn.i, cn.j + 1, R, cn.d + 1, cn.d, 0},
			{cn.i - 1, cn.j, U, 1, cn.d, 1},
			{cn.i + 1, cn.j, D, 1, cn.d, 1},
		}, ps)
	}
	errz.HardAssert(false, "invalid direction : %v", cn.v)
	return make([]*node, 0)
}

func makeNodes(ground [][]int, indexes [][]int, ps *state) []*node {
	nodes := make([]*node, 0)
	for _, idx := range indexes {
		if idx[0] < 0 || idx[0] >= len(ground) || idx[1] < 0 || idx[1] >= len(ground[0]) {
			continue
		}
		if idx[3] > ps.maxTurnDist {
			continue
		}
		if idx[5] == 1 && idx[4] > 0 && idx[4] < ps.minTurnDist {
			continue
		}
		n := ps.getNode(idx[0], idx[1], idx[2], idx[3])
		errz.HardAssert(n != nil, "invalid node : %v", idx)
		nodes = append(nodes, n)
	}
	return nodes
}

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractInt2DFromString1D(lines, "", nil, -1)
}

/***** Interfaces *****/

type node struct {
	i int
	j int
	v int // direction
	d int // distance
	c int // cost
}

func (n *node) keyStr() string {
	return nodeKeyStr(n.i, n.j, n.v, n.d)
}

func newNode(i, j, v, d int) *node {
	return &node{i: i, j: j, v: v, d: d, c: math.MaxInt}
}

func nodeKeyStr(i, j, v, d int) string {
	return fmt.Sprintf("%v,%v,%v,%v", i, j, v, d)
}

/***/

type state struct {
	heatLossMap [][][][]int
	nodeMap     map[string]*node
	visited     map[string]bool
	unvisited   map[string]*node
	minTurnDist int
	maxTurnDist int
}

func (s *state) visit(n *node) {
	s.visited[n.keyStr()] = true
	delete(s.unvisited, n.keyStr())
}

func (s *state) isVisited(n *node) bool {
	return s.visited[n.keyStr()]
}

func (s *state) getNode(i, j, v, d int) *node {
	return s.nodeMap[nodeKeyStr(i, j, v, d)]
}

func (s *state) getCost(n *node) int {
	return s.heatLossMap[n.i][n.j][n.v][n.d]
}

func (s *state) setCost(n *node, c int) {
	s.heatLossMap[n.i][n.j][n.v][n.d] = c
	n.c = c
}

func (s *state) getEndNodeMinCost() int {
	c, _ := intz.Min2D(s.heatLossMap[len(gInput)-1][len(gInput[0])-1])
	return c
}

func newState(d1, d2, d3, d4, mitd, matd int) *state {
	s := &state{}
	s.heatLossMap = intz.Init4D(len(gInput), len(gInput[0]), d3, d4, math.MaxInt)
	s.nodeMap = make(map[string]*node)
	s.visited = make(map[string]bool)
	s.unvisited = make(map[string]*node)
	s.minTurnDist = mitd
	s.maxTurnDist = matd
	for i := 0; i < d1; i++ {
		for j := 0; j < d2; j++ {
			for v := 0; v < d3; v++ {
				for d := 0; d < d4; d++ {
					n := newNode(i, j, v, d)
					s.nodeMap[n.keyStr()] = n
					s.setCost(n, math.MaxInt)
				}
			}
		}
	}
	return s
}

/****** Debug ******/

func printNodes(nodes []*node) {
	for _, n := range nodes {
		fmt.Printf("[%v,%v] %v %v = %v\n", n.i, n.j, n.v, n.d, n.c)
	}
}

func printHM(hm [][][][]int) {
	for i := 0; i < len(hm); i++ {
		for j := 0; j < len(hm[0]); j++ {
			val, _ := intz.Min2D(hm[i][j])
			if val == math.MaxInt {
				fmt.Printf("%3s ", "x")
			} else {
				fmt.Printf("%3d ", val)
			}
		}
		fmt.Println()
	}
}
