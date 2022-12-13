package main

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/outils"
)

func main() {
	demoExtract()

}

func demoPad() {
	arr := bytez.Init2D(10, 10, '.')
	for i, row := range arr {
		for j, _ := range row {
			arr[i][j] = byte((i*10+j)%10 + '0')
		}
	}

	outils.PrettyArray2DByte(arr)
	fmt.Println()

	ans := bytez.Pad2D(arr, 10, 10, 10, '.')
	outils.PrettyArray2DByte(ans)
	fmt.Println()
}

func demoExtract() {
	arr := bytez.Init2D(10, 10, '.')
	for i, row := range arr {
		for j, _ := range row {
			arr[i][j] = byte((i*10+j)%10 + '0')
		}
	}
	outils.PrettyArray2DByte(arr)
	fmt.Println()

	ans := bytez.ExtractSq2D(arr, []int{0, 0}, 3, '.')
	outils.PrettyArray2DByte(ans)
	fmt.Println()

	ans = bytez.ExtractSq2D(arr, []int{1, 1}, 3, '.')
	outils.PrettyArray2DByte(ans)
	fmt.Println()

	ans = bytez.ExtractSq2D(arr, []int{3, 3}, 3, '.')
	outils.PrettyArray2DByte(ans)
	fmt.Println()

	ans = bytez.ExtractSq2D(arr, []int{9, 9}, 3, '.')
	outils.PrettyArray2DByte(ans)
	fmt.Println()

	ans = bytez.ExtractSq2D(arr, []int{3, 3}, 6, '.')
	outils.PrettyArray2DByte(ans)
	fmt.Println()

	ans = bytez.ExtractSq2D(arr, []int{3, 3}, 2, '.')
	outils.PrettyArray2DByte(ans)
	fmt.Println()

	ans = bytez.ExtractSq2D(arr, []int{3, 3}, 1, '.')
	outils.PrettyArray2DByte(ans)
	fmt.Println()

	ans = bytez.ExtractSq2D(arr, []int{3, 3}, 0, '.')
	outils.PrettyArray2DByte(ans)
	fmt.Println()

}
