package ds

import (
	"fmt"
	"github.com/avertocle/contests/io/outils"
	"strings"
)

/***** GNode : Graph Node *****/

const keySep = "-"

type Graph struct {
	AdList map[string]map[string]int
	AdMat  map[string]int // key = v1-v2
	VMap   map[string]int
}

func NewGraph() *Graph {
	g := &Graph{
		AdList: make(map[string]map[string]int), // list is a map to easily check duplicates while adding
		AdMat:  make(map[string]int),
		VMap:   make(map[string]int),
	}
	return g
}

func (this *Graph) AddVertex(v string, adjWeightMap map[string]int) {
	this.VMap[v] = 1
	this.addVtoAdList(v, adjWeightMap)
	this.addVtoAdMat(v, adjWeightMap)
}

func (this *Graph) addVtoAdList(v string, adjWeightMap map[string]int) {
	currAdj, ok := this.AdList[v]
	if !ok {
		currAdj = make(map[string]int)
	}
	for adj, wei := range adjWeightMap {
		currAdj[adj] = wei
	}
	this.AdList[v] = currAdj
}

func (this *Graph) addVtoAdMat(v string, adjWeightMap map[string]int) {
	key := ""
	for adj, wei := range adjWeightMap {
		key = fmt.Sprintf("%v%v%v", v, keySep, adj)
		this.AdMat[key] = wei
	}
}

func (this *Graph) PrintAdList() {
	for v, adj := range this.AdList {
		fmt.Printf("%v => %v\n", v, mapToStr(adj))
	}
}

func (this *Graph) PrintAdMat() {

	matSize := len(this.VMap) + 1
	mat := make([][]string, matSize)

	vToIdxMap := make(map[string]int)
	for i := 0; i < matSize; i++ {
		mat[i] = make([]string, matSize)
		for j := 0; j < len(mat[i]); j++ {
			mat[i][j] = "0"
		}
	}

	i := 0
	for v, _ := range this.VMap {
		vToIdxMap[v] = i
		mat[0][i+1] = v
		mat[i+1][0] = v
		i++
	}

	var tokens []string
	for key, wei := range this.AdMat {
		tokens = strings.Split(key, keySep)
		mat[vToIdxMap[tokens[0]]+1][vToIdxMap[tokens[1]]+1] = fmt.Sprintf("%v", wei)
	}
	outils.PrettyArray2DString(mat)
}

func mapToStr(m map[string]int) string {
	s := ""
	for k, v := range m {
		s += fmt.Sprintf("%v(%v) ", k, v)
	}
	return s
}
