package stringz

import (
	"strconv"
	"strings"
)

func AtoiQ(s string, defVal int) int {
	s = strings.TrimSpace(s)
	i, e := strconv.Atoi(s)
	if e != nil {
		i = defVal
	}
	return i
}

func SplitMulti(s string, seps []string) []string {
	tokens := []string{s}
	var temp []string
	for _, sep := range seps {
		temp = []string{}
		for _, t := range tokens {
			temp = append(temp, strings.Split(t, sep)...)
		}
		tokens = temp
		//fmt.Printf("tokens = %v\n", strings.Join(tokens, "|"))
	}
	return tokens
}
