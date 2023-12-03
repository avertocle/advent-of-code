package day01

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"strings"
)

var gInput []string

func SolveP1() string {
	ans := 0
	x, y := 0, 0
	for i := 0; i < len(gInput); i++ {
		x = findFirstNumberP1(gInput[i])
		y = findLastNumberP1(gInput[i])
		ans += x*10 + y
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	x, y := 0, 0
	for i := 0; i < len(gInput); i++ {
		x = findFirstNumberP2(gInput[i])
		y = findLastNumberP2(gInput[i])
		ans += x*10 + y
		fmt.Println(gInput[i], x, y, x*10+y)
	}
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

/***** P1 Functions *****/

func findFirstNumberP1(line string) int {
	arr := iutils.ExtractInt1DFromString0D(line, "", -1)
	for i := 0; i < len(arr); i++ {
		if arr[i] > -1 {
			return arr[i]
		}
	}
	return 0
}

func findLastNumberP1(line string) int {
	arr := iutils.ExtractInt1DFromString0D(line, "", -1)
	for i := len(arr); i > 0; i-- {
		if arr[i-1] > -1 {
			return arr[i-1]
		}
	}
	return 0
}

/***** P2 Functions *****/

func findFirstNumberP2(line string) int {
	m := getMap()
	wordIdx, value := 10000, 0
	for k, v := range m {
		idx := strings.Index(line, k)
		if idx != -1 && idx < wordIdx {
			wordIdx = idx
			value = v
		}
	}
	numIdx := 10000
	arr := iutils.ExtractInt1DFromString0D(line, "", -1)
	for i := 0; i < len(arr); i++ {
		if arr[i] > -1 {
			numIdx = i
			break
		}
	}
	if numIdx > -1 && numIdx < wordIdx {
		return arr[numIdx]
	} else if wordIdx > -1 && wordIdx < numIdx {
		return value
	} else {
		errz.HardAssert(false, "findFirstNumberP2 %v %v %v", line, numIdx, wordIdx)
	}

	return 0
}

func findLastNumberP2(line string) int {
	m := getMap()
	wordIdx, value := -1, 0
	for k, v := range m {
		idx := strings.LastIndex(line, k)
		if idx != -1 && idx > wordIdx {
			wordIdx = idx
			value = v
		}
	}
	numIdx := -1
	arr := iutils.ExtractInt1DFromString0D(line, "", -1)
	for i := len(arr); i > 0; i-- {
		if arr[i-1] > -1 {
			numIdx = i - 1
			break
		}
	}
	if numIdx > -1 && numIdx > wordIdx {
		return arr[numIdx]
	} else if wordIdx > -1 && wordIdx > numIdx {
		return value
	} else {
		errz.HardAssert(false, "findLastNumberP2 %v %v %v", line, numIdx, wordIdx)
	}

	return 0
}

func getMap() map[string]int {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	m["six"] = 6
	m["seven"] = 7
	m["eight"] = 8
	m["nine"] = 9
	return m
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = lines
}
