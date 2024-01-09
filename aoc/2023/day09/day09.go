package day09

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"math"
)

var gInput [][]int

const DirPath = "../2023/day09"

func SolveP1() string {
	ans := 0
	for _, history := range gInput {
		ans += extrapolateP1(history)
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	for _, history := range gInput {
		ans += extrapolateP2(history)
	}
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

/***** P1 Functions *****/
func extrapolateP1(arr []int) int {
	lastElemItr := make([]int, 0)
	lastElemItr = append(lastElemItr, arr[len(arr)-1])
	for tempArrSum := -1; tempArrSum != 0; {
		tempArr := make([]int, 0)
		for i := 0; i < len(arr)-1; i++ {
			tempArr = append(tempArr, arr[i+1]-arr[i])
		}
		lastElemItr = append(lastElemItr, tempArr[len(tempArr)-1])
		tempArrSum = intz.Sum1D(tempArr)
		arr = tempArr
	}
	evalue := 0
	for i := len(lastElemItr) - 1; i >= 0; i-- {
		evalue += lastElemItr[i]
	}
	return evalue
}

/***** P2 Functions *****/
func extrapolateP2(arr []int) int {
	firstElemItr := make([]int, 0)
	firstElemItr = append(firstElemItr, arr[0])
	for tempArrSum := -1; tempArrSum != 0; {
		tempArr := make([]int, 0)
		for i := 0; i < len(arr)-1; i++ {
			tempArr = append(tempArr, arr[i+1]-arr[i])
		}
		firstElemItr = append(firstElemItr, tempArr[0])
		tempArrSum = intz.Sum1D(tempArr)
		arr = tempArr
	}
	evalue := 0
	for i := len(firstElemItr) - 1; i >= 0; i-- {
		evalue = firstElemItr[i] - evalue
	}

	return evalue

}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractInt2DFromString1D(lines, " ", nil, math.MaxInt)
}
