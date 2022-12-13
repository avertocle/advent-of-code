package day17

import (
	"github.com/avertocle/contests/aoc/testz"
	"testing"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small.txt", "45", "112"},
		{"input_final.txt", "12090", "5059"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
