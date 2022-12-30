package day03

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
)

// gInput : [][]byte : each row contains a line of iutils
var gInput [][]byte

func SolveP1() string {
	pSum := 0
	var cItems []byte
	midIdx := 0
	for _, a2d := range gInput {
		midIdx = len(a2d) / 2
		cItems = findCommonItemsIn2Lists(a2d[:midIdx], a2d[midIdx:])
		pSum += getItemPriority(cItems[0])
	}
	return fmt.Sprintf("%v", pSum)
}

func SolveP2() string {
	pSum := 0
	var cItems []byte
	for i := 0; i < len(gInput)-2; i += 3 {
		cItems = findCommonItemsInNLists(gInput[i], gInput[i+1], gInput[i+2])
		//fmt.Printf("%v vs %v vs %v = %v\n", string(gInput[i]), string(gInput[i+1]), string(gInput[i+2]), string(cItems))
		pSum += getItemPriority(cItems[0])
	}
	return fmt.Sprintf("%v", pSum)
}

/***** Common Functions *****/

func getItemPriority(item byte) int {
	if item >= 'a' && item <= 'z' {
		return int(item-'a') + 1
	}
	if item >= 'A' && item <= 'Z' {
		return int(item-'A') + 27
	}
	return 0
}

func findCommonItemsInNLists(listOfList ...[]byte) []byte {
	cItems := listOfList[0]
	for i := 1; i < len(listOfList); i++ {
		cItems = findCommonItemsIn2Lists(cItems, listOfList[i])
	}
	return cItems
}

func findCommonItemsIn2Lists(list1, list2 []byte) []byte {
	list1Map := listToMap(list1)
	cItems := make([]byte, 0)
	for _, b := range list2 {
		if _, ok := list1Map[b]; ok {
			cItems = append(cItems, b)
		}
	}
	//fmt.Printf("%v vs %v = %v\n", string(list1), string(list2), string(cItems))
	return cItems
}

// map byte => count
func listToMap(list []byte) map[byte]int {
	m := make(map[byte]int)
	for _, b := range list {
		if v, ok := m[b]; ok {
			m[b] = v + 1
		} else {
			m[b] = 1
		}
	}
	return m
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = iutils.ExtractByte2DFromString1D(lines, " ", nil, 0)
}

/***** PART 01 Functions *****/

/***** PART 02 Functions *****/
