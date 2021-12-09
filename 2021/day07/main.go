package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

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

func problem1(input []int) int {
	min, _ := io.FindMin1DInt(input)
	max, _ := io.FindMax1DInt(input)
	fuelCost := 0
	minFuelCost := math.MaxInt32
	for i := min; i <= max; i++ {
		fuelCost = 0
		for _, x := range input {
			fuelCost += int(math.Abs(float64(i - x)))
		}
		if fuelCost < minFuelCost {
			minFuelCost = fuelCost
		}
		fmt.Printf("%v %v %v\n", i, fuelCost, minFuelCost)
	}
	return minFuelCost
}

func problem2(input []int) {

}

func getInputOrDie() []int {
	lines, err := io.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}

	tokens := strings.Split(lines[0], ",")
	input := make([]int, len(tokens))
	for i, t := range tokens {
		input[i], _ = strconv.Atoi(t)
	}
	return input
}
