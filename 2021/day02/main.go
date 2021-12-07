package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/avertocle/adventofcode/io"
	"github.com/avertocle/adventofcode/metrics"
)

const inputFilePath = "input_small.txt"

func main() {
	metrics.ProgStart()
	input := getInputOrDie()
	metrics.InputLen(len(input))

	//hpos, vpos := problem1(input)

	hpos, vpos := problem2(input)

	fmt.Printf("hpos(%v) vpos(%v) ans(%v) \n", hpos, vpos, hpos*vpos)
	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func problem1(input [][]int) (int, int) {
	inputLen := len(input)
	hpos := 0
	vpos := 0
	for i := 0; i < inputLen; i++ {
		if input[i][0] == 1 {
			hpos += input[i][1]
		} else if input[i][0] == 2 {
			vpos -= input[i][1]
			if vpos < 0 {
				vpos = 0
			}
		} else {
			vpos += input[i][1]
		}
	}
	return hpos, vpos
}

func problem2(input [][]int) (int, int) {
	inputLen := len(input)
	hpos := 0
	vpos := 0
	aim := 0
	for i := 0; i < inputLen; i++ {
		if input[i][0] == 1 {
			hpos += input[i][1]
			vpos += aim * input[i][1]
		} else if input[i][0] == 2 {
			aim -= input[i][1]
			if aim < 0 {
				aim = 0
			}
		} else {
			aim += input[i][1]
		}
	}
	return hpos, vpos
}

func getInputOrDie() [][]int {
	lines, err := io.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}

	inputLen := len(lines)
	input := make([][]int, inputLen)
	tokens := make([]string, 2)
	for i := 0; i < inputLen; i++ {
		tokens = strings.Split(lines[i], " ")
		amt, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatalf("input error while strconv.Atoi (%v) (%v) (%v) | %v", i, lines[i], tokens, err)
		}
		dir := -1
		switch tokens[0] {
		case "forward":
			dir = 1
			break
		case "up":
			dir = 2
			break
		case "down":
			dir = 3
			break
		}
		if dir == -1 {
			log.Fatalf("input error while getting dir (%v) (%v) (%v) | %v", i, lines[i], tokens, err)
		}
		input[i] = []int{dir, amt}
	}
	return input
}
