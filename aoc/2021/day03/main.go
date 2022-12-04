package main

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
	"log"
	"strconv"

	"github.com/avertocle/contests/metrics"
)

const inputFilePath = "input.txt"

func main() {
	metrics.ProgStart()
	input := getInputOrDie()
	metrics.InputLen(len(input))

	gr, er := problem2(input)

	fmt.Printf("gr(%v) er(%v) ans(%v) \n", gr, er, gr*er)
	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func problem1(input [][]byte) (int, int) {
	cols := len(input[0])
	grStr := ""
	erStr := ""
	for j := 0; j < cols; j++ {
		m, l := getMostLeastCommonBitsAt(input, j)
		grStr += string(m)
		erStr += string(l)
	}
	gr, _ := strconv.ParseInt(grStr, 2, 32)
	er, _ := strconv.ParseInt(erStr, 2, 32)
	return int(gr), int(er)
}

func problem2(input [][]byte) (int, int) {
	cols := len(input[0])
	inputFilteredMost := input
	inputFilteredLeast := input
	for j := 0; j < cols; j++ {
		fmt.Printf("%v %v \n", len(inputFilteredMost), len(inputFilteredLeast))
		m, _ := getMostLeastCommonBitsAt(inputFilteredMost, j)
		if len(inputFilteredMost) > 1 {
			inputFilteredMost = filterByByteAtPos(inputFilteredMost, m, j)
		}

		_, l := getMostLeastCommonBitsAt(inputFilteredLeast, j)
		if len(inputFilteredLeast) > 1 {
			inputFilteredLeast = filterByByteAtPos(inputFilteredLeast, l, j)
		}
	}
	fmt.Println(inputFilteredMost)
	fmt.Println(inputFilteredLeast)
	gr, _ := strconv.ParseInt(string(inputFilteredMost[0]), 2, 32)
	er, _ := strconv.ParseInt(string(inputFilteredLeast[0]), 2, 32)
	return int(gr), int(er)
}

func filterByByteAtPos(input [][]byte, byt byte, pos int) [][]byte {
	rows := len(input)
	inputNew := make([][]byte, 0)
	for i := 0; i < rows; i++ {
		if input[i][pos] == byt {
			inputNew = append(inputNew, input[i])
		}
	}
	return inputNew
}

func getMostLeastCommonBitsAt(input [][]byte, pos int) (byte, byte) {
	rows := len(input)
	ctr0 := 0
	for i := 0; i < rows; i++ {
		if input[i][pos] == '0' {
			ctr0++
		}
	}
	ctr1 := rows - ctr0
	if ctr0 > ctr1 {
		return '0', '1'
	} else if ctr0 < ctr1 {
		return '1', '0'
	} else {
		return '1', '0'
	}
}

func getInputOrDie() [][]byte {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}

	inputLen := len(lines)
	input := make([][]byte, inputLen)
	for i := 0; i < inputLen; i++ {
		input[i] = []byte(lines[i])
	}
	return input
}
