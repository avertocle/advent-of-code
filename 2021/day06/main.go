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

func main() {
	metrics.ProgStart()
	input := getInputOrDie()
	metrics.InputLen(len(input))

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func problem1(input []int) {

}

func problem2(input []int) {

}

func getInputOrDie() []int {
	lines, err := io.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}

	input := make([]int, len(lines))
	tokens := strings.Split(lines[0], ",")
	for i, t := range tokens {
		input[i], _ = strconv.Atoi(t)
	}
	return input
}
