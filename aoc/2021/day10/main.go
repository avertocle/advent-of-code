package main

import (
	"fmt"
	"github.com/avertocle/contests/io/input"
	"log"
	"math"

	"github.com/avertocle/contests/io"
	"github.com/avertocle/contests/metrics"
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

const inputFilePath = "input.txt"

func main() {
	metrics.ProgStart()
	input := getInputOrDie()
	metrics.InputLen(len(input))

	ans1, ans2 := problem2(input)
	fmt.Printf("ans1 (%v) | ans2 (%v) \n", ans1, ans2)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func problem2(input [][]byte) (int, int) {
	var stack *lls.Stack
	var b interface{}
	var ok bool
	var errors []byte
	var autoComp []byte
	errorScore := 0
	var autoCompScores []int
	for _, row := range input {
		stack = lls.New()
		autoComp = make([]byte, 0)
		for i, ele := range row {
			if isOpeningBrace(ele) {
				stack.Push(ele)
			} else if b, ok = stack.Pop(); !ok {
				errors = append(errors, ele)
				errorScore += calcSynErrorScore(ele)
				break
			} else if !isMatchingBrace(ele, b.(byte)) {
				errors = append(errors, ele)
				errorScore += calcSynErrorScore(ele)
				break
			}
			if i == len(row)-1 {
				for stack.Size() > 0 {
					if b, ok := stack.Pop(); ok {
						autoComp = append(autoComp, b.(byte))
					} else {
						break
					}
				}
				autoCompScores = append(autoCompScores, calcAutoCompScore(autoComp))
			}
		}
	}
	fmt.Printf("syn-errors : %q\n", errors)
	fmt.Printf("auto-comp-scores :%v\n", autoCompScores)
	return errorScore, io.FindMiddleElem1DInt(autoCompScores)
}

func isOpeningBrace(b byte) bool {
	return (b == '(' || b == '{' || b == '[' || b == '<')
}

func isMatchingBrace(x, y byte) bool {
	z := int(math.Abs(float64(int(x - y))))
	return (z == 1 || z == 2)
}

func calcAutoCompScore(errors []byte) int {
	score := 0
	v := 0
	for _, e := range errors {
		v = 0
		switch e {
		case '(':
			v = 1
			break
		case '[':
			v = 2
			break
		case '{':
			v = 3
			break
		case '<':
			v = 4
			break
		}
		score *= 5
		score += v
	}
	return score
}

func calcSynErrorScore(errChar byte) int {
	v := 0
	switch errChar {
	case ')':
		v = 3
		break
	case ']':
		v = 57
		break
	case '}':
		v = 1197
		break
	case '>':
		v = 25137
		break
	}
	return v
}

func getInputOrDie() [][]byte {
	lines, err := input.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}
	return input.String1DToByte2D(lines)
}
