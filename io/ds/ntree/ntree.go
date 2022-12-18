package bintree

import (
	"fmt"
	"github.com/avertocle/contests/io/outils"
)

var idCtr int

type TNode struct {
	Id int
	C  map[int]*TNode
	P  *TNode
	V  int
}

func (this *TNode) IsLeaf() bool {
	return len(this.C) == 0
}

func (this *TNode) CompareTo(tn *TNode) bool {
	return this.Id == tn.Id
}

func (this *TNode) AddChildren(children []*TNode) {
	for _, c := range children {
		this.C[c.Id] = c
	}
}

func (this *TNode) AddC(c *TNode) {
	this.C[c.Id] = c
}

func NewTNode(v int, p *TNode) *TNode {
	idCtr++
	return &TNode{
		Id: idCtr,
		V:  v,
		C:  make(map[int]*TNode),
		P:  p,
	}
}

func PrintHierarchial(tn *TNode, depth int) {
	outils.PrintWithDepth(fmt.Sprintf("(%v)", tn.V), depth)
	for _, c := range tn.C {
		PrintHierarchial(c, depth+1)
	}
}

func PrintFlattenedLeafOnly(tn *TNode) {
	if tn.IsLeaf() {
		fmt.Printf("%v", tn.V)
	}
	fmt.Printf("[")
	for _, c := range tn.C {
		PrintFlattenedLeafOnly(c)
	}
	fmt.Printf("]")
}

func FlattenLeafOnly(root *TNode) []*TNode {
	flatTree := make([]*TNode, 0)
	if root.IsLeaf() {
		flatTree = append(flatTree, root)
		return flatTree
	}
	for _, c := range root.C {
		flatTree = append(flatTree, FlattenLeafOnly(c)...)
	}
	return flatTree
}
