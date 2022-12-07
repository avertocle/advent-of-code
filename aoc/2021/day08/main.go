package main

import (
	"fmt"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"log"
	"math"
	"sort"
	"strings"

	"github.com/avertocle/contests/metrics"
)

const inputFilePath = "input.txt"

var patternSizes []int // index denotes the digit

func main() {
	metrics.ProgStart()
	patterns, digits := getInputOrDie()
	metrics.InputLen(len(patterns))

	initPatternSizes()

	// ans := problem1(digits)
	// fmt.Printf("ans = %v\n", ans)

	ans := problem2(patterns, digits)
	fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func initPatternSizes() {
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

func problem2(patterns, digits [][]string) int {
	var enc *encoding
	dec := 0
	sum := 0
	for i, pattern := range patterns {
		enc = NewEncoding(pattern)
		enc.decode()
		fmt.Printf("%+v\n", enc)
		dec = enc.decodeDigits(digits[i])
		sum += dec
	}
	return sum
}

type encoding struct {
	l  []string
	m1 map[string]int
	m2 map[int]string
}

func NewEncoding(input []string) *encoding {
	e := new(encoding)
	e.l = make([]string, 10)
	e.m1 = make(map[string]int)
	e.m2 = make(map[int]string)
	for i, s := range input {
		e.l[i] = SortString(s)
	}
	return e
}

func (e *encoding) decode() {
	for _, s := range e.l {
		if len(s) == 2 {
			e.m1[s] = 1
			e.m2[1] = s
		} else if len(s) == 4 {
			e.m1[s] = 4
			e.m2[4] = s
		} else if len(s) == 3 {
			e.m1[s] = 7
			e.m2[7] = s
		} else if len(s) == 7 {
			e.m1[s] = 8
			e.m2[8] = s
		}
	}

	aeg := complement(e.m2[4])
	acf := e.m2[7]
	for _, s := range e.l {
		if len(s) == 5 && hasAll(s, aeg) {
			e.m1[s] = 2
			e.m2[2] = s
		} else if len(s) == 6 && !hasAll(s, aeg) {
			e.m1[s] = 9
			e.m2[9] = s
		} else if len(s) == 5 && hasAll(s, acf) {
			e.m1[s] = 3
			e.m2[3] = s
		} else if len(s) == 6 && !hasAll(s, acf) {
			e.m1[s] = 6
			e.m2[6] = s
		} else if len(s) == 5 {
			e.m1[s] = 5
			e.m2[5] = s
		} else if len(s) == 6 {
			e.m1[s] = 0
			e.m2[0] = s
		}
	}
}

func (e *encoding) decodeDigits(digits []string) int {
	sorted := ""
	dec := 0
	for i, digit := range digits {
		sorted = SortString(digit)
		dec += int(math.Pow10(len(digits)-i-1)) * e.m1[sorted]
		fmt.Printf("%v = %v\n", digit, dec)
	}
	return dec
}

func hasAll(str, chars string) bool {
	for _, x := range strings.Split(chars, "") {
		if !strings.Contains(str, x) {
			return false
		}
	}
	return true
}

func complement(str string) string {
	c := ""
	for _, x := range []string{"a", "b", "c", "d", "e", "f", "g"} {
		if !strings.Contains(str, x) {
			c += x
		}
	}
	return c
}

func SortString(str string) string {
	chars := strings.Split(str, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func getInputOrDie() ([][]string, [][]string) {
	lines, err := iutils.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}

	patterns := stringz.Init2D(len(lines), 10)
	digits := stringz.Init2D(len(lines), 4)
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
