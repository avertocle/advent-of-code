package day20

import (
	"fmt"
	"github.com/avertocle/contests/io/errz"
	"github.com/avertocle/contests/io/iutils"
	"github.com/avertocle/contests/io/stringz"
	"strings"
)

var gInputCons map[string][]string
var gInputMods map[string]module
var gInputP2RefArray []string

const DirPath = "../2023/day20"

const (
	FF = iota
	CJ
	BC
)

func SolveP1() string {
	ans := 0
	currState := newState()
	for i := 0; i < 1000; i++ {
		currState.startNewButtonPress()
		currState.updateForP1(false, 1)
		processBfs([]*signal{{"button", "broadcaster", false}}, currState)
	}
	//fmt.Println(currState.str())
	ans = currState.calcAnsP1()
	return fmt.Sprintf("%v", ans)
}

func SolveP2() string {
	ans := 0
	currState := newState()
	for i := 0; i < 10000; i++ {
		currState.startNewButtonPress()
		ans = currState.hasConvergedP2()
		if ans != 0 {
			break
		}
		processBfs([]*signal{{"button", "broadcaster", false}}, currState)
	}
	return fmt.Sprintf("%v", ans)
}

func processBfs(currSigs []*signal, currState *state) {
	currState.signalCtr++
	if len(currSigs) == 0 {
		return
	}
	nextSigs := make([]*signal, 0)
	for _, f := range currSigs {
		nextSigs = append(nextSigs, processOneSignal(f, currState)...)
	}
	processBfs(nextSigs, currState)
}

func processOneSignal(sig *signal, currState *state) []*signal {
	nextMap := make([]*signal, 0)
	m, ok := gInputMods[sig.dst]
	if !ok {
		return nil
	}
	m.receive(sig.src, sig.pulseType)
	emitPulse, hasEmit := m.emit()
	if destinations, ok := gInputCons[m.name()]; ok && hasEmit {
		for _, dest := range destinations {
			nextMap = append(nextMap, &signal{sig.dst, dest, emitPulse})
			currState.updateForP2(sig.dst, emitPulse)
		}
		currState.updateForP1(emitPulse, len(destinations))
	}
	return nextMap
}

/***** Common Functions *****/

/***** P1 Functions *****/

/***** P2 Functions *****/

/***** Input *****/

func ParseInput(inputFilePath string) {
	lines, err := iutils.FromFile(inputFilePath, false)
	errz.HardAssert(err == nil, "iutils error | %v", err)
	gInputMods = make(map[string]module)
	gInputCons = make(map[string][]string)
	for _, line := range lines {
		m := parseModuleFromInputLine(line)
		gInputMods[m.name()] = m
		t := strings.Fields(line)[2:]
		cons := make([]string, 0)
		for _, n := range t {
			cons = append(cons, strings.TrimSuffix(n, ","))
		}
		gInputCons[m.name()] = cons
	}

	for name, mod := range gInputMods {
		if mod.typez() != CJ {
			continue
		}
		inConns := make([]string, 0)
		for inName, cons := range gInputCons {
			if stringz.Has1D(cons, name) {
				inConns = append(inConns, inName)
			}
		}
		cjMod, _ := mod.(*cjModule)
		cjMod.initMem(inConns)
		//fmt.Println(cjMod.name(), cjMod.mem)
	}
	gInputP2RefArray = []string{"qz", "cq", "jx", "tt"}
}

func parseModuleFromInputLine(line string) module {
	t := strings.Fields(line)
	var m module
	if strings.HasPrefix(t[0], "%") {
		m = newFFModule(strings.TrimPrefix(t[0], "%"))
	} else if strings.HasPrefix(t[0], "&") {
		m = newCJModule(strings.TrimPrefix(t[0], "&"))
	} else if strings.Compare(t[0], "broadcaster") == 0 {
		m = newBCModule(t[0])
	} else {
		errz.HardAssert(false, "Invalid input line | %v", line)
	}
	return m
}

/***** Interfaces *****/

type signal struct {
	src       string
	dst       string
	pulseType bool
}

