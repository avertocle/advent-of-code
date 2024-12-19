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
var gGraph *ds.Graph
var gEndNodes map[string]bool

const StartNodeId = "start"

// 400, 388 too high
// 1019 too low

/*
Adding a '*' to the beginning of the design so that graph have a single start node '*'.
We can hence avoid maintaining a list of start nodes.
*/

func SolveP1() string {
	ans := int64(0)
	memory := make(map[string]int64)
	for _, design := range gDesigns {
		newDesign := append([]byte{'*'}, design...)
		found := countTowelArrangements(memory, newDesign, StartNodeId, false)
		ans += found
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := int64(0)
	memory := make(map[string]int64)
	for _, design := range gDesigns {
		newDesign := append([]byte{'*'}, design...)
		found := countTowelArrangements(memory, newDesign, StartNodeId, true)
		ans += found
	}
	return fmt.Sprintf("%v", ans)
}

func countTowelArrangements(memory map[string]int64, design []byte, currId string, findAll bool) int64 {
	if v, ok := memory[fmt.Sprintf("%v-%v", design, currId)]; ok {
		return v
	}
	errz.HardAssert(len(design) > 0, "design cannot be is empty")
	if gGraph.VMap[currId] != int(design[0]) {
		return 0
	}
	if len(design) == 1 {
		if gEndNodes[currId] {
			memory[fmt.Sprintf("%v-%v", design, currId)] += 1
			return 1
		} else {
			return 0
		}
	}
	nextNodes, _ := gGraph.AdList[currId]
	count := int64(0)
	for nextNodeId, _ := range nextNodes {
		count += countTowelArrangements(memory, design[1:], nextNodeId, findAll)
		if !findAll && count > 0 {
			break
		}
	}
	if count > 0 {
		memory[fmt.Sprintf("%v-%v", design, currId)] += count
	}
	return count
}

func makeGraph() (*ds.Graph, map[string]bool) {
	g := ds.NewGraph()
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
			if j == len(row)-1 {
				endNodes[vkey] = true
			}
		}
	}

	awm := make(map[string]int)
	for i, _ := range gTowels {
		awm[fmt.Sprintf("%v-%v", i, 0)] = 1
	}
	g.AddVertex(StartNodeId, int('*'), awm)

	//fmt.Println("startNodes = ", startNodes)
	//fmt.Println("endNodes = ", endNodes)
	//g.PrintAdList()
	return g, endNodes
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
	gGraph, gEndNodes = makeGraph()
}
