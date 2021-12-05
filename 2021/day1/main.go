package main

import (
	"fmt"
	"log"

	"github.com/avertocle/adventofcode/input"
	"github.com/avertocle/adventofcode/metrics"
)

const inputFilePath = "input.txt"
const window = 1

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
	input, err := input.AsIntArray(input.FromFile(inputFilePath))
	if err != nil {
		log.Fatalf("input error | %v", err)
	}
	return input
}
