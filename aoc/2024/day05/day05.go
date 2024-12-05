package day05

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"slices"
)

const DirPath = "../2024/day05"

var gInputRules, gInputPages [][]int

func SolveP1() string {
	ans := 0
	for _, page := range gInputPages {
		if checkAllRules(page) {
			ans += intz.FindMid1D(page)
		}
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	for _, page := range gInputPages {
		if !checkAllRules(page) {
			newPage := sortPageByRuleSet(page)
			ans += intz.FindMid1D(newPage)
		}
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

func sortPageByRuleSet(page []int) []int {
	sortedPage := make([]int, len(page))
	copy(sortedPage, page)
	slices.SortFunc(sortedPage, func(i, j int) int {
		for _, rule := range gInputRules {
			if rule[0] == i && rule[1] == j {
				return -1
			}
		}
		return 0
	})
	return sortedPage
}

/***** Common Functions *****/

func checkAllRules(page []int) bool {
	for _, rule := range gInputRules {
		s := intz.FindByVal1D(page, rule[0])
		e := intz.FindByVal1D(page, rule[1])
		if len(s) >= 1 && len(e) >= 1 && s[0] > e[0] {
			return false
		}
	}
	return true

}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	emptyLineIdx := stringz.FindEmpty1D(lines)[0]
	gInputRules = iutils.ExtractInt2DFromString1D(lines[0:emptyLineIdx], "|", nil, -1)
	gInputPages = iutils.ExtractInt2DFromString1D(lines[emptyLineIdx+1:], ",", nil, -1)
}
