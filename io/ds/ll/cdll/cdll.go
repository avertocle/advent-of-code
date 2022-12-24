package cdll

import "fmt"

var idCtr int

type Node struct {
	Id int
	V  int
	N  *Node
	P  *Node
}

func (n *Node) Str() string {
	sp, sn := "-", "-"
	if n.P != nil {
		sp = fmt.Sprintf("%v", n.P.V)
	}
	if n.N != nil {
		sn = fmt.Sprintf("%v", n.N.V)
	}
	return fmt.Sprintf("%v(%v,%v)", n.V, sp, sn)
}

func NewCDLLNode(v int, p, n *Node) *Node {
	idCtr++
	dlln := &Node{
		Id: idCtr,
		V:  v,
		N:  n,
		P:  p,
	}
	return dlln
}

func NavFwd(xn *Node, dis int) *Node {
	for i := 0; i < dis; i++ {
		xn = xn.N
	}
	return xn
}

func NavRev(xn *Node, dis int) *Node {
	for i := 0; i < dis; i++ {
		xn = xn.P
	}
	return xn
}

func Size(xn *Node) int {
	if xn == nil {
		return 0
	} else if xn.N == nil {
		return 1
	}
	i := 1
	for tn := xn.N; tn.Id != xn.Id; tn = tn.N {
		i++
	}
	return i
}

func AddAfterMe(me, node *Node) {
	node.P = me
	node.N = me.N
	me.N.P = node
	me.N = node
}

func DelMe(me *Node) {
	me.P.N = me.N
	me.N.P = me.P
	me.N, me.P = nil, nil
}

func FromArray(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	sn := NewCDLLNode(arr[0], nil, nil)
	tn, ln := sn, sn
	for i := 1; i < len(arr); i++ {
		tn = NewCDLLNode(arr[i], ln, nil)
		ln.N = tn
		ln = tn
	}
	sn.P = ln
	ln.N = sn
	return sn
}

func FindNodesByVal(sn *Node, val int) []*Node {
	ans := make([]*Node, 0)
	tn := sn
	for {
		if tn.V == val {
			ans = append(ans, tn)
		}
		tn = tn.N
		if tn.Id == sn.Id {
			break
		}
	}
	return ans
}

func PP(sn *Node) {
	if sn == nil {
		return
	}
	fmt.Printf("S -> ")
	tn := sn
	for {
		fmt.Printf("%v, ", tn.V)
		tn = tn.N
		if tn.Id == sn.Id {
			break
		}
	}
	fmt.Printf("S")
	fmt.Println()
}

func PPDetailed(sn *Node) {
	if sn == nil {
		return
	}
	fmt.Printf("S -> ")
	tn := sn
	for {
		fmt.Printf("%v -> ", tn.Str())
		tn = tn.N
		if tn.Id == sn.Id {
			break
		}
	}
	fmt.Printf("S")
	fmt.Println()
}
