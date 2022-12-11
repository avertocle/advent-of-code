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

func (this *Stack) Peek() string {
	if this.Top == nil {
		return ""
	}
	return this.Top.Data
}

func (this *Stack) Truncate() {
	var t *Node
	for this.Top != nil {
		t = this.Top
		this.Top = this.Top.Next
		t.Next = nil
	}
}

func (this *Stack) IsEmpty() bool {
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

/***** Util Functions ***/

func (this *Stack) PushAll(s1d []string) {
	for _, s := range s1d {
		this.Push(s)
	}
}

func (this *Stack) PushAllRev(s1d []string) {
	for i := len(s1d) - 1; i >= 0; i-- {
		this.Push(s1d[i])
	}
}

// pop N from stack or all is N < size
func (this *Stack) PopN(n int) []string {
	ans := make([]string, 0)
	x := ""
	var err error
	for i := 0; i < n; i++ {
		if x, err = this.Pop(); err != nil {
			fmt.Println(err)
			continue
		} else {
			ans = append(ans, x)
		}
	}
	return ans
}
