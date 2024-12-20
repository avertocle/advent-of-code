package day13

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
)

const DirPath = "../2024/day13"

var gPrizes []*arrz.Idx2D[int64]
var gBtnA []*arrz.Idx2D[int64]
var gBtnB []*arrz.Idx2D[int64]

func SolveP1() string {
	var ans int64
	for i, prize := range gPrizes {
		a, b := solveEquation(gBtnA[i], gBtnB[i], prize)
		ans += a*3 + b
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	var ans int64
	scale := int64(10000000000000)
	for i, prize := range gPrizes {
		scaledPrize := arrz.NewIdx2D(prize.I+scale, prize.J+scale)
		a, b := solveEquation(gBtnA[i], gBtnB[i], scaledPrize)
		ans += a*3 + b
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

func solveEquation(b1, b2, prize *arrz.Idx2D[int64]) (int64, int64) {
	/*
		Simple 2 eqn 2 variable formula
		e1 : a*btnA.I + b*btnB.I = prize.I
		e2 : a*btnA.J + b*btnB.J = prize.J
		b : (b1.I*prize.J - b1.J*prize.I) / (b2.J*b1.I - b1.J*b2.I)
		a : (prize.I - b*b2.I) / b1.I
	*/
	nb := b1.I*prize.J - b1.J*prize.I
	db := b2.J*b1.I - b1.J*b2.I
	if nb%db != 0 {
		return 0, 0
	}
	b := nb / db
	na := prize.I - b*b2.I
	da := b1.I
	if na%da != 0 {
		return 0, 0
	}
	return na / da, nb / db
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, true)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gPrizes, gBtnA, gBtnB = nil, nil, nil
	for i := 0; i < len(lines)-2; i += 3 {
		t := stringz.SplitMultiTrimSpace(lines[i], []string{" ", "+", ","})
		btnA := arrz.NewIdx2D(stringz.AtoI64(t[3], -1), stringz.AtoI64(t[5], -1))
		t = stringz.SplitMultiTrimSpace(lines[i+1], []string{" ", "+", ","})
		btnB := arrz.NewIdx2D(stringz.AtoI64(t[3], -1), stringz.AtoI64(t[5], -1))
		t = stringz.SplitMultiTrimSpace(lines[i+2], []string{" ", "=", ","})
		prize := arrz.NewIdx2D(stringz.AtoI64(t[2], -1), stringz.AtoI64(t[4], -1))
		gPrizes = append(gPrizes, prize)
		gBtnA = append(gBtnA, btnA)
		gBtnB = append(gBtnB, btnB)
	}
	//fmt.Println(arrz.Idx2DListToStr(gBtnA))
	//fmt.Println(arrz.Idx2DListToStr(gBtnB))
	//fmt.Println(arrz.Idx2DListToStr(gPrizes))
}
