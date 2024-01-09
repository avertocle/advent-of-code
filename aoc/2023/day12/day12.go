package day12

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"strings"
)

var gInputPatterns [][]byte
var gInputCounts [][]int

const DirPath = "../2023/day12"

func SolveP1() string {
	ans := 0
	for i := 0; i < len(gInputPatterns); i++ {
		matchCount := 0
		countMatches(gInputPatterns[i], gInputCounts[i], &matchCount)
		//fmt.Println(matchCount)
		ans += matchCount
	}
	//countMatches([]byte("??"), []int{2, 2}, &ans)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	for i := 0; i < len(gInputPatterns); i++ {
		matchCount := 0
		newPattern := []byte(strings.Repeat("?"+string(gInputPatterns[i]), 5))[1:]
		newCounts := intz.Repeat1D(gInputCounts[i], 5)
		fmt.Println(string(newPattern))
		fmt.Println(newCounts)
		countMatches(newPattern, newCounts, &matchCount)
		fmt.Println(matchCount)
		ans += matchCount
	}
	//countMatches([]byte("??"), []int{2, 2}, &ans)
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

/***** P1 Functions *****/

func countMatches(pattern []byte, counts []int, matchCount *int) {
	//fmt.Println(string(pattern))
	unknowns := bytez.FindAll(pattern, '?')
	if len(unknowns) == 0 {
		if isValidPattern(pattern, counts) {
			*matchCount++
			//fmt.Println("match")
		}
		//*matchCount++
		return
	}
	countMatches(copyReplace(pattern, unknowns[0], '#'), counts, matchCount)
	countMatches(copyReplace(pattern, unknowns[0], '.'), counts, matchCount)
}

func copyReplace(pattern []byte, idx int, val byte) []byte {
	pattern1 := make([]byte, len(pattern))
	copy(pattern1, pattern)
	pattern1[idx] = val
	return pattern1
}

func isValidPattern(pattern []byte, counts []int) bool {
	pCount := make([]int, 0)
	for i := 0; i < len(pattern); i++ {
		c := 0
		for ; i < len(pattern) && pattern[i] == '#'; i++ {
			c++
		}
		if c > 0 {
			pCount = append(pCount, c)
		}
	}
	return intz.Compare1D(pCount, counts) == 0
}

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInputPatterns = make([][]byte, len(lines))
	gInputCounts = make([][]int, len(lines))
	for i, line := range lines {
		tokens := strings.Fields(line)
		gInputPatterns[i] = iutils.ExtractByte1DFromString0D(tokens[0], "", -1, 0)
		gInputCounts[i] = iutils.ExtractInt1DFromString0D(tokens[1], ",", -1)
	}
	//fmt.Println(gInputPatterns)
	//fmt.Println(gInputCounts)
}
