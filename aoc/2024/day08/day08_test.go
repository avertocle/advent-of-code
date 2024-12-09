package day08

import (
	"github.com/avertocle/contests/aoc/testz"
	"testing"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small_01.txt", "14", "34"},
		{"input_small_02.txt", "3", "9"},
		{"input_final_01.txt", "364", "1231"},
		{"input_final_02.txt", "423", "1287"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
