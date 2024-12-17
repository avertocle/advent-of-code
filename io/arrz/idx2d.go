package arrz

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/numz"
)

type Idx2D[T int | int64] struct {
	I, J T
}

func NewIdx2D[T int | int64](arr ...T) *Idx2D[T] {
	return &Idx2D[T]{I: arr[0], J: arr[1]}
}

func (o *Idx2D[T]) Str() string {
	return fmt.Sprintf("[%v, %v]", o.I, o.J)
}

func (o *Idx2D[T]) Neighbours(includeDiag bool) []*Idx2D[T] {
	n := []*Idx2D[T]{NewIdx2D(o.I-1, o.J), NewIdx2D(o.I+1, o.J), NewIdx2D(o.I, o.J-1), NewIdx2D(o.I, o.J+1)}
	if includeDiag {
		n = append(n, NewIdx2D(o.I-1, o.J-1), NewIdx2D(o.I-1, o.J+1), NewIdx2D(o.I+1, o.J-1), NewIdx2D(o.I+1, o.J+1))
	}
	return n
}

func (o *Idx2D[T]) IsInBounds(rows, cols T) bool {
	return o.I >= 0 && o.I < rows && o.J >= 0 && o.J < cols
}

func (o *Idx2D[T]) ToKey() string {
	return fmt.Sprintf("%v-%v", o.I, o.J)
}

func (o *Idx2D[T]) MoveBy(i, j T) {
	o.I += i
	o.J += j
}

func (o *Idx2D[T]) MoveBounded(i, j, minI, minJ, maxI, maxJ T) {
	o.I = numz.IncBoundedV2(o.I, i, minI, maxI)
	o.J = numz.IncBoundedV2(o.J, j, minJ, maxJ)
}

func NewIdx2DFromKey[T int | int64](key string) *Idx2D[T] {
	x := iutils.ExtractInt641DFromString0D(key, "-", -1)
	return NewIdx2D[T](T(x[0]), T(x[1]))
}

func (o *Idx2D[T]) IsEqual(o1 *Idx2D[T]) bool {
	return o1 != nil && (o.I == o1.I && o.J == o1.J)
}

func (o *Idx2D[T]) Clone() *Idx2D[T] {
	return NewIdx2D(o.I, o.J)
}

func Idx2DListToStr[T int | int64](idxs []*Idx2D[T]) string {
	str := ""
	for _, idx := range idxs {
		str += idx.Str() + ","
	}
	return fmt.Sprintf("[%v]", str)
}
