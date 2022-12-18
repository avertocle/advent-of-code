package bytez

import "fmt"

func PPrint3D(arr [][][]byte) {
	for i, a2d := range arr {
		fmt.Printf("%02v => ", i)
		for _, row := range a2d {
			fmt.Printf("[%v] ", string(row))
		}
		fmt.Println()
	}
	fmt.Println()
}
