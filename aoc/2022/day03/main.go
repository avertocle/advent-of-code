package main

import (
	"fmt"
	"github.com/avertocle/contests/io"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/metrics"
	"log"
)

const inputFilePath = "input.txt"

func main() {
	metrics.ProgStart()
	input := getInputOrDie()

	io.PrettyArray2DByte(input)

	ans01 := solvePartOne(input)
	fmt.Printf("answer part-01 = %v\n", ans01)

	ans02 := solvePartTwo(input)
	fmt.Printf("answer part-02 = %v\n", ans02)
}

func solvePartOne(input [][]byte) int {
	pSum := 0
	var cItems []byte
	midIdx := 0
	for _, a2d := range input {
		midIdx = len(a2d) / 2
		cItems = findCommonItemsIn2Lists(a2d[:midIdx], a2d[midIdx:])
		pSum += getItemPriority(cItems[0])
	}
	return pSum
}

func solvePartTwo(input [][]byte) int {
	pSum := 0
	var cItems []byte
	for i := 0; i < len(input)-2; i += 3 {
		cItems = findCommonItemsInNLists(input[i], input[i+1], input[i+2])
		fmt.Printf("%v vs %v vs %v = %v\n", string(input[i]), string(input[i+1]), string(input[i+2]), string(cItems))
		pSum += getItemPriority(cItems[0])
	}
	return pSum
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

// input : [][]byte : each row contains a line of iutils
func getInputOrDie() [][]byte {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	input := iutils.ExtractByte2DFromString1D(lines, " ", nil, 0)
	return input
}

/***** PART 01 Functions *****/

/***** PART 02 Functions *****/
