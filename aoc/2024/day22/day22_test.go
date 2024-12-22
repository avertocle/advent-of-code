package day22

import (
	"github.com/avertocle/contests/aoc/testz"
	"testing"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small_01.txt", "37327623", "24"},
		{"input_small_02.txt", "37990510", "23"},
		{"input_final.txt", "17005483322", "1910"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
