package day12

import (
	"github.com/avertocle/contests/aoc/testz"
	"testing"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small_01.txt", "140", "80"},
		{"input_small_02.txt", "1930", "1206"},
		{"input_final.txt", "1477924", "841934"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
