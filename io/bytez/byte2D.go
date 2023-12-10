package bytez

import (
	"fmt"
	"github.com/avertocle/contests/io/intz"
	"math"
)

func Init2D(rows, cols int, b byte) [][]byte {
	ans := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		ans[i] = make([]byte, cols)
		for j, _ := range ans[i] {
			ans[i][j] = b
		}
	}
	return ans
}

func Copy2D(source [][]byte) [][]byte {
	return Extract2D(source, []int{0, 0}, []int{len(source) - 1, len(source[0]) - 1}, 0)
}

func CountInSection2D(arr [][]byte, boundTl, boundBr []int, v byte) int {
	return CountIf2D(arr, func(b byte, i int, j int) bool {
		if i >= boundTl[0] && i <= boundBr[0] &&
			j >= boundTl[1] && j <= boundBr[1] && b == v {
			return true
		} else {
			return false
		}
	})
}

func Find2D(grid [][]byte, v byte) [][]int {
	ans := make([][]int, 0)
	for i, row := range grid {
		for j, cell := range row {
			if cell == v {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

func Count2D(grid [][]byte, v byte) int {
	return CountIf2D(grid, func(b byte, i int, j int) bool {
		return b == v
	})
}

func CountIf2D(arr [][]byte, f func(byte, int, int) bool) int {
	count := 0
	for i, row := range arr {
		for j, cell := range row {
			if f(cell, i, j) {
				count++
			}
		}
	}
	return count
}

func ExtractSq2D(source [][]byte, center []int, size int, padding byte) [][]byte {
	ans := Init2D(size, size, padding)
	si, sj := center[0]-(size/2), center[1]-(size/2)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if si+i >= 0 && sj+j >= 0 && si+i < len(source) && sj+j < len(source[0]) {
				ans[i][j] = source[si+i][sj+j]
			}
		}
	}
	return ans
}

func Pad2D(arr [][]byte, rows, cols, padSize int, padVal byte) [][]byte {
	ans := Init2D(rows+2*padSize, cols+2*padSize, padVal)
	for i, row := range arr {
		for j, _ := range row {
			ans[i+padSize][j+padSize] = arr[i][j]
		}
	}
	return ans
}

func Extract2D(arr [][]byte, boundTL, boundBR []int, padVal byte) [][]byte {
	xlen, ylen := boundBR[0]-boundTL[0]+1, boundBR[1]-boundTL[1]+1
	ans := Init2D(xlen, ylen, padVal)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			ans[x][y] = arr[x+boundTL[0]][y+boundTL[1]]
		}
	}
	return ans
}

func Transpose2D(arr [][]byte) [][]byte {
	xlen, ylen := len(arr[0]), len(arr)
	ans := Init2D(xlen, ylen, 0)
	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			ans[x][y] = arr[y][x]
		}
	}
	return ans
}

/*
FindBounds2D
returns top-left and bottom-right bounds of the array
*/
func FindBounds2D(arr [][]byte, empty byte) ([]int, []int) {
	if len(arr) == 0 {
		return []int{}, []int{}
	}
	tli, tlj, bri, brj := math.MaxInt, math.MaxInt, -1, -1
	rlen, clen := len(arr), len(arr[0])
	for i := 0; i < rlen; i++ {
		for j := 0; j < clen; j++ {
			if arr[i][j] != empty {
				tli = intz.Min(tli, i)
				tlj = intz.Min(tlj, j)
				bri = intz.Max(bri, i)
				brj = intz.Max(brj, j)
			}
		}
	}
	return []int{tli, tlj}, []int{bri, brj}
}

func PPrint2D(arr [][]byte) {
	fmt.Println()
	for _, row := range arr {
		for _, c := range row {
			fmt.Printf("%v", string(c))
		}
		fmt.Println()
	}
}
