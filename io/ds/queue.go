package ds

import "fmt"

type Queue struct {
	Start *Node
	End   *Node
	Size  int
}

func NewQueue() *Queue {
	return &Queue{
		Start: nil,
		End:   nil,
		Size:  0,
	}
}

func (this *Queue) Push(data string) {
	n := &Node{
		Data: data,
		Next: nil,
	}
	if this.Start == nil {
		this.Start = n
	}
	if this.End == nil {
		this.End = n
	} else {
		this.End.Next = n
		this.End = n
	}
	this.Size++
}

func (this *Queue) Pop() (string, error) {
	if this.Start == nil {
		return "", fmt.Errorf("stack empty")
	}
	n := this.Start
	this.Start = this.Start.Next
	return n.Data, nil
}

func (this *Queue) Truncate() {
	var t *Node
	for this.Start != nil {
		t = this.Start
		this.Start = this.Start.Next
		t.Next = nil
	}
	this.End = nil
}

func (this *Queue) isEmpty() bool {
	return this.Size == 0
}

func (this *Queue) Print() {
	fmt.Printf("start -> ")
	for t := this.Start; t != nil; t = t.Next {
		fmt.Printf("%v -> ", t.Data)
	}
	fmt.Printf("end")
	fmt.Printf("\n")
}
