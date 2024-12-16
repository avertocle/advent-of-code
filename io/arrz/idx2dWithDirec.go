package arrz

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
)

const (
	Up = iota
	Right
	Down
	Left
)

type Idx2DD[T int | int64] struct {
	I, J T
	D    int
}

func NewIdx2DD[T int | int64](i, j T, d int) *Idx2DD[T] {
	return &Idx2DD[T]{I: i, J: j, D: d}
}

func (o *Idx2DD[T]) Str() string {
	return fmt.Sprintf("[%v, %v, %v]", o.I, o.J, o.D)
}

// can rotate in both directions and rotation is a step
func (o *Idx2DD[T]) NextStates() []*Idx2DD[T] {
	if o.D == Up {
		return []*Idx2DD[T]{NewIdx2DD(o.I-1, o.J, Up), NewIdx2DD(o.I, o.J, Right), NewIdx2DD(o.I, o.J, Left)}
	} else if o.D == Right {
		return []*Idx2DD[T]{NewIdx2DD(o.I, o.J+1, Right), NewIdx2DD(o.I, o.J, Down), NewIdx2DD(o.I, o.J, Up)}
	} else if o.D == Down {
		return []*Idx2DD[T]{NewIdx2DD(o.I+1, o.J, Down), NewIdx2DD(o.I, o.J, Left), NewIdx2DD(o.I, o.J, Right)}
	} else if o.D == Left {
		return []*Idx2DD[T]{NewIdx2DD(o.I, o.J-1, Left), NewIdx2DD(o.I, o.J, Up), NewIdx2DD(o.I, o.J, Down)}
	}
	return make([]*Idx2DD[T], 0)
}

func (o *Idx2DD[T]) IsInBounds(rows, cols T) bool {
	return o.I >= 0 && o.I < rows && o.J >= 0 && o.J < cols
}

func (o *Idx2DD[T]) ToKey() string {
	return fmt.Sprintf("%v-%v-%v", o.I, o.J, o.D)
}

func (o *Idx2DD[T]) MoveBy(i, j T) {
	o.I += i
	o.J += j
}

func NewIdx2DDFromKey[T int | int64](key string) *Idx2DD[T] {
	x := iutils.ExtractInt641DFromString0D(key, "-", -1)
	return NewIdx2DD[T](T(x[0]), T(x[1]), int(T(x[2])))
}

func (o *Idx2DD[T]) IsEqual(o1 *Idx2DD[T], noDirec bool) bool {
	if noDirec {
		return o1 != nil && (o.I == o1.I && o.J == o1.J)
	}
	return o1 != nil && (o.I == o1.I && o.J == o1.J && o.D == o1.D)
}

func (o *Idx2DD[T]) Clone() *Idx2DD[T] {
	return NewIdx2DD(o.I, o.J, o.D)
}

func Idx2DDListToStr[T int | int64](idxs []*Idx2DD[T]) string {
	str := ""
	for _, idx := range idxs {
		str += idx.Str() + ","
	}
	return fmt.Sprintf("[%v]", str)
}
