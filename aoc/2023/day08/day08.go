package day08

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/mathz"
	"github.com/avertocle/contests/io/stringz"
	"strings"
)

var gInputPath []byte
var gInputNodes [][]string
var gInputNodeMap map[string][]string

const DirPath = "../2023/day08"

func SolveP1() string {
	ans := 0
	isEndNode := func(node string) bool {
		return strings.Compare(node, "ZZZ") == 0
	}
	ans = findPathLen("AAA", isEndNode)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	snodeVals := findStartNodes()
	enodeCtrs := make([]int, len(snodeVals))
	isEndNode := func(node string) bool {
		return strings.HasSuffix(node, "Z")
	}
	for i := 0; i < len(snodeVals); i++ {
		enodeCtrs[i] = findPathLen(snodeVals[i], isEndNode)
	}
	fmt.Println(snodeVals)
	fmt.Println(enodeCtrs)
	ans := mathz.LCMArr(enodeCtrs)
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func findPathLen(snodeVal string, isEndNode func(string) bool) int {
	pathCtr := 0
	stepCtr := 0
	for pathCtr = 0; !isEndNode(snodeVal); pathCtr++ {
		//fmt.Printf("%v->", snodeVal)
		if pathCtr == len(gInputPath) {
			pathCtr = 0
		}
		if gInputPath[pathCtr] == 'L' {
			snodeVal = gInputNodeMap[snodeVal][0]
		} else {
			snodeVal = gInputNodeMap[snodeVal][1]
		}
		stepCtr++
	}
	return stepCtr
}

func getEndNodeMap() map[string]bool {
	endNodeMap := make(map[string]bool)
	for _, node := range gInputNodes {
		if strings.HasSuffix(node[0], "Z") {
			endNodeMap[node[0]] = true
		}
	}
	return endNodeMap
}

/***** P1 Functions *****/

/***** P2 Functions *****/

func findStartNodes() []string {
	startNodes := make([]string, 0)
	for _, node := range gInputNodes {
		if strings.HasSuffix(node[0], "A") {
			startNodes = append(startNodes, node[0])
		}
	}
	return startNodes
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInputPath = []byte(lines[0])
	nodes := make([][]string, 0)
	for _, line := range lines[2:] {
		nodes = append(nodes, stringz.SplitMultiTrimSpace(line, []string{" ", "=", ",", "(", ")"}))
	}
	nodeMap := make(map[string][]string)
	for _, node := range nodes {
		nodeMap[node[0]] = node[1:]
	}
	gInputNodeMap = nodeMap
	gInputNodes = nodes
}

/***** Helpers *****/

type gnode struct {
	L *gnode
	R *gnode
	V string
}

func newGNode(v string) *gnode {
	return &gnode{V: v}
}
