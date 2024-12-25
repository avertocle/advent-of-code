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
var gDesigns []string
var gGraph *ds.Graph

const (
	EndNodeId    = "end"
	StartNodeId  = "start"
	StartNodeVal = '*'
	EndNodeVal   = '#'
)

/*
Adding a '*' to the beginning of the design so that graph have a single start node '*'.
We can hence avoid maintaining a list of start nodes.
*/

func SolveP1() string {
	ans := int64(0)
	memory := make(map[string]int64)
	for _, design := range gDesigns {
		newDesign := fmt.Sprintf("%v%v%v", string(StartNodeVal), design, string(EndNodeVal))
		found := countTowelArrangements(memory, newDesign, StartNodeId, false)
		ans += found
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := int64(0)
	memory := make(map[string]int64)
	for _, design := range gDesigns {
		newDesign := fmt.Sprintf("%v%v%v", string(StartNodeVal), design, string(EndNodeVal))
		found := countTowelArrangements(memory, newDesign, StartNodeId, true)
		ans += found
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

func countTowelArrangements(memory map[string]int64, design string, currId string, findAll bool) int64 {
	errz.HardAssert(len(design) > 0, "design cannot be empty")
	if gGraph.VMap[currId] != int(design[0]) {
		return 0
	}
	if gGraph.VMap[currId] == gGraph.VMap[EndNodeId] {
		errz.HardAssert(len(design) == 1, "design will only have one character when currId is end node")
		memory[memKey(design, currId)] += 1
		return 1
	}
	if v, ok := memory[memKey(design, currId)]; ok {
		return v
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
		memory[memKey(design, currId)] += count
	}
	return count
}

func memKey(design, currId string) string {
	return currId + "-" + design
}

func makeGraph() *ds.Graph {
	g := ds.NewGraph()
	endNodes := make(map[string]bool)
	for i, row := range gTowels {
		for j, cell := range row {
			vid := fmt.Sprintf("%v-%v", i, j)
			awm := make(map[string]int)
			if j < len(row)-1 {
				awm[fmt.Sprintf("%v-%v", i, j+1)] = 1
			} else {
				for ii, _ := range gTowels {
					awm[fmt.Sprintf("%v-%v", ii, 0)] = 1
				}
			}
			if j == len(row)-1 {
				endNodes[vid] = true
				awm[EndNodeId] = 1
			}
			g.AddVertex(vid, int(cell), awm)
		}
	}

	awm := make(map[string]int)
	for i, _ := range gTowels {
		awm[fmt.Sprintf("%v-%v", i, 0)] = 1
	}
	g.AddVertex(StartNodeId, StartNodeVal, awm)
	g.AddVertex(EndNodeId, EndNodeVal, make(map[string]int))

	return g
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	sections := iutils.BreakByEmptyLineString1D(lines)
	gTowels = iutils.ExtractByte2DFromString1D(stringz.SplitMultiTrimSpace(sections[0][0], []string{","}), "", nil, 0)
	gDesigns = sections[1]
	gGraph = makeGraph()
}
