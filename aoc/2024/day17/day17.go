package day17

import (
	"fmt"
	"github.com/avertocle/contests/io/arrz"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"math"
	"strings"
)

const DirPath = "../2024/day17"

type inst func(int, *regs)

var gInput []int
var gReg []int64
var gIns []inst

func SolveP1() string {
	r := &regs{gReg[0], gReg[1], gReg[2], []int{}, -1}
	runProgram(r, false)
	ans := arrz.ToStr1D(r.o, ",")
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := int64(0)
	origProg := arrz.ToStr1D(gInput, ",")
	res := ""
	ra := int64(math.Pow(8, float64(len(gInput)-1)))
	for ctr := int64(0); res != origProg; ctr++ {
		r := &regs{ra + ctr, gReg[1], gReg[2], []int{}, -1}
		isQuine := runProgram(r, false)
		fmt.Println(ra+ctr, len(origProg), len(res), origProg, arrz.ToStr1D(r.o, ","), r.a)
		isQuine = !isQuine
		if isQuine || ctr > 10000000 {
			ans = ra + ctr
			break
		}
	}
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

func runProgram(r *regs, checkQuine bool) bool {
	for i := 0; i < len(gInput); i += 2 {
		gIns[gInput[i]](gInput[i+1], r)
		if r.j != -1 {
			i = r.j - 2
			r.j = -1
		}
		if checkQuine {
			for p := 0; p < len(r.o) && p < len(gInput); p++ {
				if r.o[p] != gInput[p] {
					return false
				}
			}
		}
	}
	return true
}

/***** P2 Functions *****/

/***** Common Functions *****/

func adv(x int, r *regs) {
	r.a = int64(float64(r.a) / math.Pow(2, float64(oprCombo(x, r))))
}

func bxl(x int, r *regs) {
	r.b = r.b ^ oprLiteral(x)
}

func bst(x int, r *regs) {
	r.b = oprCombo(x, r) % 8
}

func jnz(x int, r *regs) {
	if r.a != 0 {
		r.j = int(oprLiteral(x))
	}
}

func bxc(x int, r *regs) {
	r.b = r.b ^ r.c
}

func out(x int, r *regs) {
	r.o = append(r.o, int(oprCombo(x, r)%8))
}

func bdv(x int, r *regs) {
	r.b = int64(float64(r.a) / math.Pow(2, float64(oprCombo(x, r))))
}

func cdv(x int, r *regs) {
	r.c = int64(float64(r.a) / math.Pow(2, float64(oprCombo(x, r))))
}

func oprCombo(x int, r *regs) int64 {
	switch x {
	case 0, 1, 2, 3:
		return int64(x)
	case 4:
		return r.a
	case 5:
		return r.b
	case 6:
		return r.c
	default:
		errz.HardAssert(false, "invalid oprCombo : %v", x)
		return -1
	}
}

func runTests() {
	r := &regs{0, 0, 9, []int{}, -1}
	gIns[2](6, r)
	errz.HardAssert(r.b == 1, "test-01 : out(%v) exp(%v)", r.b, 1)

	r = &regs{0, 29, 0, []int{}, -1}
	gIns[1](7, r)
	errz.HardAssert(r.b == 26, "test-02 : out(%v) exp(%v)", r.b, 26)

	r = &regs{0, 2024, 43690, []int{}, -1}
	gIns[4](0, r)
	errz.HardAssert(r.b == 44354, "test-03 : out(%v) exp(%v)", r.b, 44354)
	fmt.Println("all tests passed")
}

func oprLiteral(x int) int64 {
	return int64(x)
}

type regs struct {
	a int64
	b int64
	c int64
	o []int
	j int
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	sets := iutils.BreakByEmptyLineString1D(lines)
	gInput = iutils.ExtractInt1DFromString0D(strings.Fields(sets[1][0])[1], ",", -1)
	gReg = []int64{
		stringz.AtoI64(strings.Fields(sets[0][0])[2], -1),
		stringz.AtoI64(strings.Fields(sets[0][1])[2], -1),
		stringz.AtoI64(strings.Fields(sets[0][2])[2], -1),
	}
	gIns = []inst{adv, bxl, bst, jnz, bxc, out, bdv, cdv}
	//fmt.Println(gInput, gReg)
}
