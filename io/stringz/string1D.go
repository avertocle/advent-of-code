package stringz

import "fmt"

func PrintWithIndex1D(arr []string) {
	for i, s := range arr {
		fmt.Printf("%2d => %v\n", i, s)
	}
}
