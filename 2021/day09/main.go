package main

import (
	"fmt"
	"log"

	"github.com/avertocle/adventofcode/io"
	"github.com/avertocle/adventofcode/metrics"
)

const inputFilePath = "input.txt"

func main() {
	metrics.ProgStart()
	input := getInputOrDie()
	metrics.InputLen(len(input))

	ans := problem1(input)
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
				fmt.Printf("%v,%v = %v\n", i, j, ele-'0')
				riskLevel += (int(ele-'0') + 1)
			}
		}
	}
	return riskLevel
}

func problem2() {

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

	input := make([][]byte, len(lines))
	for i, row := range lines {
		input[i] = []byte(row)
	}
	return input
}
