package bintree

import (
	"fmt"
	"github.com/avertocle/contests/io/outils"
	"strings"
)

var idCtr int

type TNode struct {
	Id int
	C  []*TNode // not a map because order of children matters
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
		this.AddC(c)
	}
}

func (this *TNode) AddC(c *TNode) {
	this.C = append(this.C, c)
}

func NewTNode(v int, p *TNode) *TNode {
	idCtr++
	return &TNode{
		Id: idCtr,
		V:  v,
		C:  make([]*TNode, 0),
		P:  p,
	}
}

func PrintHierarchial(tn *TNode, depth int) {
	outils.PrintWithDepth(fmt.Sprintf("(%v)", tn.V), depth)
	for _, c := range tn.C {
		PrintHierarchial(c, depth+1)
	}
}

func GetFlatStringLeafOnly(tn *TNode) string {
	if tn.IsLeaf() {
		if tn.P == nil {
			return "[]"
		} else {
			return fmt.Sprintf("%v", tn.V)
		}
	}
	s := make([]string, 0)
	for _, c := range tn.C {
		s = append(s, GetFlatStringLeafOnly(c))
	}
	return fmt.Sprintf("[%v]", strings.Join(s, ","))
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
