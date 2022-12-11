package day10

import (
	"fmt"
	"github.com/avertocle/contests/io/bytez"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/outils"
	"log"
	"strings"
)

var inpIns []string
var inpVal []int
var inpLen int

func SolveP1() string {
	cycleVals := runSim()
	ansCtrs := []int{20, 60, 100, 140, 180, 220}
	ans := 0
	for _, c := range ansCtrs {
		ans += cycleVals[c] * c
		//fmt.Printf("%v * %v = %v\n", c, cycleVals[c], c*cycleVals[c])
	}
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	cycleVals := runSim()
	scrX := 40
	scrY := 6
	screen := bytez.Init2D(scrY, scrX, '.')
	var cycPtr, pixMid int
	var isLit bool
	for y := 0; y < scrY; y++ {
		for x := 0; x < scrX; x++ {
			cycPtr = (y)*scrX + (x + 1)
			pixMid = cycleVals[cycPtr]
			isLit = (x == pixMid-1 || x == pixMid || x == pixMid+1)
			//fmt.Printf("scrPos(%v) vs cycPtr(%v) vs pixMid(%v) = %v\n", x, cycPtr, pixMid, isLit)
			if isLit {
				screen[y][x] = '#'
			}
		}
	}
	outils.PrettyArray2DByte(screen)
	ans := "0"
	return fmt.Sprintf("%v", ans)
}

func runSim() []int {
	cycVals := make([]int, 250)
	val := 1
	cycPtr := 1
	cycVals[1] = 1
	for i := 0; i < inpLen; i++ {
		cycPtr++
		cycVals[cycPtr] = val
		if strings.Compare(inpIns[i], "addx") == 0 {
			cycVals[cycPtr] = val
			cycPtr++
			val += inpVal[i]
			cycVals[cycPtr] = val
		}
		// nothing to do for noop
	}
	return cycVals
}

/***** Common Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	if err != nil {
		log.Fatalf("iutils error | %v", err)
	}
	inpLen = len(lines)
	inpVal = iutils.ExtractInt1DFromString1D(lines, " ", 1, 0)
	inpIns = iutils.ExtractString1DFromString1D(lines, " ", 0, "invalid")
}
