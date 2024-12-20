/*
- some func(s) receive error as an arg to be chainable, bad design but looks cool
- no error handling unless absolutely required
#competitive-programming #in-that-zone
*/

package iutils

import (
	"fmt"
	"github.com/avertocle/contests/io/stringz"
	"strconv"
	"strings"
)

func ExtractInt2DFromString1D(lines []string, sep string, cols []int, defaultVal int) [][]int {
	ans := make([][]int, len(lines))
	var temp []int
	for i, line := range lines {
		temp = ExtractInt1DFromString0D(line, sep, defaultVal)
		ans[i] = make([]int, len(cols))
		if cols == nil {
			ans[i] = temp
		} else {
			for j, col := range cols {
				ans[i][j] = temp[col]
			}
		}
	}
	return ans
}

func ExtractInt1DFromString1D(lines []string, sep string, col int, defaultVal int) []int {
	ans := make([]int, len(lines))
	var tokens []string
	for i, line := range lines {
		tokens = strings.Split(line, sep)
		ans[i] = defaultVal
		if col == -1 {
			ans[i] = stringz.AtoI(line, defaultVal)
		} else if col < len(tokens) {
			ans[i] = stringz.AtoI(tokens[col], defaultVal)
		}
	}
	return ans
}

// "12-13-14" ==> [12,13,14]
func ExtractInt1DFromString0D(line string, sep string, defaultVal int) []int {
	var err error
	tokens := strings.Split(line, sep)
	ans := make([]int, len(tokens))
	for i, t := range tokens {
		ans[i], err = strconv.Atoi(strings.TrimSpace(t))
		if err != nil {
			fmt.Printf("strconv.Atoi failed for line(%v) char(%v) | err = %v | using default (%v) \n", line, t, err, defaultVal)
			ans[i] = defaultVal
		}
	}
	return ans
}
