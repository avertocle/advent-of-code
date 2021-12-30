package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/avertocle/adventofcode/io"
	"github.com/avertocle/adventofcode/metrics"
)

const inputFilePath = "input.txt"

func main() {
	metrics.ProgStart()
	input := getInputOrDie()
	metrics.InputLen(len(input))

	ans := problem2(input)
	fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func problem1(input [][]byte) int {
	rows := len(input)
	cols := len(input[0])
	riskLevel := 0
	for i, row := range input {
		for j, ele := range row {
			if isValley(input, rows, cols, i, j) {
				//fmt.Printf("%v,%v = %v\n", i, j, ele-'0')
				riskLevel += (int(ele-'0') + 1)
			}
		}
	}
	return riskLevel
}

func problem2(input [][]byte) int64 {
	rows := len(input)
	cols := len(input[0])
	topBasins := make([]int, 3)
	basinSize := 0
	for i, row := range input {
		for j, _ := range row {
			visited := io.Init2DByte(rows, cols, 0)
			basinSize = getBasinSize(input, 0, i, j, rows, cols, visited)
			fmt.Printf("%v, %v, %v, %v\n", i, j, input[i][j]-'0', basinSize)
			topBasins = processForTopSlots(topBasins, basinSize)
		}
	}
	fmt.Printf("topBasins : %v\n", topBasins)
	prod := int64(1)
	for _, b := range topBasins {
		prod *= int64(b)
	}
	return prod
}

func getBasinSize(input [][]byte, base byte, i, j, rows, cols int, visited [][]byte) int {
	//fmt.Printf("%v, %v, %v\n", base-'0', i, j)
	if !io.IsValidCoord2D(i, j, rows, cols) {
		return 0
	} else if input[i][j] < base || input[i][j] == '9' {
		return 0
	} else if visited[i][j] == 1 {
		return 0
	}
	visited[i][j] = 1
	basinSize := 1 +
		getBasinSize(input, input[i][j], i+1, j, rows, cols, visited) +
		getBasinSize(input, input[i][j], i, j+1, rows, cols, visited) +
		getBasinSize(input, input[i][j], i-1, j, rows, cols, visited) +
		getBasinSize(input, input[i][j], i, j-1, rows, cols, visited)

	//fmt.Printf("%v, %v, %v\n", i, j, basinSize)
	return basinSize
}

func processForTopSlots(topBasins []int, b int) []int {
	topBasins = append(topBasins, b)
	sort.Ints(topBasins)
	return topBasins[1:]
}

func isValley(a [][]byte, rows, cols, x, y int) bool {
	v := a[x][y]
	if x == 0 && y == 0 {
		return v < a[x+1][y] && v < a[x][y+1]
	} else if x == 0 && y == cols-1 {
		return v < a[x+1][y] && v < a[x][y-1]
	} else if x == rows-1 && y == 0 {
		return v < a[x][y+1] && v < a[x-1][y]
	} else if x == rows-1 && y == cols-1 {
		return v < a[x-1][y] && v < a[x][y-1]
	} else if x == 0 {
		return v < a[x+1][y] && v < a[x][y-1] && v < a[x][y+1]
	} else if x == rows-1 {
		return v < a[x-1][y] && v < a[x][y-1] && v < a[x][y+1]
	} else if y == 0 {
		return v < a[x][y+1] && v < a[x+1][y] && v < a[x-1][y]
	} else if y == cols-1 {
		return v < a[x][y-1] && v < a[x+1][y] && v < a[x-1][y]
	} else {
		return v < a[x][y+1] && v < a[x][y-1] && v < a[x+1][y] && v < a[x-1][y]
	}
}

func getInputOrDie() [][]byte {
	lines, err := io.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}
	return io.String1DToByte2D(lines)
}
