package main

import (
	"fmt"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/avertocle/contests/metrics"
)

const inputFilePath = "input.txt"

func main() {
	metrics.ProgStart()
	input := getInputOrDie()
	metrics.InputLen(len(input))

	// ans := problemBoth(iutils, fuelCost1)
	// fmt.Printf("ans = %v\n", ans)

	ans := problemBoth(input, fuelCost2)
	fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func problemBoth(input []int, calcFuelCost func(int, int) int64) int64 {
	min, _ := intz.FindMin1D(input)
	max, _ := intz.FindMax1D(input)
	fuelCost := int64(0)
	minFuelCost := int64(math.MaxInt64)
	for i := min; i <= max; i++ {
		fuelCost = 0
		for _, x := range input {
			fuelCost += calcFuelCost(x, i)
		}
		if fuelCost < minFuelCost {
			minFuelCost = fuelCost
		}
		fmt.Printf("%v %v %v\n", i, fuelCost, minFuelCost)
	}
	return minFuelCost
}

func fuelCost1(src, des int) int64 {
	return int64(math.Abs(float64(des - src)))
}

func fuelCost2(src, des int) int64 {
	n := int64(math.Abs(float64(des - src)))
	return (n * (n + 1)) / 2
}

func getInputOrDie() []int {
	lines, err := iutils.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}

	tokens := strings.Split(lines[0], ",")
	input := make([]int, len(tokens))
	for i, t := range tokens {
		input[i], _ = strconv.Atoi(t)
	}
	return input
}
