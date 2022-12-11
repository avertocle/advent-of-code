package stringz

import "strconv"

func AtoiQ(s string, defVal int) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		i = defVal
	}
	return i
}
