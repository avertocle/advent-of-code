package day08

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"strings"
)

var gInputPath []byte
var gInputNodes [][]string

const DirPath = "../2023/day08"

func SolveP1() string {
	ans := 0
	//nodeMap := makeMapP1()
	//snodeVal := "AAA"
	//enodeVal := "ZZZ"
	//for pathCtr := 0; strings.Compare(snodeVal, enodeVal) != 0; pathCtr++ {
	//	fmt.Println(snodeVal)
	//	if pathCtr == len(gInputPath) {
	//		pathCtr = 0
	//	}
	//	if gInputPath[pathCtr] == 'L' {
	//		snodeVal = nodeMap[snodeVal][0]
	//	} else {
	//		snodeVal = nodeMap[snodeVal][1]
	//	}
	//	ans++
	//}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	nodeMap := makeMapP1()
	snodeVals := findStartNodes()
	for pathCtr := 0; ; pathCtr++ {
		if pathCtr == len(gInputPath) {
			pathCtr = 0
		}
		enodeCtr := 0
		for i := 0; i < len(snodeVals); i++ {
			if gInputPath[pathCtr] == 'L' {
				snodeVals[i] = nodeMap[snodeVals[i]][0]
			} else {
				snodeVals[i] = nodeMap[snodeVals[i]][1]
			}
			if isEndNode(snodeVals[i]) {
				enodeCtr++
			}
		}
		ans++
		if ans%10000000 == 0 {
			fmt.Printf("%v->", ans)
		}
		if enodeCtr == len(snodeVals) {
			break
		}
	}
	return fmt.Sprintf("%v", ans)
}

func findStartNodes() []string {
	startNodes := make([]string, 0)
	for _, node := range gInputNodes {
		if strings.HasSuffix(node[0], "A") {
			startNodes = append(startNodes, node[0])
		}
	}
	return startNodes
}

func isEndNode(node string) bool {
	return strings.HasSuffix(node, "Z")
}

/***** Common Functions *****/

/***** P1 Functions *****/

func makeMapP1() map[string][]string {
	m := make(map[string][]string)
	for _, node := range gInputNodes {
		m[node[0]] = node[1:]
	}
	return m
}

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInputPath = []byte(lines[0])
	for _, line := range lines[2:] {
		gInputNodes = append(gInputNodes, stringz.SplitMultiTrimSpace(line, []string{" ", "=", ",", "(", ")"}))
	}

	//fmt.Printf("gInputPath = %v\n", gInputPath)
	//fmt.Printf("gInputNodes = %v\n", gInputNodes)
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
