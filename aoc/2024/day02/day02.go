package day02

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

const DirPath = "../2024/day02"

var gInput [][]int

func SolveP1() string {
	ans := 0
	for i := 0; i < len(gInput); i++ {
		if isSafe(gInput[i]) {
			ans++
		}
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	for i := 0; i < len(gInput); i++ {
		if isSafe(gInput[i]) || isSafeWithLevelRemoval(gInput[i]) {
			ans++
		}
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

func isSafeWithLevelRemoval(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		remArray := arrz.RemoveElement1D(arr, i)
		if isSafe(remArray) {
			return true
		}
	}
	return false
}

/***** Common Functions *****/

func isSafe(arr []int) bool {
	if isAllIncreasingWithGap(arr) || isAllDecreasingWithGap(arr) {
		return true
	}
	return false
}

func isAllIncreasingWithGap(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < 1 || arr[i+1]-arr[i] > 3 {
			return false
		}
	}
	return true
}

func isAllDecreasingWithGap(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]-arr[i+1] < 1 || arr[i]-arr[i+1] > 3 {
			return false
		}
	}
	return true
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractInt2DFromString1D(lines, " ", nil, -1)
}
