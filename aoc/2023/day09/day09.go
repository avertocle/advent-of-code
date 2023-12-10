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
		ans += extrapolate(history)
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

func extrapolate(arr []int) int {
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
	fmt.Println(lastElemItr)
	evalue := 0
	for _, v := range lastElemItr {
		evalue += v
	}

	return evalue

}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractInt2DFromString1D(lines, " ", nil, math.MaxInt)
}
