package iutils

import "strings"

/*
ExtractString2DFromString1D extracts columns for a list of lines
ab,ef ; seps[",","-"] ==> [ab, cd][ef,gh]
uv,wx						[st, uv][wx,yz]
*/
func ExtractString2DFromString1D(lines []string, sep string, cols []int, defaultVal string) [][]string {
	ans := make([][]string, len(lines))
	var tokens []string
	for i, line := range lines {
		tokens = strings.Split(line, sep)
		if cols == nil {
			ans[i] = tokens
		} else {
			ans[i] = make([]string, len(cols))
			for j, c := range cols {
				ans[i][j] = defaultVal
				if c < len(tokens) {
					ans[i][j] = tokens[c]
				}
			}
		}
	}
	return ans
}
