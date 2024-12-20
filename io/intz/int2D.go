package intz

import (
	"fmt"
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

func Min2D(arr [][]int) (int, []int) {
	minn := arr[0][0]
	pos := []int{0, 0}
	for i, row := range arr {
		for j, cell := range row {
			if cell < minn {
				minn = cell
				pos = []int{i, j}
			}
		}
	}
	return minn, pos
}

func PPrint2D(arr [][]int) {
	for _, row := range arr {
		for _, cell := range row {
			//fmt.Printf("%v ", clr.Int(cell, clr.Cyan))
			if cell == 0 {
				fmt.Printf(".. ")
				continue
			}
			fmt.Printf("%02d ", cell)
		}
		fmt.Println()
	}
	fmt.Println()
}
