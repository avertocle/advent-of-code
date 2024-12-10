package day10

import (
	"github.com/avertocle/contests/aoc/testz"
	"testing"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small_01.txt", "4", "1"},
		{"input_small_02.txt", "8", "1"},
		{"input_small_03.txt", "80", "10"},
		{"input_final.txt", "6768", "351"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
