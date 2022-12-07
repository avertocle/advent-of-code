package day01

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
	"log"
	"sort"
)

// gInput : [][]int : each row contains elf calories
var gInput [][]int

func SolveP1() string {
	maxCalSum := 0
	calSum := 0
	for _, row := range gInput {
		calSum = 0
		for _, cell := range row {
			calSum += cell
		}
		if calSum > maxCalSum {
			maxCalSum = calSum
		}
	}
	return fmt.Sprintf("%v", maxCalSum)
}

func SolveP2() string {
	calSums := make([]int, len(gInput))
	for i, row := range gInput {
		calSums[i] = 0
		for _, cell := range row {
			calSums[i] += cell
		}
	}
	sort.Ints(calSums)
	l := len(calSums)
	ans := calSums[l-1] + calSums[l-2] + calSums[l-3]
	return fmt.Sprintf("%v", ans)
}

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	linesInt := iutils.ExtractInt1DFromString1D(lines, " ", 0, -1)

	//	fmt.Printf("%+v\n", linesInt)

	gInput = make([][]int, 0)
	var temp []int
	for i := 0; i < len(linesInt); i++ {
		temp = make([]int, 0)
		for ; i < len(linesInt) && linesInt[i] > -1; i++ {
			temp = append(temp, linesInt[i])
		}
		//		fmt.Printf("%+v\n", iutils)
		gInput = append(gInput, temp)
	}
}
