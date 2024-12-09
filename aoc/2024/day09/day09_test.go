package day09

import (
	"github.com/avertocle/contests/aoc/testz"
	"testing"
)

func TestAll(t *testing.T) {
	testCases := [][]string{
		{"input_small_01.txt", "60", "132"},
		{"input_small_02.txt", "1928", "2858"},
		{"input_final.txt", "6401092019345", "6431472344710"},
	}
	testz.Execute(t, testCases, ParseInput, []func() string{SolveP1, SolveP2})
}
