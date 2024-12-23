package day23

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/ds"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/mapz"
	"slices"
	"strings"
)

const DirPath = "../2024/day23"

var gInput *ds.Graph

func SolveP1() string {
	ans := 0
	allNodes := gInput.VList()
	for i := 0; i < len(allNodes); i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {
				v1, v2, v3 := allNodes[i], allNodes[j], allNodes[k]
				if hasAnyStartingWith("t", v1, v2, v3) &&
					areConnected(gInput, v1, v2, v3) {
					ans++
				}
			}
		}
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	allNodes := gInput.VList()
	maxClique := make([]string, 0)
	findMaximumClique([]string{}, allNodes, []string{}, &maxClique)
	slices.Sort(maxClique)
	ans := strings.Join(maxClique, ",")
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func areConnected(g *ds.Graph, vids ...string) bool {
	for i := 0; i < len(vids); i++ {
		for j := 0; j < i; j++ {
			if !g.AreConnected(vids[i], vids[j]) {
				return false
			}
		}
	}
	return true
}

func hasAnyStartingWith(prefix string, arr ...string) bool {
	for _, s := range arr {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

/***** P2 Functions *****/

/*
findMaximumClique uses
Bron–Kerbosch algorithm for finding all maximal cliques in an undirected graph.

BronKerbosch(R, P, X)

	if P and X are both empty then
	    report R as a maximal clique
	for each vertex v in P do
	    BronKerbosch(R ⋃ {v}, P ⋂ N(v), X ⋂ N(v))
	    P := P \ {v}
	    X := X ⋃ {v}
*/
func findMaximumClique(R, P, X []string, A *[]string) {
	//fmt.Println(R, P, X, " = ", A)
	if len(P) == 0 && len(X) == 0 {
		if len(R) > len(*A) {
			*A = append([]string{}, R...)
		}
		return
	}
	for _, v := range P {
		ruv := append(R, v)
		nv := mapz.Keys(gInput.AdList[v])
		pinv := arrz.Intersection1D[string](P, nv)
		xinv := arrz.Intersection1D[string](X, nv)
		findMaximumClique(ruv, pinv, xinv, A)
		P = arrz.RemoveElementByVal1D(P, v)
		X = append(X, v)
	}
	return
}

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = ds.NewGraph()
	for _, line := range lines {
		t := strings.Split(line, "-")
		gInput.AddVertex(t[0], 1, make(map[string]int))
		gInput.AddVertex(t[1], 1, make(map[string]int))
		gInput.AddConnection(t[0], t[1], 1)
	}
}
