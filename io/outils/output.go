package outils

import (
	"fmt"
	"strings"
)

func PrettyArray2DInt(arr [][]int) {
	for _, row := range arr {
		for _, cell := range row {
			fmt.Printf("%v ", ClrI(cell, Cyan))
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrettyArray3DByte(arr [][][]byte) {
	for i, a2d := range arr {
		fmt.Printf("%02v => ", i)
		for _, row := range a2d {
			fmt.Printf("[%v] ", string(row))
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrettyArray2DByte(arr [][]byte) {
	for _, row := range arr {
		for _, c := range row {
			fmt.Printf("%v ", string(c))
		}
		fmt.Println()
	}
}

func PrettyArray2DString(arr [][]string) {
	for _, row := range arr {
		for _, ele := range row {
			fmt.Printf("| %v ", ele)
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrettyArray3DInt(arr [][][]int) {
	for _, d1 := range arr {
		PrettyArray2DInt(d1)
		fmt.Println()
	}
	fmt.Println()
}

func PrintWithDepth(s string, d int) {
	fmt.Printf("%v- %v\n", strings.Repeat(" ", 2*d), s)
}
