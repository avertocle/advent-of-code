package day07

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/intz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
)

const DirPath = "../2024/day07"

var gResults []int64
var gEquations [][]int64
var gSlotCount int

func SolveP1() string {
	var ans int64
	ops := []byte{'+', '*'}
	allPossibleOps := generateAllOpsPerms(gSlotCount, ops)
	ans = runAllEquations(allPossibleOps)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	var ans int64
	ops := []byte{'+', '*', '@'}
	allPossibleOps := generateAllOpsPerms(gSlotCount, ops)
	ans = runAllEquations(allPossibleOps)
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Common Functions *****/

func runAllEquations(allPossibleOps [][]byte) int64 {
	var totalCaliResult int64
	for i, _ := range gEquations {
		if isEquationCorrect(i, allPossibleOps) {
			totalCaliResult += gResults[i]
		}
	}
	return totalCaliResult
}

func isEquationCorrect(eqnIdx int, allPossibleOps [][]byte) bool {
	eqn := gEquations[eqnIdx]
	result := gResults[eqnIdx]
	for _, ops := range allPossibleOps {
		if evalEquation(eqn, ops, result) {
			return true
		}
	}
	return false
}

func evalEquation(params []int64, ops []byte, result int64) bool {
	ans := params[0]
	for i := 1; i < len(params); i++ {
		if ops[i-1] == '+' {
			ans += params[i]
		} else if ops[i-1] == '*' {
			ans *= params[i]
		} else if ops[i-1] == '@' {
			ans = stringz.AtoI64(fmt.Sprintf("%v%v", ans, params[i]), 0)
		} else {
			errz.HardAssert(false, "invalid operator | %v", ops[i-1])
		}
	}
	return ans == result
}

func generateAllOpsPerms(slotCount int, ops []byte) [][]byte {
	if slotCount == 1 {
		return arrz.Upscale1D(ops)
	}
	newPerms := make([][]byte, 0)
	perms := generateAllOpsPerms(slotCount-1, ops)
	for _, perm := range perms {
		for _, op := range ops {
			newPerms = append(newPerms, arrz.Join1D(perm, []byte{op}))
		}
	}
	return newPerms
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gResults = iutils.ExtractInt641DFromString1D(lines, ":", 0, int64(-1))
	eqStr := iutils.ExtractString1DFromString1D(lines, ": ", 1, "")
	gEquations = iutils.ExtractInt642DFromString1D(eqStr, " ", nil, -1)
	gSlotCount = 0
	for _, eq := range gEquations {
		gSlotCount = intz.Max(gSlotCount, len(eq)-1)
	}
	//fmt.Println(gSlotCount)
	//fmt.Println(eqStr)
	//intz.PPrint2D(gEquations)
}
