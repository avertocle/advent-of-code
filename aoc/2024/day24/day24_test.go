package day24

import (
	"github.com/avertocle/contests/aoc/testz"
	"testing"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_final.txt", "53258032898766", ""},
		{"input_small_p1.txt", "2024", "-"},
		{"input_small_p2.txt", "9", "z0,z1,z2,z5"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
