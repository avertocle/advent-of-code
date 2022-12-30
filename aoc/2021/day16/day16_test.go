package day16

import (
	"github.com/avertocle/contests/aoc/testz"
	"testing"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small_01.txt", "16", "0"},
		{"input_small_02.txt", "12", "0"},
		{"input_small_03.txt", "23", "0"},
		{"input_small_04.txt", "31", "0"},
		{"input_final.txt", "963", "0"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
