package ds

import "fmt"

type Stack struct {
	Top  *Node
	Size int
}

func NewStack() *Stack {
	return &Stack{
		Top:  nil,
		Size: 0,
	}
}

func (this *Stack) Push(data string) {
	n := &Node{
		Data: data,
		Next: this.Top,
	}
	this.Top = n
	this.Size++
}

func (this *Stack) Pop() (string, error) {
	if this.Top == nil {
		return "", fmt.Errorf("stack empty")
	}
	n := this.Top
	this.Top = this.Top.Next
	this.Size--
	return n.Data, nil
}

func (this *Stack) Truncate() {
	var t *Node
	for this.Top != nil {
		t = this.Top
		this.Top = this.Top.Next
		t.Next = nil
	}
}

func (this *Stack) isEmpty() bool {
	return this.Size == 0
}

func (this *Stack) Print() {
	fmt.Printf("top -> ")
	for t := this.Top; t != nil; t = t.Next {
		fmt.Printf("%v -> ", t.Data)
	}
	fmt.Printf("nil")
	fmt.Printf("\n")
}
