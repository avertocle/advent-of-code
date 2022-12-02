package main

import (
	"fmt"
	input2 "github.com/avertocle/contests/io/input"
	"log"

	"github.com/avertocle/contests/metrics"
)

const inputFilePath = "input.txt"

var in *input

type input struct {
}

func main() {
	metrics.ProgStart()
	in = getInputOrDie()
	//metrics.InputLen(len(in.rules))

	ans := problem1()
	//ans := problem2()
	fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func getInputOrDie() *input {
	lines, err := input2.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}
	fmt.Printf("input-len = %v", len(lines))
	return &input{
		// poly:  poly,
		// rules: rules,
	}
}

/***** Logic Begins here *****/

const simCount = 40

func problem1() int {
	return 0
}

func problem2() int {
	return 0
}
