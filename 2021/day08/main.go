package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/avertocle/adventofcode/io"
	"github.com/avertocle/adventofcode/metrics"
)

const inputFilePath = "input.txt"

var patternSizes []int // index denotes the digit

func main() {
	metrics.ProgStart()
	patterns, digits := getInputOrDie()
	metrics.InputLen(len(patterns))

	initPatterSizes()

	ans := problem1(digits)
	fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func initPatterSizes() {
	patternSizes = make([]int, 10)
	patternSizes[0] = 6
	patternSizes[1] = 2
	patternSizes[2] = 5
	patternSizes[3] = 5
	patternSizes[4] = 4
	patternSizes[5] = 5
	patternSizes[6] = 6
	patternSizes[7] = 3
	patternSizes[8] = 7
	patternSizes[9] = 6
}

func problem1(digits [][]string) int {
	toTest := []int{1, 4, 7, 8}
	l := 0
	count := 0
	for _, row := range digits {
		for _, ele := range row {
			l = len(ele)
			for _, t := range toTest {
				if l == patternSizes[t] {
					count++
				}
			}
		}
	}
	return count
}

func problem2() {

}

func getInputOrDie() ([][]string, [][]string) {
	lines, err := io.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}

	patterns := io.Init2DString(len(lines), 10)
	digits := io.Init2DString(len(lines), 4)
	var tok1, tok2 []string
	for i, l := range lines {
		tok1 = strings.Split(l, "|")

		tok2 = strings.Split(strings.TrimSpace(tok1[0]), " ")
		for j, t := range tok2 {
			patterns[i][j] = t
		}

		tok2 = strings.Split(strings.TrimSpace(tok1[1]), " ")
		for j, t := range tok2 {
			digits[i][j] = t
		}
	}
	return patterns, digits
}
