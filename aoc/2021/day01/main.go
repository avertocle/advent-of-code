package main

import (
	"fmt"
	"github.com/avertocle/contests/io"
	"github.com/avertocle/contests/metrics"
	"log"
)

const inputFilePath = "input.txt"
const window = 3

func main() {
	metrics.ProgStart()

	input := getInputOrDie()
	var ans, sum1, sum2 int
	inputLen := len(input)
	metrics.InputLen(inputLen)
	for i := 0; i < inputLen-window; i++ {
		sum1 = 0
		sum2 = 0
		for j := 0; j < window; j++ {
			sum1 += input[i+j]
			sum2 += input[i+j+1]
		}
		if sum2 > sum1 {
			ans++
		}
	}
	fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func getInputOrDie() []int {
	lines, err := io.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}
	input := io.String1DToInt1D(lines, " ")
	return input
}
