package day24

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"github.com/avertocle/contests/io/tpz"
	"math"
	"slices"
	"strconv"
	"strings"
)

const DirPath = "../2024/day24"

var gGates []*gate
var gWires tpz.Set[string]
var gInit map[string]int

var opsMap = map[string]func(int, int) int{
	"OR":  opOr,
	"AND": opAnd,
	"XOR": opXor,
}

func SolveP1() string {
	wireState := make(map[string]int)
	for k, v := range gInit {
		wireState[k] = v
	}
	gates := cloneGates(gGates)
	_, ans := evaluateAllGates(gates, wireState)
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	/*
	   This doesn't work, just a basic attempt,
	   works for example input as there are no layers, change the op to &
	*/
	wrongWiresIdx := make([]int, 0)
	wCount := len(findWiresByPrefix("z"))
	wStart := int64(math.Pow(2, float64(wCount))) - 1
	for i := int64(0); i < wStart-1; i++ {
		sx, sy := i, i+1
		// calculate expected value
		expBin, _ := calcExpectedValue(sx, sy)
		errz.HardAssert(len(expBin) <= wCount, "expBin too long : expected(%v) found(%v) : %v", wCount, len(expBin), expBin)
		if len(expBin) < wCount {
			expBin = fmt.Sprintf("%0*s", wCount, expBin)
		}

		// calculate actual value
		gates := cloneGates(gGates)
		wireState := make(map[string]int)
		initWireState("x", sx, wireState)
		initWireState("y", sy, wireState)
		actBin, _ := evaluateAllGates(gates, wireState)

		// compare
		wrongWiresIdx = findWireMismatch(expBin, actBin)
		//fmt.Printf("%v + %v\na[%v]\ne[%v]\nwrong (%v)\n", sx, sy, actBin, expBin, wrongWiresIdx)
		if len(wrongWiresIdx) == 4 {
			break
		}
	}

	slices.Sort(wrongWiresIdx)
	ans := ""
	for _, idx := range wrongWiresIdx {
		ans = fmt.Sprintf("%v,z%v", ans, idx)
	}
	ans = strings.TrimLeft(ans, ",")
	return fmt.Sprintf("%v", ans)
}

/***** P1 Functions *****/

/***** P2 Functions *****/

func initWireState(prefix string, val int64, wireState map[string]int) {
	wCount := len(findWiresByPrefix(prefix))
	valStr := fmt.Sprintf("%0*s", wCount, strconv.FormatInt(val, 2))
	for i := 0; i < wCount; i++ {
		wireState[fmt.Sprintf("%v%02d", prefix, i)] = int(valStr[len(valStr)-i-1]) - '0'
	}
}

func calcExpectedValue(val1, val2 int64) (string, int64) {
	zValExp := val1 & val2
	zValExpStr := strconv.FormatInt(zValExp, 2)
	return zValExpStr, zValExp
}

func findWireMismatch(expBin, actBin string) []int {
	wrongWireIdx := make([]int, 0)
	for i := 0; i < len(expBin); i++ {
		if expBin[i] != actBin[i] {
			wrongWireIdx = append(wrongWireIdx, len(expBin)-i-1)
		}
	}
	return wrongWireIdx
}

/***** Common Functions *****/

func cloneGates(gates []*gate) []*gate {
	cg := make([]*gate, len(gates))
	for i, g := range gates {
		cg[i] = g.clone()
		errz.HardAssert(g.out == -1, "out is not -1")
	}
	return cg
}

func evaluateAllGates(gates []*gate, wireState map[string]int) (string, int64) {
	runForever := true
	for runForever {
		runForever = false
		for _, g := range gates {
			if g.out != -1 {
				continue
			}
			ws1, okw1 := wireState[g.w1]
			ws2, okw2 := wireState[g.w2]
			if okw1 && okw2 {
				g.out = g.op(ws1, ws2)
				wireState[g.wo] = g.out
			} else {
				runForever = true
			}
		}
	}
	zWires := findWiresByPrefix("z")
	return makeNumbersFromWires(wireState, zWires)
}

func findWiresByPrefix(prefix string) []string {
	wires := make([]string, 0)
	for k, _ := range gWires {
		if strings.HasPrefix(k, prefix) {
			wires = append(wires, k)
		}
	}
	return wires
}

func makeNumbersFromWires(wireState map[string]int, wires []string) (string, int64) {
	slices.Sort(wires)
	base2 := ""
	for i := 0; i < len(wires); i++ {
		base2 = fmt.Sprintf("%v%v", wireState[wires[i]], base2)
	}
	base10, err := strconv.ParseInt(base2, 2, 64)
	errz.HardAssert(err == nil, "strconv error : ansStr(%v) wires(%v) | %v ", base2, wires, err)
	return base2, base10
}

/***** structs *****/

type gate struct {
	w1, w2, wo string
	op         func(int, int) int
	opName     string
	out        int
}

func (g *gate) hasTerminalInput() bool {
	return strings.HasPrefix(g.w1, "x") || strings.HasPrefix(g.w1, "y") || strings.HasPrefix(g.w1, "z") ||
		strings.HasPrefix(g.w2, "x") || strings.HasPrefix(g.w2, "y") || strings.HasPrefix(g.w2, "z")
}

func (g *gate) hasTerminalOutput() bool {
	return strings.HasPrefix(g.wo, "x") || strings.HasPrefix(g.wo, "y") ||
		strings.HasPrefix(g.wo, "z")
}

func (g *gate) eval(state map[string]int) int {
	return g.op(state[g.w1], state[g.w2])
}

func (g *gate) clone() *gate {
	return &gate{w1: g.w1, w2: g.w2, wo: g.wo, op: g.op, opName: g.opName, out: g.out}
}

func (g *gate) Str() string {
	return fmt.Sprintf("[%v %v %v] = %v = %v", g.w1, g.opName, g.w2, g.wo, g.out)
}

func opOr(a, b int) int {
	return a | b
}

func opAnd(a, b int) int {
	return a & b
}

func opXor(a, b int) int {
	return a ^ b
}

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	sections := iutils.BreakByEmptyLineString1D(lines)
	gInit = make(map[string]int)
	for _, line := range sections[0] {
		t := stringz.SplitMultiTrimSpace(line, []string{":", " "})
		gInit[t[0]] = stringz.AtoI(t[1], -1)
	}

	gGates = make([]*gate, 0)
	gWires = make(tpz.Set[string])
	for _, line := range sections[1] {
		t := stringz.SplitMultiTrimSpace(line, []string{" ", "->"})
		g := &gate{w1: t[0], w2: t[2], wo: t[3], op: opsMap[t[1]], opName: t[1], out: -1}
		gGates = append(gGates, g)
		gWires[g.w1] = true
		gWires[g.w2] = true
		gWires[g.wo] = true
	}
}
