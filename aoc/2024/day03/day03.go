package day03

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"regexp"
	"slices"
	"strings"
)

const DirPath = "../2024/day03"

var gInput string

func SolveP1() string {
	ans := 0
	regexPattern := `mul\(\d{1,3},\d{1,3}\)`
	r := regexp.MustCompile(regexPattern)
	matches := r.FindAllString(gInput, -1)
	for _, match := range matches {
		ans += evaluate(match)
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	r1 := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	r2 := regexp.MustCompile(`do\(\)`)
	r3 := regexp.MustCompile(`don't\(\)`)
	var expIdxs, dos, donts [][]int
	expIdxs = append(expIdxs, r1.FindAllStringIndex(gInput, -1)...)
	dos = append(dos, r2.FindAllStringIndex(gInput, -1)...)
	donts = append(donts, r3.FindAllStringIndex(gInput, -1)...)
	sortFunc := func(a, b []int) int { return a[0] - b[0] }
	slices.SortFunc(expIdxs, sortFunc)
	slices.SortFunc(dos, sortFunc)
	slices.SortFunc(donts, sortFunc)

	for _, expIdx := range expIdxs {
		lastDoIdx := findLastBeforeMe(expIdx, dos)
		lastDontIdx := findLastBeforeMe(expIdx, donts)
		exp := gInput[expIdx[0]:expIdx[1]]
		if lastDoIdx == nil && lastDontIdx == nil {
			ans += evaluate(exp)
		} else if lastDontIdx == nil {
			ans += evaluate(exp)
		} else if lastDoIdx == nil {
			ans += 0
		} else if lastDoIdx[1] > lastDontIdx[1] {
			ans += evaluate(exp)
		}
		//fmt.Println(expIdx, exp, "lastDoIdx", lastDoIdx, "lastDontIdx", lastDontIdx, "ans", ans)
	}

	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

func findLastBeforeMe(expIdx []int, arr [][]int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i][1]-1 < expIdx[0] { // slice [a:b] is from index 'a' to 'b-1'
			return arr[i]
		}
	}
	return nil
}

/***** Common Functions *****/

func evaluate(s string) int {
	tokens := stringz.SplitMulti(s, []string{"(", ",", ")"})
	result := stringz.AtoiQ(tokens[1], 0) * stringz.AtoiQ(tokens[2], 0)
	return result
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = strings.Join(lines, "")
}
