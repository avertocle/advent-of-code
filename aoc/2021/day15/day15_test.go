package day15

import (
	"github.com/avertocle/contests/aoc/testz"
	"testing"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small.txt", "40", "315"},
		{"input_final.txt", "698", "3022"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
