package io

import "fmt"

func PrettyArray2D(arr [][]int) {
	for _, row := range arr {
		for _, cell := range row {
			fmt.Printf("%v ", cell)
		}
		fmt.Println()
	}
}

func PrettyArray3D(arr [][][]int) {
	for _, d1 := range arr {
		PrettyArray2D(d1)
		fmt.Println()
	}
}
