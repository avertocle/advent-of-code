package day21

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"strings"
)

var gInput []int

func SolveP1() string {
	p1pos, p2pos := gInput[0], gInput[1]
	p1Score, p2Score, drollCtr, dval := 0, 0, 0, 0
	drollPerRound, dvalMax, boardSize, winScore := 3, 100, 10, 1000
	for {
		p1pos, dval = playOneTurn(p1pos, dval, drollPerRound, dvalMax, boardSize)
		drollCtr += drollPerRound
		p1Score += p1pos
		if p1Score >= winScore {
			break
		}
		p2pos, dval = playOneTurn(p2pos, dval, drollPerRound, dvalMax, boardSize)
		drollCtr += drollPerRound
		p2Score += p2pos
		if p2Score >= winScore {
			break
		}
		//fmt.Printf("%v  rolls p1@%v p2@%v \n", drollCtr, p1Score, p2Score)
	}
	ans := drollCtr * p2Score
	//fmt.Printf("drolls(%v) p1Score(%v) p2Score(%v) ans(%v) \n", drollCtr, p1Score, p2Score, ans)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

/***** Common Functions *****/

//return final-pos and final-die-val
func playOneTurn(pos, dval, drolls, dvalMax, posMax int) (int, int) {
	dValTot := 0
	for i := 0; i < drolls; i++ {
		dval = intz.IncBounded(dval, 1, dvalMax)
		dValTot += dval
	}
	pos = intz.IncBounded(pos, dValTot, posMax)
	//fmt.Printf("drolls(%v) dval(%v), pos(%v) \n", drolls, dValTot, pos)
	return pos, dval
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInput = []int{
		stringz.AtoiQ(strings.Fields(lines[0])[4], -1),
		stringz.AtoiQ(strings.Fields(lines[1])[4], -1),
	}
	//fmt.Printf("start pos %v\n", gInput)
}
