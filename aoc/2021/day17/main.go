package main

import (
	"fmt"
	"log"

	"github.com/avertocle/contests/io"
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
	lines, err := io.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}
	return lines
}
