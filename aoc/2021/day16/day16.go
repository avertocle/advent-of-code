package day16

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
	"log"
)

var gInput []string

func SolveP1() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

func PrintInputMetadata(inputFilePath string) {
	fmt.Printf("input length = %v\n", len(gInput))
}

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	gInput = lines
}
