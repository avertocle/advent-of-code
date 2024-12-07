package stringz

import (
	"strconv"
	"strings"
)

func AtoI(s string, defVal int) int {
	s = strings.TrimSpace(s)
	i, e := strconv.Atoi(s)
	if e != nil {
		i = defVal
	}
	return i
}

func AtoI64(s string, defVal int64) int64 {
	s = strings.TrimSpace(s)
	i, e := strconv.ParseInt(s, 10, 64)
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

func SplitMultiTrimSpace(s string, seps []string) []string {
	tokens := []string{s}
	var temp []string
	for _, sep := range seps {
		temp = []string{}
		for _, t := range tokens {
			t = strings.TrimSpace(t)
			for _, t2 := range strings.Split(t, sep) {
				t2 = strings.TrimSpace(t2)
				if len(t2) > 0 {
					temp = append(temp, t2)
				}
			}
		}
		tokens = temp
		//fmt.Printf("tokens = %v\n", strings.Join(tokens, "|"))
	}
	return tokens
}
