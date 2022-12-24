package dll

import "fmt"

type DLLNode struct {
	V int
	N *DLLNode
	P *DLLNode
}

func (n *DLLNode) Str() string {
	sp, sn := "-", "-"
	if n.P != nil {
		sp = fmt.Sprintf("%v", n.P.V)
	}
	if n.N != nil {
		sn = fmt.Sprintf("%v", n.N.V)
	}
	return fmt.Sprintf("%v(%v,%v)", n.V, sp, sn)
}

func NewDLLNode(v int, p, n *DLLNode) *DLLNode {
	dlln := &DLLNode{
		V: v,
		N: n,
		P: p,
	}
	return dlln
}

/*
AddAfter
use idx = -1 for adding to start
returns start node
intentionally crashes if idx >= length of list
*/
func AddAfter(snode, node *DLLNode, idx int) *DLLNode {
	if idx == -1 { // add at start
		node.P = nil
		node.N = snode
		snode.P = node
		return node
	}
	temp := NavTo(snode, idx)
	node.P = temp
	node.N = temp.N
	if temp.N != nil { // add at middle
		temp.N.P = node
	}
	temp.N = node
	return snode
}

func NavTo(snode *DLLNode, idx int) *DLLNode {
	temp := snode
	for i := 0; i < idx; i++ {
		temp = temp.N
	}
	return temp
}

func NavToEnd(node *DLLNode) *DLLNode {
	for node.N != nil {
		node = node.N
	}
	return node
}

func Size(sn *DLLNode) int {
	i := 0
	for ; sn != nil; sn = sn.N {
		i++
	}
	return i
}

/*
DelAt
returns start node or nil if list is empty
*/
func DelAt(snode *DLLNode, idx int) *DLLNode {
	var temp *DLLNode
	if idx == 0 { //deleting at start
		if snode.N == nil { // deleting start in an empty list
			return nil
		}
		temp = snode
		snode = snode.N
		snode.P = nil
		temp.N = nil
		return snode
	}
	temp = NavTo(snode, idx)
	temp.P.N = temp.N
	if temp.N != nil {
		temp.N.P = temp.P
	}
	temp.N, temp.P = nil, nil
	return snode
}

func AddAfterMe(snode, target, node *DLLNode) *DLLNode {
	if target == nil { // add at start
		node.P = nil
		node.N = snode
		snode.P = node
		return node
	}
	node.P = target
	node.N = target.N
	if target.N != nil { // add at middle
		target.N.P = node
	}
	target.N = node
	return snode
}

func DelMe(snode, node *DLLNode) *DLLNode {
	if node.P == nil { //deleting at start
		if node.N == nil { // deleting start in an empty list
			return nil
		}
		snode = node.N
		snode.P = nil
		node.N, node.P = nil, nil
		return snode
	}
	node.P.N = node.N
	if node.N != nil {
		node.N.P = node.P
	}
	node.N, node.P = nil, nil
	return snode
}

func FromArray(arr []int) *DLLNode {
	if len(arr) == 0 {
		return nil
	}
	sn := NewDLLNode(arr[0], nil, nil)
	tn, ln := sn, sn
	for i := 1; i < len(arr); i++ {
		tn = NewDLLNode(arr[i], ln, nil)
		ln.N = tn
		ln = ln.N
	}
	return sn
}

func PP(snode *DLLNode) {
	t := snode
	fmt.Printf("S -> ")
	for t != nil {
		fmt.Printf("[%v] -> ", t.V)
		t = t.N
	}
	fmt.Printf("nil")
	fmt.Println()
}

func PPDetailed(snode *DLLNode) {
	t := snode
	fmt.Printf("S -> ")
	for t != nil {
		fmt.Printf("[%v] -> ", t.Str())
		t = t.N
	}
	fmt.Printf("nil")
	fmt.Println()
}
