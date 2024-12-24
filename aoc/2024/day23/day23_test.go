package day23

import (
	"github.com/avertocle/contests/aoc/testz"
	"testing"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small.txt", "7", "co,de,ka,ta"},
		{"input_final.txt", "893", "cw,dy,ef,iw,ji,jv,ka,ob,qv,ry,ua,wt,xz"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
