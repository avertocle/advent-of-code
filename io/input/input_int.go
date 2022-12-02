/*
- some funcs recieve error as an arg to be chainable, bad design but looks cool
- no error handling unless absolutely required
#competitive-programming #in-that-zone
*/

package input

import (
	"fmt"
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
	var temp []int
	for i, line := range lines {
		temp = ExtractInt1DFromString0D(line, sep, defaultVal)
		ans[i] = temp[col]
	}
	return ans
}

// "12-13-14" ==> [12,13,14]
func ExtractInt1DFromString0D(line string, sep string, defaultVal int) []int {
	var err error
	tokens := strings.Split(line, sep)
	ans := make([]int, len(tokens))
	for i, t := range tokens {
		ans[i], err = strconv.Atoi(t)
		if err != nil {
			fmt.Printf("strconv.Atoi failed for (%v) | err = %v | using default (%v) \n", t, err, defaultVal)
			ans[i] = defaultVal
		}
	}
	return ans
}
