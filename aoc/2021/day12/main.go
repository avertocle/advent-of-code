package main

import (
	"fmt"
	input2 "github.com/avertocle/contests/io/input"
	"log"
	"strings"

	"github.com/avertocle/contests/metrics"
)

const inputFilePath = "input.txt"

var in *input

type input struct {
	poly    []byte
	rules   map[string]byte
	pairMap map[string]int
}

func main() {
	metrics.ProgStart()
	in = getInputOrDie()
	metrics.InputLen(len(in.rules))

	//ans := problem1()
	//ans := problem2()
	//fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func getInputOrDie() *input {
	lines, err := input2.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}
	poly := []byte(lines[0])
	rules := make(map[string]byte)
	var tokens []string
	for i := 1; i < len(lines); i++ {
		tokens = strings.Split(lines[i], "->")
		rules[strings.TrimSpace(tokens[0])] = []byte(strings.TrimSpace(tokens[1]))[0]
	}
	return &input{
		poly:  poly,
		rules: rules,
	}
}

/***** Logic Begins here *****/

const simCount = 40
