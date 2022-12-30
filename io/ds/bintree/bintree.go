package bintree

import (
	"fmt"
	"strings"
)

var idCtr int

type TNode struct {
	Id int
	L  *TNode
	R  *TNode
	P  *TNode
	V  int
}

func (this *TNode) IsLeaf() bool {
	return this.L == nil && this.R == nil
}

func (this *TNode) IsFull() bool {
	return this.L != nil && this.R != nil
}

func (this *TNode) CompareTo(tn *TNode) bool {
	return this.Id == tn.Id
}

func NewTNode(v int, p *TNode) *TNode {
	idCtr++
	return &TNode{
		Id: idCtr,
		V:  v,
		L:  nil,
		R:  nil,
		P:  p,
	}
}

func PrintHierarchial(root *TNode, depth int) {
	if root == nil {
		return
	}
	fmt.Printf("%v- %v\n", strings.Repeat(" ", 2*depth), fmt.Sprintf("(%v)", root.V))
	PrintHierarchial(root.L, depth+1)
	PrintHierarchial(root.R, depth+1)
}

func PrintInorder(root *TNode) {
	if root == nil {
		return
	}
	PrintInorder(root.L)
	fmt.Printf("%v", root.V)
	fmt.Printf(",")
	PrintInorder(root.R)
}

func PrintInorderLeafOnly(root *TNode) {
	if root == nil {
		return
	}
	if !root.IsLeaf() {
		fmt.Printf("[")
	}
	PrintInorderLeafOnly(root.L)
	if root.IsLeaf() {
		fmt.Printf("%v", root.V)
	} else {
		fmt.Printf(",")
	}
	PrintInorderLeafOnly(root.R)
	if !root.IsLeaf() {
		fmt.Printf("]")
	}
}

func FlattenLeafOnly(root *TNode) []*TNode {
	flatTree := make([]*TNode, 0)
	if root == nil {
		return flatTree
	}
	//if !root.IsLeaf() {
	//	fmt.Printf("[")
	//}
	flatTree = append(FlattenLeafOnly(root.L), flatTree...)
	if root.IsLeaf() {
		flatTree = append(flatTree, root)
		//fmt.Printf("%v", root.V)
	}
	flatTree = append(flatTree, FlattenLeafOnly(root.R)...)
	//if !root.IsLeaf() {
	//	fmt.Printf("]")
	//}
	return flatTree
}
