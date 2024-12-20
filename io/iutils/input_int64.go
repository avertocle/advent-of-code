/*
- some func(s) receive error as an arg to be chainable, bad design but looks cool
- no error handling unless absolutely required
#competitive-programming #in-that-zone
*/

package iutils

import (
	"github.com/avertocle/contests/io/stringz"
	"strings"
)

func ExtractInt642DFromString1D(lines []string, sep string, cols []int, defaultVal int64) [][]int64 {
	ans := make([][]int64, len(lines))
	var temp []int64
	for i, line := range lines {
		temp = ExtractInt641DFromString0D(line, sep, defaultVal)
		ans[i] = make([]int64, len(cols))
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

func ExtractInt641DFromString1D(lines []string, sep string, col int, defaultVal int64) []int64 {
	ans := make([]int64, len(lines))
	var tokens []string
	for i, line := range lines {
		tokens = strings.Split(line, sep)
		ans[i] = defaultVal
		if col == -1 {
			ans[i] = stringz.AtoI64(line, defaultVal)
		} else if col < len(tokens) {
			ans[i] = stringz.AtoI64(tokens[col], defaultVal)
		}
	}
	return ans
}

func ExtractInt641DFromString0D(line string, sep string, defaultVal int64) []int64 {
	tokens := strings.Split(line, sep)
	ans := make([]int64, len(tokens))
	for i, t := range tokens {
		ans[i] = stringz.AtoI64(t, defaultVal)
	}
	return ans
}
