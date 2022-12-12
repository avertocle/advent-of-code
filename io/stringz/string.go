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
