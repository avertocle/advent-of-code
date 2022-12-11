package day21

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
	"log"
)

var input []string

func ParseInput(inputFilePath string) {
	input = getInputOrDie(inputFilePath)
}

func SolveP1() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

func PrintInputMetadata(inputFilePath string) {
	fmt.Printf("input length = %v\n", len(input))
}

/***** Common Functions *****/

/***** Input *****/

func getInputOrDie(inputFilePath string) []string {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	return lines
}
