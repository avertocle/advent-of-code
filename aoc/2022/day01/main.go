package main

import (
	"fmt"
	"github.com/avertocle/contests/io"
	"github.com/avertocle/contests/metrics"
	"log"
	"sort"
)

const inputFilePath = "input.txt"

func main() {
	metrics.ProgStart()
	input := getInputOrDie()

	io.PrettyArray2DInt(input)

	ans01 := solvePartOne(input)
	fmt.Printf("answer part-01 = %v\n", ans01)

	ans02 := solvePartTwo(input)
	fmt.Printf("answer part-02 = %v\n", ans02)

}

func solvePartOne(input [][]int) int {
	maxCalSum := 0
	calSum := 0
	for _, row := range input {
		calSum = 0
		for _, cell := range row {
			calSum += cell
		}
		if calSum > maxCalSum {
			maxCalSum = calSum
		}
	}
	return maxCalSum
}

func solvePartTwo(input [][]int) int {
	calSums := make([]int, len(input))
	for i, row := range input {
		calSums[i] = 0
		for _, cell := range row {
			calSums[i] += cell
		}
	}
	sort.Ints(calSums)
	l := len(calSums)
	return calSums[l-1] + calSums[l-2] + calSums[l-3]
}

// input : [][]int : each row contains elf calories
func getInputOrDie() [][]int {
	lines, err := io.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}
	linesInt := io.ExtractInt1DFromString1D(lines, " ", 0, -1)

	//	fmt.Printf("%+v\n", linesInt)

	input := make([][]int, 0)
	var temp []int
	for i := 0; i < len(linesInt); i++ {
		temp = make([]int, 0)
		for ; i < len(linesInt) && linesInt[i] > -1; i++ {
			temp = append(temp, linesInt[i])
		}
		//		fmt.Printf("%+v\n", input)
		input = append(input, temp)
	}
	return input
}
