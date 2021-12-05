package main

import (
	"fmt"

	"github.com/avertocle/adventofcode/metrics"
	"github.com/avertocle/adventofcode/utils"
)

const inputFilePath = "input.txt"

func main() {
	metrics.ProgStart()

	input, err := utils.FetchInputFromFileAsIntArray(inputFilePath)
	if err != nil {
		fmt.Printf("utils.FetchInputFromFileAsIntArray failed | %v", err)
		return
	}

	var ans, window, sum1, sum2 int
	inputLen := len(input)
	window = 3
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

	fmt.Println(inputLen)
	fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}
