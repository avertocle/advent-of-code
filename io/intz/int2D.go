package intz

import (
	"fmt"
	"github.com/avertocle/contests/io/clr"
)

func Init2D(rows, cols, val int) [][]int {
	ans := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			ans[i][j] = val
		}
	}
	return ans
}

func PPrint2D(arr [][]int) {
	for _, row := range arr {
		for _, cell := range row {
			fmt.Printf("%v ", clr.Int(cell, clr.Cyan))
		}
		fmt.Println()
	}
	fmt.Println()
}
