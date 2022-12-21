package outils

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/stringz"
	"strings"
)

func PrettyArray2DInt(arr [][]int) {
	intz.PPrint2D(arr)
}

func PrettyArray2DByte(arr [][]byte) {
	bytez.PPrint2D(arr)
}

func PrettyArray2DString(arr [][]string) {
	stringz.PPrint2D(arr)
}

func PrintWithDepth(s string, d int) {
	fmt.Printf("%v- %v\n", strings.Repeat(" ", 2*d), s)
}
