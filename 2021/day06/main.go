package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/avertocle/adventofcode/io"
	"github.com/avertocle/adventofcode/metrics"
)

const inputFilePath = "input.txt"
const maxFishAge = 9
const simDays = 256

//caches fish-count produced by one fish starting at [init-age] on a [current-day]
var simCache [][]int

func main() {
	metrics.ProgStart()
	input := getInputOrDie()
	metrics.InputLen(len(input))

	simCache = io.Init2DInt(0, maxFishAge, simDays+1)

	ans := problemBoth(input, simDays)
	fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func problemBoth(input []int, totalDays int) int64 {
	var count int64
	for _, x := range input {
		count += int64(getFishCount(x, totalDays))
	}
	return count
}

func getFishCount(currAge, simDaysLeft int) int {
	fishCount := 0
	if simCache[currAge][simDaysLeft] > 0 {
		fishCount = simCache[currAge][simDaysLeft]
	} else if currAge >= simDaysLeft {
		fishCount = 1
	} else if currAge > 0 {
		fishCount = getFishCount(currAge-1, simDaysLeft-1)
		simCache[currAge][simDaysLeft] = fishCount
		return fishCount
	} else {
		fishCount = getFishCount(6, simDaysLeft-1) + getFishCount(8, simDaysLeft-1)
		simCache[currAge][simDaysLeft] = fishCount
	}
	fmt.Printf("getFishCount %v @ %v = %v\n", currAge, simDaysLeft, fishCount)
	return fishCount
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
