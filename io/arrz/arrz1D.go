package arrz

import (
	"fmt"
	"github.com/avertocle/contests/io/clr"
	"github.com/avertocle/contests/io/cmz"
)

func Key1D[T cmz.Primitive](keys []T) string {
	return fmt.Sprintf("%v", keys)
}

func PPrint1D[T cmz.PrimitivePlus](arr []T, withIndex bool) {
	for i, cell := range arr {
		if withIndex {
			fmt.Printf("(%v:%v) | ", clr.Int(i, clr.Cyan), clr.Gen(cell, clr.Cyan))
		} else {
			fmt.Printf("(%v) | ", clr.Gen(cell, clr.Cyan))
		}
	}
	fmt.Println()
}
