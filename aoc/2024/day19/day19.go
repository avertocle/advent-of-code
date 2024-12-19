package day19

import (
	"fmt"
	"github.com/avertocle/contests/io/ds"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
)

const DirPath = "../2024/day19"

var gTowels [][]byte
var gDesigns [][]byte

// 400, 388 too high

func SolveP1() string {
	ans := 0
	graph, startNodes, endNodes := makeGraph()
	for _, design := range gDesigns {
		found := findDesignInGraph(graph, startNodes, endNodes, design)
		fmt.Println("===> ", string(design), found)
		if found {
			ans++
		}
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	return fmt.Sprintf("%v", ans)
}

func findDesignInGraph(g *ds.Graph, startNodes, endNodes map[string]bool, design []byte) bool {
	allNodes := g.FindVertexesByValue(int(design[0]))
	starts := make([]string, 0)
	for _, n := range allNodes {
		if startNodes[n] {
			starts = append(starts, n)
		}
	}
	for _, node := range starts {
		//fmt.Printf("\n\n--%v %v--\n\n", node, starts)
		if isDesignInGraph(g, startNodes, endNodes, design, node) {
			return true
		}
	}
	return false
}

func isDesignInGraph(g *ds.Graph, startNodes, endNodes map[string]bool, design []byte, currId string) bool {
	//fmt.Printf("%v | %q vs %v,%q | processing \n", string(design), string(design[0]), currId, g.VMap[currId])
	if len(design) == 0 {
		fmt.Println("what the hell")
		return true
	}
	if g.VMap[currId] != int(design[0]) {
		//fmt.Printf("==> %v | %q vs %v,%q | not matched \n", string(design), string(design[0]), currId, g.VMap[currId])
		return false
	}
	if len(design) == 1 {
		if endNodes[currId] {
			//fmt.Printf("==> %v | %q vs %v,%q | matched end-node \n", string(design), string(design[0]), currId, g.VMap[currId])
			return true
		} else {
			//fmt.Printf("==> %v | %q vs %v,%q | matched non-end-node \n", string(design), string(design[0]), currId, g.VMap[currId])
			return false
		}
	}
	nextNodes, _ := g.AdList[currId]
	//fmt.Printf("==> %v | recur to find %v starting in %v \n", string(design), string(design[1:]), g.MapToStr(nextNodes))
	for nextId, _ := range nextNodes {
		if isDesignInGraph(g, startNodes, endNodes, design[1:], nextId) {
			return true
		}
	}
	return false
}

func makeGraph() (*ds.Graph, map[string]bool, map[string]bool) {
	g := ds.NewGraph()
	startNodes := make(map[string]bool)
	endNodes := make(map[string]bool)
	for i, row := range gTowels {
		for j, cell := range row {
			vkey := fmt.Sprintf("%v-%v", i, j)
			awm := make(map[string]int)
			if j < len(row)-1 {
				awm[fmt.Sprintf("%v-%v", i, j+1)] = 1
			} else {
				for ii, _ := range gTowels {
					awm[fmt.Sprintf("%v-%v", ii, 0)] = 1
				}
			}
			g.AddVertex(vkey, int(cell), awm)
			if j == 0 {
				startNodes[vkey] = true
			}
			if j == len(row)-1 {
				endNodes[vkey] = true
			}
		}
	}
	//fmt.Println("startNodes = ", startNodes)
	//fmt.Println("endNodes = ", endNodes)
	//g.PrintAdList()
	return g, startNodes, endNodes
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	sections := iutils.BreakByEmptyLineString1D(lines)
	gTowels = iutils.ExtractByte2DFromString1D(stringz.SplitMultiTrimSpace(sections[0][0], []string{","}), "", nil, 0)
	gDesigns = iutils.ExtractByte2DFromString1D(sections[1], "", nil, 0)
	//arrz.PPrint2D(gTowels)
	//fmt.Println(len(gTowels))
	//fmt.Println(len(gDesigns))
	//arrz.PPrint2D(gDesigns)
}
