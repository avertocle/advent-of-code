package day15

import (
	"github.com/avertocle/contests/aoc/testz"
	"testing"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small_01.txt", "2028", "0"},
		{"input_small_02.txt", "10092", "0"},
		{"input_final.txt", "1476771", "0"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