type state struct {
	lowPulseCtr  int
	highPulseCtr int
	stepCounts   map[string][][]int
	buttonCtr    int
	signalCtr    int
}

func newState() *state {
	s := &state{0, 0, nil, 0, 0}
	s.stepCounts = make(map[string][][]int)
	for _, src := range gInputP2RefArray {
		s.stepCounts[src] = make([][]int, 0)
	}
	return s
}

func (s *state) updateForP1(pulseType bool, dstCount int) {
	if pulseType {
		s.highPulseCtr += dstCount
	} else {
		s.lowPulseCtr += dstCount
	}
}

func (s *state) updateForP2(src string, pulseType bool) {
	if !pulseType || !stringz.Has1D(gInputP2RefArray, src) {
		return
	}
	//fmt.Println(src, pulseType, s.buttonCtr, s.signalCtr)
	s.stepCounts[src][s.buttonCtr-1] = append(s.stepCounts[src][s.buttonCtr-1], s.signalCtr)
}

func (s *state) hasConvergedP2() int {
	convArr := make([]int, len(gInputP2RefArray))
	for refIdx, src := range gInputP2RefArray {
		for btIdx, arr := range s.stepCounts[src] {
			if convArr[refIdx] != 0 {
				break
			}
			for i := 0; i < len(arr); i++ {
				if arr[i] != 0 {
					convArr[refIdx] = btIdx + 1
					break
				}
			}
		}
	}
	return convArr[0] * convArr[1] * convArr[2] * convArr[3]
}

func (s *state) startNewButtonPress() {
	for _, src := range gInputP2RefArray {
		s.stepCounts[src] = append(s.stepCounts[src], []int{})
	}
	s.buttonCtr++
	s.signalCtr = 0
}

func (s *state) calcAnsP1() int {
	return s.lowPulseCtr * s.highPulseCtr
}

func (s *state) calcAnsP2() int {
	return 0
}

func (s *state) str() string {
	return fmt.Sprintf("%v, %v, %v", s.lowPulseCtr, s.highPulseCtr, s.stepCounts)
}

type module interface {
	name() string
	typez() int
	receive(string, bool)
	emit() (bool, bool) // [pulse, hasEmit]
}

type ffModule struct {
	n   string
	t   int
	mem []bool // [curr-state, prev-state]
}

func newFFModule(name string) *ffModule {
	return &ffModule{n: name, t: FF, mem: []bool{false, false}}
}

func (o *ffModule) name() string {
	return o.n
}

func (o *ffModule) typez() int {
	return o.t
}

func (o *ffModule) receive(sender string, pulse bool) {
	if pulse {
		o.mem = []bool{o.mem[0], o.mem[0]} // high pulse is ignored
	} else {
		o.mem = []bool{!o.mem[0], o.mem[0]}
	}
}

func (o *ffModule) emit() (bool, bool) {
	if !o.mem[1] && o.mem[0] {
		return true, true
	} else if o.mem[1] && !o.mem[0] {
		return false, true
	}
	return false, false
}

type cjModule struct {
	n   string
	t   int
	mem map[string]bool
}

func newCJModule(name string) *cjModule {
	return &cjModule{n: name, t: CJ, mem: make(map[string]bool)}
}

func (o *cjModule) name() string {
	return o.n
}

func (o *cjModule) typez() int {
	return o.t
}

func (o *cjModule) receive(sender string, pulse bool) {
	o.mem[sender] = pulse
}

func (o cjModule) emit() (bool, bool) {
	for _, v := range o.mem {
		if !v {
			return true, true
		}
	}
	return false, true
}

func (o *cjModule) initMem(inConns []string) {
	for _, c := range inConns {
		o.mem[c] = false
	}
}

type bcModule struct {
	n   string
	t   int
	mem bool
}

func newBCModule(name string) *bcModule {
	return &bcModule{n: name, t: BC}
}
func (o *bcModule) name() string {
	return o.n
}

func (o *bcModule) typez() int {
	return o.t
}

func (o *bcModule) receive(sender string, pulse bool) {
	o.mem = pulse
}

func (o *bcModule) emit() (bool, bool) {
	return o.mem, true
}
