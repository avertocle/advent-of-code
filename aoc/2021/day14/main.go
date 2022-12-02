package main

import (
	"fmt"
	input2 "github.com/avertocle/contests/io/input"
	"log"
	"math"
	"strings"

	"github.com/avertocle/contests/io"
	"github.com/avertocle/contests/metrics"
)

const inputFilePath = "input.txt"

var in *input

type input struct {
	poly    []byte
	rules   map[string]byte
	pairMap map[string]int
}

func main() {
	metrics.ProgStart()
	in = getInputOrDie()
	metrics.InputLen(len(in.rules))

	//ans := problem1()
	ans := problem2()
	fmt.Printf("ans = %v\n", ans)

	metrics.ProgEnd()
	fmt.Printf("metrics : [%v]", metrics.ToString())
}

func getInputOrDie() *input {
	lines, err := input2.FromFile(inputFilePath, true)
	if err != nil {
		log.Fatalf("input error | %v", err)
	}
	poly := []byte(lines[0])
	rules := make(map[string]byte)
	var tokens []string
	for i := 1; i < len(lines); i++ {
		tokens = strings.Split(lines[i], "->")
		rules[strings.TrimSpace(tokens[0])] = []byte(strings.TrimSpace(tokens[1]))[0]
	}
	return &input{
		poly:  poly,
		rules: rules,
	}
}

/***** Logic Begins here *****/

const simCount = 40

func problem2() int64 {
	in.pairMap = make(map[string]int)

	// populate initial pair map
	key := ""
	for i := 0; i < len(in.poly)-1; i++ {
		key = string([]byte{in.poly[i], in.poly[i+1]})
		in.pairMap[key]++
	}

	for i := 0; i < simCount; i++ {
		iterate2()
		//fmt.Printf("%+v\n", in.pairMap)
	}
	min, max := findMinMaxCount2()
	fmt.Printf("min(%v), max(%v)\n", min, max)
	return max - min
}

func iterate2() {
	pairMapNew := make(map[string]int)
	k1, k2 := "", ""
	for pat, rep := range in.rules {
		if v, ok := in.pairMap[pat]; ok {
			k1 = string([]byte{pat[0], rep})
			k2 = string([]byte{rep, pat[1]})
			pairMapNew[k1] += v
			pairMapNew[k2] += v
		}
	}
	in.pairMap = pairMapNew
}

func findMinMaxCount2() (int64, int64) {
	charCounts := make([]int64, 26)
	for k, v := range in.pairMap {
		charCounts[int(k[0]-'A')] += int64(v)
		//charCounts[int(k[1]-'A')] += int64(v)
	}
	charCounts[in.poly[len(in.poly)-1]-'A']++
	//	fmt.Printf("%+v\n", charCounts)

	min := int64(math.MaxInt64)
	max := int64(0)
	for _, v := range charCounts {
		if v == 0 {
			continue
		}
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func problem1() int {
	for i := 0; i < simCount; i++ {
		iterate()
	}
	min, max := findMinMaxCount()
	return max - min
}

func iterate() {
	repMap := make(map[int]byte)
	var matches []int
	for pat, rep := range in.rules {
		matches = io.Find1DByteIn1DByte(in.poly, []byte(pat))
		for _, m := range matches {
			repMap[m+1] = rep
		}
	}
	newPoly := make([]byte, len(in.poly)+len(repMap))
	k := 0
	for i := 0; i < len(in.poly); i++ {
		if v, ok := repMap[i]; ok {
			newPoly[k] = v
			k++
		}
		newPoly[k] = in.poly[i]
		k++
	}
	//fmt.Printf("%v => %v\n", string(in.poly), string(newPoly))
	in.poly = newPoly
}

func findMinMaxCount() (int, int) {
	min := math.MaxInt32
	max := 0
	m := io.CountUniqByteIn1DByte(in.poly)
	for _, v := range m {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func libTest() {
	fmt.Printf("%v \n", io.Find1DByteIn1DByte([]byte("VHCKBFOVCHHKOHBPNCKO"), []byte("CY")))

}
