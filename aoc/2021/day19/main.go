package main

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
	"log"

	"github.com/avertocle/contests/metrics"
)

const inputFilePath = "input.txt"

func main() {
	metrics.ProgStart()
	input := getInputOrDie()
	metrics.InputLen(len(input))

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func problem1() {

}

func problem2() {

}

func getInputOrDie() []string {
	lines, err := iutils.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	return lines
}
