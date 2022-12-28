package day25

import (
	"github.com/avertocle/contests/aoc/testz"
	"testing"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small.txt", "2=-1=0", "0"},
		{"input_final.txt", "2-0==21--=0==2201==2", "0"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
