package day07

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
)

var input []string

func SolveP1() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		fmt.Printf("iutils error | %v", err)
	}
	input = lines
}
