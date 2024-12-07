package arrz

import "fmt"

type Idx2D struct {
	I, J int
}

func NewIdxP2D(i, j int) *Idx2D {
	return &Idx2D{I: i, J: j}
}

func (o *Idx2D) Str() string {
	return fmt.Sprintf("[%v, %v]", o.I, o.J)
}

func (o *Idx2D) Neighbours(includeDiag bool) []*Idx2D {
	n := []*Idx2D{NewIdxP2D(o.I-1, o.J), NewIdxP2D(o.I+1, o.J), NewIdxP2D(o.I, o.J-1), NewIdxP2D(o.I, o.J+1)}
	if includeDiag {
		n = append(n, NewIdxP2D(o.I-1, o.J-1), NewIdxP2D(o.I-1, o.J+1), NewIdxP2D(o.I+1, o.J-1), NewIdxP2D(o.I+1, o.J+1))
	}
	return n
}

func (o *Idx2D) ToKey() string {
	return fmt.Sprintf("%v-%v", o.I, o.J)
}

func (o *Idx2D) IsEqual(o1 *Idx2D) bool {
	return o1 != nil && (o.I == o1.I && o.J == o1.J)
}

func (o *Idx2D) Clone() *Idx2D {
	return NewIdxP2D(o.I, o.J)
}

func (o *Idx2D) North() *Idx2D {
	return NewIdxP2D(o.I-1, o.J)
}

func (o *Idx2D) South() *Idx2D {
	return NewIdxP2D(o.I+1, o.J)
}

func (o *Idx2D) East() *Idx2D {
	return NewIdxP2D(o.I, o.J+1)
}

func (o *Idx2D) West() *Idx2D {
	return NewIdxP2D(o.I, o.J-1)
}

func Idx2DListToStr(idxs []*Idx2D) string {
	str := ""
	for _, idx := range idxs {
		str += idx.Str() + ","
	}
	return fmt.Sprintf("[%v]", str)
}
