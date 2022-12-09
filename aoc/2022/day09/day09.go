package day09

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
	"log"
	"math"
)

var input [][]byte

func SolveP1() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

func nextHPos(x, y, cmdIdx int) (int, int) {
	direc := input[cmdIdx][0]
	steps := input[cmdIdx][1]
	switch direc {
	case 'L':
		return x - steps, y
	case 'R':
		return x + steps, y
	case 'U':
		return x - steps, y
	case 'D':
		return x - steps, y
	default:
		fmt.Printf("wrong move direction for H : %v, cmd(%+v)", direc, input[cmdIdx])
		return math.MaxInt, math.MaxInt
	}
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

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	input = iutils.ExtractByte2DFromString1D(lines, " ", nil, 0)
	//input = make([][]int, len(lines)
	//var tok []string
	//for i, l := range lines {
	//	tok = strings.Split(l, " ")
	//	input[i] = []int{
	//}
	//return lines
}
